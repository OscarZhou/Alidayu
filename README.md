# Alidayu
Alidayu communication **Go** library. 

The current version supports the `sms` function.  

*API documents:*
1. [Signature method](http://open.taobao.com/doc.htm?docId=101617&docType=1)
2. [Sms sending](http://open.taobao.com/api.htm?docId=25450&docType=2)

**Example Code: **

```
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

```

 *Similar API:* [Aliyuncs API](https://github.com/OscarZhou/Aliyuncs)