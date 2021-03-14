package chart

import (
	"encoding/json"
	"sort"

	"github.com/spiegel-im-spiegel/cov19jpn/entity"
	"github.com/spiegel-im-spiegel/cov19jpn/values/date"
)

//HistData class
type HistData struct {
	Period       date.Period
	Cases        float64
	Deaths       float64
	Hospitalized float64
}

//HistList class is list of HistData
type HistList struct {
	list []*HistData
}

//New functions return new HistList instance.
func New(start, end date.Date, step int, elist *entity.EntityList) *HistList {
	if start.IsZero() && end.IsZero() {
		start = date.Yesterday().AddDay(-7 * 4)
		end = date.Yesterday().AddDay(7)
	}
	if start.IsZero() {
		start = end.AddDay(-7 * 5)
	}
	if end.IsZero() {
		end = start.AddDay(7 * 5)
	}
	if end.Before(start) {
		start, end = end, start
	}
	list := []*HistData{}
	last := end
	if step < 1 {
		step = 1
	}
	for {
		if start.After(last) {
			break
		}
		begin := last.AddDay(1 - step)
		list = append(list, &HistData{Period: date.NewPeriod(begin, last), Cases: 0.0, Deaths: 0.0})
		last = begin.AddDay(-1)
	}
	hlist := &HistList{list: list}
	for i := 0; i < elist.Size(); i++ {
		hlist.addData(elist.Entity(i))
	}
	return hlist.sort()
}

func (hl *HistList) sort() *HistList {
	if hl == nil {
		return hl
	}
	if len(hl.list) < 2 {
		return hl
	}
	sort.Slice(hl.list, func(i, j int) bool {
		return hl.list[i].Period.Start.Before(hl.list[j].Period.Start)
	})
	return hl
}

func (hl *HistList) addData(e *entity.Entity) {
	if hl.Size() == 0 {
		return
	}
	for i := 0; i < hl.Size(); i++ {
		h := hl.list[i]
		if h.Period.Contains(e.TargetPredictionDate) {
			h.Cases += jsonNUmberToFloat64(e.NewConfirmedValue())
			h.Deaths += jsonNUmberToFloat64(e.NewDeathsValue())
			h.Hospitalized = max(h.Hospitalized, jsonNUmberToFloat64(e.HospitalizedPatientsValue()))
		}
	}
}

//Size method returns list size.
func (hl *HistList) Size() int {
	if hl == nil {
		return 0
	}
	return len(hl.list)
}

//Data method returns HistData in list.
func (hl *HistList) Data(i int) *HistData {
	if hl == nil {
		return nil
	}
	if i < 0 || i >= hl.Size() {
		return nil
	}
	return hl.list[i]

}

func jsonNUmberToFloat64(num *json.Number) float64 {
	if num == nil {
		return 0.0
	}
	f, err := num.Float64()
	if err != nil {
		return 0.0
	}
	return f
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
