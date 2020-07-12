package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const templatePath = "./templates"

func main() {
	s, err := generateReactFcTemplateStr(templatePath+"/fc.tsx", "Test")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(s)
}

func generateReactFcTemplateStr(templateFilePath string, name string) (string, error) {
	f, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	result := strings.Replace(string(f), "App", name, -1)
	return result, nil
}
