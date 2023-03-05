package validate

import (
	"errors"
	"golang.org/x/exp/constraints"
	"regexp"
	"strconv"
)

type Rule func() error

func Between[T constraints.Ordered](name string, value, min, max T) Rule {
	return func() error {
		if value < min || value > max {
			return errors.New(name + " is out of range")
		}
		return nil
	}
}

func Exists(name string, value interface{}) Rule {
	return func() error {
		if value == nil {
			return errors.New(name + " is required")
		}
		return nil
	}
}

func IsInt(name string, value string, bitSize int) Rule {
	return func() error {
		if _, err := strconv.ParseInt(value, 10, bitSize); err != nil {
			return errors.New(name + " is not an integer")
		}
		return nil
	}
}

func Length[T any](name string, value []T, max int) Rule {
	return func() error {
		if len(value) > max {
			return errors.New(name + " is out of range")
		}
		return nil
	}
}

func Matches(name string, value string, regex *regexp.Regexp) Rule {
	return func() error {
		if !regex.MatchString(value) {
			return errors.New(name + " is invalid")
		}
		return nil
	}
}

func NotEmpty(name string, value []interface{}) Rule {
	return func() error {
		if len(value) == 0 {
			return errors.New(name + " is required")
		}
		return nil
	}
}
