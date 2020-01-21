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

	goFiles, err := w.collectGoFiles(root)
	if err != nil {
		return nil, err
	}

	entityInfo, err := w.searchEntity(goFiles)
	if err != nil {
		return nil, err
	}
	return entityInfo, nil
}

func (w Walker) collectGoFiles(root string) ([]string, error) {
	var goFiles []string
	err := filepath.Walk(root, w.visit(&goFiles))
	if err != nil {
		return nil, err
	}
	return goFiles, nil
}

func (w Walker) visit(goFiles *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if w.canSearch(path, info) {
			*goFiles = append(*goFiles, path)
		}
		return nil
	}
}

func (w Walker) searchEntity(goFiles []string) (*entityInfo, error) {
	for _, gfPath := range goFiles {
		words := w.scanWords(gfPath)
		for index, word := range words {
			if w.isEntity(words, index, word) {
				return &entityInfo{
					Name:            w.entityName,
					Package:         words[1],
					FullPackagePath: w.resolveFullPackageName(gfPath),
				}, nil
			}
		}
	}
	return nil, errors.Errorf("can't find given entity: %s", w.entityName)
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

func (w Walker) scanWords(path string) []string {

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
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
