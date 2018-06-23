package Alidayu

import (
	"testing"
)

func TestSMS(t *testing.T) {
	smsConfig := SMSConfig{
		APPKey:       "",
		APPSecret:    "",
		FreeSignName: "阿里大于",
		Param: map[string]string{
			"code":    "1234",
			"product": "alidayu",
		},
		PhoneNumber:  "13000000000",
		TemplateCode: "SMS_585014",
	}
	sms, err := NewSMS(smsConfig)
	if err != nil {
		t.Error(err)
	}

	statusCode, err := sms.SendSMS()
	if err != nil {
		t.Errorf("status code is %d, error is %s", statusCode, err.Error())
	}
}
