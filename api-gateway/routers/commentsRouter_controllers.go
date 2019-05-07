package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"] = append(beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"] = append(beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"] = append(beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"] = append(beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"] = append(beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"],
        beego.ControllerComments{
            Method: "ListAll",
            Router: `/info/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"] = append(beego.GlobalControllerRouter["github.com/emiyalee/video-process/api-gateway/controllers:FileController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/info/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
