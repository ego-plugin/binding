package fields

import (
	"bytes"
	"reflect"
	"strings"
	"unicode"
)

// Columns 传入一个结构体反回数据库 column
func Columns(a interface{}) []string {
	typeof := reflect.TypeOf(a)
	if typeof.Kind() == reflect.Ptr {
		typeof = typeof.Elem()
	}
	columns := make([]string, 0)
	for i := 0; i < typeof.NumField(); i++ {
		column := typeof.Field(i).Tag.Get("db")
		columnList := strings.SplitN(column, ",", 1)
		column = columnList[0]
		if column == "-" {
			continue
		}

		// 存在跳过
		if column != "" {
			columns = append(columns, column)
			continue
		}

		column = camelCaseToSnakeCase(typeof.Field(i).Name)
		columns = append(columns, column)
	}
	return columns
}

func ContainsColumns(columns []string, chars []string) bool {
	for _, column := range columns {
		ok := false
		for _, v := range chars {
			if !strings.Contains(column, v) {
				continue
			}
			ok = true
		}
		if !ok {
			return false
		}
	}
	return true
}

func camelCaseToSnakeCase(name string) string {
	buf := new(bytes.Buffer)

	runes := []rune(name)
	for i := 0; i < len(runes); i++ {
		buf.WriteRune(unicode.ToLower(runes[i]))
		if i != len(runes)-1 && unicode.IsUpper(runes[i+1]) &&
			(unicode.IsLower(runes[i]) || unicode.IsDigit(runes[i]) ||
				(i != len(runes)-2 && unicode.IsLower(runes[i+2]))) {
			buf.WriteRune('_')
		}
	}

	return buf.String()
}
