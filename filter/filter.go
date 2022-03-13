package filter

import (
	"github.com/goark/cov19jpn/values/date"
	"github.com/goark/cov19jpn/values/prefcodejpn"
)

//Filter clsss is filtering conditions
type Filter struct {
	period    date.Period
	prefcodes map[prefcodejpn.Code]struct{}
}

//NewFilter function creates Filter instance.
func New(pcs ...prefcodejpn.Code) *Filter {
	return (&Filter{period: date.NewPeriod(date.Zero, date.Zero), prefcodes: map[prefcodejpn.Code]struct{}{}}).AddPrefCode(pcs...)
}

//SetPeriod method sets date.Period condition in Filter instance.
func (f *Filter) SetPeriod(p date.Period) *Filter {
	if f == nil {
		f = New()
	}
	f.period = p
	return f
}

//AddPrefCode method adds prefcodejpn.Code conditions in Filter instance.
func (f *Filter) AddPrefCode(pcs ...prefcodejpn.Code) *Filter {
	if f == nil {
		f = New()
	}
	for _, pc := range pcs {
		f.prefcodes[pc] = struct{}{}
	}
	return f
}

//Match method returns true if entity.Entity matches in Filter conditions.
func (f *Filter) Match(d date.Date, c prefcodejpn.Code) bool {
	if f == nil {
		return true
	}
	if !f.period.Contains(d) {
		return false
	}
	if len(f.prefcodes) == 0 {
		return true
	}
	for k := range f.prefcodes {
		if k == c {
			return true
		}
	}
	return false
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
