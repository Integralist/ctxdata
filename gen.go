package ctxdata

import (
	"time"
)

// Code generated by generate-helpers.fish; DO NOT EDIT.

// GetStringDefault returns the value of key as type string.
// If key is not set, or its value is not of type string,
// then GetStringDefault returns defaultValue.
func (d *Data) GetStringDefault(key string, defaultValue string) string {
	val, err := d.Get(key)
	if err != nil {
		return defaultValue
	}

	res, ok := val.(string)
	if !ok {
		return defaultValue
	}

	return res
}

// GetIntDefault returns the value of key as type int.
// If key is not set, or its value is not of type int,
// then GetIntDefault returns defaultValue.
func (d *Data) GetIntDefault(key string, defaultValue int) int {
	val, err := d.Get(key)
	if err != nil {
		return defaultValue
	}

	res, ok := val.(int)
	if !ok {
		return defaultValue
	}

	return res
}

// GetInt64Default returns the value of key as type int64.
// If key is not set, or its value is not of type int64,
// then GetInt64Default returns defaultValue.
func (d *Data) GetInt64Default(key string, defaultValue int64) int64 {
	val, err := d.Get(key)
	if err != nil {
		return defaultValue
	}

	res, ok := val.(int64)
	if !ok {
		return defaultValue
	}

	return res
}

// GetUint64Default returns the value of key as type uint64.
// If key is not set, or its value is not of type uint64,
// then GetUint64Default returns defaultValue.
func (d *Data) GetUint64Default(key string, defaultValue uint64) uint64 {
	val, err := d.Get(key)
	if err != nil {
		return defaultValue
	}

	res, ok := val.(uint64)
	if !ok {
		return defaultValue
	}

	return res
}

// GetFloat64Default returns the value of key as type float64.
// If key is not set, or its value is not of type float64,
// then GetFloat64Default returns defaultValue.
func (d *Data) GetFloat64Default(key string, defaultValue float64) float64 {
	val, err := d.Get(key)
	if err != nil {
		return defaultValue
	}

	res, ok := val.(float64)
	if !ok {
		return defaultValue
	}

	return res
}

// GetBoolDefault returns the value of key as type bool.
// If key is not set, or its value is not of type bool,
// then GetBoolDefault returns defaultValue.
func (d *Data) GetBoolDefault(key string, defaultValue bool) bool {
	val, err := d.Get(key)
	if err != nil {
		return defaultValue
	}

	res, ok := val.(bool)
	if !ok {
		return defaultValue
	}

	return res
}

// GetDurationDefault returns the value of key as type time.Duration.
// If key is not set, or its value is not of type time.Duration,
// then GetDurationDefault returns defaultValue.
func (d *Data) GetDurationDefault(key string, defaultValue time.Duration) time.Duration {
	val, err := d.Get(key)
	if err != nil {
		return defaultValue
	}

	res, ok := val.(time.Duration)
	if !ok {
		return defaultValue
	}

	return res
}

// GetTimeDefault returns the value of key as type time.Time.
// If key is not set, or its value is not of type time.Time,
// then GetTimeDefault returns defaultValue.
func (d *Data) GetTimeDefault(key string, defaultValue time.Time) time.Time {
	val, err := d.Get(key)
	if err != nil {
		return defaultValue
	}

	res, ok := val.(time.Time)
	if !ok {
		return defaultValue
	}

	return res
}

// GetErrorDefault returns the value of key as type error.
// If key is not set, or its value is not of type error,
// then GetErrorDefault returns defaultValue.
func (d *Data) GetErrorDefault(key string, defaultValue error) error {
	val, err := d.Get(key)
	if err != nil {
		return defaultValue
	}

	res, ok := val.(error)
	if !ok {
		return defaultValue
	}

	return res
}

// GetString returns the value of key as type string.
// If key is not set, or its value is not of type string,
// then GetString returns the zero value for type string.
func (d *Data) GetString(key string) string {
	var ret string

	val, err := d.Get(key)
	if err != nil {
		return ret
	}

	res, ok := val.(string)
	if !ok {
		return ret
	}

	return res
}

// GetInt returns the value of key as type int.
// If key is not set, or its value is not of type int,
// then GetInt returns the zero value for type int.
func (d *Data) GetInt(key string) int {
	var ret int

	val, err := d.Get(key)
	if err != nil {
		return ret
	}

	res, ok := val.(int)
	if !ok {
		return ret
	}

	return res
}

// GetInt64 returns the value of key as type int64.
// If key is not set, or its value is not of type int64,
// then GetInt64 returns the zero value for type int64.
func (d *Data) GetInt64(key string) int64 {
	var ret int64

	val, err := d.Get(key)
	if err != nil {
		return ret
	}

	res, ok := val.(int64)
	if !ok {
		return ret
	}

	return res
}

// GetUint64 returns the value of key as type uint64.
// If key is not set, or its value is not of type uint64,
// then GetUint64 returns the zero value for type uint64.
func (d *Data) GetUint64(key string) uint64 {
	var ret uint64

	val, err := d.Get(key)
	if err != nil {
		return ret
	}

	res, ok := val.(uint64)
	if !ok {
		return ret
	}

	return res
}

// GetFloat64 returns the value of key as type float64.
// If key is not set, or its value is not of type float64,
// then GetFloat64 returns the zero value for type float64.
func (d *Data) GetFloat64(key string) float64 {
	var ret float64

	val, err := d.Get(key)
	if err != nil {
		return ret
	}

	res, ok := val.(float64)
	if !ok {
		return ret
	}

	return res
}

// GetBool returns the value of key as type bool.
// If key is not set, or its value is not of type bool,
// then GetBool returns the zero value for type bool.
func (d *Data) GetBool(key string) bool {
	var ret bool

	val, err := d.Get(key)
	if err != nil {
		return ret
	}

	res, ok := val.(bool)
	if !ok {
		return ret
	}

	return res
}

// GetDuration returns the value of key as type time.Duration.
// If key is not set, or its value is not of type time.Duration,
// then GetDuration returns the zero value for type time.Duration.
func (d *Data) GetDuration(key string) time.Duration {
	var ret time.Duration

	val, err := d.Get(key)
	if err != nil {
		return ret
	}

	res, ok := val.(time.Duration)
	if !ok {
		return ret
	}

	return res
}

// GetTime returns the value of key as type time.Time.
// If key is not set, or its value is not of type time.Time,
// then GetTime returns the zero value for type time.Time.
func (d *Data) GetTime(key string) time.Time {
	var ret time.Time

	val, err := d.Get(key)
	if err != nil {
		return ret
	}

	res, ok := val.(time.Time)
	if !ok {
		return ret
	}

	return res
}

// GetError returns the value of key as type error.
// If key is not set, or its value is not of type error,
// then GetError returns the zero value for type error.
func (d *Data) GetError(key string) error {
	var ret error

	val, err := d.Get(key)
	if err != nil {
		return ret
	}

	res, ok := val.(error)
	if !ok {
		return ret
	}

	return res
}
