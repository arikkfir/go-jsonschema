// Code generated by github.com/arikkfir/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"

type Object struct {
	// MyObject corresponds to the JSON schema field "myObject".
	MyObject *ObjectMyObject `json:"myObject,omitempty" yaml:"myObject,omitempty" mapstructure:"myObject,omitempty"`
}

type ObjectMyObject struct {
	// MyString corresponds to the JSON schema field "myString".
	MyString string `json:"myString" yaml:"myString" mapstructure:"myString"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectMyObject) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["myString"]; raw != nil && !ok {
		return fmt.Errorf("field myString in ObjectMyObject: required")
	}
	type Plain ObjectMyObject
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ObjectMyObject(plain)
	return nil
}
