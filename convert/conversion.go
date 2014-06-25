package convert

import (
	"fmt"
	"strconv"
	"time"
)

var TimeFormats = []string{
	"2006-01-02",
}

func ToString(object interface{}) string {
	if object == nil {
		return ""
	}

	switch val := object.(type) {
	case []byte:
		return ToString(string(val))
	default:
		return fmt.Sprintf("%v", val)
	}
}

func ToBool(object interface{}) bool {
	switch val := object.(type) {
	case bool:
		return val

	case string:
		boolVal, _ := strconv.ParseBool(val)
		return boolVal

	case []byte:
		return ToBool(string(val))

	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64:
		return val != 0

	case float32, float64:
		return val != 0.0

	default:
		return false
	}
}

func ToFloat(object interface{}) float64 {
	switch val := object.(type) {

	case string:
		floatVal, _ := strconv.ParseFloat(val, 64)
		return floatVal

	case []byte:
		return ToFloat(string(val))

	case float64:
		return val

	case float32:
		return float64(val)

	case uint:
		return float64(val)

	case uint8:
		return float64(val)

	case uint16:
		return float64(val)

	case uint32:
		return float64(val)

	case uint64:
		return float64(val)

	case int:
		return float64(val)

	case int8:
		return float64(val)

	case int16:
		return float64(val)

	case int32:
		return float64(val)

	case int64:
		return float64(val)

	case bool:
		if val {
			return 1.0
		} else {
			return 0.0
		}

	default:
		return 0.0
	}
}

func ToInt(object interface{}) int64 {
	switch val := object.(type) {

	case string:
		floatVal, _ := strconv.ParseFloat(val, 64)
		return int64(floatVal)

	case []byte:
		return ToInt(string(val))

	case float64:
		return int64(val)

	case float32:
		return int64(val)

	case uint:
		return int64(val)

	case uint8:
		return int64(val)

	case uint16:
		return int64(val)

	case uint32:
		return int64(val)

	case uint64:
		return int64(val)

	case int:
		return int64(val)

	case int8:
		return int64(val)

	case int16:
		return int64(val)

	case int32:
		return int64(val)

	case int64:
		return val

	case bool:
		if val {
			return 1
		} else {
			return 0
		}

	default:
		return 0
	}
}

func ToTime(object interface{}) time.Time {
	switch val := object.(type) {

	case time.Time:
		return val

	case string:
		for _, format := range TimeFormats {
			stringVal, err := time.Parse(format, val)
			if err == nil {
				return stringVal
			}
		}
		// Nothing worked, return Zero time
		return time.Time{}

	case []byte:
		return ToTime(string(val))

	default:
		return time.Time{}
	}
}
