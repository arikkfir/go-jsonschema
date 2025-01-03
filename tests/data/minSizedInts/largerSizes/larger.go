// Code generated by github.com/arikkfir/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"

type Larger struct {
	// I16L corresponds to the JSON schema field "i16l".
	I16L *int16 `json:"i16l,omitempty" yaml:"i16l,omitempty" mapstructure:"i16l,omitempty"`

	// I16U corresponds to the JSON schema field "i16u".
	I16U *int16 `json:"i16u,omitempty" yaml:"i16u,omitempty" mapstructure:"i16u,omitempty"`

	// I32L corresponds to the JSON schema field "i32l".
	I32L *int32 `json:"i32l,omitempty" yaml:"i32l,omitempty" mapstructure:"i32l,omitempty"`

	// I32U corresponds to the JSON schema field "i32u".
	I32U *int32 `json:"i32u,omitempty" yaml:"i32u,omitempty" mapstructure:"i32u,omitempty"`

	// I64L corresponds to the JSON schema field "i64l".
	I64L *int64 `json:"i64l,omitempty" yaml:"i64l,omitempty" mapstructure:"i64l,omitempty"`

	// I64U corresponds to the JSON schema field "i64u".
	I64U *int64 `json:"i64u,omitempty" yaml:"i64u,omitempty" mapstructure:"i64u,omitempty"`

	// U16 corresponds to the JSON schema field "u16".
	U16 uint16 `json:"u16" yaml:"u16" mapstructure:"u16"`

	// U32 corresponds to the JSON schema field "u32".
	U32 uint32 `json:"u32" yaml:"u32" mapstructure:"u32"`

	// U64 corresponds to the JSON schema field "u64".
	U64 uint64 `json:"u64" yaml:"u64" mapstructure:"u64"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Larger) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["u16"]; raw != nil && !ok {
		return fmt.Errorf("field u16 in Larger: required")
	}
	if _, ok := raw["u32"]; raw != nil && !ok {
		return fmt.Errorf("field u32 in Larger: required")
	}
	if _, ok := raw["u64"]; raw != nil && !ok {
		return fmt.Errorf("field u64 in Larger: required")
	}
	type Plain Larger
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if plain.I16L != nil && 127 < *plain.I16L {
		return fmt.Errorf("field %s: must be <= %v", "i16l", 127)
	}
	if plain.I16L != nil && -129 > *plain.I16L {
		return fmt.Errorf("field %s: must be >= %v", "i16l", -129)
	}
	if plain.I16U != nil && 128 < *plain.I16U {
		return fmt.Errorf("field %s: must be <= %v", "i16u", 128)
	}
	if plain.I16U != nil && -128 > *plain.I16U {
		return fmt.Errorf("field %s: must be >= %v", "i16u", -128)
	}
	if plain.I32L != nil && 32767 < *plain.I32L {
		return fmt.Errorf("field %s: must be <= %v", "i32l", 32767)
	}
	if plain.I32L != nil && -32769 > *plain.I32L {
		return fmt.Errorf("field %s: must be >= %v", "i32l", -32769)
	}
	if plain.I32U != nil && 32768 < *plain.I32U {
		return fmt.Errorf("field %s: must be <= %v", "i32u", 32768)
	}
	if plain.I32U != nil && -32768 > *plain.I32U {
		return fmt.Errorf("field %s: must be >= %v", "i32u", -32768)
	}
	if plain.I64L != nil && 2147483647 < *plain.I64L {
		return fmt.Errorf("field %s: must be <= %v", "i64l", 2147483647)
	}
	if plain.I64L != nil && -2147483649 > *plain.I64L {
		return fmt.Errorf("field %s: must be >= %v", "i64l", -2147483649)
	}
	if plain.I64U != nil && 2147483648 < *plain.I64U {
		return fmt.Errorf("field %s: must be <= %v", "i64u", 2147483648)
	}
	if plain.I64U != nil && -2147483648 > *plain.I64U {
		return fmt.Errorf("field %s: must be >= %v", "i64u", -2147483648)
	}
	if 256 < plain.U16 {
		return fmt.Errorf("field %s: must be <= %v", "u16", 256)
	}
	if 65536 < plain.U32 {
		return fmt.Errorf("field %s: must be <= %v", "u32", 65536)
	}
	if 4294967296 < plain.U64 {
		return fmt.Errorf("field %s: must be <= %v", "u64", 4294967296)
	}
	*j = Larger(plain)
	return nil
}
