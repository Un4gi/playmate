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
