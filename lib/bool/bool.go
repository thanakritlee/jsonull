package bool

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Bool struct.
type Bool struct {
	Value bool
	Valid bool
	Set   bool
}

// UnmarshalJSON unmarshaler implementation for Bool.
func (j *Bool) UnmarshalJSON(data []byte) error {
	// If this method is call, the value is set.
	// Value could be set to either null or some non-null value.
	j.Set = true

	if string(data) == "null" {
		// The value is set to null.
		j.Valid = false
		return nil
	}

	// The value isn't set to null.
	var temp bool
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	j.Value = temp
	j.Valid = true
	return nil
}

// MarshalJSON marshaler implementation for Bool.
func (j Bool) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("")
	if j.Valid {
		buffer.WriteString(strconv.FormatBool(j.Value))
	} else {
		buffer.WriteString("null")
	}

	return buffer.Bytes(), nil
}
