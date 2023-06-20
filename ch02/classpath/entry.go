package classpath

import (
	"os"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) {
	// if strings.Contains(path, pathListSeparator) {
	// 	return

	// }
}
