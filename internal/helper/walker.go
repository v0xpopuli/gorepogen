package helper

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

var (
	entityPattern = "type %s struct"
	excludedDirs  = []string{".git", ".idea", ".vscode"}
)

type EntityInfo struct {
	Name            string
	Package         string
	FullPackagePath string
}

func Search(whereToSearch, entityName string) (*EntityInfo, error) {

	var entityInfo EntityInfo
	err := filepath.Walk(
		whereToSearch,
		search(filepath.Base(whereToSearch), entityName, &entityInfo),
	)
	if err != nil {
		return nil, err
	}
	if entityInfo.Package == "" {
		return nil, errors.Errorf("can't find given entity: %s", entityName)
	}

	return &entityInfo, nil
}

func search(projectDir, entityName string, entityInfo *EntityInfo) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if canSearch(path, info) {
			words, _ := scanWords(path)
			for index, word := range words {
				if isEntity(words, index, word, entityName) {
					entityInfo.Name = entityName
					entityInfo.Package = words[1]
					entityInfo.FullPackagePath = resolveFullPackageName(path, projectDir)
					return nil
				}
			}
		}
		return nil
	}
}

func canSearch(path string, info os.FileInfo) bool {
	return !info.IsDir() && filepath.Ext(info.Name()) == ".go" && !isDirExcluded(path)
}

func isEntity(words []string, index int, word, entityName string) bool {
	if index >= 1 && index < len(words)-1 {
		signature := []string{words[index-1], word, words[index+1]}
		return fmt.Sprintf(entityPattern, entityName) == strings.Join(signature, " ")
	}
	return false
}

func scanWords(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, nil
}

func resolveFullPackageName(path string, projectDir string) string {
	dir := filepath.Dir(path)
	fullPackageName := dir[strings.Index(dir, projectDir):]
	return strings.Replace(fullPackageName, "\\", "/", -1)
}

func isDirExcluded(path string) bool {
	for _, e := range excludedDirs {
		if strings.Contains(path, e) {
			return true
		}
	}
	return false
}
