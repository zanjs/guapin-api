package controllers

import (
	"fmt"
	"log"
	"mugg/guapin/app/conf"
	"mugg/guapin/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type (
	// UploadController is
	UploadController struct {
		BaseController
	}
)

// NewUpload is
func NewUpload() *UploadController {
	return &UploadController{}
}

// Create is
func (s UploadController) Create(c *gin.Context) {
	// Upload := &model.Upload{}

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	fileNameAll := utils.Upload(file.Filename)

	if err := c.SaveUploadedFile(file, fileNameAll); err != nil {
		s.ErrorJSON(c, "upload file err:"+err.Error())
		return
	}

	s.SuccessJSONData(c, fileNameAll)
}

// FileInfo is
type FileInfo struct {
	Path string `json:"path"`
}

func substring(source string, start int, end int) string {
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

// Delete is
func (s UploadController) Delete(c *gin.Context) {
	fileInfo := &FileInfo{}

	if err := c.BindJSON(fileInfo); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	if fileInfo.Path == "" {
		s.ErrorJSON(c, "file path?")
		return
	}

	fileUpFix := conf.Config.File.Path

	fmt.Println(len(fileUpFix))

	xxLoad := substring(fileInfo.Path, 0, len(fileUpFix))

	fmt.Println(xxLoad)

	if xxLoad != fileUpFix {
		s.ErrorJSON(c, "xxload error")
		return
	}

	err := os.Remove(fileInfo.Path) //删除文件
	if err != nil {
		//如果删除失败则输出 file remove Error!
		s.ErrorJSON(c, "file remove Error!"+err.Error())
		return
	}
	//如果删除成功则输出 file remove OK!
	fmt.Print("file remove OK!")

	s.SuccessJSONData(c, fileInfo)
}
