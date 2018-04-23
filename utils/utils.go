package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego"
)

var CrossCustomerSendServiceCharge float64
var DailyTransferLimit float64
var PaymentApprovalURI string

func InitConfigSetting() {
	InitCrossCustomerSendServiceCharge()
	InitDailyTransferLimit()
	InitPaymentApprovalURI()
}

func InitCrossCustomerSendServiceCharge() {
	CrossCustomerSendServiceCharge, _ = beego.AppConfig.Float("crosscustomerservicecharge")
}

func InitDailyTransferLimit() {
	DailyTransferLimit, _ = beego.AppConfig.Float("dailytransferlimit")
}

func InitPaymentApprovalURI() {
	PaymentApprovalURI = beego.AppConfig.String("paymentapprovalURI")
}

func SendHttpError(w http.ResponseWriter, message, footPrint string, httpStatusCode int) {
	http.Error(w, fmt.Sprintf("message: %v, footPrint: %v", message, footPrint), http.StatusNotFound)
}

func GetTodayStart() time.Time {
	return time.Now().UTC()
}

func GetTomorrowStart(today time.Time) time.Time {
	return today.AddDate(0, 0, 1).UTC()
}
