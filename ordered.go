// Package orderedjson allows you to unmarshal JSON objects while
// maintaining the order of the keys.
//
// Given this input:
// ```json
//
//	{
//	  "0": null,
//	  "1": 0,
//	  "2": "s",
//	  "3": [null, 0, "string", [], {}],
//	  "4": {"0": null, "1": 0, "2": "s", "3": [], "4": {}},
//	}
//
// ```
//
// Use the `orderedjson.Map` type to unmarshal:
//
// ```go
//
// var object orderedjson.Map
// err := json.Unmarshal(input, &object)
// ```
//
// The content of `object` will then be:
//
// ```go
//
//	object := orderedjson.Map{
//	  {Key: json.RawMessage(`"0"`), Value: json.RawMessage(`null`)},
//	  {Key: json.RawMessage(`"1"`), Value: json.RawMessage(`0`)},
//	  {Key: json.RawMessage(`"2"`), Value: json.RawMessage(`"s"`)},
//	  {Key: json.RawMessage(`"3"`), Value: json.RawMessage(`[null, 0, "string", [], {}]`)},
//	  {Key: json.RawMessage(`"4"`), Value: json.RawMessage(`{"0": null, "1": 0, "2": "s", "3": [], "4": {}}`)},
//	}
//
// ```
//
// And you can continue unmarshalling `Key` and `Value` however you wish.
package orderedjson

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/aybabtme/flatjson"
)

// Map contains the ordered keys that are parsed from the data.
type Map []MapEntry

// MapEntry is a key-value in the unmarshaled JSON object.
type MapEntry struct {
	Key   json.RawMessage
	Value json.RawMessage
}

func (m *Map) UnmarshalJSON(data []byte) error {
	_, found, err := flatjson.ScanObject(data, 0, &flatjson.Callbacks{
		OnRaw: func(prefixes flatjson.Prefixes, name flatjson.Prefix, value flatjson.Pos) {
			entry := MapEntry{
				Key:   json.RawMessage(name.Bytes(data)),
				Value: json.RawMessage(value.Bytes(data)),
			}
			(*m) = append((*m), entry)
		},
	})
	if err != nil {
		return err
	}
	if !found {
		return errors.New("expected an object but none found")
	}
	return nil
}

func (m Map) MarshalJSON() ([]byte, error) {
	out := bytes.NewBuffer(nil)
	out.WriteRune('{')
	for i, kv := range m {
		if i != 0 {
			out.WriteRune(',')
		}
		out.Write(kv.Key)
		out.WriteRune(':')
		out.Write(kv.Value)
	}
	out.WriteRune('}')
	return out.Bytes(), nil
}
