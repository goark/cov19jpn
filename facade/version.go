package facade

import (
	"strings"

	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

var (
	usage = []string{ //output message of version
		Name + " " + Version,
		"repository: https://github.com/goark/cov19jpn",
	}
)

//newVersionCmd returns cobra.Command instance for show sub-command
func newVersionCmd(ui *rwi.RWI) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"ver", "v"},
		Short:   "Print the version number",
		Long:    "Print the version number of " + Name,
		RunE: func(cmd *cobra.Command, args []string) error {
			return ui.OutputErrln(strings.Join(usage, "\n"))
		},
	}

	return versionCmd
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
