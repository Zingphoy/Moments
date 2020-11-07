package utils

import "strconv"

// Convert string to specific type
type Str string

func (s Str) toString() string {
	return string(s)
}

func (s Str) Int() (int, error) {
	v, err := strconv.ParseInt(s.toString(), 10, 0)
	return int(v), err
}

func (s Str) Int8() (int8, error) {
	v, err := strconv.ParseUint(s.toString(), 10, 8)
	return int8(v), err
}

func (s Str) Int64() (int64, error) {
	v, err := strconv.ParseInt(s.toString(), 10, 64)
	return v, err
}

func (s Str) Int32() (int32, error) {
	v, err := strconv.ParseInt(s.toString(), 10, 32)
	return int32(v), err
}

// MustInt only for the case when must be int
func (s Str) MustInt() int {
	v, _ := s.Int()
	return v
}

// MustInt8 only for the case when must be unit8
func (s Str) MustUint8() int8 {
	v, _ := s.Int8()
	return v
}

// MustInt64 only for the case when must be int64
func (s Str) MustInt64() int64 {
	v, _ := s.Int64()
	return v
}

// MustInt32 only for the case when must be int32
func (s Str) MustInt32() int32 {
	v, _ := s.Int32()
	return v
}
