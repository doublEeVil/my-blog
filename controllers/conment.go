package controllers

import (
	"my-blog/utils"
	"my-blog/models"
	"my-blog/services"
	"time"
)

type ConmentController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

func (this *ConmentController) Post() {
	blogid := this.GetString("id")
	commentid := utils.UUIDgen()
	author := this.GetString("author")
	mail := this.GetString("mail")
	tel := this.GetString("tel")
	content := this.GetString("content")
	postTime := time.Now()
	conment := models.BlogComment{commentid, author, mail, tel, content, postTime}
	services.AddComment(blogid, conment)
	this.Redirect("/single.html?id=" + blogid, 302)
}
