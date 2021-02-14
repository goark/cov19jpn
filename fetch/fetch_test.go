package fetch

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/spiegel-im-spiegel/cov19jpn/entity"
	"github.com/spiegel-im-spiegel/cov19jpn/filter"
	"github.com/spiegel-im-spiegel/cov19jpn/values/date"
	"github.com/spiegel-im-spiegel/cov19jpn/values/prefcodejpn"
	fch "github.com/spiegel-im-spiegel/fetch"
)

type mockResp struct {
	rc io.ReadCloser
}

func (mr *mockResp) Request() *http.Request            { return nil }
func (mr *mockResp) Header() http.Header               { return nil }
func (mr *mockResp) Body() io.ReadCloser               { return mr.rc }
func (mr *mockResp) Close()                            { mr.rc.Close() }
func (mr *mockResp) DumpBodyAndClose() ([]byte, error) { return nil, nil }

type mockClient struct {
	path string
}

func (mc *mockClient) Get(u *url.URL, opts ...fch.RequestOpts) (fch.Response, error) {
	r, err := File(mc.path)
	return &mockResp{r}, err
}

func (mc *mockClient) Post(u *url.URL, payload io.Reader, opts ...fch.RequestOpts) (fch.Response, error) {
	return mc.Get(u, opts...)
}

func getReference(path string) []byte {
	r, err := File(path)
	if err != nil {
		return nil
	}
	defer r.Close()
	buf := &bytes.Buffer{}
	_, _ = io.Copy(buf, r)
	return buf.Bytes()
}

func TestWeb(t *testing.T) {
	testCases := []struct {
		input  string
		output string
		pref   prefcodejpn.Code
	}{
		{input: "testdata/input.csv", output: "testdata/output-shimane.csv", pref: prefcodejpn.SHIMANE},
	}

	for _, tc := range testCases {
		func() {
			r, err := fetchWeb(context.Background(), &mockClient{tc.input})
			if err != nil {
				t.Errorf("File() is \"%v\", want nil.", err)
				return
			}
			defer r.Close()

			es, err := Import(
				r,
				filter.New(tc.pref).SetPeriod(date.NewPeriod(date.Zero, date.FromString("2021-01-24").AddDay(7+1))),
			)
			if err != nil {
				t.Errorf("Import() is \"%v\", want nil.", err)
				return
			}
			ref := getReference(tc.output)

			list := entity.NewList(es)
			b := list.EncodeCSV()
			if !bytes.Equal(ref, b) {
				t.Errorf("Import() result \"%v\", want \"%v\".", string(b), string(ref))
			}
		}()
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
