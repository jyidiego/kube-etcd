package storage

import (
	"bytes"
	"crypto/rand"
	"math"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/coreos/etcd/storage/storagepb"
)

func TestRange(t *testing.T) {
	s := newStore("test")
	defer os.Remove("test")

	s.Put([]byte("foo"), []byte("bar"))
	s.Put([]byte("foo1"), []byte("bar1"))
	s.Put([]byte("foo2"), []byte("bar2"))
	kvs := []storagepb.KeyValue{
		{Key: []byte("foo"), Value: []byte("bar"), CreateIndex: 1, ModIndex: 1, Version: 1},
		{Key: []byte("foo1"), Value: []byte("bar1"), CreateIndex: 2, ModIndex: 2, Version: 1},
		{Key: []byte("foo2"), Value: []byte("bar2"), CreateIndex: 3, ModIndex: 3, Version: 1},
	}

	tests := []struct {
		key, end []byte
		rev      int64

		wrev int64
		wkvs []storagepb.KeyValue
	}{
		{
			[]byte("foo"), []byte("foo3"), 0,
			3, kvs,
		},
		{
			[]byte("foo"), []byte("foo1"), 0,
			3, kvs[:1],
		},
		{
			[]byte("foo"), []byte("foo3"), 1,
			1, kvs[:1],
		},
		{
			[]byte("foo"), []byte("foo3"), 2,
			2, kvs[:2],
		},
	}

	for i, tt := range tests {
		kvs, rev, err := s.Range(tt.key, tt.end, 0, tt.rev)
		if err != nil {
			t.Fatal(err)
		}
		if rev != tt.wrev {
			t.Errorf("#%d: rev = %d, want %d", i, tt.rev, tt.wrev)
		}
		if !reflect.DeepEqual(kvs, tt.wkvs) {
			t.Errorf("#%d: kvs = %+v, want %+v", i, kvs, tt.wkvs)
		}
	}
}

func TestRangeBadRev(t *testing.T) {
	s := newStore("test")
	defer os.Remove("test")

	s.Put([]byte("foo"), []byte("bar"))
	s.Put([]byte("foo1"), []byte("bar1"))
	s.Put([]byte("foo2"), []byte("bar2"))
	if err := s.Compact(3); err != nil {
		t.Fatalf("compact error (%v)", err)
	}

	tests := []struct {
		rev  int64
		werr error
	}{
		{2, ErrCompacted},
		{3, ErrCompacted},
		{4, ErrFutureRev},
	}
	for i, tt := range tests {
		_, _, err := s.Range([]byte("foo"), []byte("foo3"), 0, tt.rev)
		if err != tt.werr {
			t.Errorf("#%d: error = %v, want %v", i, err, tt.werr)
		}
	}
}

func TestRangeLimit(t *testing.T) {
	s := newStore("test")
	defer os.Remove("test")

	s.Put([]byte("foo"), []byte("bar"))
	s.Put([]byte("foo1"), []byte("bar1"))
	s.Put([]byte("foo2"), []byte("bar2"))
	s.DeleteRange([]byte("foo1"), nil)
	kvs := []storagepb.KeyValue{
		{Key: []byte("foo"), Value: []byte("bar"), CreateIndex: 1, ModIndex: 1, Version: 1},
		{Key: []byte("foo2"), Value: []byte("bar2"), CreateIndex: 3, ModIndex: 3, Version: 1},
	}

	tests := []struct {
		limit int64
		wkvs  []storagepb.KeyValue
	}{
		// no limit
		{0, kvs},
		{1, kvs[:1]},
		{2, kvs},
		{3, kvs},
	}
	for i, tt := range tests {
		kvs, _, err := s.Range([]byte("foo"), []byte("foo3"), tt.limit, 0)
		if err != nil {
			t.Fatalf("#%d: range error (%v)", i, err)
		}
		if !reflect.DeepEqual(kvs, tt.wkvs) {
			t.Errorf("#%d: kvs = %+v, want %+v", i, kvs, tt.wkvs)
		}
	}
}

func TestSimpleDeleteRange(t *testing.T) {
	tests := []struct {
		key, end []byte

		wrev int64
		wN   int64
	}{
		{
			[]byte("foo"), []byte("foo1"),
			4, 1,
		},
		{
			[]byte("foo"), []byte("foo2"),
			4, 2,
		},
		{
			[]byte("foo"), []byte("foo3"),
			4, 3,
		},
		{
			[]byte("foo3"), []byte("foo8"),
			3, 0,
		},
	}

	for i, tt := range tests {
		s := newStore("test")

		s.Put([]byte("foo"), []byte("bar"))
		s.Put([]byte("foo1"), []byte("bar1"))
		s.Put([]byte("foo2"), []byte("bar2"))

		n, rev := s.DeleteRange(tt.key, tt.end)
		if n != tt.wN {
			t.Errorf("#%d: n = %d, want %d", i, n, tt.wN)
		}
		if rev != tt.wrev {
			t.Errorf("#%d: rev = %d, wang %d", i, rev, tt.wrev)
		}

		os.Remove("test")
	}
}

func TestRangeInSequence(t *testing.T) {
	s := newStore("test")
	defer os.Remove("test")

	s.Put([]byte("foo"), []byte("bar"))
	s.Put([]byte("foo1"), []byte("bar1"))
	s.Put([]byte("foo2"), []byte("bar2"))

	// remove foo
	n, rev := s.DeleteRange([]byte("foo"), nil)
	if n != 1 || rev != 4 {
		t.Fatalf("n = %d, index = %d, want (%d, %d)", n, rev, 1, 4)
	}

	// before removal foo
	kvs, rev, err := s.Range([]byte("foo"), []byte("foo3"), 0, 3)
	if err != nil {
		t.Fatal(err)
	}
	if len(kvs) != 3 {
		t.Fatalf("len(kvs) = %d, want %d", len(kvs), 3)
	}

	// after removal foo
	kvs, rev, err = s.Range([]byte("foo"), []byte("foo3"), 0, 4)
	if err != nil {
		t.Fatal(err)
	}
	if len(kvs) != 2 {
		t.Fatalf("len(kvs) = %d, want %d", len(kvs), 2)
	}

	// remove again -> expect nothing
	n, rev = s.DeleteRange([]byte("foo"), nil)
	if n != 0 || rev != 4 {
		t.Fatalf("n = %d, rev = %d, want (%d, %d)", n, rev, 0, 4)
	}

	// remove foo1
	n, rev = s.DeleteRange([]byte("foo"), []byte("foo2"))
	if n != 1 || rev != 5 {
		t.Fatalf("n = %d, rev = %d, want (%d, %d)", n, rev, 1, 5)
	}

	// after removal foo1
	kvs, rev, err = s.Range([]byte("foo"), []byte("foo3"), 0, 5)
	if err != nil {
		t.Fatal(err)
	}
	if len(kvs) != 1 {
		t.Fatalf("len(kvs) = %d, want %d", len(kvs), 1)
	}

	// remove foo2
	n, rev = s.DeleteRange([]byte("foo2"), []byte("foo3"))
	if n != 1 || rev != 6 {
		t.Fatalf("n = %d, rev = %d, want (%d, %d)", n, rev, 1, 6)
	}

	// after removal foo2
	kvs, rev, err = s.Range([]byte("foo"), []byte("foo3"), 0, 6)
	if err != nil {
		t.Fatal(err)
	}
	if len(kvs) != 0 {
		t.Fatalf("len(kvs) = %d, want %d", len(kvs), 0)
	}
}

func TestOneTnx(t *testing.T) {
	s := newStore("test")
	defer os.Remove("test")

	id := s.TnxBegin()
	for i := 0; i < 3; i++ {
		s.TnxPut(id, []byte("foo"), []byte("bar"))
		s.TnxPut(id, []byte("foo1"), []byte("bar1"))
		s.TnxPut(id, []byte("foo2"), []byte("bar2"))

		// remove foo
		n, rev, err := s.TnxDeleteRange(id, []byte("foo"), nil)
		if err != nil {
			t.Fatal(err)
		}
		if n != 1 || rev != 1 {
			t.Fatalf("n = %d, rev = %d, want (%d, %d)", n, rev, 1, 1)
		}

		kvs, rev, err := s.TnxRange(id, []byte("foo"), []byte("foo3"), 0, 0)
		if err != nil {
			t.Fatal(err)
		}
		if len(kvs) != 2 {
			t.Fatalf("len(kvs) = %d, want %d", len(kvs), 2)
		}

		// remove again -> expect nothing
		n, rev, err = s.TnxDeleteRange(id, []byte("foo"), nil)
		if err != nil {
			t.Fatal(err)
		}
		if n != 0 || rev != 1 {
			t.Fatalf("n = %d, rev = %d, want (%d, %d)", n, rev, 0, 1)
		}

		// remove foo1
		n, rev, err = s.TnxDeleteRange(id, []byte("foo"), []byte("foo2"))
		if err != nil {
			t.Fatal(err)
		}
		if n != 1 || rev != 1 {
			t.Fatalf("n = %d, rev = %d, want (%d, %d)", n, rev, 1, 1)
		}

		// after removal foo1
		kvs, rev, err = s.TnxRange(id, []byte("foo"), []byte("foo3"), 0, 0)
		if err != nil {
			t.Fatal(err)
		}
		if len(kvs) != 1 {
			t.Fatalf("len(kvs) = %d, want %d", len(kvs), 1)
		}

		// remove foo2
		n, rev, err = s.TnxDeleteRange(id, []byte("foo2"), []byte("foo3"))
		if err != nil {
			t.Fatal(err)
		}
		if n != 1 || rev != 1 {
			t.Fatalf("n = %d, rev = %d, want (%d, %d)", n, rev, 1, 1)
		}

		// after removal foo2
		kvs, rev, err = s.TnxRange(id, []byte("foo"), []byte("foo3"), 0, 0)
		if err != nil {
			t.Fatal(err)
		}
		if len(kvs) != 0 {
			t.Fatalf("len(kvs) = %d, want %d", len(kvs), 0)
		}
	}
	err := s.TnxEnd(id)
	if err != nil {
		t.Fatal(err)
	}

	// After tnx
	kvs, rev, err := s.Range([]byte("foo"), []byte("foo3"), 0, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(kvs) != 0 {
		t.Fatalf("len(kvs) = %d, want %d", len(kvs), 0)
	}
	if rev != 1 {
		t.Fatalf("rev = %d, want %d", rev, 1)
	}
}

func TestCompaction(t *testing.T) {
	s := newStore("test")
	defer os.Remove("test")

	s.Put([]byte("foo"), []byte("bar"))
	s.Put([]byte("foo1"), []byte("bar1"))
	s.Put([]byte("foo2"), []byte("bar2"))
	s.Put([]byte("foo"), []byte("bar11"))
	s.Put([]byte("foo1"), []byte("bar12"))
	s.Put([]byte("foo2"), []byte("bar13"))
	s.Put([]byte("foo1"), []byte("bar14"))
	s.DeleteRange([]byte("foo"), []byte("foo200"))
	s.Put([]byte("foo4"), []byte("bar4"))

	err := s.Compact(4)
	if err != nil {
		t.Errorf("unexpect compact error %v", err)
	}

	err = s.Compact(4)
	if err != ErrCompacted {
		t.Errorf("err = %v, want %v", err, ErrCompacted)
	}

	_, _, err = s.Range([]byte("foo"), nil, 0, 4)
	if err != ErrCompacted {
		t.Errorf("err = %v, want %v", err, ErrCompacted)
	}

	// compact should not compact the last value of foo
	kvs, rev, err := s.Range([]byte("foo"), nil, 0, 5)
	if err != nil {
		t.Errorf("unexpected range error %v", err)
	}
	if !bytes.Equal(kvs[0].Value, []byte("bar11")) {
		t.Errorf("value = %s, want %s", string(kvs[0].Value), "bar11")
	}
	if rev != 5 {
		t.Errorf("rev = %d, want %d", rev, 5)
	}

	// compact everything
	err = s.Compact(8)
	if err != nil {
		t.Errorf("unexpect compact error %v", err)
	}

	kvs, rev, err = s.Range([]byte("foo"), []byte("fop"), 0, 0)
	if err != nil {
		t.Errorf("unexpected range error %v", err)
	}
	if len(kvs) != 1 {
		t.Errorf("len(kvs) = %d, want %d", len(kvs), 1)
	}
	if !bytes.Equal(kvs[0].Value, []byte("bar4")) {
		t.Errorf("value = %s, want %s", string(kvs[0].Value), "bar4")
	}
	if rev != 9 {
		t.Errorf("rev = %d, want %d", rev, 9)
	}
}

func TestRestore(t *testing.T) {
	s0 := newStore("test")
	defer os.Remove("test")

	s0.Put([]byte("foo"), []byte("bar"))
	s0.Put([]byte("foo1"), []byte("bar1"))
	s0.Put([]byte("foo2"), []byte("bar2"))
	s0.Put([]byte("foo"), []byte("bar11"))
	s0.Put([]byte("foo1"), []byte("bar12"))
	s0.Put([]byte("foo2"), []byte("bar13"))
	s0.Put([]byte("foo1"), []byte("bar14"))
	s0.Put([]byte("foo3"), []byte("bar3"))
	s0.DeleteRange([]byte("foo3"), nil)
	s0.Put([]byte("foo3"), []byte("bar31"))
	s0.DeleteRange([]byte("foo3"), nil)

	mink := newRevBytes()
	revToBytes(reversion{main: 0, sub: 0}, mink)
	maxk := newRevBytes()
	revToBytes(reversion{main: math.MaxInt64, sub: math.MaxInt64}, maxk)
	s0kvs, _, err := s0.rangeKeys(mink, maxk, 0, 0)
	if err != nil {
		t.Fatalf("rangeKeys on s0 error (%v)", err)
	}

	s0.Close()

	s1 := newStore("test")
	s1.Restore()

	if !s0.Equal(s1) {
		t.Errorf("not equal!")
	}
	s1kvs, _, err := s1.rangeKeys(mink, maxk, 0, 0)
	if err != nil {
		t.Fatalf("rangeKeys on s1 error (%v)", err)
	}
	if !reflect.DeepEqual(s1kvs, s0kvs) {
		t.Errorf("s1kvs = %+v, want %+v", s1kvs, s0kvs)
	}
}

func TestRestoreContinueUnfinishedCompaction(t *testing.T) {
	s0 := newStore("test")
	defer os.Remove("test")

	s0.Put([]byte("foo"), []byte("bar"))
	s0.Put([]byte("foo"), []byte("bar1"))
	s0.Put([]byte("foo"), []byte("bar2"))

	// write scheduled compaction, but not do compaction
	rbytes := newRevBytes()
	revToBytes(reversion{main: 2}, rbytes)
	tx := s0.b.BatchTx()
	tx.Lock()
	tx.UnsafePut(metaBucketName, scheduledCompactKeyName, rbytes)
	tx.Unlock()

	s0.Close()

	s1 := newStore("test")
	s1.Restore()

	// wait for scheduled compaction to be finished
	time.Sleep(100 * time.Millisecond)

	if _, _, err := s1.Range([]byte("foo"), nil, 0, 2); err != ErrCompacted {
		t.Errorf("range on compacted rev error = %v, want %v", err, ErrCompacted)
	}
	// check the key in backend is deleted
	revbytes := newRevBytes()
	// TODO: compact should delete main=2 key too
	revToBytes(reversion{main: 1}, revbytes)
	tx = s1.b.BatchTx()
	tx.Lock()
	ks, _ := tx.UnsafeRange(keyBucketName, revbytes, nil, 0)
	if len(ks) != 0 {
		t.Errorf("key for rev %+v still exists, want deleted", bytesToRev(revbytes))
	}
	tx.Unlock()
}

func BenchmarkStorePut(b *testing.B) {
	s := newStore("test")
	defer os.Remove("test")

	// prepare keys
	keys := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = make([]byte, 64)
		rand.Read(keys[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Put(keys[i], []byte("foo"))
	}
}
