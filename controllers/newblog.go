package controllers

import (
	"github.com/astaxie/beego"
	"my-blog/models"
	"my-blog/utils"
	"time"
	"my-blog/services"
)

type NewBlogController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

func (this *NewBlogController) Get() {
	this.TplName = "newblog.html"
}

func (this *NewBlogController) Post()  {
	author := "admin"
	postTime := time.Now()
	title := this.GetString("title")
	tag := this.GetString("tag")
	content := this.GetString("content")
	beego.Debug(title,tag, content)
	id := utils.UUIDgen()
	blog := models.Blog{id, author, postTime, title, tag, content, nil}
	go services.AddNewBlog(blog)
	this.Redirect("single.html?id=" + id, 302)
}