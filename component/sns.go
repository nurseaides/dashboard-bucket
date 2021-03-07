package component

import (
	"encoding/json"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711" //引入sms
)

//发送短信
func SentSns(mobile string, snsContent string) (err error) {
	fmt.Println(config)
	credential := common.NewCredential(config.SecretID, config.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	cpf.SignMethod = "HmacSHA1"
	client, _ := sms.NewClient(credential, config.Region, cpf)
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppid = common.StringPtr(config.SdkAppID)
	request.Sign = common.StringPtr(config.SmsSign)
	request.TemplateID = common.StringPtr(config.CaptchaSmsTemplateID)
	request.TemplateParamSet = common.StringPtrs([]string{snsContent})
	request.PhoneNumberSet = common.StringPtrs([]string{mobile})

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err2 := client.SendSms(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err2)
		return err2
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		fmt.Println("send sms error", err2)
		return err2
	}

	// 打印返回的json字符串
	b, _ := json.Marshal(response.Response)
	fmt.Println("send sms response:", string(b))
	// fmt.Println("send sms content :", mobile, snsContent)
	return err
}
