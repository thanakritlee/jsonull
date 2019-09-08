package uint

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Uint32 struct.
type Uint32 struct {
	Value uint32
	Valid bool
	Set   bool
}

// UnmarshalJSON unmarshaler implementation for Uint32.
func (j *Uint32) UnmarshalJSON(data []byte) error {
	// If this method is call, the value is set.
	// Value could be set to either null or some non-null value.
	j.Set = true

	if string(data) == "null" {
		// The value is set to null.
		j.Valid = false
		return nil
	}

	// The value isn't set to null.
	var temp uint32
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	j.Value = temp
	j.Valid = true
	return nil
}

// MarshalJSON marshaler implementation for Uint32.
func (j Uint32) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("")
	if j.Valid {
		buffer.WriteString(fmt.Sprintf("%d", j.Value))
	} else {
		buffer.WriteString("null")
	}

	return buffer.Bytes(), nil
}
