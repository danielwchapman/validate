package validate

import (
    "errors"
    "regexp"
    "testing"
)

func Test_Between_Int(t *testing.T) {
    testcases := []struct {
        test      string
        name      string
        value     int
        min       int
        max       int
        expectErr bool
    }{
        {"Valid", "value", 5, 1, 10, false},
        {"Too Large", "value", 11, 1, 10, true},
        {"Too Small", "value", 0, 1, 10, true},
    }

    for _, tc := range testcases {
        t.Run(tc.test, func(t *testing.T) {
            err := Between(tc.name, tc.value, tc.min, tc.max)
            if tc.expectErr && err == nil {
                t.Error("Expected error, got nil")
            }
            if !tc.expectErr && err != nil {
                t.Errorf("Expected no error, got %v", err)
            }
        })
    }
}

func Test_Between_Float(t *testing.T) {
    testcases := []struct {
        test      string
        name      string
        value     float64
        min       float64
        max       float64
        expectErr bool
    }{
        {"Valid", "value", 5, 1, 10, false},
        {"Too Large", "value", 11, 1, 10, true},
        {"Too Small", "value", 0, 1, 10, true},
    }

    for _, tc := range testcases {
        t.Run(tc.test, func(t *testing.T) {
            err := Between(tc.name, tc.value, tc.min, tc.max)
            if tc.expectErr && err == nil {
                t.Error("Expected error, got nil")
            }
            if !tc.expectErr && err != nil {
                t.Errorf("Expected no error, got %v", err)
            }
        })
    }
}

func Test_Exists(t *testing.T) {
    testcases := []struct {
        test      string
        name      string
        value     interface{}
        expectErr bool
    }{
        {"Valid", "value", "test", false},
        {"Invalid", "value", nil, true},
    }

    for _, tc := range testcases {
        t.Run(tc.test, func(t *testing.T) {
            err := Exists(tc.name, tc.value)
            if tc.expectErr && err == nil {
                t.Error("Expected error, got nil")
            }
            if !tc.expectErr && err != nil {
                t.Errorf("Expected no error, got %v", err)
            }
        })
    }
}

func Test_IsInt64(t *testing.T) {
    testcases := []struct {
        test      string
        name      string
        value     string
        bitSize   int
        expectErr bool
    }{
        {"Valid", "value", "123", 32, false},
        {"Invalid", "value", "test", 64, true},
    }

    for _, tc := range testcases {
        t.Run(tc.test, func(t *testing.T) {
            err := IsInt(tc.name, tc.value, tc.bitSize)
            if tc.expectErr && err == nil {
                t.Error("Expected error, got nil")
            }
            if !tc.expectErr && err != nil {
                t.Errorf("Expected no error, got %v", err)
            }
        })
    }
}

func Test_Matches(t *testing.T) {
    testcases := []struct {
        test      string
        name      string
        value     string
        pattern   string
        expectErr bool
    }{
        {"Valid", "value", "test", "t[a-z]st", false},
        {"Invalid", "value", "test", "bad", true},
    }

    for _, tc := range testcases {
        t.Run(tc.test, func(t *testing.T) {
            regexp, _ := regexp.Compile(tc.pattern)
            err := Matches(tc.name, tc.value, regexp)
            if tc.expectErr && err == nil {
                t.Error("Expected error, got nil")
            }
            if !tc.expectErr && err != nil {
                t.Errorf("Expected no error, got %v", err)
            }
        })
    }
}

func Test_Length(t *testing.T) {
    testcases := []struct {
        test      string
        name      string
        value     []interface{}
        max       int
        expectErr bool
    }{
        {"Valid", "value", []interface{}{"test"}, 10, false},
        {"Too Large", "value", []interface{}{"test", "test"}, 1, true},
    }

    for _, tc := range testcases {
        t.Run(tc.test, func(t *testing.T) {
            err := Length(tc.name, tc.value, tc.max)
            if tc.expectErr && err == nil {
                t.Error("Expected error, got nil")
            }
            if !tc.expectErr && err != nil {
                t.Errorf("Expected no error, got %v", err)
            }
        })
    }
}

func Test_NotEmpty(t *testing.T) {
    testcases := []struct {
        test      string
        name      string
        value     []interface{}
        expectErr bool
    }{
        {"Valid", "value", []interface{}{"test"}, false},
        {"Invalid", "value", []interface{}{}, true},
        {"Invalid", "value", nil, true},
    }

    for _, tc := range testcases {
        t.Run(tc.test, func(t *testing.T) {
            err := NotEmpty(tc.name, tc.value)
            if tc.expectErr && err == nil {
                t.Error("Expected error, got nil")
            }
            if !tc.expectErr && err != nil {
                t.Errorf("Expected no error, got %v", err)
            }
        })
    }
}

func Test_ValidationChain(t *testing.T) {
    t.Run("No Errors", func(t *testing.T) {
        err := errors.Join(
            Exists("exists", "test"),
            NotEmpty("notempty", []string{"test"}),
            Between("between", 5, 1, 10),
        )

        if err != nil {
            t.Error("Expected no error, got", err)
        }
    })

    t.Run("One Error", func(t *testing.T) {
        err := errors.Join(
            Exists("exists", "test"),
            NotEmpty("notempty", []string{}),
            Between("between", 5, 1, 10),
        )

        if err == nil {
            t.Error("Expected error")
        }
    })

    t.Run("Many Error", func(t *testing.T) {
        err := errors.Join(
            Exists("exists", nil),
            NotEmpty("notempty", []string{}),
            Between("between", 11, 1, 10),
        )

        if err == nil {
            t.Error("Expected error")
        }
    })
}
