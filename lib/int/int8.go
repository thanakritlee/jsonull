package int

import (
	"encoding/json"
	"bytes"
	"strconv"
)

// Int8JSON struct.
type Int8JSON struct {
	Value int8
	Valid bool
	Set bool
}

// UnmarshalJSON unmarshaler implementation for Int8JSON (int8).
func (j *Int8JSON) UnmarshalJSON(data []byte) error {
	// If this method is call, the value is set.
	// Value could be set to either null or some non-null value.
	j.Set = true

	if string(data) == "null" {
		// The value is set to null.
		j.Valid = false
		return nil
	}

	// The value isn't set to null.
	var temp int8
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	j.Value = temp
	j.Valid = true
	return nil
}

// MarshalJSON marshaler implementation for Int8JSON (int8).
func (j Int8JSON) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("")
	if j.Valid {
		buffer.WriteString(strconv.FormatInt(int64(j.Value), 10))
	} else {
		buffer.WriteString("null")
	}

	return buffer.Bytes(), nil
}