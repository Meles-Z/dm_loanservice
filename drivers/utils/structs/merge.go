package structs

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/aarondl/null/v8"
)

func DifSqlSet(src, req interface{}) {
	s := reflect.ValueOf(src)
	r := reflect.ValueOf(req)
	if s.Kind() != reflect.Ptr || r.Kind() != reflect.Ptr {
		return
	}
	for i := 0; i < s.Elem().NumField(); i++ {
		v := s.Elem().Field(i)
		fieldName := s.Elem().Type().Field(i).Name
		vr := r.Elem().FieldByName(fieldName)
		skip := s.Elem().Type().Field(i).Tag.Get("json")
		if skip == "-" {
			continue
		}
		if !vr.IsValid() {
			continue
		}
		if CheckNil(vr) {
			if v.Kind() > reflect.Float64 &&
				v.Kind() != reflect.String &&
				v.Kind() != reflect.Struct &&
				v.Kind() != reflect.Ptr &&
				v.Kind() != reflect.Slice {
				continue
			}
			r.Elem().FieldByName(fieldName).Set(v)
		}
	}
}

func ConvertValue(val reflect.Value) interface{} {
	v := val.Interface()
	switch v.(type) {
	case *int:
		return *v.(*int)
	case *int32:
		return *v.(*int32)
	case *int64:
		return *v.(*int64)
	case *float32:
		return *v.(*float32)
	case *float64:
		return *v.(*float64)
	case *string:
		return *v.(*string)
	case int:
		return v.(int)
	case int32:
		return v.(int32)
	case int64:
		return v.(int64)
	case float32:
		return v.(float32)
	case float64:
		return v.(float64)
	default:
		return v.(string)
	}

}

func CheckNil(val reflect.Value) bool {
	v := val.Interface()
	switch v.(type) {
	case *int:
		if v.(*int) != nil {
			return false
		} else {
			return true
		}
	case *int32:
		if v.(*int32) != nil {
			return false
		} else {
			return true
		}
	case *int64:
		if v.(*int64) != nil {
			return false
		} else {
			return true
		}
	case *float32:
		if v.(*float32) != nil {
			return false
		} else {
			return true
		}
	case *float64:
		if v.(*float64) != nil {
			return false
		} else {
			return true
		}
	case *string:
		if v.(*string) != nil {
			return false
		} else {
			return true
		}
	case []*string:
		if len(v.([]*string)) > 0 {
			return false
		} else {
			return true
		}
	case []*int:
		if len(v.([]*int)) > 0 {
			return false
		} else {
			return true
		}
	case []*int32:
		if len(v.([]*int32)) > 0 {
			return false
		} else {
			return true
		}
	case []*int64:
		if len(v.([]*int64)) > 0 {
			return false
		} else {
			return true
		}
	case []*float32:
		if len(v.([]*float32)) > 0 {
			return false
		} else {
			return true
		}
	case **float64:
		if len(v.([]*float64)) > 0 {
			return false
		} else {
			return true
		}
	case *bool:
		return false
	case *time.Time:
		if v.(*time.Time) != nil {
			return false
		}
		return true
	case int:
		if v.(int) > 0 {
			return false
		} else {
			return true
		}
	case int32:
		if v.(int32) > 0 {
			return false
		} else {
			return true
		}
	case int64:
		if v.(int64) > 0 {
			return false
		} else {
			return true
		}
	case float32:
		if v.(float32) > 0 {
			return false
		} else {
			return true
		}
	case float64:
		if v.(float64) != 0 {
			return false
		} else {
			return true
		}
	case []string:
		if len(v.([]string)) > 0 {
			return false
		} else {
			return true
		}
	case []int:
		if len(v.([]int)) > 0 {
			return false
		} else {
			return true
		}
	case []int32:
		if len(v.([]int32)) > 0 {
			return false
		} else {
			return true
		}
	case []int64:
		if len(v.([]int64)) > 0 {
			return false
		} else {
			return true
		}
	case []float32:
		if len(v.([]float32)) > 0 {
			return false
		} else {
			return true
		}
	case []float64:
		if len(v.([]float64)) > 0 {
			return false
		} else {
			return true
		}
	case bool:
		return false
	case time.Time:
		return v.(time.Time).IsZero()
	case null.String:
		data, err := json.Marshal(v)
		temp := &null.String{}
		err = json.Unmarshal(data, temp)
		if err != nil {
			return false
		}
		return !temp.Valid
	case null.Time:
		data, err := json.Marshal(v)
		temp := &null.Time{}
		err = json.Unmarshal(data, temp)
		if err != nil {
			return false
		}
		return !temp.Valid
	case null.Int64:
		data, err := json.Marshal(v)
		temp := &null.Int64{}
		err = json.Unmarshal(data, temp)
		if err != nil {
			return false
		}
		return !temp.Valid
	default:
		if len(v.(string)) > 0 {
			return false
		} else {
			return true
		}
	}
}
