package helper

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//GetTheme downloads the zip form of the theme from github
func GetTheme(themeName string) {
	ToBuild()
	fileURL := "https://github.com/resumic/theme-" + themeName + "/archive/master.zip"
	err := os.Chdir("themes")
	if err != nil {
		log.Fatal(err)
	}
	err = DownloadFile("theme-"+themeName+".zip", fileURL)
	if err != nil {
		log.Fatal(err)
	}
	err = Unzip("theme-"+themeName+".zip", "")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("theme downloaded successfully")
		log.Println("now you can use the downloaded theme using: resumic serve --theme", themeName)
	}
	DeleteFile("theme-" + themeName + ".zip")

	os.Rename("theme-"+themeName+"-master", themeName)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Unzip unzips the downloaded theme to the destination folder
func Unzip(src string, dest string) error {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return err
			}

		}
	}
	return nil
}

//DeleteFile deletes the file in the path specified
func DeleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}

//ListTheme lists the resumic themes available for download
func ListTheme() {
	fmt.Print("1-minimal\n2-simple-blue")
}
