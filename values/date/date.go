package date

import (
	"strconv"
	"strings"
	"time"
)

//Time is wrapper class of time.Time.
type Date struct {
	time.Time
}

var (
	defaultDateForm = "2006-01-02"
	timeTemplate    = []string{
		defaultDateForm,
		time.RFC3339,
	}
	Zero = FromTime(time.Time{}) //Zero value of Date
)

//New returns Date instance from year/month/day.
func New(year int, month time.Month, day int) Date {
	return FromTime(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

//FromString returns Date instance from formatted date string.
func FromString(s string) Date {
	if len(s) == 0 || strings.EqualFold(s, "null") {
		return Zero
	}
	for _, tmplt := range timeTemplate {
		if tm, err := time.Parse(tmplt, s); err == nil {
			return FromTime(tm)
		}
	}
	return Zero
}

//FromTime returns Date instance from fime.Time.
func FromTime(tm time.Time) Date {
	if tm.IsZero() {
		return Date{tm}
	}
	_, offset := tm.Zone()
	return Date{time.Unix(((tm.Unix()+int64(offset))/86400)*86400, 0).In(time.UTC)}
}

//Today function retuens Date instance of today
func Today() Date {
	return FromTime(time.Now())
}

//Yesterday function retuens Date instance of yesterday
func Yesterday() Date {
	return Today().AddDay(-1)
}

//UnmarshalJSON returns result of Unmarshal for json.Unmarshal().
func (t *Date) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	*t = FromString(s)
	return nil
}

//MarshalJSON returns time string with RFC3339 format for json.Marshal().
func (t *Date) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`""`), nil
	}
	if t.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(strconv.Quote(t.String())), nil
}

func (t Date) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Format(defaultDateForm)
}

//Equal reports whether t and dt represent the same time instant.
func (t Date) Equal(dt Date) bool {
	return t.Year() == dt.Year() && t.Month() == dt.Month() && t.Day() == dt.Day()
}

//Before reports whether the Date instant t is before dt.
func (t Date) Before(dt Date) bool {
	return !t.Equal(dt) && t.Time.Before(dt.Time)
}

//After reports whether the Date instant t is after dt.
func (t Date) After(dt Date) bool {
	return !t.Equal(dt) && t.Time.After(dt.Time)
}

//AddDay method adds n days and returns new Date instance.
func (t Date) AddDay(days int) Date {
	return FromTime(t.Time.AddDate(0, 0, days))
}

/* Copyright 2021 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
