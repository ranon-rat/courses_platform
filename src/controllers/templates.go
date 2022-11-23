package controllers

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func InitializeTemplates(path string, funcs template.FuncMap) (*template.Template, error) {
	var temp = template.New("").Funcs(funcs)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".html") {

			return err
		}

		_, err = temp.ParseFiles(path)

		return err
	})

	return temp, err
}
