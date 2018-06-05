package utils

import (
	"errors"
	"fmt"
	"mugg/guapin/app/conf"
	"path"
	"strings"

	"github.com/houndgo/houndgo/ifile"
	"github.com/houndgo/houndgo/itime"
	"github.com/houndgo/suuid"
)

// Storage is
type Storage struct {
	Name       string
	AllName    string
	URL        string
	FileSuffix string
	Type       string
}

// Upload is
func (s Storage) Upload(filename string) Storage {

	var filenameWithSuffix string
	filenameWithSuffix = path.Base(filename)

	fmt.Println("filenameWithSuffix =", filenameWithSuffix)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)
	fmt.Println("fileSuffix =", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	fmt.Println("filenameOnly =", filenameOnly)

	today := itime.Today()
	todayPath := conf.Config.File.Path + "/" + today
	checkBool := ifile.CheckFileIsExist(todayPath)
	if !checkBool {
		ifile.Mkdir(todayPath)
	}
	s.Name = suuid.New().String() + fileSuffix
	s.AllName = todayPath + "/" + s.Name
	s.Type = fileSuffix
	s.URL = conf.Config.File.Host + s.AllName
	return s
}

// Substring is
func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

// FilePatFixRemove is
func FilePatFixRemove(path string, fileUpFix string) (string, error) {

	fmt.Println(len(fileUpFix))

	xxLoad := Substring(path, 0, len(fileUpFix))

	if xxLoad != fileUpFix {
		return xxLoad, errors.New("err")
	}

	return xxLoad, nil
}
