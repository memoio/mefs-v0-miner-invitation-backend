package email
import (
	"strings"
	"crypto/tls"
	
  "gopkg.in/gomail.v2"
)

// 协议
type treaty struct{
	// 地址
	addr string
	// 端口
	port int
}
type Mail struct{
  // 名称
  name string
  // 地址
  addr string
  // 密码
	pwd string
	// smtp
	smtp treaty
}
func New(name, addr, pwd string) Mail {
	var smtp treaty
	if strings.HasSuffix(addr,"@163.com") {
		// smtp = treaty{"smtp.163.com",25,}
		smtp = treaty{"smtp.163.com",465,}
	}else if strings.HasSuffix(addr,"@qq.com") {
		smtp = treaty{"smtp.qq.com",25,}
	}else if strings.HasSuffix(addr,"@126.com") {
		// smtp = treaty{"smtp.126.com",25,}
		smtp = treaty{"smtp.126.com",465,}
	}else if strings.HasSuffix(addr,"@memolabs.io") {
		smtp = treaty{"smtp.exmail.qq.com",465,}
	}
	return Mail{
		name:name,
		addr:addr,
		pwd:pwd,
		smtp:smtp,
	}
}
func (this Mail) Send(to, title, body string) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", this.addr, this.name)
	m.SetHeader("To", to)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", body)
	// 附件
	// m.Embed("a.png")
	d := gomail.NewDialer(this.smtp.addr, this.smtp.port, this.addr, this.pwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}