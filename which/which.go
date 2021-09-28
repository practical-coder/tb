package which

import (
	"fmt"
	"os"
	"path/filepath"
)

func Path(envKey string) string {
	return os.Getenv(envKey)
}

func PathList() []string {
	return filepath.SplitList(Path("PATH"))
}

func ExecutableExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err == nil {
		mode := fileInfo.Mode()
		return mode.IsRegular() && mode&0111 != 0
	}
	return false
}

func Find(names ...string) {
	for _, directory := range PathList() {
		for _, name := range names {
			filePath := filepath.Join(directory, name)
			if ExecutableExists(filePath) {
				fmt.Println(filePath)
			}
		}
	}
}
