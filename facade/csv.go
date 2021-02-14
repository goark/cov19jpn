package facade

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/cov19jpn/entity"
	"github.com/spiegel-im-spiegel/cov19jpn/fetch"
	"github.com/spiegel-im-spiegel/cov19jpn/filter"
	"github.com/spiegel-im-spiegel/cov19jpn/values/prefcodejpn"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newVersionCmd returns cobra.Command instance for show sub-command
func newCSVCmd(ui *rwi.RWI) *cobra.Command {
	csvCmd := &cobra.Command{
		Use:     "csv [flags] [pref name [pref name]...]",
		Aliases: []string{"c"},
		Short:   "Output text formatting CSV",
		Long:    "Output text formatting CSV.",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			out, err := cmd.Flags().GetString("output")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err))
			}
			//Output stream
			w := ui.Writer()
			if len(out) > 0 {
				file, err := os.Create(out)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, errs.WithContext("output", out)))
				}
				defer file.Close()
				w = file
			}

			prefcodes := []prefcodejpn.Code{}
			for _, arg := range args {
				code := prefcodejpn.FromPrefName(arg)
				if code == prefcodejpn.UNKNOWN {
					return debugPrint(ui, errs.New("illegal pref. name: "+arg, errs.WithCause(err), errs.WithContext("arg", arg)))
				}
				prefcodes = append(prefcodes, code)
			}

			r, err := fetch.Web(context.Background(), &http.Client{})
			if err != nil {
				return debugPrint(ui, errs.Wrap(err))
			}
			defer r.Close()
			es, err := fetch.Import(r, filter.New(prefcodes...))
			if err != nil {
				return debugPrint(ui, errs.Wrap(err))
			}
			list := entity.NewList(es)

			list.Sort()
			if _, err := io.Copy(w, bytes.NewReader(list.EncodeCSV())); err != nil {
				return debugPrint(ui, errs.Wrap(err))
			}
			return nil
		},
	}
	csvCmd.Flags().StringP("output", "o", "", "path of CSV file")

	return csvCmd
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
