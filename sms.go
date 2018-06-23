package Alidayu

import (
	"errors"
	"net/url"
	"strings"
	"time"
)

type SMS struct {
	URL       string
	APPKey    string
	APPSecret string
	SMSParam  SMSParam
}

// NewSMS creates a sms handler
func NewSMS(config SMSConfig) (*SMS, error) {
	if config.APPKey == "" || config.APPSecret == "" {
		return nil, errors.New("key and secret can not be empty")
	}

	smsParam := SMSParam{
		Type:           "normal",
		FreeSignName:   config.FreeSignName,
		Param:          config.Param,
		ReceiverNumber: config.PhoneNumber,
		TemplateCode:   config.TemplateCode,
		Method:         "alibaba.aliqin.fc.sms.num.send",
		APPKey:         config.APPKey,
		TimeStamp:      time.Now().Format("2006-01-02 15:04:05"),
		SignMethod:     "md5",
		Format:         "json",
		V:              "2.0",
	}
	sms := &SMS{
		URL:       "https://eco.taobao.com/router/rest",
		APPKey:    config.APPKey,
		APPSecret: config.APPSecret,
		SMSParam:  smsParam,
	}
	return sms, nil
}

// SendSMS sends sms request
func (sms *SMS) SendSMS() (int, error) {
	body, err := sms.GetURLQuery()
	if err != nil {
		return 500, err
	}
	return DoRequest("POST", sms.URL, []byte(body))
}

// GetURLQuery generates the body that the http request needs
func (sms *SMS) GetURLQuery() (string, error) {
	keys, err := extractNotNullKeys(sms.SMSParam)
	if err != nil {
		return "", err
	}

	pMap, err := generateMap(sms.SMSParam, keys)
	if err != nil {
		return "", err
	}
	sms.SMSParam.Sign, err = signTopRequest(pMap, sms.APPSecret, sms.SMSParam.SignMethod)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("sign", strings.ToUpper(sms.SMSParam.Sign))
	for k, v := range pMap {
		params.Set(k, v)
	}
	return params.Encode(), nil
}
