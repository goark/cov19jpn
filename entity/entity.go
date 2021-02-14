package entity

import (
	"encoding/json"

	"github.com/spiegel-im-spiegel/cov19jpn/ecode"
	"github.com/spiegel-im-spiegel/cov19jpn/values/date"
	"github.com/spiegel-im-spiegel/cov19jpn/values/prefcodejpn"
	"github.com/spiegel-im-spiegel/errs"
)

const (
	// CSV header string for "Google COVID-19 Public Forecasts in Japan"
	CSVHeader = `japan_prefecture_code,prefecture_name,target_prediction_date,cumulative_confirmed,cumulative_confirmed_q0025,cumulative_confirmed_q0975,cumulative_deaths,cumulative_deaths_q0025,cumulative_deaths_q0975,hospitalized_patients,hospitalized_patients_q0025,hospitalized_patients_q0975,recovered,recovered_q0025,recovered_q0975,cumulative_confirmed_ground_truth,cumulative_deaths_ground_truth,hospitalized_patients_ground_truth,recovered_ground_truth,forecast_date,new_deaths,new_confirmed,new_deaths_ground_truth,new_confirmed_ground_truth,prefecture_name_kanji`
	CSVCols   = 25
)

// Entity is entity class for "Google COVID-19 Public Forecasts in Japan"
type Entity struct {
	JapanPrefectureCode             prefcodejpn.Code `json:"japan_prefecture_code"`
	PrefectureName                  string           `json:"prefecture_name"`
	TargetPredictionDate            date.Date        `json:"target_prediction_date"`
	CumulativeConfirmed             *json.Number     `json:"cumulative_confirmed,omitempty"`
	CumulativeConfirmedQ0025        *json.Number     `json:"cumulative_confirmed_q0025,omitempty"`
	CumulativeConfirmedQ0975        *json.Number     `json:"cumulative_confirmed_q0975,omitempty"`
	CumulativeDeaths                *json.Number     `json:"cumulative_deaths,omitempty"`
	CumulativeDeathsQ0025           *json.Number     `json:"cumulative_deaths_q0025,omitempty"`
	CumulativeDeathsQ0975           *json.Number     `json:"cumulative_deaths_q0975,omitempty"`
	HospitalizedPatients            *json.Number     `json:"hospitalized_patients,omitempty"`
	HospitalizedPatientsQ0025       *json.Number     `json:"hospitalized_patients_q0025,omitempty"`
	HospitalizedPatientsQ0975       *json.Number     `json:"hospitalized_patients_q0975,omitempty"`
	Recovered                       *json.Number     `json:"recovered,omitempty"`
	RecoveredQ0025                  *json.Number     `json:"recovered_q0025,omitempty"`
	RecoveredQ0975                  *json.Number     `json:"recovered_q0975,omitempty"`
	CumulativeConfirmedGroundTruth  *json.Number     `json:"cumulative_confirmed_ground_truth,omitempty"`
	CumulativeDeathsGroundTruth     *json.Number     `json:"cumulative_deaths_ground_truth,omitempty"`
	HospitalizedPatientsGroundTruth *json.Number     `json:"hospitalized_patients_ground_truth,omitempty"`
	RecoveredGroundTruth            *json.Number     `json:"recovered_ground_truth,omitempty"`
	ForecastDate                    date.Date        `json:"forecast_date"`
	NewDeaths                       *json.Number     `json:"new_deaths,omitempty"`
	NewConfirmed                    *json.Number     `json:"new_confirmed,omitempty"`
	NewDeathsGroundTruth            *json.Number     `json:"new_deaths_ground_truth,omitempty"`
	NewConfirmedGroundTruth         *json.Number     `json:"new_confirmed_ground_truth,omitempty"`
	PrefectureNameKanji             string           `json:"prefecture_name_kanji,omitempty"`
}

func Decode(elements []string) (*Entity, error) {
	if len(elements) != CSVCols {
		return nil, errs.Wrap(ecode.ErrInvalidRecord)
	}
	return &Entity{
		JapanPrefectureCode:             prefcodejpn.FromCodeString(elements[0]),
		PrefectureName:                  elements[1],
		TargetPredictionDate:            date.FromString(elements[2]),
		CumulativeConfirmed:             toJSONNumber(elements[3]),
		CumulativeConfirmedQ0025:        toJSONNumber(elements[4]),
		CumulativeConfirmedQ0975:        toJSONNumber(elements[5]),
		CumulativeDeaths:                toJSONNumber(elements[6]),
		CumulativeDeathsQ0025:           toJSONNumber(elements[7]),
		CumulativeDeathsQ0975:           toJSONNumber(elements[8]),
		HospitalizedPatients:            toJSONNumber(elements[9]),
		HospitalizedPatientsQ0025:       toJSONNumber(elements[10]),
		HospitalizedPatientsQ0975:       toJSONNumber(elements[11]),
		Recovered:                       toJSONNumber(elements[12]),
		RecoveredQ0025:                  toJSONNumber(elements[13]),
		RecoveredQ0975:                  toJSONNumber(elements[14]),
		CumulativeConfirmedGroundTruth:  toJSONNumber(elements[15]),
		CumulativeDeathsGroundTruth:     toJSONNumber(elements[16]),
		HospitalizedPatientsGroundTruth: toJSONNumber(elements[17]),
		RecoveredGroundTruth:            toJSONNumber(elements[18]),
		ForecastDate:                    date.FromString(elements[19]),
		NewDeaths:                       toJSONNumber(elements[20]),
		NewConfirmed:                    toJSONNumber(elements[21]),
		NewDeathsGroundTruth:            toJSONNumber(elements[22]),
		NewConfirmedGroundTruth:         toJSONNumber(elements[23]),
		PrefectureNameKanji:             elements[24],
	}, nil
}

func (e *Entity) EncodeStrings() []string {
	elements := make([]string, CSVCols)
	elements[0] = e.JapanPrefectureCode.String()
	elements[1] = e.PrefectureName
	elements[2] = e.TargetPredictionDate.String()
	elements[3] = jsonNumberToString(e.CumulativeConfirmed)
	elements[4] = jsonNumberToString(e.CumulativeConfirmedQ0025)
	elements[5] = jsonNumberToString(e.CumulativeConfirmedQ0975)
	elements[6] = jsonNumberToString(e.CumulativeDeaths)
	elements[7] = jsonNumberToString(e.CumulativeDeathsQ0025)
	elements[8] = jsonNumberToString(e.CumulativeDeathsQ0975)
	elements[9] = jsonNumberToString(e.HospitalizedPatients)
	elements[10] = jsonNumberToString(e.HospitalizedPatientsQ0025)
	elements[11] = jsonNumberToString(e.HospitalizedPatientsQ0975)
	elements[12] = jsonNumberToString(e.Recovered)
	elements[13] = jsonNumberToString(e.RecoveredQ0025)
	elements[14] = jsonNumberToString(e.RecoveredQ0975)
	elements[15] = jsonNumberToString(e.CumulativeConfirmedGroundTruth)
	elements[16] = jsonNumberToString(e.CumulativeDeathsGroundTruth)
	elements[17] = jsonNumberToString(e.HospitalizedPatientsGroundTruth)
	elements[18] = jsonNumberToString(e.RecoveredGroundTruth)
	elements[19] = e.ForecastDate.String()
	elements[20] = jsonNumberToString(e.NewDeaths)
	elements[21] = jsonNumberToString(e.NewConfirmed)
	elements[22] = jsonNumberToString(e.NewDeathsGroundTruth)
	elements[23] = jsonNumberToString(e.NewConfirmedGroundTruth)
	elements[24] = e.PrefectureNameKanji
	return elements
}

func (e *Entity) IsForecast() bool {
	return e.ForecastDate.Before(e.TargetPredictionDate)
}

//CumulativeConfirmedValue method returns value of CumulativeConfirmed.
func (e *Entity) CumulativeConfirmedValue() *json.Number {
	if e == nil {
		return nil
	}
	if e.IsForecast() {
		return e.CumulativeConfirmed
	}
	return e.CumulativeConfirmedGroundTruth
}

//CumulativeDeathsValue method returns value of CumulativeDeaths.
func (e *Entity) CumulativeDeathsValue() *json.Number {
	if e == nil {
		return nil
	}
	if e.IsForecast() {
		return e.CumulativeDeaths
	}
	return e.CumulativeDeathsGroundTruth
}

//NewConfirmedValue method returns value of NewConfirmed.
func (e *Entity) NewConfirmedValue() *json.Number {
	if e == nil {
		return nil
	}
	if e.IsForecast() {
		return e.NewConfirmed
	}
	return e.NewConfirmedGroundTruth
}

//NewDeathsValue method returns value of NewDeaths.
func (e *Entity) NewDeathsValue() *json.Number {
	if e == nil {
		return nil
	}
	if e.IsForecast() {
		return e.NewDeaths
	}
	return e.NewDeathsGroundTruth
}

//HospitalizedPatientsValue method returns value of HospitalizedPatients.
func (e *Entity) HospitalizedPatientsValue() *json.Number {
	if e == nil {
		return nil
	}
	if e.IsForecast() {
		return e.HospitalizedPatients
	}
	return e.HospitalizedPatientsGroundTruth
}

//RecoveredValue method returns value of Recovered.
func (e *Entity) RecoveredValue() *json.Number {
	if e == nil {
		return nil
	}
	if e.IsForecast() {
		return e.Recovered
	}
	return e.RecoveredGroundTruth
}

func toJSONNumber(s string) *json.Number {
	if len(s) == 0 {
		return nil
	}
	return (*json.Number)(&s)
}

func jsonNumberToString(jn *json.Number) string {
	if jn == nil {
		return ""
	}
	return jn.String()
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
