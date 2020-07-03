package statik

import (
	"github.com/gentwolf-shen/gohelper/logger"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	instance http.FileSystem
)

func initInstance() {
	if instance == nil {
		var err error
		instance, err = fs.New()
		if err != nil {
			panic(err.Error())
		}
	}
}

func Read(filename string) []byte {
	initInstance()

	file, err := instance.Open(filename)
	if err != nil {
		logger.Error("open file error: " + filename)
		panic(err)
	}
	defer file.Close()

	b, err1 := ioutil.ReadAll(file)
	if err1 != nil {
		logger.Error(err)
	}

	return b
}

func ReadDir(filename string) []os.FileInfo {
	initInstance()
	
	file, err := instance.Open(filename)
	if err != nil {
		logger.Error("open file error: " + filename)
		panic(err)
	}
	defer file.Close()

	files, err := file.Readdir(-1)
	if err != nil {
		logger.Error(err)
		return nil
	}

	return files
}
