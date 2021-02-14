package csvdata

import (
	"encoding/csv"
	"io"

	"github.com/spiegel-im-spiegel/cov19jpn/ecode"
	"github.com/spiegel-im-spiegel/errs"
)

//Reader is class of CSV reader
type Reader struct {
	reader        *csv.Reader
	cols          int
	headerFlag    bool
	headerStrings []string
}

//New function creates a new Reader instance.
func New(r io.Reader, cols int, headerFlag bool) *Reader {
	cr := csv.NewReader(r)
	cr.Comma = ','
	cr.LazyQuotes = true       // a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.
	cr.TrimLeadingSpace = true // leading
	return &Reader{reader: cr, cols: cols, headerFlag: headerFlag}
}

//readRecord method returns a new record.
func (r *Reader) readRecord() ([]string, error) {
	elms, err := r.reader.Read()
	if err != nil {
		if errs.Is(err, io.EOF) {
			return nil, errs.Wrap(ecode.ErrNoData, errs.WithCause(err))
		}
		return nil, errs.Wrap(ecode.ErrInvalidRecord, errs.WithCause(err))
	}
	if len(elms) < r.cols {
		return nil, errs.Wrap(ecode.ErrInvalidRecord, errs.WithContext("record", elms))
	}
	return elms, nil
}

//Header method returns header strings.
func (r *Reader) Header() ([]string, error) {
	var err error
	if r.headerFlag {
		r.headerFlag = false
		r.headerStrings, err = r.readRecord()
	}
	return r.headerStrings, errs.Wrap(err)
}

//Next method returns a next record.
func (r *Reader) Next() ([]string, error) {
	if r.headerFlag {
		if _, err := r.Header(); err != nil {
			return nil, errs.Wrap(err)
		}
	}
	elms, err := r.readRecord()
	return elms, errs.Wrap(err)
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
