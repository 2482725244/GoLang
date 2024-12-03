package constant

import (
	beego "github.com/beego/beego/v2/server/web"
)

var (
	ContractAddress string
	Abi             string
	ContractName    string
	User            string
	WeBASEUrl       string
)

func init() {
	ContractAddress = beego.AppConfig.DefaultString("webase::address", "0xf1c52579430f33a67a010667687c32b85d60f840")
	Abi = beego.AppConfig.DefaultString("webase::abi", "")
	ContractName = beego.AppConfig.DefaultString("webase::name", "")
	User = beego.AppConfig.DefaultString("webase::user", "")
	WeBASEUrl = beego.AppConfig.DefaultString("webase::url", "")
}
