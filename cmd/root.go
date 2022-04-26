/*
Copyright © 2022 Tony West

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"runtime"

	"github.com/spf13/cobra"
)

var (
	architectureFlag string
	dllNameFlag      string
	exeNameFlag      string
	formatFlag       string
	root             string
)

var banner string = `.------..------..------..------..------..------..------..------.
|P.--. ||L.--. ||A.--. ||Y.--. ||M.--. ||A.--. ||T.--. ||E.--. |
| :/\: || :/\: || (\/) || (\/) || (\/) || (\/) || :/\: || (\/) |
| (__) || (__) || :\/: || :\/: || :\/: || :\/: || (__) || :\/: |
| '--'P|| '--'L|| '--'A|| '--'Y|| '--'M|| '--'A|| '--'T|| '--'E|
'------''------''------''------''------''------''------''------'
playmate v0.1.0
`
var platform string = runtime.GOOS

var rootCmd = &cobra.Command{
	Use:   "playmate",
	Short: "An automated ISO generation tool",
	Long: banner + `
Playmate is a CLI tool designed to automate the workflow of 
creating autoplay CDs for Penetration Testing or Red Team use.

Often for physical penetration tests or red team engagements, the
need to insert a CD or USB for autoplay purposes will arise. This
tool seeks to automate the boring part of creating the autorun CD
in order to allow more time to be spent on the part that matters.

Use the "build" command to build an ISO containing an autorun file
for the operating system of your choice. To get more information,
use the "playmate build --help" command.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
