package cast

import (
	"database/sql"
	"strconv"
)

func ToString(value interface{}) string {
	var s string
	if value == nil {
		return s
	}
	switch value.(type) {
	case sql.NullString:
		v := value.(sql.NullString)
		if v.Valid {
			s = v.String
		} else {
			s = ""
		}
	case bool:
		if value.(bool) {
			s = "yes"
		} else {
			s = "no"
		}
	case int:
		s = strconv.Itoa(value.(int))
	case int64:
		s = strconv.FormatInt(value.(int64), 10)
	case float64:
		s = strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case string:
		s = value.(string)
	case sql.NullBool:
		v := value.(sql.NullBool)
		if v.Valid && v.Bool {
			s = "yes"
		} else {
			s = "no"
		}
	case sql.NullInt64:
		v := value.(sql.NullInt64)
		if v.Valid {
			s = strconv.FormatInt(v.Int64, 10)
		} else {
			s = ""
		}
	case sql.NullFloat64:
		v := value.(sql.NullFloat64)
		if v.Valid {
			s = strconv.FormatFloat(v.Float64, 'f', -1, 64)
		} else {
			s = ""
		}
	default:
		s = ""
	}

	return s
}

func ToNullString(value interface{}) sql.NullString {
	var s sql.NullString
	if value == nil {
		return s
	}
	switch value.(type) {
	case sql.NullString:
		s = value.(sql.NullString)
	case sql.NullInt64:
		v := value.(sql.NullInt64)
		if v.Valid {
			s.String = strconv.FormatInt(v.Int64, 10)
			s.Valid = true
		} else {
			s.String = ""
			s.Valid = false
		}
	case sql.NullFloat64:
		v := value.(sql.NullFloat64)
		if v.Valid {
			s.String = strconv.FormatFloat(v.Float64, 'f', -1, 64)
			s.Valid = true
		} else {
			s.String = ""
			s.Valid = false
		}
	case sql.NullBool:
		v := value.(sql.NullBool)
		if v.Valid {
			if v.Bool {
				s.String = "yes"
			} else {
				s.String = "no"
			}
			s.Valid = true
		} else {
			s.String = ""
			s.Valid = false
		}
	case bool:
		if value.(bool) {
			s.String = "yes"
		} else {
			s.String = "no"
		}
	case int:
		s.String = strconv.Itoa(value.(int))
		s.Valid = true
	case int64:
		s.String = strconv.FormatInt(value.(int64), 10)
		s.Valid = true
	case float64:
		s.String = strconv.FormatFloat(value.(float64), 'f', -1, 64)
		s.Valid = true
	case string:
		s.String = value.(string)
		s.Valid = true
	default:
		s.Valid = false
	}

	return s
}

func ToBool(value interface{}) bool {
	var b bool
	if value == nil {
		return b
	}

	switch value.(type) {
	case string:
		v := value.(string)
		b = v == "1" || v == "true" || v == "yes"
	case sql.NullString:
		v := value.(sql.NullString)
		b = v.Valid && (v.String == "1" || v.String == "true" || v.String == "yes")
	case sql.NullBool:
		v := value.(sql.NullBool)
		b = v.Valid && v.Bool
	case sql.NullInt64:
		v := value.(sql.NullInt64)
		b = v.Valid && v.Int64 != 0
	case sql.NullFloat64:
		v := value.(sql.NullFloat64)
		b = v.Valid && v.Float64 != 0
	case bool:
		b = value.(bool)
	case int:
		b = value.(int) != 0
	case int64:
		b = value.(int64) != 0
	case float64:
		b = value.(float64) != 0
	default:
		b = false
	}

	return b
}

func ToInt(number interface{}) int {
	var i int
	if number == nil {
		return i
	}
	switch number.(type) {
	case sql.NullString:
		v := number.(sql.NullString)
		if v.Valid {
			i, _ = strconv.Atoi(v.String)
		} else {
			i = 0
		}
	case sql.NullInt64:
		v := number.(sql.NullInt64)
		if v.Valid {
			i = int(v.Int64)
		} else {
			i = 0
		}
	case sql.NullFloat64:
		v := number.(sql.NullFloat64)
		if v.Valid {
			i = int(v.Float64)
		} else {
			i = 0
		}
	case sql.NullBool:
		v := number.(sql.NullBool)
		if v.Valid && v.Bool {
			i = 1
		} else {
			i = 0
		}
	case bool:
		if number.(bool) {
			i = 1
		} else {
			i = 0
		}
	case int:
		i = number.(int)
	case int64:
		i = int(number.(int64))
	case float64:
		i = int(number.(float64))
	case string:
		i, _ = strconv.Atoi(number.(string))
	default:
		i = 0
	}

	return i
}

func ToInt64(number interface{}) int64 {
	var i int64
	if number == nil {
		return i
	}
	switch number.(type) {
	case sql.NullString:
		v := number.(sql.NullString)
		if v.Valid {
			i, _ = strconv.ParseInt(v.String, 10, 64)
		} else {
			i = 0
		}
	case sql.NullInt64:
		v := number.(sql.NullInt64)
		if v.Valid {
			i = v.Int64
		} else {
			i = 0
		}
	case sql.NullFloat64:
		v := number.(sql.NullFloat64)
		if v.Valid {
			i = int64(v.Float64)
		} else {
			i = 0
		}
	case sql.NullBool:
		v := number.(sql.NullBool)
		if v.Valid && v.Bool {
			i = 1
		} else {
			i = 0
		}
	case bool:
		if number.(bool) {
			i = 1
		} else {
			i = 0
		}
	case int:
		i = int64(number.(int))
	case int64:
		i = number.(int64)
	case float64:
		i = int64(number.(float64))
	case string:
		i, _ = strconv.ParseInt(number.(string), 10, 64)
	default:
		i = 0
	}

	return i
}

func ToFloat64(number interface{}) float64 {
	var f float64
	if number == nil {
		return f
	}
	switch number.(type) {
	case sql.NullString:
		v := number.(sql.NullString)
		if v.Valid {
			f, _ = strconv.ParseFloat(v.String, 64)
		} else {
			f = 0
		}
	case sql.NullInt64:
		v := number.(sql.NullInt64)
		if v.Valid {
			f = float64(v.Int64)
		} else {
			f = 0
		}
	case sql.NullFloat64:
		v := number.(sql.NullFloat64)
		if v.Valid {
			f = v.Float64
		} else {
			f = 0
		}
	case sql.NullBool:
		v := number.(sql.NullBool)
		if v.Valid && v.Bool {
			f = 1
		} else {
			f = 0
		}
	case bool:
		if number.(bool) {
			f = 1
		} else {
			f = 0
		}
	case int:
		f = float64(number.(int))
	case int64:
		f = float64(number.(int64))
	case float64:
		f = number.(float64)
	case string:
		f, _ = strconv.ParseFloat(number.(string), 64)
	default:
		f = 0
	}

	return f
}

func ToNullInt64(number interface{}, zeroToNull ...bool) sql.NullInt64 {
	var i sql.NullInt64
	if number == nil {
		return i
	}

	var err error
	switch number.(type) {
	case sql.NullString:
		v := number.(sql.NullString)
		if v.Valid {
			i.Int64, err = strconv.ParseInt(v.String, 10, 64)
			if err != nil {
				i.Int64 = 0
				i.Valid = false
			} else {
				i.Valid = true
			}
		} else {
			i.Int64 = 0
			i.Valid = false
		}
	case sql.NullInt64:
		v := number.(sql.NullInt64)
		i.Int64 = v.Int64
		i.Valid = v.Valid
	case sql.NullFloat64:
		v := number.(sql.NullFloat64)
		if v.Valid {
			i.Int64 = int64(v.Float64)
			i.Valid = true
		} else {
			i.Int64 = 0
			i.Valid = false
		}
	case sql.NullBool:
		v := number.(sql.NullBool)
		if v.Valid && v.Bool {
			i.Int64 = 1
			i.Valid = true
		} else {
			i.Int64 = 0
			i.Valid = true
		}
	case bool:
		if number.(bool) {
			i.Int64 = 1
			i.Valid = true
		} else {
			i.Int64 = 0
			i.Valid = true
		}
	case int:
		i.Int64 = int64(number.(int))
		i.Valid = true
	case int64:
		i.Int64 = number.(int64)
		i.Valid = true
	case float64:
		i.Int64 = int64(number.(float64))
		i.Valid = true
	case string:
		i.Int64, _ = strconv.ParseInt(number.(string), 10, 64)
		i.Valid = true
	default:
		i.Int64 = 0
		i.Valid = false
	}

	if len(zeroToNull) > 0 && zeroToNull[0] && i.Int64 == 0 {
		i.Valid = false
	}

	return i
}

func ToNullFloat64(number interface{}) sql.NullFloat64 {
	var f sql.NullFloat64
	if number == nil {
		return f
	}

	var err error
	switch number.(type) {
	case sql.NullString:
		v := number.(sql.NullString)
		if v.Valid {
			f.Float64, err = strconv.ParseFloat(v.String, 64)
			if err != nil {
				f.Float64 = 0
				f.Valid = false
			} else {
				f.Valid = true
			}
		} else {
			f.Float64 = 0
			f.Valid = false
		}
	case sql.NullInt64:
		v := number.(sql.NullInt64)
		if v.Valid {
			f.Float64 = float64(v.Int64)
			f.Valid = true
		} else {
			f.Float64 = 0
			f.Valid = false
		}
	case sql.NullFloat64:
		v := number.(sql.NullFloat64)
		f.Float64 = v.Float64
		f.Valid = v.Valid
	case sql.NullBool:
		v := number.(sql.NullBool)
		if v.Valid && v.Bool {
			f.Float64 = 1
			f.Valid = true
		} else {
			f.Float64 = 0
			f.Valid = true
		}
	case bool:
		if number.(bool) {
			f.Float64 = 1
			f.Valid = true
		} else {
			f.Float64 = 0
			f.Valid = true
		}
	case int:
		f.Float64 = float64(number.(int))
		f.Valid = true
	case int64:
		f.Float64 = float64(number.(int64))
		f.Valid = true
	case float64:
		f.Float64 = number.(float64)
		f.Valid = true
	case string:
		f.Float64, _ = strconv.ParseFloat(number.(string), 64)
		f.Valid = true
	default:
		f.Float64 = 0
		f.Valid = false
	}

	return f

}

func ToIntArray(value interface{}) []int {
	var arr []int
	if value == nil {
		return arr
	}

	switch value.(type) {
	case []int:
		arr = value.([]int)
	case []int64:
		for _, v := range value.([]int64) {
			arr = append(arr, int(v))
		}
	case []float64:
		for _, v := range value.([]float64) {
			arr = append(arr, int(v))
		}
	case []string:
		for _, v := range value.([]string) {
			i, _ := strconv.Atoi(v)
			arr = append(arr, i)
		}
	case []sql.NullInt64:
		for _, v := range value.([]sql.NullInt64) {
			if v.Valid {
				arr = append(arr, int(v.Int64))
			}
		}
	case []sql.NullFloat64:
		for _, v := range value.([]sql.NullFloat64) {
			if v.Valid {
				arr = append(arr, int(v.Float64))
			}
		}
	case []sql.NullString:
		for _, v := range value.([]sql.NullString) {
			if v.Valid {
				i, _ := strconv.Atoi(v.String)
				arr = append(arr, i)
			}
		}
	}

	return arr
}

func ToStringArray(value interface{}) []string {
	var arr []string
	if value == nil {
		return arr
	}

	switch value.(type) {
	case []string:
		arr = value.([]string)
	case []int:
		for _, v := range value.([]int) {
			arr = append(arr, strconv.Itoa(v))
		}
	case []int64:
		for _, v := range value.([]int64) {
			arr = append(arr, strconv.FormatInt(v, 10))
		}
	case []float64:
		for _, v := range value.([]float64) {
			arr = append(arr, strconv.FormatFloat(v, 'f', -1, 64))
		}
	case []sql.NullInt64:
		for _, v := range value.([]sql.NullInt64) {
			if v.Valid {
				arr = append(arr, strconv.FormatInt(v.Int64, 10))
			}
		}
	case []sql.NullFloat64:
		for _, v := range value.([]sql.NullFloat64) {
			if v.Valid {
				arr = append(arr, strconv.FormatFloat(v.Float64, 'f', -1, 64))
			}
		}
	case []sql.NullString:
		for _, v := range value.([]sql.NullString) {
			if v.Valid {
				arr = append(arr, v.String)
			}
		}
	}

	return arr
}

func StringToIntArray(value string) []int {
	var arr []int
	if value == "" {
		return arr
	}

	for _, v := range value {
		i, _ := strconv.Atoi(string(v))
		arr = append(arr, i)
	}

	return arr
}
