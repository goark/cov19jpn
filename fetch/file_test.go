package fetch

import (
	"bytes"
	"io"
	"testing"

	"github.com/goark/cov19jpn/entity"
	"github.com/goark/cov19jpn/filter"
)

func TestFile(t *testing.T) {
	testCases := []struct {
		path string
	}{
		{path: "testdata/input.csv"},
	}

	for _, tc := range testCases {
		func(path string) {
			r, err := File(tc.path)
			if err != nil {
				t.Errorf("File() is \"%v\", want nil.", err)
				return
			}
			defer r.Close()

			buf := &bytes.Buffer{}
			es, err := Import(io.TeeReader(r, buf), filter.New())
			if err != nil {
				t.Errorf("Import() is \"%v\", want nil.", err)
				return
			}
			ref := buf.Bytes()

			list := entity.NewList(es)
			b := list.EncodeCSV()
			if !bytes.Equal(ref, b) {
				t.Errorf("Import() result \"%v\", want \"%v\".", string(b), string(ref))
			}
		}(tc.path)
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
