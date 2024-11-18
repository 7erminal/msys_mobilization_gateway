package functions

import (
	beego "github.com/beego/beego/v2/server/web"
)

func HostMapping(clientid string) (host string) {
	tHost := ""

	if clientid == "889999" {
		tHost, _ = beego.AppConfig.String("mizpaApiBaseUrl")
	}
	if clientid == "889901" {
		tHost, _ = beego.AppConfig.String("crystalCOPApiBaseUrl")
	}

	return tHost
}
