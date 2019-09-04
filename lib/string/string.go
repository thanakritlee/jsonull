package string

import (
	"bytes"
	"encoding/json"
)

// String struct.
type String struct {
	Value string
	Valid bool
	Set   bool
}

// UnmarshalJSON unmarshaler implementation for String.
func (j *String) UnmarshalJSON(data []byte) error {
	// If this method is call, the value is set.
	// Value could be set to either null or some non-null value.
	j.Set = true

	if string(data) == "null" {
		// The value is set to null.
		j.Valid = false
		return nil
	}

	// The value isn't set to null.
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	j.Value = temp
	j.Valid = true
	return nil
}

// MarshalJSON marshaler implementation for String.
func (j String) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("")
	if j.Valid {
		buffer.WriteString(`"` + j.Value + `"`)
	} else {
		buffer.WriteString("null")
	}

	return buffer.Bytes(), nil
}
