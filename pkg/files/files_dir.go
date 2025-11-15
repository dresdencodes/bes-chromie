package files

import (
	"os"
	"sync"
	"errors"
	"strings"
)

var directoryRecord = map[string]struct{}{}
var directoryRecordMu = &sync.Mutex{}

func MakeDirectory(path string) error {

	// check directory record
	if _, ok := directoryRecord[path]; ok {return nil} 
	
	// split file off
	spl := strings.Split(path, "/")
	spl = spl[:len(spl)-1]
	dirPath := strings.Join(spl, "/")
	
	// eval dir path
	if dirPath == "" || dirPath == "/" || len(spl) == 0 {
		return errors.New("invalid directory path passed")
	}
 
	// os mkdir all
	err := os.MkdirAll(dirPath, 0700)
	if err==nil {directoryRecord[path] = struct{}{}}
	return err
}

func ResetDir(path string) error {

	// os remove all
	err := os.RemoveAll(path)
	if err!=nil {
		return err
	}

	// make dir
	return MakeDirectory(path)

}
