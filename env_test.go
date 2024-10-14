// SPDX-License-Identifier: Apache-2.0
package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-obvious/env"
)

func TestGetArray(t *testing.T) {
	tests := []struct {
		name        string
		envValue    string
		sep         string
		allowQuotes bool
		trimQuoted  bool
		expected    []string
	}{
		{
			name:        "simple comma separated values",
			envValue:    "val1,val2,val3",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    []string{"val1", "val2", "val3"},
		},
		{
			name:        "values with spaces",
			envValue:    " val1 , val2 , val3 ",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    []string{"val1", "val2", "val3"},
		},
		{
			name:        "empty array",
			envValue:    "",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    nil,
		},
		{
			name:        "array with brackets",
			envValue:    "[val1,val2,val3]",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    []string{"val1", "val2", "val3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("TEST_ENV", tt.envValue)
			defer os.Unsetenv("TEST_ENV")

			result := env.GetArray("TEST_ENV", tt.sep, tt.allowQuotes, tt.trimQuoted)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMustGet(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected string
		panicMsg string
	}{
		{
			name:     "existing environment variable",
			envValue: "test_value",
			expected: "test_value",
			panicMsg: "",
		},
		{
			name:     "environment variable with quotes",
			envValue: `"test_value"`,
			expected: "test_value",
			panicMsg: "",
		},
		{
			name:     "environment variable with nested quotes",
			envValue: `"'test_value'"`,
			expected: "test_value",
			panicMsg: "",
		},
		{
			name:     "missing environment variable",
			envValue: "",
			expected: "",
			panicMsg: "missing env var - TEST_ENV",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv("TEST_ENV", tt.envValue)
				defer os.Unsetenv("TEST_ENV")
			} else {
				os.Unsetenv("TEST_ENV")
			}

			if tt.panicMsg != "" {
				defer func() {
					if r := recover(); r != nil {
						assert.Equal(t, tt.panicMsg, r)
					} else {
						t.Errorf("expected panic with message: %s", tt.panicMsg)
					}
				}()
			}

			result := env.MustGet("TEST_ENV")
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected bool
	}{
		{
			name:     "existing environment variable",
			envValue: "test_value",
			expected: true,
		},
		{
			name:     "missing environment variable",
			envValue: "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv("TEST_ENV", tt.envValue)
				defer os.Unsetenv("TEST_ENV")
			} else {
				os.Unsetenv("TEST_ENV")
			}

			result := env.Exists("TEST_ENV")
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetBoolOr(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		dval     bool
		expected bool
	}{
		{
			name:     "existing environment variable true",
			envValue: "true",
			dval:     false,
			expected: true,
		},
		{
			name:     "existing environment variable false",
			envValue: "false",
			dval:     true,
			expected: false,
		},
		{
			name:     "missing environment variable default true",
			envValue: "",
			dval:     true,
			expected: true,
		},
		{
			name:     "missing environment variable default false",
			envValue: "",
			dval:     false,
			expected: false,
		},
		{
			name:     "case insensitive true",
			envValue: "TrUe",
			dval:     false,
			expected: true,
		},
		{
			name:     "case insensitive false",
			envValue: "FaLsE",
			dval:     true,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv("TEST_BOOL_ENV", tt.envValue)
				defer os.Unsetenv("TEST_BOOL_ENV")
			} else {
				os.Unsetenv("TEST_BOOL_ENV")
			}

			result := env.GetBoolOr("TEST_BOOL_ENV", tt.dval)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetUintOr(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		dval     uint
		expected uint
	}{
		{
			name:     "existing environment variable",
			envValue: "42",
			dval:     0,
			expected: 42,
		},
		{
			name:     "missing environment variable",
			envValue: "",
			dval:     100,
			expected: 100,
		},
		{
			name:     "invalid environment variable",
			envValue: "invalid",
			dval:     50,
			expected: 50, // This will panic, so we expect the default value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv("TEST_UINT_ENV", tt.envValue)
				defer os.Unsetenv("TEST_UINT_ENV")
			} else {
				os.Unsetenv("TEST_UINT_ENV")
			}

			if tt.envValue == "invalid" {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected panic for invalid uint value")
					}
				}()
			}

			result := env.GetUintOr("TEST_UINT_ENV", tt.dval)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMustGetArray(t *testing.T) {
	tests := []struct {
		name        string
		envValue    string
		sep         string
		allowQuotes bool
		trimQuoted  bool
		expected    []string
		panicMsg    string
	}{
		{
			name:        "simple comma separated values",
			envValue:    "val1,val2,val3",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    []string{"val1", "val2", "val3"},
			panicMsg:    "",
		},
		{
			name:        "values with spaces",
			envValue:    " val1 , val2 , val3 ",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    []string{"val1", "val2", "val3"},
			panicMsg:    "",
		},
		{
			name:        "empty array",
			envValue:    "",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    nil,
			panicMsg:    "missing env var (empty array) - TEST_ENV",
		},
		{
			name:        "array with brackets",
			envValue:    "[val1,val2,val3]",
			sep:         ",",
			allowQuotes: false,
			trimQuoted:  false,
			expected:    []string{"val1", "val2", "val3"},
			panicMsg:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv("TEST_ENV", tt.envValue)
				defer os.Unsetenv("TEST_ENV")
			} else {
				os.Unsetenv("TEST_ENV")
			}

			if tt.panicMsg != "" {
				defer func() {
					if r := recover(); r != nil {
						assert.Equal(t, tt.panicMsg, r)
					} else {
						t.Errorf("expected panic with message: %s", tt.panicMsg)
					}
				}()
			}

			result := env.MustGetArray("TEST_ENV", tt.sep, tt.allowQuotes, tt.trimQuoted)
			assert.Equal(t, tt.expected, result)
		})
	}
}
