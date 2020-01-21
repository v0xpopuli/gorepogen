package generator

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

type entityInfo struct {
	Name            string
	Package         string
	FullPackagePath string
}

type Walker struct {
	projectDir    string
	entityName    string
	entityPattern string
	excludedDirs  []string
}

func NewWalker(projectDir, entityName string) *Walker {
	return &Walker{
		projectDir:    projectDir,
		entityName:    entityName,
		entityPattern: "type %s struct",
		excludedDirs:  []string{".git", ".idea", ".vscode"},
	}
}

// Walk searching for entity by given entity name
// from directory where program was ran
func (w Walker) Walk(root string) (*entityInfo, error) {

	var entityInfo entityInfo
	err := filepath.Walk(root, w.search(&entityInfo))
	if err != nil {
		return nil, err
	}
	if entityInfo.Package == "" {
		return nil, errors.Errorf("can't find given entity: %s", w.entityName)
	}

	return &entityInfo, nil
}

func (w Walker) search(entityInfo *entityInfo) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if w.canSearch(path, info) {
			words, _ := w.scanWords(path)
			for index, word := range words {
				if w.isEntity(words, index, word) {
					entityInfo.Name = w.entityName
					entityInfo.Package = words[1]
					entityInfo.FullPackagePath = w.resolveFullPackageName(path)
					return nil
				}
			}
		}
		return nil
	}
}

func (w Walker) canSearch(path string, info os.FileInfo) bool {
	return !info.IsDir() && filepath.Ext(info.Name()) == ".go" && !w.isDirExcluded(path)
}

func (w Walker) isEntity(words []string, index int, word string) bool {
	if index >= 1 && index < len(words)-1 {
		signature := []string{words[index-1], word, words[index+1]}
		return fmt.Sprintf(entityPattern, w.entityName) == strings.Join(signature, " ")
	}
	return false
}

func (w Walker) scanWords(path string) ([]string, error) {

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

func (w Walker) resolveFullPackageName(path string) string {
	dir := filepath.Dir(path)
	fullPackageName := dir[strings.Index(dir, w.projectDir):]
	return strings.Replace(fullPackageName, "\\", "/", -1)
}

func (w Walker) isDirExcluded(path string) bool {
	for _, e := range excludedDirs {
		if strings.Contains(path, e) {
			return true
		}
	}
	return false
}
