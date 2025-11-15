package files

import (
	"os"
	"time"
    "io/ioutil"
)

func Open(file string) (string, error) {

	// read file 	
    body, err := ioutil.ReadFile(file)
    if err != nil {
		return "", err
    }

	return string(body), nil

}

func Exists(file string) bool {
	_, err := Open(file)
	if err!=nil {
		return false
	}
	return true 
}

func OpenIfUpdated(file string, fileRecord map[string]time.Time) (string, bool, error) {

	// Get file information using os.Stat()
	fileInfo, err := os.Stat(file)
	if err != nil {
		return "", true, err
	}

	// set file record (defered)
	defer func(){
		fileRecord[file] = fileInfo.ModTime()
	}()
	
	// open
	content, err := Open(file)

	// file record currently empty
	if _, ok := fileRecord[file]; !ok {
		return content, true, err
	}

	// new file
	newFile := fileRecord[file].Unix() != fileInfo.ModTime().Unix()
	
	// file record not empty
	return content, newFile, err

}