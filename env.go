// SPDX-FileCopyrightText: Copyright (c) 2024, Obviously, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Get(name string) string {
	return os.Getenv(name)
}

func MustGet(name string) string {
	val := Get(name)
	if val == "" {
		panic(fmt.Sprintf("missing env var - %s", name))
	}

	for i := 0; i < 3; i++ {
		// Crazy things happen in our configs...
		// Doing this 3 times because I've already seen two levels of nesting...
		val = strings.Trim(val, `"'`)
	}

	return val
}

func Exists(name string) bool {
	return Get(name) != ""
}

func GetOr(name, dval string) string {
	value := os.Getenv(name)
	if value == "" {
		value = dval
	}
	value = strings.Trim(value, "\"")
	return value
}

func GetBoolOr(name string, dval bool) bool {
	str := GetOr(name, fmt.Sprintf("%v", dval))
	return strings.ToLower(str) == "true"
}

func GetUintOr(name string, dval uint) uint {
	str := GetOr(name, fmt.Sprintf("%v", dval))
	return mustParseUint32(str)
}

func mustParseUint32(str string) uint {
	val, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		panic(err)
	}

	return uint(val)
}

func GetArray(name, sep string, allowQuotes, trimQuoted bool) []string {
	raw := GetOr(name, "")
	if raw == "" {
		return nil
	}

	if sep == "" {
		sep = ","
	}

	raw = strings.TrimSpace(raw)
	if len(raw) == 0 {
		return nil
	}

	if raw[0] == '[' && raw[len(raw)-1] == ']' {
		raw = strings.Trim(raw, "[]")
		raw = strings.TrimSpace(raw)
		raw = strings.TrimRight(raw, sep)
		raw = strings.TrimSpace(raw)
	}

	parts := strings.Split(raw, sep)

	for idx := range parts {
		parts[idx] = strings.TrimSpace(parts[idx])
	}

	if !allowQuotes {
		for idx := range parts {
			parts[idx] = strings.Trim(parts[idx], `"`)
		}

		if trimQuoted {
			for idx := range parts {
				parts[idx] = strings.TrimSpace(parts[idx])
			}
		}
	}

	var result []string
	for idx := range parts {
		if parts[idx] != "" {
			result = append(result, parts[idx])
		}
	}

	return result
}

func MustGetArray(name, sep string, allowQuotes, trimQuoted bool) []string {
	val := GetArray(name, sep, allowQuotes, trimQuoted)
	if len(val) == 0 {
		panic(fmt.Sprintf("missing env var (empty array) - %s", name))
	}

	return val
}
