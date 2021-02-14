package date

import "fmt"

//Period is class of time scope
type Period struct {
	Start Date
	End   Date
}

//NewPeriod function creates new Period instance
func NewPeriod(from, to Date) Period {
	if from.IsZero() || to.IsZero() {
		return Period{Start: from, End: to}
	}
	if from.After(to) {
		return Period{Start: to, End: from}
	}
	return Period{Start: from, End: to}
}

//Equal method returns true if all elemnts in Period instance equal
func (lp Period) Equal(rp Period) bool {
	return lp.Start == rp.Start && lp.End == rp.End
}

//Contains method returns true if scape of this contains date of parameter.
func (p Period) Contains(dt Date) bool {
	if dt.IsZero() {
		return false
	}
	if p.IsZero() {
		return true
	}
	if p.IsZeroEnd() {
		return !p.Start.After(dt)
	}
	if p.IsZeroStart() {
		return !p.End.Before(dt)
	}
	return !p.Start.After(dt) && !p.End.Before(dt)
}

//IsZeroStart method returns trus if Start element is zero values.
func (p Period) IsZeroStart() bool {
	return p.Start.IsZero()
}

//IsZeroEnd method returns trus if End element is zero values.
func (p Period) IsZeroEnd() bool {
	return p.End.IsZero()
}

//IsZero method returns trus if all elements are zero values.
func (p Period) IsZero() bool {
	return p.IsZeroStart() && p.IsZeroEnd()
}

//StringStart method returns string of Period.Start
func (p Period) StringStart() string {
	if p.IsZeroStart() {
		return ""
	}
	return p.Start.String()
}

//StringEnd method returns string of Period.End
func (p Period) StringEnd() string {
	if p.IsZeroEnd() {
		return ""
	}
	return p.End.String()
}

//String method is fmt.Stringer for Period
func (p Period) String() string {
	if p.IsZero() {
		return ""
	}
	if p.IsZeroStart() {
		return fmt.Sprintf("> %s", p.StringEnd())
	}
	if p.IsZeroEnd() {
		return fmt.Sprintf("%s >", p.StringStart())
	}
	return fmt.Sprintf("%s > %s", p.StringStart(), p.StringEnd())
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
