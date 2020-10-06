// Package filesprovider provide functional for prepare and store files for parsing.
// This package detects directories in a file list and merge there content to file list for providing.
package filesprovider

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const supportedExt = ".dbf"

// FileProvider is structure for provide files for parsing
type FileProvider struct {
	Files          []string
	maxDirDeep     uint8
	currentDirDeep uint8
}

// NewFileProvider is the constructor for FileProvider structure
func NewFileProvider(fileList []string) (*FileProvider, error) {
	fp := &FileProvider{maxDirDeep: 1}

	if err := fp.processedFiles(fileList); err != nil {
		return nil, err
	}

	return fp, nil
}

func (fp *FileProvider) processedFiles(fileList []string) error {
	for _, path := range fileList {
		if err := fp.processedFile(path); err != nil {
			return err
		}
	}

	return nil
}

func (fp *FileProvider) processedFile(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		if err := fp.processedDir(path); err != nil {
			return err
		}
	case mode.IsRegular():
		if ext := filepath.Ext(path); ext != supportedExt {
			return fmt.Errorf("extension for file %s must be %s", path, supportedExt)
		}
		fp.Files = append(fp.Files, path)
	}

	return nil
}

func (fp *FileProvider) processedDir(path string) error {
	fp.currentDirDeep++
	if fp.currentDirDeep <= fp.maxDirDeep {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		for _, f := range files {
			if err := fp.processedFile(filepath.Join(path, f.Name())); err != nil {
				return err
			}
		}
	}
	fp.currentDirDeep--

	return nil
}
