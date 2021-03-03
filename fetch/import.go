package fetch

import (
	"io"

	"github.com/spiegel-im-spiegel/cov19jpn/entity"
	"github.com/spiegel-im-spiegel/cov19jpn/filter"
	"github.com/spiegel-im-spiegel/csvdata"
	"github.com/spiegel-im-spiegel/errs"
)

//Import function returns slice of entity.Entity
func Import(r io.Reader, f *filter.Filter) ([]*entity.Entity, error) {
	list := entity.NewList(nil)
	cr := csvdata.New(r, true).WithFieldsPerRecord(entity.CSVCols)
	for {
		if err := cr.Next(); err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return nil, errs.Wrap(err)
		}
		e, err := entity.Decode(cr.Row())
		if err != nil {
			return nil, errs.Wrap(err)
		}
		if f.Match(e.TargetPredictionDate, e.JapanPrefectureCode) {
			list.Add(e)
		}
	}
	return list.Entities(), nil
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
