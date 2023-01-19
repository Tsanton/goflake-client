package utilities

import (
	"fmt"
	"strings"
	"time"
)

// CustomTime provides an example of how to declare a new time Type with a custom formatter.
// Note that time.Time methods are not available, if needed you can add and cast like the String method does
// Otherwise, only use in the json struct at marshal/unmarshal time.
type SnowTime time.Time

const ctLayout = "2006-01-02 15:04:05.999999-07:00" //20

// UnmarshalJSON Parses the json string in the custom format
func (ct *SnowTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(ctLayout, s)
	*ct = SnowTime(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct SnowTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

// String returns the time in the custom format
func (ct *SnowTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(ctLayout))
}
