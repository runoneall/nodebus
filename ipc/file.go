package ipc

import (
	"fmt"
	"os"
	"path/filepath"
)

func getIPCFile(ipcName string) string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	dirPath := filepath.Join(configDir, "nodebus")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0750); err != nil {
			panic(err)
		}
	}

	return filepath.Join(dirPath, fmt.Sprintf("%s.ipc", ipcName))
}

func cleanIPCFile(filePath string) {
	if _, err := os.Stat(filePath); os.IsExist(err) {
		if err := os.RemoveAll(filePath); err != nil {
			panic(err)
		}
	}
}
