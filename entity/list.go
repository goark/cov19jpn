package entity

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/goark/cov19jpn/filter"
	"github.com/goark/cov19jpn/values/date"
)

//EntityList class is list of Entities.
type EntityList struct {
	list []*Entity
}

//NewList function returns new EntityList instance.
func NewList(list []*Entity) *EntityList {
	if list == nil {
		return &EntityList{[]*Entity{}}
	}
	return &EntityList{list}
}

//Add method adds Entities in EntityList.
func (el *EntityList) Add(e ...*Entity) {
	if el == nil {
		return
	}
	if el.list == nil {
		el.list = []*Entity{}
	}
	el.list = append(el.list, e...)
}

//Entities returns list of Entity.
func (el *EntityList) Entities() []*Entity {
	if el == nil {
		return []*Entity{}
	}
	return el.list
}

//Size method returns size of list.
func (el *EntityList) Size() int {
	if el == nil {
		return 0
	}
	return len(el.list)
}

//Entity method return Entity in EntityList.
func (el *EntityList) Entity(i int) *Entity {
	if el == nil {
		return nil
	}
	if i < 0 || i >= el.Size() {
		return nil
	}
	return el.list[i]
}

//StartDayMeasure method returns date.Date of start day (Measurements)
func (el *EntityList) StartDayMeasure() date.Date {
	if el.Size() == 0 {
		return date.Zero
	}
	dt := date.Zero
	for _, e := range el.list {
		if !e.IsForecast() {
			if dt.IsZero() {
				dt = e.TargetPredictionDate
			} else if dt.After(e.TargetPredictionDate) {
				dt = e.TargetPredictionDate
			}
		}
	}
	return dt
}

//EndDayMeasure method returns date.Date of end day (Measurements)
func (el *EntityList) EndDayMeasure() date.Date {
	if el.Size() == 0 {
		return date.Zero
	}
	dt := date.Zero
	for _, e := range el.list {
		if !e.IsForecast() {
			if dt.IsZero() {
				dt = e.TargetPredictionDate
			} else if dt.Before(e.TargetPredictionDate) {
				dt = e.TargetPredictionDate
			}
		}
	}
	return dt
}

//StartDayForecast method returns date.Date of start day (Forecasts)
func (el *EntityList) StartDayForecast() date.Date {
	if el.Size() == 0 {
		return date.Zero
	}
	dt := date.Zero
	for _, e := range el.list {
		if e.IsForecast() {
			if dt.IsZero() {
				dt = e.TargetPredictionDate
			} else if dt.After(e.TargetPredictionDate) {
				dt = e.TargetPredictionDate
			}
		}
	}
	return dt
}

//EndDayForecast method returns date.Date of end day (Forecasts)
func (el *EntityList) EndDayForecast() date.Date {
	if el.Size() == 0 {
		return date.Zero
	}
	dt := date.Zero
	for _, e := range el.list {
		if e.IsForecast() {
			if dt.IsZero() {
				dt = e.TargetPredictionDate
			} else if dt.Before(e.TargetPredictionDate) {
				dt = e.TargetPredictionDate
			}
		}
	}
	return dt
}

//Filer method returns filtering entity list
func (el *EntityList) Filer(f *filter.Filter) *EntityList {
	es := []*Entity{}
	if el.Size() < 1 {
		return NewList(es)
	}
	for _, e := range el.list {
		if f.Match(e.TargetPredictionDate, e.JapanPrefectureCode) {
			es = append(es, e)
		}
	}
	return NewList(es)
}

//Sort method sorts entity list
func (el *EntityList) Sort() {
	if el.Size() < 2 {
		return
	}
	sort.Slice(el.list, func(i, j int) bool {
		if el.Entity(i).JapanPrefectureCode.Equal(el.Entity(j).JapanPrefectureCode) {
			return el.Entity(i).TargetPredictionDate.Before(el.Entity(j).TargetPredictionDate)
		}
		return el.Entity(i).JapanPrefectureCode.Less(el.Entity(j).JapanPrefectureCode)
	})
}

//Sort method sorts entity list
func (el *EntityList) EncodeCSV() []byte {
	buf := &bytes.Buffer{}
	fmt.Fprintln(buf, CSVHeader)
	if el.Size() == 0 {
		return buf.Bytes()
	}
	for _, e := range el.list {
		fmt.Fprintln(buf, strings.Join(e.EncodeStrings(), ","))
	}
	return buf.Bytes()
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
