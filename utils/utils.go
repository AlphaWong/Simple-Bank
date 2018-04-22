package utils

import "github.com/astaxie/beego"

var CrossCustomerSendServiceCharge float64

func InitConfigSetting() {
	InitCrossCustomerSendServiceCharge()
}

func InitCrossCustomerSendServiceCharge() {
	CrossCustomerSendServiceCharge, _ = beego.AppConfig.Float("crosscustomerservicecharge")
}
