package dto

import (
	"fmt"
	"strings"
	"time"
)

type MonthYear time.Time

func (ct *MonthYear) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" || s == "null" {
		return nil
	}

	t, err := time.Parse("01-2006", s)
	if err == nil {
		*ct = MonthYear(t)
		return nil
	}

	return fmt.Errorf("invalid time format: %s", s)
}

func (ct MonthYear) MarshalJSON() ([]byte, error) {
	t := time.Time(ct)
	return []byte(`"` + t.Format("01-2006") + `"`), nil
}
