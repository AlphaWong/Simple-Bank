package utils

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

var CrossCustomerSendServiceCharge float64
var PaymentApprovalURI string

func InitConfigSetting() {
	InitCrossCustomerSendServiceCharge()
	InitPaymentApprovalURI()
}

func InitCrossCustomerSendServiceCharge() {
	CrossCustomerSendServiceCharge, _ = beego.AppConfig.Float("crosscustomerservicecharge")
}

func InitPaymentApprovalURI() {
	PaymentApprovalURI = beego.AppConfig.String("paymentapprovalURI")
}

func SendHttpError(w http.ResponseWriter, message, footPrint string, httpStatusCode int) {
	http.Error(w, fmt.Sprintf("message: %v, footPrint: %v", message, footPrint), http.StatusNotFound)
}
