package models

import "fmt"

type ServiceTypeEnum string

const (
	FREE    ServiceTypeEnum = "FREE"    // The user uses the service without restrictions
	MONTHLY ServiceTypeEnum = "MONTHLY" // The user uses the service with a monthly usage limit
	TOTALLY ServiceTypeEnum = "TOTALLY" // The user uses the service with a limited Rx-TX
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
