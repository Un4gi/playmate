package stage

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kdomanski/iso9660"
)

func StageFilesLinux(src, dst string) {
	fmt.Printf("[*] Staging files for ISO creation...\r\n")
	os.Link(src+"autorun.sh", dst+"autorun.sh")
	fmt.Printf("[+] Staging complete. Files staged to %s.\r\n", dst)
}

func StageFilesWindows(src, dst string) {
	fmt.Printf("[*] Staging files for ISO creation...\r\n")
	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		if strings.Contains(file.Name(), ".exe") || strings.Contains(file.Name(), ".dll") {
			fname := fmt.Sprint(file.Name())
			os.Rename(fname, dst+fname)
		} else if strings.Contains(file.Name(), ".obj") || strings.Contains(file.Name(), ".exp") || strings.Contains(file.Name(), ".lib") {
			fname := fmt.Sprint(file.Name())
			os.Remove(fname)
		}
	}
	os.Link(src+"autorun.inf", dst+"autorun.inf")
	os.Link(src+"folder.ico", dst+"folder.ico")
	fmt.Printf("[+] Staging complete. Files staged to %s.\r\n", dst)
}

func CreateISO(dst string) {
	writer, err := iso9660.NewWriter()
	if err != nil {
		log.Fatalf("failed to create writer: %s", err)
	}
	defer writer.Cleanup()

	items, _ := ioutil.ReadDir(dst)
	for _, item := range items {
		if !item.IsDir() {
			fname := fmt.Sprint(item.Name())
			f, err := os.Open(dst + fname)
			if err != nil {
				log.Fatalf("failed to open file: %s", err)
			}
			defer f.Close()

			err = writer.AddFile(f, fname)
			if err != nil {
				log.Fatalf("failed to add file: %s", err)
			}

			outputFile, err := os.OpenFile(dst+"playmate.iso", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
			if err != nil {
				log.Fatalf("failed to create the ISO file: %s", err)
			}

			err = writer.WriteTo(outputFile, "PlayVol")
			if err != nil {
				log.Fatalf("failed to write ISO image: %s", err)
			}

			err = outputFile.Close()
			if err != nil {
				log.Fatalf("failed to close output file: %s", err)
			}
		}
	}
}
