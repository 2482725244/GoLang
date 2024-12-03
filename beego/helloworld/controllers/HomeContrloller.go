package controllers

import beego "github.com/beego/beego/v2/server/web"

type HomeContrloller struct {
	beego.Controller
}

func (c *HomeContrloller) Get() {
	c.Data["name"] = "3221"
}
