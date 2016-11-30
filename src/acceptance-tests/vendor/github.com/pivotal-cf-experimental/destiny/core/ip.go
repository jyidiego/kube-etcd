package core

import (
	"fmt"
	"strconv"
	"strings"
)

type IP struct {
	ip int
}

func ParseIP(ip string) (IP, error) {
	const IP_PARTS = 4
	const MAX_IP_PART = 256

	ipParts := strings.Split(ip, ".")

	if len(ipParts) != IP_PARTS {
		return IP{}, fmt.Errorf(`'%s' is not a valid ip address`, ip)
	}

	ipValue := 0
	for _, ipPart := range ipParts {
		ipPartInt, err := strconv.Atoi(ipPart)
		if err != nil {
			return IP{}, err
		}

		if ipPartInt < 0 || ipPartInt >= MAX_IP_PART {
			return IP{}, fmt.Errorf("invalid ip, %s has values out of range", ip)
		}

		ipValue = ipValue*MAX_IP_PART + ipPartInt
	}

	return IP{
		ip: ipValue,
	}, nil
}

func (i IP) Add(offset int) IP {
	return IP{
		ip: i.ip + offset,
	}
}

func (i IP) Subtract(offset int) IP {
	return IP{
		ip: i.ip - offset,
	}
}

func (i IP) To(ip IP) []string {
	var firstIP IP
	var secondIP IP
	if i.ip < ip.ip {
		firstIP = i
		secondIP = ip
	} else {
		firstIP = ip
		secondIP = i
	}

	ips := []string{}
	delta := secondIP.ip - firstIP.ip
	for c := 0; c <= delta; c++ {
		ips = append(ips, firstIP.Add(c).String())
	}
	return ips
}

func (i IP) String() string {
	first := i.ip & 0xff000000 >> 24
	second := i.ip & 0xff0000 >> 16
	third := i.ip & 0xff00 >> 8
	fourth := i.ip & 0xff
	return fmt.Sprintf("%v.%v.%v.%v", first, second, third, fourth)
}
