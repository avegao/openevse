package util

import "strconv"

func ParseInt(valueToParse string) (int, error) {
	value, err := strconv.ParseInt(valueToParse, 10, 0)

	return int(value), err
}

func ParseHexInt(valueToParse string) (int, error) {
	value, err := strconv.ParseInt(valueToParse, 16, 0)

	return int(value), err
}
