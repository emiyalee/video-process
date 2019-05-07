package models

import (
	"errors"
	"io"
	"os"
)

var (
	//ErrCanNotFind ...
	ErrCanNotFind = errors.New("can not find file")

	//ErrExistFileWithSameSHA ...
	ErrExistFileWithSameSHA = errors.New("file with the same sha")

	//ErrNotExist ...
	ErrNotExist = errors.New("file not exist")

	//ErrTODO ...
	ErrTODO = errors.New("not implement")
)

//FileInfo ...
type FileInfo struct {
	FileID   string
	FileName string
	FilePath string
	//SHA1     string
}

//FileInfos ...
type FileInfos []*FileInfo

//FileStorage ...
type FileStorage interface {
	CreateFile(filename string, r io.Reader, enableSameSHA bool) (string, error)
	GetFile(id string) (*os.File, error)
	ListFile(id string) (*FileInfo, error)
	ListAllFiles() (FileInfos, error)
	DeleteFile(id string) error
	DeleteAllFile() error
}
