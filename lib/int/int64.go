package int

import (
	"encoding/json"
	"bytes"
	"strconv"
)

// Int64JSON struct.
type Int64JSON struct {
	Value int64
	Valid bool
	Set bool
}

// UnmarshalJSON unmarshaler implementation for Int64JSON (int64).
func (j *Int64JSON) UnmarshalJSON(data []byte) error {
	// If this method is call, the value is set.
	// Value could be set to either null or some non-null value.
	j.Set = true

	if string(data) == "null" {
		// The value is set to null.
		j.Valid = false
		return nil
	}

	// The value isn't set to null.
	var temp int64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	j.Value = temp
	j.Valid = true
	return nil
}

// MarshalJSON marshaler implementation for Int64JSON (int64).
func (j Int64JSON) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("")
	if j.Valid {
		buffer.WriteString(strconv.FormatInt(j.Value, 10))
	} else {
		buffer.WriteString("null")
	}

	return buffer.Bytes(), nil
}