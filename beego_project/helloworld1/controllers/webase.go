package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"helloworld/constant"
	"helloworld/utils"
)

type WeBASEController struct {
	beego.Controller
}

func (c *WeBASEController) Set() {
	value := c.Ctx.Input.Query("value")
	funcParam := []interface{}{value}
	body := utils.CommonEq("set", constant.ContractName, constant.ContractAddress, constant.Abi, funcParam)
	msg := utils.GetJsonVal(body, "message")
	var data map[string]string
	if msg == "Success" {
		data = map[string]string{
			"code": "200",
			"msg":  "ok",
		}
	} else {
		data = map[string]string{
			"code": "400",
			"msg":  "fail",
		}

	}
	fmt.Println(msg)

	c.Data["json"] = data
	c.ServeJSON()
}

func (c *WeBASEController) Get() {
	funcParam := []interface{}{}
	body := utils.CommonEq("get", constant.ContractName, constant.ContractAddress, constant.Abi, funcParam)
	fmt.Println(body)
	data := map[string]string{
		"code": "200",
		"msg":  body,
	}
	c.Data["json"] = data
	c.ServeJSON()
}
