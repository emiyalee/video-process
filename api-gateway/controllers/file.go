package controllers

import (
	"path"

	"github.com/emiyalee/video-process/api-gateway/models"

	"github.com/astaxie/beego"
)

// FileController operations for File
type FileController struct {
	beego.Controller
	FileStorage models.FileStorage
}

// URLMapping ...
func (c *FileController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ListAll", c.ListAll)
	c.Mapping("List", c.List)
}

type uploadResponse struct {
	ErrorCode string `json:"error_code"`
	FileID    string `json:"file_id"`
}

// Post ...
// @Title Create
// @Description create File
// @Param	body		body 	models.File	true		"body for File content"
// @Success 201 {object} models.File
// @Failure 403 body is empty
// @router / [post]
func (c *FileController) Post() {
	f, h, err := c.GetFile("file")
	if err != nil {
		c.Data["json"] = &uploadResponse{err.Error(), ""}
		c.ServeJSON()
		return
	}
	defer f.Close()

	fileID, err := c.FileStorage.CreateFile(h.Filename, f, false)
	if err != nil {
		c.Data["json"] = &uploadResponse{err.Error(), ""}
		c.ServeJSON()
		return
	}

	c.Data["json"] = &uploadResponse{"success", fileID}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get File by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.File
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FileController) GetOne() {
	fileInfo, err := c.FileStorage.ListFile(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Data["json"] = &listResponse{err.Error(), nil}
		c.ServeJSON()
		return
	}
	c.Ctx.Output.Download(fileInfo.FilePath, fileInfo.FileName)
}

// Put ...
// @Title Put
// @Description update the File
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.File	true		"body for File content"
// @Success 200 {object} models.File
// @Failure 403 :id is not int
// @router /:id [put]
func (c *FileController) Put() {

}

//FileInfo ...
type FileInfo struct {
	FileID   string `json:"file_id"`
	FileName string `json:"file_name"`
	URL      string `json:"url"`
}

type listAllResponse struct {
	ErrorCode string      `json:"error_code"`
	FileInfos []*FileInfo `json:"file_infos"`
}

// ListAll ...
// @router /info/ [get]
func (c *FileController) ListAll() {
	fileInfos, err := c.FileStorage.ListAllFiles()
	if err != nil {
		c.Data["json"] = &listAllResponse{err.Error(), nil}
		c.ServeJSON()
		return
	}

	rsp := &listAllResponse{ErrorCode: "success"}
	rsp.FileInfos = make([]*FileInfo, 0)

	for index := 0; index < len(fileInfos); index++ {
		var fileInfo FileInfo
		fileInfo.FileID = fileInfos[index].FileID
		fileInfo.FileName = fileInfos[index].FileName
		fileInfo.URL = path.Join("/v1/file", fileInfo.FileID)
		rsp.FileInfos = append(rsp.FileInfos, &fileInfo)
	}

	c.Data["json"] = rsp
	c.ServeJSON()
}

type listResponse struct {
	ErrorCode string    `json:"error_code"`
	FileInfo  *FileInfo `json:"file_info"`
}

// List ...
// @router /info/:id [get]
func (c *FileController) List() {
	fileInfo, err := c.FileStorage.ListFile(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Data["json"] = &listResponse{err.Error(), nil}
		c.ServeJSON()
		return
	}

	rsp := &listResponse{ErrorCode: "success", FileInfo: &FileInfo{}}
	rsp.FileInfo.FileID = fileInfo.FileID
	rsp.FileInfo.FileName = fileInfo.FileName
	rsp.FileInfo.URL = path.Join("/v1/file", rsp.FileInfo.FileID)

	c.Data["json"] = rsp
	c.ServeJSON()
}

type deleteResponse struct {
	ErrorCode string `json:"error_code"`
}

// Delete ...
// @Title Delete
// @Description delete the File
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FileController) Delete() {
	err := c.FileStorage.DeleteFile(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Data["json"] = &deleteResponse{err.Error()}
	} else {
		c.Data["json"] = &deleteResponse{"success"}
	}

	c.ServeJSON()
	return

}
