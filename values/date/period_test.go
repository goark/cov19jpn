package date

import (
	"fmt"
	"testing"
	"time"
)

func TestPeriod(t *testing.T) {
	testCases := []struct {
		start, end, dt Date
		res            bool
	}{
		{start: New(2020, time.March, 1), end: New(2020, time.May, 31), dt: New(2020, time.March, 0), res: false},
		{start: New(2020, time.March, 1), end: New(2020, time.May, 31), dt: New(2020, time.March, 1), res: true},
		{start: New(2020, time.May, 31), end: New(2020, time.March, 1), dt: New(2020, time.May, 31), res: true},
		{start: New(2020, time.May, 31), end: New(2020, time.March, 1), dt: New(2020, time.May, 32), res: false},
		{start: New(2020, time.March, 1), end: Zero, dt: New(2020, time.March, 0), res: false},
		{start: New(2020, time.March, 1), end: Zero, dt: New(2020, time.March, 1), res: true},
		{start: Zero, end: New(2020, time.May, 31), dt: New(2020, time.May, 31), res: true},
		{start: Zero, end: New(2020, time.May, 31), dt: New(2020, time.May, 32), res: false},
		{start: Zero, end: Zero, dt: New(2020, time.January, 1), res: true},
		{start: Zero, end: Zero, dt: Zero, res: false},
	}

	for _, tc := range testCases {
		p := NewPeriod(tc.start, tc.end)
		if res := p.Contains(tc.dt); res != tc.res {
			t.Errorf("{%v}.Contains(%v) = %v, want %v.", p, tc.dt, res, tc.res)
		} else {
			fmt.Println(p, "in", tc.dt, "=>", p.Contains(tc.dt))
		}
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
