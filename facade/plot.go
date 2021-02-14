package facade

import (
	"context"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/cov19jpn/chart"
	"github.com/spiegel-im-spiegel/cov19jpn/entity"
	"github.com/spiegel-im-spiegel/cov19jpn/fetch"
	"github.com/spiegel-im-spiegel/cov19jpn/filter"
	"github.com/spiegel-im-spiegel/cov19jpn/values/prefcodejpn"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newVersionCmd returns cobra.Command instance for show sub-command
func newPlotCmd(ui *rwi.RWI) *cobra.Command {
	plotCmd := &cobra.Command{
		Use:     "plot [flags] [pref name [pref name]...]",
		Aliases: []string{"p"},
		Short:   "Output chart images",
		Long:    "Output chart images.",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			out, err := cmd.Flags().GetString("output")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err))
			}
			if len(out) == 0 {
				out = "./cov19-chart.png"
			}

			prefcodes := []prefcodejpn.Code{}
			if len(args) > 0 {
				for _, arg := range args {
					code := prefcodejpn.FromPrefName(arg)
					if code == prefcodejpn.UNKNOWN {
						return debugPrint(ui, errs.New("illegal pref. name: "+arg, errs.WithCause(err), errs.WithContext("arg", arg)))
					}
					prefcodes = append(prefcodes, code)
				}
			} else {
				for i := uint(1); ; i++ {
					c := prefcodejpn.Code(i)
					prefcodes = append(prefcodes, c)
					if c == prefcodejpn.PREFCODE_MAX {
						break
					}
				}
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
			dir, filename := filepath.Split(out)
			for _, c := range prefcodes {
				sublist := list.Filer(filter.New(c))
				hlist := chart.New(sublist.StartDayMeasure(), sublist.EndDayMeasure().AddDay(7), 7, sublist)
				if err := chart.MakeHistChart(hlist, c.Title(), filepath.Join(dir, strings.Join([]string{c.Name(), filename}, "-"))); err != nil {
					_ = debugPrint(ui, errs.Wrap(err))
				}
			}
			return nil
		},
	}
	plotCmd.Flags().StringP("output", "o", "", "path of CSV file")

	return plotCmd
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
