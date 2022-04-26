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
	"log"

	"github.com/spf13/cobra"
	"github.com/un4gi/playmate/compile"
	"github.com/un4gi/playmate/stage"
)

var buildCmd = &cobra.Command{
	Use:   "build [source folder] [destination folder]",
	Short: "Build an ISO prepped for autoplay",
	Long: banner + `
-- Windows --
To build an autoplay CD for Windows, you will first need to stage
your payload in the "examples/windows" directory. At this time, only C++ 
payloads are supported. The "examples/windows" directory currently includes
two example payloads, "dwmapi-winexec.cpp" and "LaunchU3.cpp". The
"dwmapi-winexec.cpp" file will be used to create a .dll, and the
"LaunchU3.cpp" file can be used optionally as a launcher for the dll.
Note that if you change the name of you payload, you will need to 
specify the new name using the "-d" and/or "-e" flags.

The examples/windows/ directory also contains an example "autorun.inf" file
and "folder.ico" file. The "autorun.inf" contains directives for what
actions will be available upon insertion of the CD. The "folder.ico" 
file will change the appearance of the icon in the autoplay menu when 
the CD is inserted into the target computer.

-- Linux --
To build an autoplay CD for Linux, you will only need a single file
staged: the "examples/linux/autorun.sh" file. This file should contain
the commands you wish to run on the target computer upon insertion of
the CD. If you would like to deploy additional files with the "autorun.sh"
file, you will need to modify the code accordingly, as this is not supported
as an automated feature at this time.

-- Usage --
To use playmate, you will need to provide at minimum a source folder 
and a destination folder. The source folder is the location of your
payload/autorun files. The destination folder is the location where 
you wish the ISO to be stored after creation.

You can also supply various flags when building (Windows only):
-a, --arch			This specifies the architecture of the target system (x86/x64).
-d, --dll-name		This is the name of the dll to be compiled. This should match 
						the name of your staged C++ (DLL) payload.
-e, --exe-name		This is the name of the exe (launcher) to be compiled. This
						should match the name of your staged C++ launcher payload.
-f, --format		This is the payload format you would like to compile. If you
						would like to build a standalone dll, choose "dll". If
						building a dll with an exe launcher, choose "exe".
`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		switch platform {
		case "windows":
			if formatFlag != "dll" && formatFlag != "exe" {
				log.Fatalf("[-] Unable to compile with format \"%s\". Expected \"dll\" or \"exe\".", formatFlag)
			}

			if architectureFlag == "" {
				log.Fatalf("[-] An architecture must be specified (x86/x64).")
			}

			if architectureFlag != "x86" && architectureFlag != "x64" {
				log.Fatalf("[-] Unrecognized architecture \"%s\" specified. Try x86 or x64.", architectureFlag)
			}

			if dllNameFlag == "dwmapi-winexec" {
				log.Printf("[!] Default dll name (%s) is being used. This is your warning that this is not very OPSEC friendly.", dllNameFlag)
			}

			if formatFlag == "exe" && exeNameFlag == "LaunchU3" {
				log.Printf("[!] Default exe name (%s) is being used. This is your warning that this is not very OPSEC friendly.", exeNameFlag)
			}

			src := args[0]
			dst := args[1]
			if src[len(src)-1:] != "/" && src[len(src)-1:] != "\\" {
				src = src + "/"
			}
			if dst[len(dst)-1:] != "/" && dst[len(dst)-1:] != "\\" {
				dst = dst + "/"
			}
			root = "C:\\"
			compile.CompileFromSource(architectureFlag, root, dllNameFlag, exeNameFlag, formatFlag, src, dst)
			stage.StageFilesWindows(src, dst)
			stage.CreateISO(dst)
		case "linux":
			src := args[0]
			dst := args[1]
			if src[len(src)-1:] != "/" && src[len(src)-1:] != "\\" {
				src = src + "/"
			}
			log.Println(src)
			if dst[len(dst)-1:] != "/" && dst[len(dst)-1:] != "\\" {
				dst = dst + "/"
			}
			log.Println(dst)
			root = "/"
			stage.StageFilesLinux(src, dst)
			stage.CreateISO(dst)
		default:
			log.Fatalf("[-] Unable to run on %s. Use a Windows or Linux system.", platform)
		}
	},
}

func init() {
	buildCmd.Flags().StringVarP(&architectureFlag, "arch", "a", "", "the architecture of the system")
	buildCmd.Flags().StringVarP(&dllNameFlag, "dll-name", "d", "dwmapi-winexec", "the name of the dll to be compiled.")
	buildCmd.Flags().StringVarP(&exeNameFlag, "exe-name", "e", "LaunchU3", "the name of the exe to be compiled.")
	buildCmd.Flags().StringVarP(&formatFlag, "format", "f", "", `the payload output format (standalone dll [-f dll] 
	or dll launched via executable [-f exe])`)
	rootCmd.AddCommand(buildCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
