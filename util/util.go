package util

import "strconv"

func ParseInt(valueToParse string) (int, error) {
	value, err := strconv.ParseInt(valueToParse, 10, 0)

	return int(value), err
}
