package helper

import "path/filepath"

func GetRootDirPath() (string, error) {
	return filepath.Abs("../..")
}
