package main

import (
	"log"

	"github.com/emiyalee/video-process/api-gateway/controllers"
	"github.com/emiyalee/video-process/api-gateway/models"
	_ "github.com/emiyalee/video-process/api-gateway/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	fileStorage, err := models.NewFileStorageDB("root", "123456", "./filestorage")
	if err != nil {
		log.Fatalln("error : ", err)
	}

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/file",
			beego.NSInclude(
				&controllers.FileController{FileStorage: fileStorage},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Run()
}
