package types

import (
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5/pgtype"
)

type SocialProvider string

const (
	KAKAO  SocialProvider = "KAKAO"
	GOOGLE SocialProvider = "GOOGLE"
)

type Timestamp struct {
	pgtype.Timestamp
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	if y := t.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	if !t.Valid {
		return json.Marshal(nil)
	}
	str := t.Time.Format("2006-01-02 15:04:05")
	return json.Marshal(str)
}
