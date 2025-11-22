package chrome

import (
	"os"
	"log"
	"strings"
	"path/filepath"
)

type ProfileOpts struct {



}

func MakeProfile() (string, func(), error) {

	// remove profiles
	removeProfileDirs("./tmp")

	// Ensure tmp folder exists
	os.MkdirAll("./tmp", 0755)

	// Create a dedicated temp profile folder
	profileDir, err := os.MkdirTemp("./tmp", "chrome-profile-*")
	if err != nil {
		return "", func(){}, err
	}

	// write prefs 
	writePreferences(profileDir)

	// make defer fn
	deferFn := func() {
		os.RemoveAll(profileDir)
	}

	return profileDir, deferFn, nil

}

func writePreferences(profileDir string) {
	defaultDir := filepath.Join(profileDir, "Default")
	os.MkdirAll(defaultDir, 0755)

	prefs := `{
	}`

	os.WriteFile(filepath.Join(defaultDir, "Preferences"), []byte(prefs), 0644)
}


func removeProfileDirs(base string) error {
	entries, err := os.ReadDir(base)
	if err != nil {
		// If the directory doesn't exist, nothing to delete
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, e := range entries {
		if e.IsDir() && strings.HasPrefix(e.Name(), "chrome-profile-") {
			fullPath := filepath.Join(base, e.Name())
			log.Println("Deleting:", fullPath)
			if err := os.RemoveAll(fullPath); err != nil {
				return err
			}
		}
	}

	return nil
}