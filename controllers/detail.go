package controllers

import (
	"github.com/astaxie/beego"
	"my-blog/services"
	"fmt"
)

type DetailController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

func (this *DetailController) Get() {
	id := this.GetString("id")
	beego.Info("id is ", id)
	blog := services.GetBlog(id)
	this.Data["UUID"] = blog.UUID
	this.Data["Title"] = blog.Title
	this.Data["Author"] = blog.Author
	this.Data["PostTime"] = blog.PostTime
	this.Data["Content"] = blog.Content
	this.Data["Conment"] = blog.Comment
	fmt.Println(len(blog.Comment))
	this.TplName = "single.html"
}