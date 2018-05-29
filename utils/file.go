package utils

import (
	"fmt"
	"mugg/guapin/app/conf"
	"path"
	"strings"

	"github.com/houndgo/houndgo/ifile"
	"github.com/houndgo/houndgo/itime"
	"github.com/houndgo/suuid"
)

// Upload is
func Upload(filename string) string {

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
	sfileName := suuid.New().String() + fileSuffix
	fileNameAll := todayPath + "/" + sfileName

	return fileNameAll
}
