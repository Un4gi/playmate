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
package compile

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var cl string
var vcvars string = "vcvarsall.bat"

func CompileFromSource(arch, root, d, e, f, src, dst string) {
	pwd, _ := os.Getwd()
	d = d + ".cpp"
	e = e + ".cpp"

	fileWalk(root)
	log.Printf("[*] Setting up environment variables...\n")
	dir, file := filepath.Split(vcvars)

	log.Println("[*] Compiling binary...")

	switch f {
	case "exe":
		cmd := exec.Command("cmd.exe", "/K", file, arch, "&&", "cd", pwd, "&&", cl, "/EHsc", "/LD", src+d, "&&", cl, "/EHsc", src+e)
		cmd.Dir = dir
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("[+] Compiled.")
		}
	case "dll":
		cmd := exec.Command("cmd.exe", "/K", file, arch, "&&", "cd", pwd, "&&", cl, "/EHsc", "/LD", src+d)
		cmd.Dir = dir
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("[+] Compiled.")
		}
	default:
		fmt.Println("Format must be <dll> or <exe>.")
		os.Exit(1)
	}
}

func fileWalk(root string) {
	fmt.Printf("[*] Searching for %s in %s\n", vcvars, root)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}

		if d.Name() == vcvars {
			vcvars = path
			fmt.Printf("[!] vcvarsall.bat found: %s\r\n\r\n", vcvars)
			return io.EOF
		}

		return nil
	})

	if err == io.EOF {
		err = nil
	}

	fmt.Printf("[*] Searching for cl.exe in %s\n", root)
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}

		if d.Name() == "cl.exe" {
			cl = path
			fmt.Printf("[!] Compiler found: %s\r\n\r\n", cl)
			return io.EOF
		}

		return nil
	})

	if err == io.EOF {
		err = nil
	}
}
