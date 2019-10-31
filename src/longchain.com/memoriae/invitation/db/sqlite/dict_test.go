package sqlite
import(
	"testing"
)

func Test(t *testing.T) {
	res := NewDict().Get("emailAcc")
	t.Log(res)
}