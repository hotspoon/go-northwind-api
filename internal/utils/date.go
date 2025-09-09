// utils/date.go
package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// DateOnly represents a date without time
type DateOnly struct {
	time.Time
}

const layout = "2006-01-02"

// Scan implements sql.Scanner so DateOnly can be read from the database
func (d *DateOnly) Scan(value any) error {
	var t time.Time
	switch v := value.(type) {
	case time.Time:
		t = v
	case []byte:
		parsed, err := time.Parse(layout, string(v))
		if err != nil {
			return fmt.Errorf("failed to parse DateOnly from []byte: %w", err)
		}
		t = parsed
	default:
		return fmt.Errorf("unsupported type for DateOnly: %T", value)
	}
	d.Time = t
	return nil
}

// Value implements driver.Valuer to write DateOnly to the database
func (d DateOnly) Value() (driver.Value, error) {
	return d.Time, nil
}

// MarshalJSON implements json.Marshaler
func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format(layout) + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (d *DateOnly) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" {
		return nil
	}
	t, err := time.Parse(`"`+layout+`"`, s)
	if err != nil {
		return fmt.Errorf("failed to parse DateOnly from JSON: %w", err)
	}
	d.Time = t
	return nil
}
