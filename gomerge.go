// The MIT License (MIT)
//
// Copyright (c) 2016 Fredy Wijaya
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var (
	templatePath string
	outputPath   string
)

func getAbsolutePath(path string) string {
	absPath, _ := filepath.Abs(path)
	return absPath
}

func createFileFromTemplate(templatePath, outputPath string, funcMap map[string]interface{}) error {
	log.Printf("Creating file: %s from template file: %s\n", getAbsolutePath(outputPath),
		getAbsolutePath(templatePath))

	templateBytes, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}
	t := template.New("template").Funcs(template.FuncMap(funcMap))
	template.Must(t.Parse(string(templateBytes)))
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	err = t.Execute(outputFile, nil)
	if err != nil {
		return err
	}
	return nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func mergeFunc(outputPath string) map[string]interface{} {
	funcMap := map[string]interface{}{}
	funcMap["Merge"] = func(path string) string {
		if fileExists(path) {
			log.Printf("Applying content of %s to file: %s\n",
				getAbsolutePath(path), getAbsolutePath(outputPath))
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			return string(bytes)
		} else {
			log.Printf("File: %s does not exist, ignoring it\n", getAbsolutePath(path))
		}
		return ""
	}
	return funcMap
}

func merge(templatePath, outputPath string) error {
	return createFileFromTemplate(templatePath, outputPath, mergeFunc(outputPath))
}

func init() {
	flag.StringVar(&templatePath, "template", "", "Template file")
	flag.StringVar(&outputPath, "output", "", "Output file")
}

func validateArgs() {
	if len(templatePath) == 0 {
		log.Fatalf("--template option is required")
	}
	if len(outputPath) == 0 {
		log.Fatalf("--output option is rquired")
	}
	if !fileExists(templatePath) {
		log.Fatalf("Template: %s does not exist\n", getAbsolutePath(templatePath))
	}
}

func main() {
	flag.Parse()
	validateArgs()
	err := merge(templatePath, outputPath)
	if err != nil {
		log.Fatal(err)
	}
}
