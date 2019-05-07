package models

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//FileStorageDB ...
type FileStorageDB struct {
	user            string
	password        string
	fileStoragePath string
}

func generateDBConnectURL(user, password string) string {
	return fmt.Sprintf("%s:%s@/videoprocess?charset=utf8&parseTime=True&loc=Local", user, password)
}

//File ...
type File struct {
	gorm.Model
	FileID   string `gorm:"primary_key;column:file_id"`
	FilePath string `gorm:"not null;column:file_path"`
	FileName string `gorm:"not null;column:file_name"`
	//SHA1     string `gorm:"not null;unique"`
}

//New ...
func NewFileStorageDB(user, password, pathPrefix string) (FileStorage, error) {
	url := generateDBConnectURL(user, password)
	log.Println(url)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	db.AutoMigrate(&File{})
	return &FileStorageDB{user, password, pathPrefix}, nil
}

//CreateFile ...
func (f *FileStorageDB) CreateFile(filename string, r io.Reader, enableSameSHA bool) (string, error) {
	fileID := uuid.New().String()
	filepath := path.Join(f.fileStoragePath, fileID)

	file, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, r)

	if err != nil {
		return "", err
	}

	url := generateDBConnectURL(f.user, f.password)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		os.Remove(filepath)
		return "", err
	}
	defer db.Close()

	err = db.Create(&File{
		FileID:   fileID,
		FilePath: filepath,
		FileName: filename,
	}).Error

	if err != nil {
		os.Remove(filepath)
		return "", err
	}

	return fileID, nil
}

//GetFile ...
func (f *FileStorageDB) GetFile(id string) (*os.File, error) {
	return nil, ErrTODO

}

//ListFile ...
func (f *FileStorageDB) ListFile(id string) (*FileInfo, error) {
	url := generateDBConnectURL(f.user, f.password)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.Select([]string{"file_id", "file_name", "file_path"}).Where(&File{FileID: id}).Find(&File{}).Row()
	if row == nil {
		return nil, ErrNotExist
	}

	var fileID string
	var fileName string
	var filePath string
	err = row.Scan(&fileID, &fileName, &filePath)
	if err != nil {
		return nil, err
	}

	return &FileInfo{fileID, fileName, filePath}, nil
}

//ListAllFiles ...
func (f *FileStorageDB) ListAllFiles() (FileInfos, error) {
	url := generateDBConnectURL(f.user, f.password)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Select([]string{"file_id", "file_name", "file_path"}).Find(&File{}).Rows()
	if err != nil {
		return nil, err
	}

	fileInfos := make([]*FileInfo, 0)
	for rows.Next() {
		var fileID string
		var fileName string
		var filePath string
		err := rows.Scan(&fileID, &fileName, &filePath)
		if err != nil {
			return nil, err
		}
		fileInfos = append(fileInfos, &FileInfo{fileID, fileName, filePath})
	}

	return fileInfos, nil
}

//DeleteFile ...
func (f *FileStorageDB) DeleteFile(id string) error {
	url := generateDBConnectURL(f.user, f.password)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return err
	}
	defer db.Close()

	var file File
	err = db.Where(&File{FileID: id}).First(&file).Error
	if err != nil {
		return err
	}
	os.Remove(file.FilePath)

	return db.Delete(&file).Error
}

//DeleteAllFile ...
func (f *FileStorageDB) DeleteAllFile() error {
	return ErrTODO
}
