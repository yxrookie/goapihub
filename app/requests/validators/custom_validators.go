// Package validators 存放自定义规则和验证器
package validators

import "goapihub/pkg/captcha"

// ValidateCaptcha: use custom rules to validate the picture captcha
func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}