package shared

import (
	"bytes"
	"encoding/json"
)

// DecodeJSON decodes the input data into
// a JSON structure, which is stored in out
func DecodeJSON(data []byte, out interface{}) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	err := d.Decode(out)

	return err
}
