package controllers

import (
	"fmt"
	"log"
	"mugg/guapin/app/conf"
	"mugg/guapin/app/service"
	"mugg/guapin/model"
	"mugg/guapin/utils"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

type (
	// ExcelController is
	ExcelController struct {
		BaseController
		StorageService service.Storage
		StorageModel   model.Storage
	}
)

// NewExcel is
func NewExcel() *ExcelController {
	return &ExcelController{}
}

// Create is
func (s ExcelController) Create(c *gin.Context) {
	// Excel := &model.Excel{}

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	fileStorage := utils.Storage{}.Upload(file.Filename)

	if err := c.SaveUploadedFile(file, fileStorage.AllName); err != nil {
		s.ErrorJSON(c, "upload file err:"+err.Error())
		return
	}
	// confFile := conf.Config.File

	// storage := &s.StorageModel
	// storage.Name = file.Filename
	// storage.Type = fileStorage.Type
	// storage.URL = utils.Substring(fileStorage.URL, len(confFile.Host), len(fileStorage.URL))
	// storage.UID = suuid.New().String()
	// go func() {
	// 	err := s.StorageService.Create(storage)
	// 	if err != nil {
	// 		fmt.Println("file storage create Error!" + err.Error())
	// 	}
	// }()

	xlsx, err := excelize.OpenFile(fileStorage.AllName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	// cell := xlsx.GetCellValue("Sheet1", "B2")
	// fmt.Println(cell)
	// Get all the rows in the Sheet1.
	// rows := xlsx.GetRows("销售单报表")
	var arrarmo = []model.ExcelMarket{}
	rows := xlsx.GetRows("Sheet1")
	for indexY, row := range rows {
		fmt.Println(indexY)
		var excelMarket = model.ExcelMarket{}
		for index2, colCell := range row {
			fmt.Println(index2)
			fmt.Print(colCell, "\t")
			switch index2 {
			case 0:
				excelMarket.ExternalNumber = colCell
			case 1:
				excelMarket.ShopName = colCell
			case 2:
				price, _ := strconv.ParseFloat(colCell, 64)
				excelMarket.PriceGoods = price
			default:
				fmt.Printf("Default")
			}
		}
		fmt.Println(row)
		arrarmo = append(arrarmo, excelMarket)

	}

	s.SuccessJSONData(c, arrarmo)
}

// ConfigInfo is
func (s ExcelController) ConfigInfo(c *gin.Context) {
	fileInfo := conf.Config.File

	s.SuccessJSONData(c, fileInfo)
}

// Delete is
func (s ExcelController) Delete(c *gin.Context) {
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
