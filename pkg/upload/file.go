package upload

import (
	"io"

	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/util"
)

// FileType for upload file type
type FileType int

const (
	// TypeImage for image type
	TypeImage FileType = iota + 1
	// TypeExcel for excel type
	TypeExcel
	// TypeTxt for txt type
	TypeTxt
)

// GetFileExt returns the file extension of the given name.
//
// name: the name of the file.
// Returns: the file extension as a string.
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath return save path
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// CheckSavePath check save path
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// GetFileNmae return md5string name
func GetFileNmae(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// CheckAllowExt checks if the given file type is allowed based on the file extension.
//
// It takes two parameters: t, which is the FileType enum that represents the type of the file,
// and name, which is the string representing the name of the file.
//
// It returns a boolean value indicating whether the file type is allowed or not.
func CheckAllowExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

// CheckOverMaxSize checks if the size of the file is over the maximum allowed size.
//
// It takes two parameters: t FileType and f multipart.File.
// It returns a boolean indicating whether the size is over the maximum allowed size.
func CheckOverMaxSize(t FileType, f multipart.File) bool {
	content, _ := io.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// CreateSavePath creates a directory at the given destination path with the specified permissions.
//
// Parameters:
// - dst: the destination path where the directory will be created.
// - perm: the permissions to be set for the created directory.
// Return type: error.
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// CheckPermission checks if the given destination has proper permission.
//
// dst: the destination to check permission for.
// bool: true if the permission is granted, false otherwise.
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// SaveFile saves a file to the specified destination.
//
// It takes a *multipart.FileHeader as the file to be saved and a string
// 'dst' as the destination path. It returns an error if any error occurs
// during the process.
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}
