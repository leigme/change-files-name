package config

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

const (
	DefaultConfig  = "cfn.json"
	FileNamePrefix = "文件名称前缀"
)

type Content struct {
	Path   string `json:"path"`
	Scan   bool   `json:"scan"`
	Expr   string `json:"expr,omitempty"`
	Prefix string `json:"prefix,omitempty"`
	Suffix string `json:"suffix,omitempty"`
	Ignore bool   `json:"ignore,omitempty"`
}

func crete() Content {
	mode := Content{
		Path:   "需要批量更换文件名的父目录的绝对路径",
		Expr:   "[S|s]{1}[0-9]{2}[e|E]{1}[0-9]{2}",
		Prefix: FileNamePrefix,
		Suffix: "mp4",
		Ignore: true,
		Scan:   false,
	}
	return mode
}

func Init() {
	if _, err := os.Stat(DefaultConfig); err != nil && os.IsNotExist(err) {
		if cf, err := os.Create(DefaultConfig); err == nil {
			defer cf.Close()
			mode := crete()
			if jsonStr, err := json.Marshal(mode); err == nil {
				bcf := bufio.NewWriter(cf)
				bcf.Write(jsonStr)
				bcf.Flush()
			}
		}
	}
}

func Parse(path string) Content {
	m := crete()
	f, err := os.Open(path)
	if err != nil {
		return m
	}
	defer f.Close()
	if data, err := io.ReadAll(f); err == nil {
		json.Unmarshal(data, &m)
	}
	return m
}
