package component

type BucketConfig struct {
	Type       string `json:"bucket_type"`
	BucketUrl  string `json:"bucket_url"`
	BucketName string `json:"bucket_name"`
	SecretID   string `json:"secret_id"`
	SecretKey  string `json:"secret_key"`

	Region               string `json:"region"`
	SmsSign              string `json:"sms_sign"`
	SdkAppID             string `json:"sdk_app_id"`
	CaptchaSmsTemplateID string `json:"captcha_sms_template_id"`
}
