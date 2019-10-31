package util
import(
	"github.com/rs/xid"
)

// 获取唯一ID
func GetId() string{
	id := xid.New()
	return id.String()
}