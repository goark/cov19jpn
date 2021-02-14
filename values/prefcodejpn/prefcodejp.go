package prefcodejpn

import (
	"fmt"
	"strconv"
	"strings"
)

//Code is prefecture code in Japan
type Code uint

const (
	UNKNOWN Code = iota
	HOKKAIDO
	AOMORI
	IWATE
	MIYAGI
	AKITA
	YAMAGATA
	FUKUSHIMA
	IBARAKI
	TOCHIGI
	GUNMA
	SAITAMA
	CHIBA
	TOKYO
	KANAGAWA
	NIIGATA
	TOYAMA
	ISHIKAWA
	FUKUI
	YAMANASHI
	NAGANO
	GIFU
	SHIZUOKA
	AICHI
	MIE
	SHIGA
	KYOTO
	OSAKA
	HYOGO
	NARA
	WAKAYAMA
	TOTTORI
	SHIMANE
	OKAYAMA
	HIROSHIMA
	YAMAGUCHI
	TOKUSHIMA
	KAGAWA
	EHIME
	KOCHI
	FUKUOKA
	SAGA
	NAGASAKI
	KUMAMOTO
	OITA
	MIYAZAKI
	KAGOSHIMA
	OKINAWA
)

var (
	PREFCODE_MIN = HOKKAIDO
	PREFCODE_MAX = OKINAWA
	pcNames      = map[Code]string{
		HOKKAIDO:  "hokkaido",
		AOMORI:    "aomori",
		IWATE:     "iwate",
		MIYAGI:    "miyagi",
		AKITA:     "akita",
		YAMAGATA:  "yamagata",
		FUKUSHIMA: "fukushima",
		IBARAKI:   "ibaraki",
		TOCHIGI:   "tochigi",
		GUNMA:     "gunma",
		SAITAMA:   "saitama",
		CHIBA:     "chiba",
		TOKYO:     "tokyo",
		KANAGAWA:  "kanagawa",
		NIIGATA:   "niigata",
		TOYAMA:    "toyama",
		ISHIKAWA:  "ishikawa",
		FUKUI:     "fukui",
		YAMANASHI: "yamanashi",
		NAGANO:    "nagano",
		GIFU:      "gifu",
		SHIZUOKA:  "shizuoka",
		AICHI:     "aichi",
		MIE:       "mie",
		SHIGA:     "shiga",
		KYOTO:     "kyoto",
		OSAKA:     "osaka",
		HYOGO:     "hyogo",
		NARA:      "nara",
		WAKAYAMA:  "wakayama",
		TOTTORI:   "tottori",
		SHIMANE:   "shimane",
		OKAYAMA:   "okayama",
		HIROSHIMA: "hiroshima",
		YAMAGUCHI: "yamaguchi",
		TOKUSHIMA: "tokushima",
		KAGAWA:    "kagawa",
		EHIME:     "ehime",
		KOCHI:     "kochi",
		FUKUOKA:   "fukuoka",
		SAGA:      "saga",
		NAGASAKI:  "nagasaki",
		KUMAMOTO:  "kumamoto",
		OITA:      "oita",
		MIYAZAKI:  "miyazaki",
		KAGOSHIMA: "kagoshima",
		OKINAWA:   "okinawa",
	}
)

func New(cd uint) Code {
	if cd < uint(PREFCODE_MIN) || cd > uint(PREFCODE_MAX) {
		return UNKNOWN
	}
	return Code(cd)
}

func FromCodeString(s string) Code {
	s = strings.TrimPrefix(strings.ToUpper(strings.TrimSpace(s)), "JP-")
	n, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return UNKNOWN
	}
	return New(uint(n))
}

func FromPrefName(s string) Code {
	for k, v := range pcNames {
		if strings.EqualFold(v, s) {
			return k
		}
	}
	return UNKNOWN
}

//Equal method returns true if lpc == rpc
func (lpc Code) Equal(rpc Code) bool {
	return lpc == rpc
}

//Less method returns true if lpc < rpc
func (lpc Code) Less(rpc Code) bool {
	return uint(lpc) < uint(rpc)
}

func (pc Code) String() string {
	return fmt.Sprintf("JP-%02d", uint(pc))
}

func (pc Code) Name() string {
	for k, v := range pcNames {
		if pc == k {
			return v
		}
	}
	return "unknown"
}

func (pc Code) NameUpper() string {
	return strings.ToUpper(pc.Name())
}

func (pc Code) Title() string {
	return strings.Title(pc.Name())
}

//UnmarshalJSON method returns result of Unmarshal for json.Unmarshal().
func (pc *Code) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	*pc = FromCodeString(s)
	return nil
}

//MarshalJSON method returns string for json.Marshal().
func (pc Code) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(pc.String())), nil
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
