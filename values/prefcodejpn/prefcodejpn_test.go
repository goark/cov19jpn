package prefcodejpn

import "testing"

func TestPrefCodeJpn(t *testing.T) {
	testCases := []struct {
		code    Code
		name    string
		codeStr string
	}{
		{code: Code(0), name: "unknown", codeStr: "JP-00"},
		{code: HOKKAIDO, name: "hokkaido", codeStr: "JP-01"},
		{code: AOMORI, name: "aomori", codeStr: "JP-02"},
		{code: IWATE, name: "iwate", codeStr: "JP-03"},
		{code: MIYAGI, name: "miyagi", codeStr: "JP-04"},
		{code: AKITA, name: "akita", codeStr: "JP-05"},
		{code: YAMAGATA, name: "yamagata", codeStr: "JP-06"},
		{code: FUKUSHIMA, name: "fukushima", codeStr: "JP-07"},
		{code: IBARAKI, name: "ibaraki", codeStr: "JP-08"},
		{code: TOCHIGI, name: "tochigi", codeStr: "JP-09"},
		{code: GUNMA, name: "gunma", codeStr: "JP-10"},
		{code: SAITAMA, name: "saitama", codeStr: "JP-11"},
		{code: CHIBA, name: "chiba", codeStr: "JP-12"},
		{code: TOKYO, name: "tokyo", codeStr: "JP-13"},
		{code: KANAGAWA, name: "kanagawa", codeStr: "JP-14"},
		{code: NIIGATA, name: "niigata", codeStr: "JP-15"},
		{code: TOYAMA, name: "toyama", codeStr: "JP-16"},
		{code: ISHIKAWA, name: "ishikawa", codeStr: "JP-17"},
		{code: FUKUI, name: "fukui", codeStr: "JP-18"},
		{code: YAMANASHI, name: "yamanashi", codeStr: "JP-19"},
		{code: NAGANO, name: "nagano", codeStr: "JP-20"},
		{code: GIFU, name: "gifu", codeStr: "JP-21"},
		{code: SHIZUOKA, name: "shizuoka", codeStr: "JP-22"},
		{code: AICHI, name: "aichi", codeStr: "JP-23"},
		{code: MIE, name: "mie", codeStr: "JP-24"},
		{code: SHIGA, name: "shiga", codeStr: "JP-25"},
		{code: KYOTO, name: "kyoto", codeStr: "JP-26"},
		{code: OSAKA, name: "osaka", codeStr: "JP-27"},
		{code: HYOGO, name: "hyogo", codeStr: "JP-28"},
		{code: NARA, name: "nara", codeStr: "JP-29"},
		{code: WAKAYAMA, name: "wakayama", codeStr: "JP-30"},
		{code: TOTTORI, name: "tottori", codeStr: "JP-31"},
		{code: SHIMANE, name: "shimane", codeStr: "JP-32"},
		{code: OKAYAMA, name: "okayama", codeStr: "JP-33"},
		{code: HIROSHIMA, name: "hiroshima", codeStr: "JP-34"},
		{code: YAMAGUCHI, name: "yamaguchi", codeStr: "JP-35"},
		{code: TOKUSHIMA, name: "tokushima", codeStr: "JP-36"},
		{code: KAGAWA, name: "kagawa", codeStr: "JP-37"},
		{code: EHIME, name: "ehime", codeStr: "JP-38"},
		{code: KOCHI, name: "kochi", codeStr: "JP-39"},
		{code: FUKUOKA, name: "fukuoka", codeStr: "JP-40"},
		{code: SAGA, name: "saga", codeStr: "JP-41"},
		{code: NAGASAKI, name: "nagasaki", codeStr: "JP-42"},
		{code: KUMAMOTO, name: "kumamoto", codeStr: "JP-43"},
		{code: OITA, name: "oita", codeStr: "JP-44"},
		{code: MIYAZAKI, name: "miyazaki", codeStr: "JP-45"},
		{code: KAGOSHIMA, name: "kagoshima", codeStr: "JP-46"},
		{code: OKINAWA, name: "okinawa", codeStr: "JP-47"},
	}

	for _, tc := range testCases {
		pc := FromPrefName(tc.name)
		if pc != tc.code {
			t.Errorf("FromPrefName(%v) = \"%v\", want \"%v\".", tc.name, pc.String(), tc.code.String())
		}
		if pc.Name() != tc.name {
			t.Errorf("FromPrefName(%v) = \"%v\", want \"%v\".", tc.name, pc.String(), tc.name)
		}
		pc = FromCodeString(tc.codeStr)
		if pc != tc.code {
			t.Errorf("FromCodeString(%v) = \"%v\", want \"%v\".", tc.codeStr, pc.String(), tc.code.String())
		}
		if pc.String() != tc.codeStr {
			t.Errorf("FromPrefName(%v) = \"%v\", want \"%v\".", tc.name, pc.String(), tc.codeStr)
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
