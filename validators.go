package validate

import (
    "errors"
    "golang.org/x/exp/constraints"
    "regexp"
    "strconv"
)

type Rule func() error

// Between checks if a value is between min and max.
func Between[T constraints.Ordered](name string, value, min, max T) Rule {
    return func() error {
        if value < min || value > max {
            return errors.New(name + " is out of range")
        }
        return nil
    }
}

// Exists checks if a value is not nil.
func Exists(name string, value interface{}) Rule {
    return func() error {
        if value == nil {
            return errors.New(name + " is required")
        }
        return nil
    }
}

// IsInt checks if a string value is an integer.
func IsInt(name string, value string, bitSize int) Rule {
    return func() error {
        if _, err := strconv.ParseInt(value, 10, bitSize); err != nil {
            return errors.New(name + " is not an integer")
        }
        return nil
    }
}

// Length checks if a slice is of a certain length.
func Length[T any](name string, value []T, max int) Rule {
    return func() error {
        if len(value) > max {
            return errors.New(name + " is out of range")
        }
        return nil
    }
}

// Matches checks if a string value matches a regular expression.
func Matches(name string, value string, regex *regexp.Regexp) Rule {
    return func() error {
        if !regex.MatchString(value) {
            return errors.New(name + " is invalid")
        }
        return nil
    }
}

// NotEmpty checks if a slice is not empty.
func NotEmpty(name string, value []interface{}) Rule {
    return func() error {
        if len(value) == 0 {
            return errors.New(name + " is required")
        }
        return nil
    }
}
