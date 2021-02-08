package component

type BucketConfig struct {
	Type      string `json:"bucket_type"`
	BucketUrl string `json:"bucket_url"`
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
}
