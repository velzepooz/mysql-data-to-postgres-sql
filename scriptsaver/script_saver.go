package scriptsaver

import (
	"errors"
	"log"
	"os"
)

func SaveScriptToFile(targetDir string, fileName string, script string) {
	createDir(targetDir)
	createFile(targetDir, fileName, script)
}

func createDir(dirName string) {
	if _, err := os.Stat(dirName); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func createFile(targetDir string, tableName string, script string) {
	if err := os.WriteFile(targetDir+"/"+tableName+".sql", []byte(script), 0666); err != nil {
		log.Fatal(err)
	}
}
