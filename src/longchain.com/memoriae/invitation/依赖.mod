module longchain.com/memoriae/invitation

go 1.12

require (
	// web框架
	github.com/kataras/iris v11.2.2
  // 查询json数据
  // github.com/json-iterator/go v1.1.7
  // email
  "gopkg.in/gomail.v2" v2.0.0
  // 生成唯一id
  "github.com/rs/xid" v1.2.1
  // 操作sqlite数据库
  // "github.com/mattn/go-sqlite3" v1.11.0
  // orm
  "github.com/jinzhu/gorm"
  "github.com/jinzhu/gorm/dialects/sqlite"
)