package main
import(

	"github.com/henrylee2cn/pholcus/exec"
	//_ "test/lib"

)
//func main() {
//	//qrcode.WriteFile("http://www.flysnow.org/",qrcode.Medium,256,"./blog_qrcode.png")
//	code, err := qrcode.New("http://www.flysnow.org/", qrcode.Medium)
//	if err != nil {
//		log.Fatal(err)
//	}else {
//		code.BackgroundColor = color.RGBA{
//			R: 50,
//			G: 205,
//			B: 50,
//			A: 255,
//		}
//		code.ForegroundColor = color.White
//		code.WriteFile(256,"./blog_qrcode.png")
//	}
//}

//func main()  {
//	reader := bufio.NewReader(os.Stdin)
//	for  {
//		c, err := reader.ReadString('\n')
//		if err == nil{
//			c = strings.Replace(c,"吗","",-1)
//			c = strings.Replace(c,"? ","!",-1)
//			c = strings.Replace(c,"?","!",-1)
//			fmt.Println(c)
//		}
//	}
//}
func main() {

	exec.DefaultRun("web")

}