package service
import(
	// "fmt"
	"strings"

	db "longchain.com/memoriae/invitation/db/sqlite"
	"longchain.com/memoriae/invitation/email"
	"longchain.com/memoriae/invitation/util"
	"longchain.com/memoriae/invitation/log"
)
var (
	subscribe = db.NewSubscribe()
	dict = db.NewDict()
)
// 申请订阅
func ApplySubscribe(name,emailStr,role string) bool {
	res := subscribe.SelectByEmailAndRole(db.Subscribe{
		Email:emailStr,
		Role:role,
	})
	// 已有此邮箱则修改为待确认
	if res.Id != "" {
		subscribe.UpdateByEmailAndRole(db.Subscribe{
			Email:emailStr,
			Role:role,
			Status:db.Enum.Subscribe.Unconfirmed,
		})
	}else{
		subscribe.Insert(db.Subscribe{util.GetId(),name,emailStr,role,db.Enum.Subscribe.Unconfirmed})
	}

	emailTitle := dict.Get("emailTitle")
	content := dict.Get("emailContent")
	emailUrl := dict.Get("emailUrl")
	content = strings.ReplaceAll(content,"${name}",role)
	content = strings.ReplaceAll(content,"${href}",emailUrl+"?email="+emailStr+"&role="+role)
	
	return send(emailStr,emailTitle,content)
}
// 确认订阅
func ConfirmSubscribe(email,role string) bool {
	subscribe.UpdateByEmailAndRole(db.Subscribe{
		Email:email,
		Role:role,
		Status:db.Enum.Subscribe.Subscribed,
	})
	content := ""
	switch{
		case role == db.Enum.Role.User:
			content = dict.Get("emailUserInfo")
		case role == db.Enum.Role.Keeper:
			content = dict.Get("emailKeeperInfo")
		case role == db.Enum.Role.Provider:
			content = dict.Get("emailProviderInfo")
	}
	if content != "" {
		send(email,"文档",content)
	}
	return true
}

func send(emailStr , emailTitle, content string) bool {
	name := dict.Get("emailName")
	acc := dict.Get("emailAcc")
	pwd := dict.Get("emailPwd")
	e := email.New(name,acc,pwd)
	if err := e.Send(emailStr,emailTitle,content); err != nil{
		log.Error(err)
		return false
	}
	return true
}