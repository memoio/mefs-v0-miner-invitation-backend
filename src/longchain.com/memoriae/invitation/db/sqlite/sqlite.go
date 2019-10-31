package sqlite
import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	
	"longchain.com/memoriae/invitation/config"
)
var (
	db *gorm.DB
)
func init() {
	var err error
	db, err = gorm.Open("sqlite3", config.DbHome)
  if err != nil {
    panic("连接数据库失败")
  }
}