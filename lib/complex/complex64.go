package complex

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// Complex64JSON struct.
type Complex64JSON struct {
	Value complex64
	Valid bool
	Set bool
}

// UnmarshalJSON unmarshaler implementation for Complex64JSON (complex64).
func (j *Complex64JSON) UnmarshalJSON(data []byte) error {
	// If this method is call, the value is set.
	// Value could be set to either null or some non-null value.
	j.Set = true

	if string(data) == "null" {
		// The value is set to null.
		j.Valid = false
		return nil
	}

	// The value isn't set to null.
	var temp complex64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	j.Value = temp
	j.Valid = true
	return nil
}

// MarshalJSON marshaler implementation for Complex64JSON (complex64).
func (j Complex64JSON) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("")
	if j.Valid {
		buffer.WriteString(fmt.Sprintf("f", j.Value))
	} else {
		buffer.WriteString("null")
	}

	return buffer.Bytes(), nil
}