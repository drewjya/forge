package template

func ULIDRoot() string {
	return `package types

import (
		"encoding/json"

		"github.com/oklog/ulid/v2"
)

type ULID string

func ParseULID(s string) (ULID, error) {
		_, err := ulid.Parse(s)
		if err != nil {
			return "", err
		}
		return ULID(s), nil
}

func (u ULID) String() string {
		return string(u)
}

func (u *ULID) UnmarshalJSON(data []byte) error {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}

		id, err := ParseULID(s)
		if err != nil {
			return err
		}

		*u = id
		return nil
}

// Optional but nice
func (u ULID) MarshalJSON() ([]byte, error) {
		return json.Marshal(string(u))
}

func (u *ULID) UnmarshalText(text []byte) error {
		id, err := ParseULID(string(text))
		if err != nil {
			return err
		}
		*u = id
		return nil
}
`
}
