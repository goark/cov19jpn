package date

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type ForTestStruct struct {
	DateTaken Date `json:"date_taken,omitempty"`
}

func TestUnmarshal(t *testing.T) {
	testCases := []struct {
		s   string
		str string
		jsn string
	}{
		{s: `{"date_taken": "2005-03-26 00:00:00 UTC"}`, str: "2005-03-26", jsn: `{"date_taken":"2005-03-26"}`},
		{s: `{"date_taken": "2005-03-26T00:00:00+09:00"}`, str: "2005-03-26", jsn: `{"date_taken":"2005-03-26"}`},
		{s: `{"date_taken": "2005-03-26"}`, str: "2005-03-26", jsn: `{"date_taken":"2005-03-26"}`},
		{s: `{"date_taken": ""}`, str: "", jsn: `{"date_taken":""}`},
		{s: `{}`, str: "", jsn: `{"date_taken":""}`},
	}

	for _, tc := range testCases {
		tst := &ForTestStruct{}
		if err := json.Unmarshal([]byte(tc.s), tst); err != nil {
			t.Errorf("json.Unmarshal() is \"%v\", want nil.", err)
			continue
		}
		str := tst.DateTaken.String()
		if str != tc.str {
			t.Errorf("Date = \"%v\", want \"%v\".", str, tc.str)
		}
		b, err := json.Marshal(tst)
		if err != nil {
			t.Errorf("json.Marshal() is \"%v\", want nil.", err)
			continue
		}
		str = string(b)
		if str != tc.jsn {
			t.Errorf("ForTestStruct = \"%v\", want \"%v\".", str, tc.jsn)
		}
	}
}

func TestUnmarshalErr(t *testing.T) {
	data := `{"date_taken": "2005/03/26"}`
	tst := &ForTestStruct{}
	if err := json.Unmarshal([]byte(data), tst); err != nil {
		t.Error("Unmarshal() error is not nil, want nil.")
	} else {
		fmt.Printf("Info: %+v\n", err)
	}
}

func TestEqual(t *testing.T) {
	dt1 := New(2020, time.March, 31)
	dt2 := FromString("2020-03-31")
	dt3 := FromString("2020-04-01")
	if !dt1.Equal(dt2) {
		t.Error("Date.Equal() is true, want false.")
	}
	if dt1.Equal(dt3) {
		t.Error("Date.Equal() is true, want false.")
	}
	if !dt1.Before(dt3) {
		t.Error("Date.Before() is false, want true.")
	}
	if dt1.After(dt3) {
		t.Error("Date.After() is true, want false.")
	}
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
