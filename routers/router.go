package routers

import (
	"my-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.SetStaticPath("/", "static")
	//beego.SetStaticPath("/images","images")
	//beego.SetStaticPath("/css","css")
	//beego.SetStaticPath("/js","js")
    	beego.Router("/", &controllers.IndexController{})
	beego.Router("/index.html", &controllers.IndexController{})
    	beego.Router("/single.html", &controllers.DetailController{})
    	beego.Router("/contact.html", &controllers.ContactController{})
    	beego.Router("/newblog.html", &controllers.NewBlogController{})
    	beego.Router("/terms.html", &controllers.MoreController{})
    	beego.Router("/conment", &controllers.ConmentController{})

}
