package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

//type DynamicFloat64 struct {
//	Value *float64
//}
//
//func (d *DynamicFloat64) UnmarshalJSON(data []byte) error {
//	if string(data) == "null" {
//		d.Value = nil
//		return nil
//	}
//
//	var f float64
//	if err := json.Unmarshal(data, &f); err == nil {
//		d.Value = &f
//		return nil
//	}
//
//	var s string
//	if err := json.Unmarshal(data, &s); err == nil {
//		if s == "" {
//			d.Value = nil
//			return nil
//		}
//		if f, err := strconv.ParseFloat(s, 64); err == nil {
//			d.Value = &f
//			return nil
//		}
//	}
//
//	return fmt.Errorf("invalid value for DynamicFloat64: %s", string(data))
//}

type DynamicFloat64 float64

func (d *DynamicFloat64) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		*d = DynamicFloat64(f)
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		if s == "" {
			*d = DynamicFloat64(0)
			return nil
		}
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			*d = DynamicFloat64(f)
			return nil
		}
	}

	return fmt.Errorf("invalid value for DynamicFloat64: %s", string(data))
}
