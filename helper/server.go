package helper

import (
	"log"
	"os"
	"runtime"

	"github.com/gohugoio/hugo/commands"
)

//ResumeServer will start up the hugo server
func ResumeServer(args []string) {
	ToBuild()
	CreateDirIfNotExist("content")
	CreateDirIfNotExist("themes")
	CreateDirIfNotExist("data")
	CreateConfigFile()
	if _, err := os.Stat("../resume.json"); os.IsExist(err) {
		if _, err := os.Stat("data/resume.json"); os.IsNotExist(err) {
			err = os.Link("../resume.json", "data/resume.json")
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	resp := commands.Execute(args)

	if resp.Err != nil {
		if resp.IsUserError() {
			resp.Cmd.Println("")
			resp.Cmd.Println(resp.Cmd.UsageString())
		}
		os.Exit(-1)
	}

}

//CreateDirIfNotExist creates a new directory if it does not already exists
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//CreateConfigFile creates a config file that is watched by hugo
func CreateConfigFile() {
	// detect if file exists
	var _, err = os.Stat("config.toml")

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create("config.toml")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

}

// ToBuild moves
func ToBuild() {
	CreateDirIfNotExist("build")

	err := os.Chdir("build/")
	if err != nil {
		log.Fatal(err)
	}
}
