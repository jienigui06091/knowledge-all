package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JSONB 是 PostgreSQL JSONB 字段的 GORM 映射类型。
type JSONB json.RawMessage

func (j *JSONB) Scan(value any) error {
	switch data := value.(type) {
	case nil:
		*j = nil
		return nil
	case []byte:
		*j = append((*j)[:0], data...)
	case string:
		*j = append((*j)[:0], data...)
	default:
		return fmt.Errorf("cannot scan %T into JSONB", value)
	}

	if !json.Valid(*j) {
		return fmt.Errorf("invalid JSONB value")
	}
	return nil
}

func (j JSONB) Value() (driver.Value, error) {
	if len(j) == 0 {
		return []byte("{}"), nil
	}
	if !json.Valid(j) {
		return nil, fmt.Errorf("invalid JSONB value")
	}
	return []byte(j), nil
}
