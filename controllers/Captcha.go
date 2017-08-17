package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"strings"
	"strconv"
)

const HTML_TPL string = `
<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Captcha by Golang</title>
</head>

<body>
<form method="post">
    <p><img src="/captcha/{{.CaptchaId}}.png" /></p>
    <p><input name="captcha" placeholder="请输入验证码" type="text" /></p>
    <input name="captcha_id" type="hidden" value="{{.CaptchaId}}" />
    <input type="submit" />
</form>
</body>
</html>`

// Main 控制器
type CaptchaController struct {
	beego.Controller
}

func (this *CaptchaController) Get() {
	captchaId := captcha.NewLen(4) //验证码长度为6
	html := strings.Replace(HTML_TPL, "{{.CaptchaId}}", captchaId, -1)
	this.Ctx.WriteString(html)
}

func (this *CaptchaController) Post() {
	id, value := this.GetString("captcha_id"), this.GetString("captcha")
	b := captcha.VerifyString(id, value) //验证码校验
	this.Ctx.WriteString(strconv.FormatBool(b))
}
