package fetch

import (
	"context"
	"io"
	"net/http"

	"github.com/spiegel-im-spiegel/errs"
	fch "github.com/spiegel-im-spiegel/fetch"
)

//NewWeb returns new Import instance
func Web(ctx context.Context, c *http.Client) (io.ReadCloser, error) {
	r, err := fetchWeb(ctx, fch.New(fch.WithHTTPClient(c)))
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return r, nil
}

func fetchWeb(ctx context.Context, c fch.Client) (io.ReadCloser, error) {
	u, err := fch.URL("https://storage.googleapis.com/covid-external/forecast_JAPAN_PREFECTURE_28.csv")
	if err != nil {
		return nil, errs.Wrap(err)
	}
	resp, err := c.Get(u, fch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return resp.Body(), nil

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
