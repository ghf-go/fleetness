package storage

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/ghf-go/fleetness/core/conf"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
)

type Qiniu struct {
	client     *auth.Credentials
	bucket     string
	cdnHost    string
	uploadHost string
}

func newQiniu(con conf.StorageConfig) Qiniu {

	return Qiniu{
		client:     credentials.NewCredentials(con.Ak, con.Sk),
		bucket:     con.Bucket,
		cdnHost:    con.CdnHost,
		uploadHost: con.UploadHost,
	}
}
func (s Qiniu) BuildToken(fileKey, fileName string) map[string]any {
	ext := path.Ext(fileName)
	fileName = path.Base(fileName)
	fkey := fmt.Sprintf("%s/%s%s", time.Now().Format("2006/01/02"), fileKey, ext)
	putPolicy, err := uptoken.NewPutPolicyWithKey(s.bucket, fkey, time.Now().Add(1*time.Hour))
	if err != nil {
		return map[string]any{}
	}

	putPolicy.SetReturnBody(`{"key":"$(key)","url":"$(x:url)","file_size":$(fsize),"file_key":"$(x:file_key)","file_name":"$(x:file_name)"}`)
	upToken, err := uptoken.NewSigner(putPolicy, s.client).GetUpToken(context.Background())
	if err != nil {
		return map[string]any{}
	}

	return map[string]any{
		"upload_host": s.uploadHost,
		"data": map[string]any{
			"token":       upToken,
			"key":         fkey,
			"x:file_key":  fileKey,
			"x:file_name": fileName,
			"x:url":       s.cdnHost + fkey,
		},
	}
}
