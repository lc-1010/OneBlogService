package service

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

type UploadRequest struct {
	FormName     string `json:"file,omitempty"`
	FormFileType string `json:"type,omitempty"`
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File,
	fileHeader *multipart.FileHeader) (*FileInfo, error) {

	fileName := upload.GetFileNmae(fileHeader.Filename)
	if !upload.CheckAllowExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckOverMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximux file limit")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("falied to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}
	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName

	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
