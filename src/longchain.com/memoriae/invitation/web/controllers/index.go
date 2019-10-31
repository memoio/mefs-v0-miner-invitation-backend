package controllers
import(
	"github.com/kataras/iris"

	rm "longchain.com/memoriae/invitation/web/resultModel"
	"longchain.com/memoriae/invitation/web/service"
)

func New(app *iris.Application) {
	// 申请订阅
	// email 邮箱
	app.Get("/applySubscribe/{name}/{email}/{role}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		email := ctx.Params().Get("email")
		role := ctx.Params().Get("role")
		if service.ApplySubscribe(name,email,role) {
			ctx.JSON(rm.Ok("成功"))
		}else {
			ctx.JSON(rm.Error("失败"))
		}
	})
	// 确认订阅
	app.Get("/confirmSubscribe/{email}/{role}", func(ctx iris.Context) {
		email := ctx.Params().Get("email")
		role := ctx.Params().Get("role")
		if service.ConfirmSubscribe(email,role) {
			ctx.JSON(rm.Ok("成功"))
		}else {
			ctx.JSON(rm.Error("失败"))
		}
	})
}