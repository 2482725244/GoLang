package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"helloworld/services"
)

type SdkController struct {
	beego.Controller
}

func (c *SdkController) Set() {
	value := c.Ctx.Input.Query("value")
	services.SetServiceImpl(value)

	c.Data["json"] = "ok"
	c.ServeJSON()
}

func (c *SdkController) Get() {
	value := services.GetServiceImpl()
	c.Data["json"] = value
	c.ServeJSON()
}
