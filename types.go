package main

import (
	"database/sql/driver"
	"strings"
)

type Strlist []string

func (m *Strlist) Scan(val interface{}) error {
	s := val.(string)
	if len(s) > 0 {
		ss := strings.Split(s, "|")
		*m = ss
	} else {
		*m = Strlist{}
	}
	return nil
}

func (m Strlist) Value() (driver.Value, error) {
	str := strings.Join(m, "|")
	return str, nil
}
