package service

import (
	"context"
	"fmt"
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/revel/revel"
	"os"
	"strconv"
	"time"
)

type UploadsService struct{}

func (this *UploadsService) QiniuUploads(format string, filePath string) models.Res {
	accessKey := revel.Config.StringDefault("qiniu.access_key", "")
	secretKey := revel.Config.StringDefault("qiniu.secret_key", "")
	bucket := revel.Config.StringDefault("qiniu.bucket", "")
	domain := revel.Config.StringDefault("qiniu.domain", "")
	https := revel.Config.StringDefault("qiniu.https", "false")
	localSave := revel.Config.StringDefault("file.local.save", "false")

	putPolicy := storage.PutPolicy{Scope: bucket}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	qiniuCfg := storage.Config{Zone: &storage.ZoneHuadong,
		UseHTTPS: true}
	formUploader := storage.NewFormUploader(&qiniuCfg)
	ret := storage.PutRet{}

	nowDate := time.Now().Format("2006-01-02")
	key := fmt.Sprintf("%s/%s.%s", nowDate, utils.NewGuid(), format)
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, filePath, nil)
	if err != nil {
		return models.Res{Status: 500, Info: err.Error()}
	}

	//删除掉本地文件
	if localSaveBool, localSaveErr := strconv.ParseBool(localSave); localSaveErr == nil && !localSaveBool {
		if delErr := os.Remove(filePath); delErr != nil {
			revel.AppLog.Errorf("Delete local file failed: %s", delErr.Error())
		}
	}

	var protocol = ""
	if httpsBool, httpsBoolErr := strconv.ParseBool(https); httpsBoolErr == nil && httpsBool {
		protocol = "s"
	}
	item := fmt.Sprintf("http%s://%s/%s", protocol, domain, key)
	return models.Res{Status: 200, Item: item, Info: key}
}
