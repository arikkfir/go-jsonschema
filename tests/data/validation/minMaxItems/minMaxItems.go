// Code generated by github.com/arikkfir/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"

type MinMaxItems struct {
	// MyNestedArray corresponds to the JSON schema field "myNestedArray".
	MyNestedArray [][]interface{} `json:"myNestedArray,omitempty" yaml:"myNestedArray,omitempty" mapstructure:"myNestedArray,omitempty"`

	// MyStringArray corresponds to the JSON schema field "myStringArray".
	MyStringArray []string `json:"myStringArray,omitempty" yaml:"myStringArray,omitempty" mapstructure:"myStringArray,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MinMaxItems) UnmarshalJSON(b []byte) error {
	type Plain MinMaxItems
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if plain.MyNestedArray != nil && len(plain.MyNestedArray) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "myNestedArray", 1)
	}
	if len(plain.MyNestedArray) > 5 {
		return fmt.Errorf("field %s length: must be <= %d", "myNestedArray", 5)
	}
	for i1 := range plain.MyNestedArray {
		if plain.MyNestedArray[i1] != nil && len(plain.MyNestedArray[i1]) < 1 {
			return fmt.Errorf("field %s length: must be >= %d", fmt.Sprintf("myNestedArray[%d]", i1), 1)
		}
		if len(plain.MyNestedArray[i1]) > 5 {
			return fmt.Errorf("field %s length: must be <= %d", fmt.Sprintf("myNestedArray[%d]", i1), 5)
		}
	}
	if plain.MyStringArray != nil && len(plain.MyStringArray) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "myStringArray", 1)
	}
	if len(plain.MyStringArray) > 3 {
		return fmt.Errorf("field %s length: must be <= %d", "myStringArray", 3)
	}
	*j = MinMaxItems(plain)
	return nil
}
