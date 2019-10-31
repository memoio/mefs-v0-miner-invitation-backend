package log
import(
	logM "log"
	"os"
)
var(
	log = logM.New(os.Stdout,"",logM.Ldate | logM.Ltime | logM.Llongfile)
)

func Debug(msg ...interface{}){
	log.Print("debug: ",msg)
}
func Info(msg ...interface{}){
	log.Print(" info: ",msg)
}
func Warn(msg ...interface{}){
	log.Print(" warn: ",msg)
}
func Error(msg ...interface{}){
	log.Print("error: ",msg)
}
func Fatal(msg ...interface{}){
	log.Fatal("fatal: ",msg)
}
func Panic(msg ...interface{}){
	log.Panic("panic: ",msg)
}
