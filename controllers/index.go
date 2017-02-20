package controllers

import (
	"github.com/beego/i18n"
	"github.com/astaxie/beego"
	"my-blog/services"
)

type baseController struct {
	beego.Controller // Embed struct that has stub implementation of the interface.
	i18n.Locale      // For i18n usage when process data and render template.
}

type IndexController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

func (this *IndexController) Get() {
	blogs := services.GetAllBlogs()
	this.Data["Blogs"] = blogs
	this.TplName = "index.html"
}