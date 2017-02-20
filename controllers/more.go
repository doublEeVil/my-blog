package controllers

type MoreController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

func (this *MoreController) Get() {
	this.TplName = "terms.html"
}
