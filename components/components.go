package components

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var dashReg = regexp.MustCompile(`-[a-zA-Z]`)

// name 是用短线连接的
// TODO: 具体路径还没实现
func GenerateReactFc(templateFilePath string, outDir string, name string) error {
	var err error
	f, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		fmt.Print(err)
		return err
	}
	result := strings.Replace(string(f), "App", getComponentName(name), -1)
	result = strings.Replace(result, "style.css", getCssFileName(name), -1)
	pathPrefix := outDir + "/" + getFolderName(name)
	err = os.Mkdir(pathPrefix, os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(
		outDir+"/"+getFolderName(name)+"/"+getFileName(name),
		[]byte(result),
		0644,
	)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(
		outDir+"/"+getFolderName(name)+"/"+getCssFileName(name),
		[]byte(""),
		0644,
	)
	return nil
}

func getFileName(name string) string {
	return "index.tsx"
}

func getCssFileName(name string) string {
	return "style.css"
}

func getFolderName(name string) string {
	return dashToPascal(name)
}

func getComponentName(name string) string {
	return dashToPascal(name)
}

func dashToCamel(name string) string {
	return dashReg.ReplaceAllStringFunc(name, func(s string) string {
		letter := s[1:2]
		return strings.ToUpper(letter)
	})
}

func dashToPascal(name string) string {
	firstLetter := name[0:1]
	rest := name[1:]
	return strings.ToUpper(firstLetter) + dashToCamel(rest)
}
