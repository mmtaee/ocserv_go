package repository

import (
	"context"
	"fmt"
	"os"
	"reflect"
)

// var ocservGroupDir = "/etc/ocserv/groups"
var ocservGroupDir = "/tmp"

type OcsGroupRepository struct{}

type OcservGroupRepositoryInterface interface {
	GroupList() []string
	GroupCreateOrUpdate(context.Context) error
	GroupDelete(string) error
}

func NewOcservGroupRepository() *OcsGroupRepository {
	return &OcsGroupRepository{}
}

func (g *OcsGroupRepository) GroupList() []string {
	var groups []string
	groupsFiles, err := os.ReadDir(ocservGroupDir)
	if err != nil {
		return groups
	}
	for _, file := range groupsFiles {
		groups = append(groups, file.Name())
	}
	return groups
}

func (g *OcsGroupRepository) GroupCreateOrUpdate(ctx context.Context) (err error) {
	ch := make(chan error, 1)
	name := ctx.Value("name").(string)
	value := ctx.Value("config").(reflect.Value)
	filePath := fmt.Sprintf("%s/%s", ocservGroupDir, name)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

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

	select {
	case err = <-ch:
		if err != nil {
			cancel()
			_ = os.Remove(filePath)
			return err
		}
		return nil
	case <-ctx.Done():
		_ = os.Remove(filePath)
		return ctx.Err()
	}
}

func (g *OcsGroupRepository) GroupDelete(groupName string) (err error) {
	return os.Remove(fmt.Sprintf("%s/%s", ocservGroupDir, groupName))
}
