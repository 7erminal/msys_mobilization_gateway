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

	if clientid == "889902" {
		tHost, _ = beego.AppConfig.String("mizpaSusuApiBaseUrl")
	}

	if clientid == "889903" {
		tHost, _ = beego.AppConfig.String("secureApiBaseUrl")
	}

	if clientid == "889906" {
		tHost, _ = beego.AppConfig.String("JONApiBaseUrl")
	}

	if clientid == "889905" {
		tHost, _ = beego.AppConfig.String("StMargaretApiBaseUrl")
	}

	if clientid == "889904" {
		tHost, _ = beego.AppConfig.String("AkuapemApiBaseUrl")
	}
	if clientid == "889907" {
		tHost, _ = beego.AppConfig.String("TrueLifeApiBaseUrl")
	}
	if clientid == "889908" {
		tHost, _ = beego.AppConfig.String("TwedieCommunityCoopApiBaseUrl")
	}
	if clientid == "889909" {
		tHost, _ = beego.AppConfig.String("KwanwomaCoopApiBaseUrl")
	}
	if clientid == "889910" {
		tHost, _ = beego.AppConfig.String("FAKACoopCreditUnionApiBaseUrl")
	}
	if clientid == "889911" {
		tHost, _ = beego.AppConfig.String("ZionPraiseHospitalCoopApiBaseUrl")
	}

	return tHost
}
