/*
文件操作相关工具
*/
package goutils

import (
	"errors"
	"io"
	"os"
	"path"
)

// 文件大小获取
//获取文件大小的接口
type Size interface {
	Size() int64
}

func File_GetSize(r io.Reader) (error, int64) {
	sizeInterface, ok := r.(Size)
	if !ok {
		return errors.New("Unrecongnizable file"), 0
	}
	return nil, sizeInterface.Size()
}

// 文件信息获取
//获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

func File_GetInfo(r io.Reader) (error, os.FileInfo) {
	statInterface, ok := r.(Stat)
	if !ok {
		return errors.New("Unrecongnizable file"), nil
	}

	info, err := statInterface.Stat()
	if err != nil {
		return err, nil
	}

	return nil, info
}

//检查路径是否存在，若不存在自动创建文件路径
func Path_CheckAndCreate(strPath string) error {
	if Path_IsExist(strPath) {
		return nil
	}
	return os.MkdirAll(path.Dir(strPath), os.ModePerm)
}

// 检查路径是否存在
func Path_IsExist(strPath string) bool {
	_, err := os.Stat(strPath)
	return err == nil || os.IsExist(err)
}

// 路径创建

// 路径深度创建
