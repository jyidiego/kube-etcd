package gomegamatchers

import (
	"fmt"

	"github.com/cloudfoundry-incubator/candiedyaml"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/pivotal-cf-experimental/gomegamatchers/internal/deepequal"
	"github.com/pivotal-cf-experimental/gomegamatchers/internal/prettyprint"
)

func MatchYAML(expected interface{}) types.GomegaMatcher {
	return &MatchYAMLMatcher{
		YAMLToMatch: expected,
	}
}

type MatchYAMLMatcher struct {
	YAMLToMatch interface{}
}

func (matcher *MatchYAMLMatcher) Match(actual interface{}) (success bool, err error) {
	equal, _, err := matcher.equal(matcher.YAMLToMatch, actual)
	if err != nil {
		return false, err
	}

	return equal, nil
}

func (matcher *MatchYAMLMatcher) FailureMessage(actual interface{}) (message string) {
	_, message, err := matcher.equal(matcher.YAMLToMatch, actual)
	if err != nil {
		return err.Error()
	}

	return message
}

func (matcher *MatchYAMLMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	actualString, _ := matcher.prettyPrint(actual)
	expectedString, _ := matcher.prettyPrint(matcher.YAMLToMatch)
	return format.Message(actualString, "not to match YAML of", expectedString)
}

func (matcher *MatchYAMLMatcher) equal(expected interface{}, actual interface{}) (bool, string, error) {
	actualString, err := matcher.prettyPrint(actual)
	if err != nil {
		return false, "", err
	}

	expectedString, err := matcher.prettyPrint(expected)
	if err != nil {
		return false, "", err
	}

	var actualValue interface{}
	var expectedValue interface{}

	// this is guarded by prettyPrint
	candiedyaml.Unmarshal([]byte(actualString), &actualValue)
	candiedyaml.Unmarshal([]byte(expectedString), &expectedValue)

	equal, difference := deepequal.Compare(expectedValue, actualValue)

	return equal, prettyprint.ExpectationFailure(difference), nil
}

func (matcher *MatchYAMLMatcher) prettyPrint(input interface{}) (formatted string, err error) {
	inputString, ok := toString(input)
	if !ok {
		return "", fmt.Errorf("MatchYAMLMatcher matcher requires a string or stringer.  Got:\n%s", format.Object(input, 1))
	}

	var data interface{}
	if err := candiedyaml.Unmarshal([]byte(inputString), &data); err != nil {
		return "", err
	}
	buf, _ := candiedyaml.Marshal(data)

	return string(buf), nil
}

func toString(value interface{}) (string, bool) {
	valueString, isString := value.(string)
	if isString {
		return valueString, true
	}

	valueBytes, isBytes := value.([]byte)
	if isBytes {
		return string(valueBytes), true
	}

	valueStringer, isStringer := value.(fmt.Stringer)
	if isStringer {
		return valueStringer.String(), true
	}

	return "", false
}
