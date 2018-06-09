package controllers

import (
	"bytes"
	"fmt"
	"mugg/guapin/model"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/houndgo/houndgo/ifile"
)

type (
	// LogsController is
	LogsController struct {
		BaseController
	}
)

// NewLogs is
func NewLogs() *LogsController {
	return &LogsController{}
}

// Create is
func (s LogsController) Create(c *gin.Context) {
	Logs := &model.Log{}

	if err := c.BindJSON(Logs); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if Logs.AppID == "" || Logs.Content == "" {
		s.ErrorJSON(c, "err.Error()")
		return
	}

	go func() {

		logFileDir := "./log/" + Logs.AppID + "/"
		ifile.Mkdir(logFileDir)

		t := time.Now()
		dayStr := t.Format("2006-01-02")
		dayTimeStr := t.Format("2006-01-02 15:04:05")
		dayFile := logFileDir + dayStr + ".log"
		fmt.Println(t)
		fmt.Println(dayTimeStr)

		var buffer bytes.Buffer
		buffer.WriteString("time: " + dayTimeStr)
		buffer.WriteString(" content: " + Logs.Content)

		f, err := os.OpenFile(dayFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0766)
		if err != nil {
			fmt.Println(err.Error())
		}
		f.Write([]byte(buffer.String() + "\r\n"))
		f.Close()
		// fmt.Println(dayFile)
		// if !ifile.CheckFileIsExist(dayFile) {
		// 	fmt.Print("no")
		// 	file3, error := os.Create(dayFile)
		// 	if error != nil {
		// 		fmt.Println(error)
		// 	}
		// 	fmt.Println(file3)
		// 	file3.Close()
		// }

	}()

	s.SuccessJSON(c)
}
