package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasPrefix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasPrefix(path, ".jar") || strings.HasPrefix(path, ".JAR") ||
		strings.HasPrefix(path, ".zip") || strings.HasPrefix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
