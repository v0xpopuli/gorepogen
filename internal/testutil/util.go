package testutil

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// CreateTempFile provide ability to
// create temporary files during testing
func CreateTempFile(t *testing.T, cd, content string) string {

	filePath := filepath.Join(cd, fmt.Sprintf("temp_%d.go", time.Now().Unix()))
	file, err := os.Create(filePath)
	if err != nil {
		t.Error("Failed to create the temp file: "+filePath, err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		t.Error("Failed to write to temp file", err)
	}
	return filePath
}

// DeleteTempFile provide ability to
// delete temporary files after test execution
func DeleteTempFile(t *testing.T, filePath string) {
	if err := os.RemoveAll(filePath); err != nil {
		t.Error("Failed to delete temp file:"+filePath, err)
	}
}
