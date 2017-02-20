package controllers

type ContactController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

func (this *ContactController) Get() {
	this.TplName = "contact.html"
}