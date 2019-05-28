package controllers

import (
	"lcc/blog/models"
	"lcc/blog/syserror"
)

type NoteController struct {
	BaseController
}

///note
// @router /save [post]
func (this *NoteController) Save() {
	this.MustLogin()
	nkey := this.UUID()
	title := this.GetMustString("title", "请输入标题！")
	content := this.GetMustString("content", "请输入文章内容！")

	var n models.Note
	n = models.Note{
		NKey:    nkey,
		Title:   title,
		Content: content,
		UserID:  int(this.User.ID),
		User:    this.User,
	}

	if err := models.SaveNote(&n); err != nil {
		this.Abort500(syserror.New("保存失败3", err))
	}
	this.JsonOk("保存成功")
}

///note
// @router /edit [post]
func (this *NoteController) Edit() {
	this.MustLogin()
	nkey := this.GetMustString("nkey", "nkey不能为空")
	title := this.GetMustString("title", "title不能为空")
	content := this.GetMustString("content", "content不能为空")
	//查询是否有该文章以及对应的用户
	_, err := models.UpdatedNoteByKeyAndUserId(nkey, title, content, int(this.User.ID))
	if err != nil {
		this.Abort500(syserror.New("文章不存在", err))
	}
	this.JsonOk("更新成功")
}
