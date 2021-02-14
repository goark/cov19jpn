package csvdata

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/spiegel-im-spiegel/cov19jpn/ecode"
)

const (
	csvheader  = "japan_prefecture_code,prefecture_name,target_prediction_date,cumulative_confirmed,cumulative_confirmed_q0025,cumulative_confirmed_q0975,cumulative_deaths,cumulative_deaths_q0025,cumulative_deaths_q0975,hospitalized_patients,hospitalized_patients_q0025,hospitalized_patients_q0975,recovered,recovered_q0025,recovered_q0975,cumulative_confirmed_ground_truth,cumulative_deaths_ground_truth,hospitalized_patients_ground_truth,recovered_ground_truth,forecast_date,new_deaths,new_confirmed,new_deaths_ground_truth,new_confirmed_ground_truth,prefecture_name_kanji\n"
	csvrecord1 = "JP-01,HOKKAIDO,2021-01-26,16818.171875,16756.521484375,17568.35546875,578.6916503906,575.696472168,604.0839233398,1376.7945556641,1365.6677246094,1491.4765625,14849.623046875,14803.416015625,15458.55859375,,,,,2021-01-25,3.6916503905999889,103.171875,,,北海道\n"
	csvrecord2 = "JP-01,HOKKAIDO,2021-01-27,16896.451171875,16778.654296875,18325.40234375,582.3676757812,576.6384887695,630.9478759766,1338.0850830078,1317.5307617188,1549.9569091797,14959.9169921875,14871.4765625,16125.6875,,,,,2021-01-25,3.6760253905999889,78.279296875,,,北海道\n"
	csvString  = csvheader + csvrecord1 + csvrecord2
)

func TestPrefCodeJpn(t *testing.T) {
	testCases := []struct {
		cols    int
		csv     string
		header  string
		record1 string
		record2 string
	}{
		{cols: 25, csv: csvString, header: csvheader, record1: csvrecord1, record2: csvrecord2},
	}

	for _, tc := range testCases {
		r := New(strings.NewReader(tc.csv), tc.cols, true)
		elm, err := r.Next()
		if errors.Is(err, io.EOF) || errors.Is(err, ecode.ErrInvalidRecord) {
			t.Errorf("Next() is \"%v\", want nil.", err)
		} else {
			line := strings.Join(elm, ",") + "\n"
			if line != tc.record1 {
				t.Errorf("Next() = \"%v\", want \"%v\".", line, tc.record1)
			}
		}
		elm, err = r.Next()
		if errors.Is(err, io.EOF) || errors.Is(err, ecode.ErrInvalidRecord) {
			t.Errorf("Next() is \"%v\", want nil.", err)
		} else {
			line := strings.Join(elm, ",") + "\n"
			if line != tc.record2 {
				t.Errorf("Next() = \"%v\", want \"%v\".", line, tc.record2)
			}
		}
		elm, _ = r.Header()
		line := strings.Join(elm, ",") + "\n"
		if line != tc.header {
			t.Errorf("Header() = \"%v\", want \"%v\".", line, tc.header)
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
