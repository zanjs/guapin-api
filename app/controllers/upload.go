package controllers

import (
	"fmt"
	"log"
	"mugg/guapin/app/conf"
	"mugg/guapin/app/service"
	"mugg/guapin/model"
	"mugg/guapin/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/houndgo/suuid"
)

type (
	// UploadController is
	UploadController struct {
		BaseController
		StorageService service.Storage
		StorageModel   model.Storage
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

	fileStorage := utils.Storage{}.Upload(file.Filename)

	if err := c.SaveUploadedFile(file, fileStorage.AllName); err != nil {
		s.ErrorJSON(c, "upload file err:"+err.Error())
		return
	}
	confFile := conf.Config.File

	storage := &s.StorageModel
	storage.Name = file.Filename
	storage.Type = fileStorage.Type
	storage.URL = utils.Substring(fileStorage.URL, len(confFile.Host), len(fileStorage.URL))
	storage.UID = suuid.New().String()
	go func() {
		err := s.StorageService.Create(storage)
		if err != nil {
			fmt.Println("file storage create Error!" + err.Error())
		}
	}()

	s.SuccessJSONData(c, fileStorage.URL)
}

// ConfigInfo is
func (s UploadController) ConfigInfo(c *gin.Context) {
	fileInfo := conf.Config.File

	s.SuccessJSONData(c, fileInfo)
}

// FileInfo is
type FileInfo struct {
	Path string `json:"path"`
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

	xxLoad, err := utils.FilePatFixRemove(fileInfo.Path, conf.Config.File.Path)
	if err != nil {
		s.ErrorJSON(c, "file remove Error!"+err.Error())
		return
	}

	fmt.Println(xxLoad)

	err = os.Remove(fileInfo.Path) //删除文件
	if err != nil {
		//如果删除失败则输出 file remove Error!
		s.ErrorJSON(c, "file remove Error!"+err.Error())
		return
	}
	//如果删除成功则输出 file remove OK!
	fmt.Print("file remove OK!")

	s.SuccessJSONData(c, fileInfo)
}
