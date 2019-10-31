package email
import (
	"testing"
)

func TestCheck(t *testing.T) {
	email := New("test","bqmqdlf@163.com","密码")
	if err := email.Send("mayichichong@163.com","标题","内容"); err != nil{
		t.Log(err)
	}
}