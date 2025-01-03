// Code generated by github.com/arikkfir/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"

type ObjectAdditionalProperties struct {
	// Foo corresponds to the JSON schema field "foo".
	Foo ObjectAdditionalPropertiesFoo `json:"foo,omitempty" yaml:"foo,omitempty" mapstructure:"foo,omitempty"`
}

type ObjectAdditionalPropertiesFoo map[string]string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectAdditionalProperties) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain ObjectAdditionalProperties
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["foo"]; !ok || v == nil {
		plain.Foo = map[string]string{}
	}
	*j = ObjectAdditionalProperties(plain)
	return nil
}
