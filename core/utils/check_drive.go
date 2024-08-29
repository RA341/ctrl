package utils

import "os"

func IsExternalDriveAccessible(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false // The path does not exist
	}
	return err == nil // If err is nil, the path is accessible
}
