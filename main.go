package main

import (
	"change-files-name/config"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	PREFIX      = "."
	SuffixSplit = "|"
)

var (
	configPath    string
	configContent config.Content
)

func init() {
	flag.StringVar(&configPath, "c", config.DefaultConfig, "配置文件的绝对路径")
	flag.Parse()
}

func main() {
	config.Init()
	configContent = config.Parse(configPath)
	if reg, err := regexp.Compile(configContent.Expr); err == nil {
		configContent.Path = strings.TrimSpace(configContent.Path)
		fis, err := os.ReadDir(configContent.Path)
		if err != nil {
			fmt.Println(err)
		}
		for _, fi := range fis {
			fn := fi.Name()
			if !ignoreHide(fn) && hasFileSuffix(fn) {
				nfn := strings.ToUpper(reg.FindString(fn)) + path.Ext(fn)
				if !strings.EqualFold(configContent.Prefix, "") && !strings.EqualFold(configContent.Prefix, config.FileNamePrefix) {
					nfn = fmt.Sprintf("%s_%s\n", configContent.Prefix, nfn)
				}
				if configContent.Scan {
					fmt.Println(nfn)
					fmt.Println(filepath.Join(configContent.Path, nfn))
				} else {
					fna := filepath.Join(configContent.Path, fi.Name())
					err := os.Rename(fna, filepath.Join(configContent.Path, nfn))
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	} else {
		fmt.Println(err)
	}
}

func ignoreHide(filename string) bool {
	if strings.HasPrefix(filename, PREFIX) {
		return configContent.Ignore
	} else {
		return false
	}
}

func hasFileSuffix(path string) bool {
	if strings.EqualFold(configContent.Suffix, "*") {
		return true
	}
	suffixes := strings.Split(configContent.Suffix, SuffixSplit)
	for _, s := range suffixes {
		if strings.HasSuffix(path, s) {
			return true
		}
	}
	return false
}
