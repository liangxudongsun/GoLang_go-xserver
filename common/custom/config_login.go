package custom

// ConfigLogin : 配置 login 节
type ConfigLogin struct {
	Sign1 string `default:"5UY6$f$h" desc:"用于登录验证的部分签名段"`
	Sign2 string `default:"3wokZB%q" desc:"用于登录验证的部分签名段"`
	Sign3 string `default:"%2Fi9TRf" desc:"用于登录验证的部分签名段"`
}
