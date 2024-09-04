package models

import "fmt"

type ServiceTypeEnum string

const (
	FREE    ServiceTypeEnum = "FREE"
	MONTHLY ServiceTypeEnum = "MONTHLY"
	TOTALLY ServiceTypeEnum = "TOTALLY"
)

func (s *ServiceTypeEnum) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid service type")
	}
	*s = ServiceTypeEnum(str)
	return nil
}

func (s ServiceTypeEnum) Value() (interface{}, error) {
	return string(s), nil
}
