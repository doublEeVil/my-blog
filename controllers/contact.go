package controllers

import (
	"my-blog/models"
	"my-blog/services"
)

type ContactController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

func (this *ContactController) Get() {
	this.TplName = "contact.html"
}

func (this *ContactController) Post()  {
	name := this.GetString("name")
	email := this.GetString("email")
	tel := this.GetString("tel")
	msg := this.GetString("msg")
	contactInfo := models.ContactMe{name, email, tel, msg}
	services.AddContactMsg(contactInfo)
	this.Redirect("index.html" , 302)
}