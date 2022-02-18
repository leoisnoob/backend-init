package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

type UploadFileToLocalOsgReq struct {
	BucketName string                         `json:"bucket_name,omitempty"`
	ObjectInfo UploadFileToLocalOsgObjectInfo `json:"object_info,omitempty"`
	Blob       string                         `json:"blob,omitempty"`
}

type UploadFileToLocalOsgObjectInfo struct {
	Key      string `json:"key,omitempty"`
	Metadata struct {
		Size         int    `json:"size,omitempty"`
		Etag         string `json:"etag,omitempty"`
		ContentType  string `json:"content_type,omitempty"`
		CreationTime string `json:"creation_time,omitempty"`
		Attrs        struct {
		} `json:"attrs,omitempty"`
	} `json:"metadata,omitempty"`
}

type UploadFileToLocalOsgObjectResp struct {
	Error      string                         `json:"error"`
	Code       int                            `json:"code"`
	ObjectInfo UploadFileToLocalOsgObjectInfo `json:"object_info"`
}

// 把一张图片存入 osg, putObject
func uploadFileToLocalOsg(client *http.Client, url string, key string, content io.Reader) (fileKey string, err error) {
	data := UploadFileToLocalOsgReq{
		BucketName: "hitl_data",
	}
	data.ObjectInfo.Key = key

	dataBytes, err := ioutil.ReadAll(content)
	if err != nil {
		logger.Error("读取上传的文件错误", zap.Error(err))
		return
	}
	data.ObjectInfo.Metadata.Size = len(dataBytes)
	data.Blob = base64.StdEncoding.EncodeToString(dataBytes)
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Error("marshal failed", zap.Error(err))
		return
	}
	req, err := http.NewRequest(http.MethodPost, url+"/put_object", bytes.NewReader(jsonData))
	if err != nil {
		logger.Error("new requeset failed", zap.Error(err))
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("do requeset failed ", zap.Error(err))
		return
	}
	cont, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("read resp failed", zap.Error(err))
		return
	}
	resp.Body.Close()

	var v UploadFileToLocalOsgObjectResp
	if err = json.Unmarshal(cont, &v); err != nil {
		logger.Error("unmarshal resp failed", zap.Error(err))
		return
	}
	if v.Error != "" {
		logger.Error("upload failed", zap.Error(err))
		return "", errors.New("存储错误，上传失败")
	}
	fileKey = v.ObjectInfo.Key
	return
}
