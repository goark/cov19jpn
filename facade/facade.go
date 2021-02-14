package facade

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/cov19jpn/ecode"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	//Name is applicatin name
	Name = "cov19jpn"
	//Version is version for applicatin
	Version = "developer version"
)
var (
	debugFlag bool //debug flag
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   Name,
		Short: "Plotting COVID-2019 in Japan",
		Long:  "Plotting COVID-2019 in Japan\nImporting Google COVID-19 Public Forecasts",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugPrint(ui, errs.Wrap(ecode.ErrNoCommand))
		},
	}
	rootCmd.SilenceUsage = true
	rootCmd.SetArgs(args)            //arguments of command-line
	rootCmd.SetIn(ui.Reader())       //Stdin
	rootCmd.SetOut(ui.ErrorWriter()) //Stdout -> Stderr
	rootCmd.SetErr(ui.ErrorWriter()) //Stderr
	rootCmd.AddCommand(
		newVersionCmd(ui),
		newCSVCmd(ui),
		newPlotCmd(ui),
	)

	//global options
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "for debug")

	return rootCmd
}

func debugPrint(ui *rwi.RWI, err error) error {
	if debugFlag && err != nil {
		fmt.Fprintf(ui.Writer(), "%+v\n", err)
	}
	return err
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
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
