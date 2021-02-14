package fetch

import (
	"io"

	"github.com/spiegel-im-spiegel/cov19jpn/csvdata"
	"github.com/spiegel-im-spiegel/cov19jpn/entity"
	"github.com/spiegel-im-spiegel/cov19jpn/filter"
	"github.com/spiegel-im-spiegel/errs"
)

//Import function returns slice of entity.Entity
func Import(r io.Reader, f *filter.Filter) ([]*entity.Entity, error) {
	es := []*entity.Entity{}
	cr := csvdata.New(r, entity.CSVCols, true)
	for {
		elms, err := cr.Next()
		if err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return nil, errs.Wrap(err)
		}
		e, err := entity.Decode(elms)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		if f.Match(e.TargetPredictionDate, e.JapanPrefectureCode) {
			es = append(es, e)
		}
	}
	return es, nil
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