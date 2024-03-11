// Package validators 存放自定义规则和验证器
package validators

import (
	"goapihub/pkg/captcha"
	"goapihub/pkg/verifycode"
)

// ValidateCaptcha: use custom rules to validate the picture captcha
func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

// ValidatePasswordConfirm: custom rules to judge whether two password is same
func ValidatePasswordConfirm(password, password_confirm string, errs map[string][]string) map[string][]string {
	if password != password_confirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配")
	}
	return errs
}

// ValidateVerifyCode 自定义规则， 验证[手机/邮箱验证码]
func ValidateVerifyCode(key, answer string, errs map[string][]string) map[string][]string {
	if ok := verifycode.NewVerifyCode().CheckAnswer(key, answer); !ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
	}
	return errs
}