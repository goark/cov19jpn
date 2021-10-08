package entity

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/spiegel-im-spiegel/cov19jpn/ecode"
	"github.com/spiegel-im-spiegel/cov19jpn/filter"
	"github.com/spiegel-im-spiegel/cov19jpn/values/date"
	"github.com/spiegel-im-spiegel/csvdata"
)

const (
	csvOrg = `japan_prefecture_code,prefecture_name,target_prediction_date,cumulative_confirmed,cumulative_confirmed_q0025,cumulative_confirmed_q0975,cumulative_deaths,cumulative_deaths_q0025,cumulative_deaths_q0975,hospitalized_patients,hospitalized_patients_q0025,hospitalized_patients_q0975,recovered,recovered_q0025,recovered_q0975,cumulative_confirmed_ground_truth,cumulative_deaths_ground_truth,hospitalized_patients_ground_truth,recovered_ground_truth,forecast_date,new_deaths,new_confirmed,new_deaths_ground_truth,new_confirmed_ground_truth,prefecture_name_kanji
JP-32,SHIMANE,2021-01-26,246.7317199707,245.9146270752,258.7432861328,4.5689563751,4.5458364487,4.7951383591,10.5724658966,10.4934711456,11.4544467926,231.5160980225,230.8079528809,242.4129638672,,,,,2021-01-25,4.5689563751,4.7317199706999986,,,島根県
JP-32,SHIMANE,2021-01-27,246.8562927246,245.2927398682,269.6721496582,4.5829758644,4.5389103889,5.0144014359,10.0738105774,9.9307956696,11.6712617874,232.0361328125,230.688079834,252.795791626,,,,,2021-01-25,0.014019489299999854,0.12457275389999722,,,島根県
JP-32,SHIMANE,2021-01-28,247.0369262695,244.7880859375,279.5713806152,4.6005039215,4.5374093056,5.21860075,9.543586731,9.3503313065,11.7029485703,232.6273345947,230.7003326416,262.3211669922,,,,,2021-01-25,0.017528057099999828,0.18063354489999028,,,島根県
JP-32,SHIMANE,2021-01-29,247.2489318848,244.3757781982,288.4949035645,4.6192407608,4.5388202667,5.406701088,9.0601406097,8.8272609711,11.6619529724,233.220413208,230.769317627,270.9718933105,,,,,2021-01-25,0.018736839300000696,0.21200561530000073,,,島根県
JP-32,SHIMANE,2021-01-30,247.4383087158,244.0038146973,296.5071411133,4.6387639046,4.5425610542,5.5803847313,8.6011285782,8.3377904892,11.5428380966,233.7939910889,230.8684234619,278.8346252441,,,,,2021-01-25,0.019523143799999865,0.18937683100000413,,,島根県
JP-32,SHIMANE,2021-01-31,247.632019043,243.6917877197,303.7306213379,4.6583194733,4.5477347374,5.7402420044,8.1748638153,7.8883895874,11.3744478226,234.3528137207,230.9974517822,285.9873657227,,,,,2021-01-25,0.019555568699999526,0.1937103272000229,,,島根県
JP-32,SHIMANE,2021-02-01,247.8811340332,243.4769134521,310.2735900879,4.6775584221,4.5538425446,5.8864283562,7.816740036,7.5115981102,11.2230625153,234.8936767578,231.1488647461,292.4465942383,,,,,2021-01-25,0.019238948799999989,0.24911499019998473,,,島根県
JP-32,SHIMANE,2020-12-29,,,,,,,,,,,,,208,0,31,177,2021-01-25,,,,,島根県
JP-32,SHIMANE,2020-12-30,,,,,,,,,,,,,208,0,31,177,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2020-12-31,,,,,,,,,,,,,207,0,29,178,2021-01-25,,,0,-1,島根県
JP-32,SHIMANE,2021-01-01,,,,,,,,,,,,,209,0,28,181,2021-01-25,,,0,2,島根県
JP-32,SHIMANE,2021-01-02,,,,,,,,,,,,,213,0,31,182,2021-01-25,,,0,4,島根県
JP-32,SHIMANE,2021-01-03,,,,,,,,,,,,,214,0,27,187,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-04,,,,,,,,,,,,,214,0,25,189,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-05,,,,,,,,,,,,,215,0,22,193,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-06,,,,,,,,,,,,,216,0,20,196,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-07,,,,,,,,,,,,,218,0,16,202,2021-01-25,,,0,2,島根県
JP-32,SHIMANE,2021-01-08,,,,,,,,,,,,,221,0,16,205,2021-01-25,,,0,3,島根県
JP-32,SHIMANE,2021-01-09,,,,,,,,,,,,,224,0,20,204,2021-01-25,,,0,3,島根県
JP-32,SHIMANE,2021-01-10,,,,,,,,,,,,,227,0,20,207,2021-01-25,,,0,3,島根県
JP-32,SHIMANE,2021-01-11,,,,,,,,,,,,,227,0,21,206,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-12,,,,,,,,,,,,,227,0,20,207,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-13,,,,,,,,,,,,,228,0,19,209,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-14,,,,,,,,,,,,,230,0,16,214,2021-01-25,,,0,2,島根県
JP-32,SHIMANE,2021-01-15,,,,,,,,,,,,,233,0,20,213,2021-01-25,,,0,3,島根県
JP-32,SHIMANE,2021-01-16,,,,,,,,,,,,,234,0,15,219,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-17,,,,,,,,,,,,,235,0,14,221,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-18,,,,,,,,,,,,,235,0,16,219,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-19,,,,,,,,,,,,,237,0,14,223,2021-01-25,,,0,2,島根県
JP-32,SHIMANE,2021-01-20,,,,,,,,,,,,,239,0,15,224,2021-01-25,,,0,2,島根県
JP-32,SHIMANE,2021-01-21,,,,,,,,,,,,,242,0,15,227,2021-01-25,,,0,3,島根県
JP-32,SHIMANE,2021-01-22,,,,,,,,,,,,,242,0,14,228,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-23,,,,,,,,,,,,,242,0,12,230,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-24,,,,,,,,,,,,,242,0,,,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-25,,,,,,,,,,,,,242,0,,,2021-01-25,,,0,0,島根県
`
	csvMeasure = `japan_prefecture_code,prefecture_name,target_prediction_date,cumulative_confirmed,cumulative_confirmed_q0025,cumulative_confirmed_q0975,cumulative_deaths,cumulative_deaths_q0025,cumulative_deaths_q0975,hospitalized_patients,hospitalized_patients_q0025,hospitalized_patients_q0975,recovered,recovered_q0025,recovered_q0975,cumulative_confirmed_ground_truth,cumulative_deaths_ground_truth,hospitalized_patients_ground_truth,recovered_ground_truth,forecast_date,new_deaths,new_confirmed,new_deaths_ground_truth,new_confirmed_ground_truth,prefecture_name_kanji
JP-32,SHIMANE,2020-12-29,,,,,,,,,,,,,208,0,31,177,2021-01-25,,,,,島根県
JP-32,SHIMANE,2020-12-30,,,,,,,,,,,,,208,0,31,177,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2020-12-31,,,,,,,,,,,,,207,0,29,178,2021-01-25,,,0,-1,島根県
JP-32,SHIMANE,2021-01-01,,,,,,,,,,,,,209,0,28,181,2021-01-25,,,0,2,島根県
JP-32,SHIMANE,2021-01-02,,,,,,,,,,,,,213,0,31,182,2021-01-25,,,0,4,島根県
JP-32,SHIMANE,2021-01-03,,,,,,,,,,,,,214,0,27,187,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-04,,,,,,,,,,,,,214,0,25,189,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-05,,,,,,,,,,,,,215,0,22,193,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-06,,,,,,,,,,,,,216,0,20,196,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-07,,,,,,,,,,,,,218,0,16,202,2021-01-25,,,0,2,島根県
JP-32,SHIMANE,2021-01-08,,,,,,,,,,,,,221,0,16,205,2021-01-25,,,0,3,島根県
JP-32,SHIMANE,2021-01-09,,,,,,,,,,,,,224,0,20,204,2021-01-25,,,0,3,島根県
JP-32,SHIMANE,2021-01-10,,,,,,,,,,,,,227,0,20,207,2021-01-25,,,0,3,島根県
JP-32,SHIMANE,2021-01-11,,,,,,,,,,,,,227,0,21,206,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-12,,,,,,,,,,,,,227,0,20,207,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-13,,,,,,,,,,,,,228,0,19,209,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-14,,,,,,,,,,,,,230,0,16,214,2021-01-25,,,0,2,島根県
JP-32,SHIMANE,2021-01-15,,,,,,,,,,,,,233,0,20,213,2021-01-25,,,0,3,島根県
JP-32,SHIMANE,2021-01-16,,,,,,,,,,,,,234,0,15,219,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-17,,,,,,,,,,,,,235,0,14,221,2021-01-25,,,0,1,島根県
JP-32,SHIMANE,2021-01-18,,,,,,,,,,,,,235,0,16,219,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-19,,,,,,,,,,,,,237,0,14,223,2021-01-25,,,0,2,島根県
JP-32,SHIMANE,2021-01-20,,,,,,,,,,,,,239,0,15,224,2021-01-25,,,0,2,島根県
JP-32,SHIMANE,2021-01-21,,,,,,,,,,,,,242,0,15,227,2021-01-25,,,0,3,島根県
JP-32,SHIMANE,2021-01-22,,,,,,,,,,,,,242,0,14,228,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-23,,,,,,,,,,,,,242,0,12,230,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-24,,,,,,,,,,,,,242,0,,,2021-01-25,,,0,0,島根県
JP-32,SHIMANE,2021-01-25,,,,,,,,,,,,,242,0,,,2021-01-25,,,0,0,島根県
`
	csvForecasts = `japan_prefecture_code,prefecture_name,target_prediction_date,cumulative_confirmed,cumulative_confirmed_q0025,cumulative_confirmed_q0975,cumulative_deaths,cumulative_deaths_q0025,cumulative_deaths_q0975,hospitalized_patients,hospitalized_patients_q0025,hospitalized_patients_q0975,recovered,recovered_q0025,recovered_q0975,cumulative_confirmed_ground_truth,cumulative_deaths_ground_truth,hospitalized_patients_ground_truth,recovered_ground_truth,forecast_date,new_deaths,new_confirmed,new_deaths_ground_truth,new_confirmed_ground_truth,prefecture_name_kanji
JP-32,SHIMANE,2021-01-26,246.7317199707,245.9146270752,258.7432861328,4.5689563751,4.5458364487,4.7951383591,10.5724658966,10.4934711456,11.4544467926,231.5160980225,230.8079528809,242.4129638672,,,,,2021-01-25,4.5689563751,4.7317199706999986,,,島根県
JP-32,SHIMANE,2021-01-27,246.8562927246,245.2927398682,269.6721496582,4.5829758644,4.5389103889,5.0144014359,10.0738105774,9.9307956696,11.6712617874,232.0361328125,230.688079834,252.795791626,,,,,2021-01-25,0.014019489299999854,0.12457275389999722,,,島根県
JP-32,SHIMANE,2021-01-28,247.0369262695,244.7880859375,279.5713806152,4.6005039215,4.5374093056,5.21860075,9.543586731,9.3503313065,11.7029485703,232.6273345947,230.7003326416,262.3211669922,,,,,2021-01-25,0.017528057099999828,0.18063354489999028,,,島根県
JP-32,SHIMANE,2021-01-29,247.2489318848,244.3757781982,288.4949035645,4.6192407608,4.5388202667,5.406701088,9.0601406097,8.8272609711,11.6619529724,233.220413208,230.769317627,270.9718933105,,,,,2021-01-25,0.018736839300000696,0.21200561530000073,,,島根県
JP-32,SHIMANE,2021-01-30,247.4383087158,244.0038146973,296.5071411133,4.6387639046,4.5425610542,5.5803847313,8.6011285782,8.3377904892,11.5428380966,233.7939910889,230.8684234619,278.8346252441,,,,,2021-01-25,0.019523143799999865,0.18937683100000413,,,島根県
JP-32,SHIMANE,2021-01-31,247.632019043,243.6917877197,303.7306213379,4.6583194733,4.5477347374,5.7402420044,8.1748638153,7.8883895874,11.3744478226,234.3528137207,230.9974517822,285.9873657227,,,,,2021-01-25,0.019555568699999526,0.1937103272000229,,,島根県
JP-32,SHIMANE,2021-02-01,247.8811340332,243.4769134521,310.2735900879,4.6775584221,4.5538425446,5.8864283562,7.816740036,7.5115981102,11.2230625153,234.8936767578,231.1488647461,292.4465942383,,,,,2021-01-25,0.019238948799999989,0.24911499019998473,,,島根県
`
)

func TestEntity(t *testing.T) {
	testCases := []struct {
		input string
		res1  []byte
		res2  []byte
	}{
		{input: csvOrg, res1: []byte(csvMeasure), res2: []byte(csvForecasts)},
	}

	for _, tc := range testCases {
		r := csvdata.NewRows(csvdata.New(strings.NewReader(tc.input)).WithFieldsPerRecord(CSVCols), true)
		list := NewList(nil)
		for {
			if err := r.Next(); err != nil {
				break
			}
			e, err := Decode(r.Row())
			if errors.Is(err, ecode.ErrInvalidRecord) {
				t.Errorf("Decode() is \"%v\", want nil.", err)
				continue
			}
			list.Add(e)
		}
		list0 := NewList(list.Entities())
		list0.Sort()
		list1 := list0.Filer(filter.New().SetPeriod(date.NewPeriod(list0.StartDayMeasure(), list0.EndDayMeasure())))
		b := list1.EncodeCSV()
		if !bytes.Equal(tc.res1, b) {
			t.Errorf("Filer() result \"%v\", want \"%v\".", string(b), string(tc.res1))
		}
		list1 = list0.Filer(filter.New().SetPeriod(date.NewPeriod(list0.StartDayForecast(), list0.EndDayForecast())))
		b = list1.EncodeCSV()
		if !bytes.Equal(tc.res2, b) {
			t.Errorf("Filer() result \"%v\", want \"%v\".", string(b), string(tc.res2))
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
