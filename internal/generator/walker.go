package generator

import (
	"fmt"
	"io/ioutil"
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

func (w Walker) collectGoFiles(root string) (map[string]string, error) {

	goFiles := make(map[string]string)
	err := filepath.Walk(root, w.visit(goFiles))
	if err != nil {
		return nil, err
	}
	return goFiles, nil
}

func (w Walker) visit(goFiles map[string]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if w.canSearch(path, info) {
			content, _ := ioutil.ReadFile(path)
			goFiles[path] = string(content)
		}
		return nil
	}
}

func (w Walker) searchEntity(goFiles map[string]string) (*entityInfo, error) {
	for path, content := range goFiles {
		if w.isEntity(content) {
			return &entityInfo{
				Name:            w.entityName,
				Package:         w.resolvePackageName(content),
				FullPackagePath: w.resolveFullPackageName(path),
			}, nil
		}
	}
	return nil, errors.Errorf("can't find given entity: %s", w.entityName)
}

func (w Walker) canSearch(path string, info os.FileInfo) bool {
	// TODO: refactor needed
	return !info.IsDir() && filepath.Ext(info.Name()) == ".go" && !strings.Contains(path, "_test.go") && !w.isDirExcluded(path)
}

func (w Walker) isEntity(content string) bool {
	// TODO: improve determining of entity
	return strings.Contains(content, fmt.Sprintf(entityPattern, w.entityName))
}

func (w Walker) resolvePackageName(content string) string {
	trimmed := strings.TrimSuffix(content, "\n")
	splitted := strings.Split(trimmed, "\n")[0]
	return strings.Split(splitted, " ")[1]
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
