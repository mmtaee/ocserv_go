package repository

import (
	"fmt"
	"os"
	"reflect"
)

func updateConfigFile(filePath string, value reflect.Value) chan error {
	ch := make(chan error, 1)
	go func() {
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			ch <- err
			return
		}

		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				ch <- err
				return
			}
		}(file)

		t := value.Type()
		validConfigString := ""
		for i := 0; i < t.NumField(); i++ {
			fieldValue := value.Field(i)
			fieldType := t.Field(i)
			jsonKey := fieldType.Tag.Get("json")
			switch fieldValue.Kind() {
			case reflect.Int64:
				if fieldValue.Int() != 0 {
					validConfigString += fmt.Sprintf("%s=%d\n", jsonKey, fieldValue.Int())
				}
			case reflect.String:
				if fieldValue.String() != "" {
					validConfigString += fmt.Sprintf("%s=%s\n", jsonKey, fieldValue.String())
				}
			case reflect.Bool:
				if fieldValue.Bool() {
					validConfigString += fmt.Sprintf("%s=%t\n", jsonKey, fieldValue.Bool())
				}
			case reflect.Slice:
				if fieldValue.Len() > 0 {
					for i := 0; i < fieldValue.Len(); i++ {
						validConfigString += fmt.Sprintf("%s=%s\n", jsonKey, fieldValue.Index(i))
					}
				}
			default:
				continue
			}
		}
		_, err = file.WriteString(validConfigString)
		if err != nil {
			ch <- err
			return
		}
		ch <- nil
	}()
	return ch
}
