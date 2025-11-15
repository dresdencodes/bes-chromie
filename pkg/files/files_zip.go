package files

import (
	"os"
	"strings"
	"io/ioutil"
    "archive/zip"
)

func Zip(zipFilename string, files []string) error {

    // Get a Buffer to Write To
    outFile, err := os.Create(zipFilename)
    if err != nil { return err }
    defer outFile.Close()

    // Create a new zip archive.
    zipWriter := zip.NewWriter(outFile)

    // Add some files to the archive.
	for _, filename := range files {
		err := addFiles(zipWriter, filename)
		if err != nil { 
			return err 
		}
	}

    // Make sure to check the error on Close.
	return zipWriter.Close()

}

func addFiles(w *zip.Writer, file string) error {

	// read file
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// strip dir
	fileSpl := strings.Split(file, "/")
	
	// Add some files to the archive.
	f, err := w.Create(fileSpl[len(fileSpl)-1])
	if err != nil {
		return err
	}

	// write
	_, err = f.Write(dat)
	if err != nil {
		return err
	}
	
	return nil
}
