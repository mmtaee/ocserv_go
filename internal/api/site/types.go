package site

type Data struct {
	CaptchaSiteKey   string  `json:"captcha_site_key"  binding:"omitempty"`
	CaptchaSecretKey string  `json:"captcha_secret_key"  binding:"omitempty"`
	DefaultTraffic   float64 `json:"default_traffic" binding:"required"`
}

type DataUpdate struct {
	CaptchaSiteKey   string  `json:"captcha_site_key"  binding:"omitempty"`
	CaptchaSecretKey string  `json:"captcha_secret_key"  binding:"omitempty"`
	DefaultTraffic   float64 `json:"default_traffic"  binding:"omitempty"`
}
