package timestamp

import (
	"strconv"
	"time"
)

// Simple timestamp utility functions to ensure we always work in UTC

func Now() time.Time {
	return time.Now().UTC()
}

func Nano() int64 {
	return time.Now().UTC().UnixNano()
}

func Milli() int64 {
	return Nano() / int64(time.Millisecond)
}

func Sec() int64 {
	return time.Now().UTC().Unix()
}

func MilliFrom(nsec int64) int64 {
	return nsec / 1e6
}

func NanoStringToTime(input string) string {
	val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return ""
	}

	seconds := int64(val / int64(time.Second))
	nseconds := seconds * int64(time.Second)
	nleft := val - nseconds
	t := time.Unix(seconds, nleft)
	return t.UTC().Format(time.RFC3339Nano)
}

func ToEpoche(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}

func FromEpoche(epoche string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, epoche)
}
