package walker

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/v0xpopuli/gorepogen/internal/param"
)

type walker struct {
	projectDir    string
	entityName    string
	entityPattern string
	excludedDirs  []string
}

func New(projectDir, entityName string) *walker {
	return &walker{
		projectDir:    projectDir,
		entityName:    entityName,
		entityPattern: "type %s struct",
		excludedDirs:  []string{".git", ".idea", ".vscode"},
	}
}

// Walk searching for entity by given entity name
// from directory where program was ran
func (w walker) Walk(root string) (*param.EntityDefinition, error) {

	goFiles, err := w.collectGoFiles(root)
	if err != nil {
		return nil, err
	}

	entityDef, err := w.searchEntity(goFiles)
	if err != nil {
		return nil, err
	}
	return entityDef, nil
}

func (w walker) collectGoFiles(root string) (map[string]string, error) {
	goFiles := make(map[string]string)
	err := filepath.Walk(root, w.visit(goFiles))
	if err != nil {
		return nil, err
	}
	return goFiles, nil
}

func (w walker) visit(goFiles map[string]string) filepath.WalkFunc {
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

func (w walker) searchEntity(goFiles map[string]string) (*param.EntityDefinition, error) {
	for path, content := range goFiles {
		if w.isEntity(content) {
			info := param.EntityInfo{
				EntityName:      w.entityName,
				EntityPackage:   w.resolvePackageName(content),
				FullPackagePath: w.resolveFullPackageName(path),
			}
			return &param.EntityDefinition{info: nil}, nil
		}
	}
	return nil, errors.Errorf("Can't find given entity: %s", w.entityName)
}

func (w walker) canSearch(path string, info os.FileInfo) bool {
	return !info.IsDir() && w.isGoFile(info.Name()) && !w.isDirExcluded(path)
}

func (w walker) isGoFile(name string) bool {
	return filepath.Ext(name) == ".go" && !strings.Contains(name, "_test.go")
}

func (w walker) isEntity(content string) bool {
	return strings.Contains(content, fmt.Sprintf(w.entityPattern, w.entityName))
}

func (w walker) resolvePackageName(content string) string {
	trimmed := strings.TrimSuffix(content, "\n")
	splitted := strings.Split(trimmed, "\n")[0]
	return strings.Split(splitted, " ")[1]
}

func (w walker) resolveFullPackageName(path string) string {
	dir := filepath.Dir(path)
	fullPackageName := dir[strings.Index(dir, w.projectDir):]
	return strings.Replace(fullPackageName, "\\", "/", -1)
}

func (w walker) isDirExcluded(path string) bool {
	for _, e := range w.excludedDirs {
		if strings.Contains(path, e) {
			return true
		}
	}
	return false
}
