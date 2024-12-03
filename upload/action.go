package upload

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/storage"
	"github.com/ghf-go/fleetness/upload/model"
	"gorm.io/gorm"
)

type getTokenActionParam struct {
	FileKey  string `json:"key"`
	FileName string `json:"file_name"`
}

// 获取上传凭证
func getTokenAction(c *core.GContent) {
	p := &getTokenActionParam{}
	if e := c.BindJson(p); e != nil || p.FileKey == "" || len(p.FileKey) != 32 || p.FileName == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	row := &model.UploadFile{}
	getDB(c).First(row, "file_key=?", p.FileKey)
	if row.ID > 0 { //文件已经上传
		getDB(c).Model(row).Where(row.ID).Update("upload_times", gorm.Expr("upload_times+1"))
		c.SuccessJson(map[string]any{
			"is_exists": true,
			"url":       row.Url,
		})
		return
	}

	ret := storage.GetStorage(c).BuildToken(p.FileKey, p.FileName)
	ret["is_exists"] = false
	c.SuccessJson(ret)
}

// 上传成功
func uploadSuccessAction(c *core.GContent) {
	p := &model.UploadFile{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	p.CreateIP = c.GetIP()
	p.UpdateIP = c.GetIP()
	p.UploadTimes = 1
	getDB(c).Save(p)
	c.SuccessJson("success")
}

// 上传文件
func uploadFileAction(c *core.GContent) {}
