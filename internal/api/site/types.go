package site

type CreateData struct {
	CaptchaSiteKey   string  `json:"captcha_site_key"  binding:"omitempty"`
	CaptchaSecretKey string  `json:"captcha_secret_key"  binding:"omitempty"`
	DefaultTraffic   float64 `json:"default_traffic" binding:"required"`
}

type UpdateData struct {
	CaptchaSiteKey   string  `json:"captcha_site_key"  binding:"omitempty"`
	CaptchaSecretKey string  `json:"captcha_secret_key"  binding:"omitempty"`
	DefaultTraffic   float64 `json:"default_traffic"  binding:"omitempty"`
}
