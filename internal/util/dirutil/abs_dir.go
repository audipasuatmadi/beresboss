package dirutil

import "path/filepath"

func GetAbsDirectory(directory string) (string, error) {
	var err error
	absDir := directory
	if !filepath.IsAbs(absDir) {
		absDir, err = filepath.Abs(absDir)
		return absDir, err
	}
	return absDir, nil
}
