package controllers

import (
	"liteblog/syserrors"
	"strings"
)

type NoteController struct {
	BaseController
}

func (ctx *NoteController) NestPrepare() {
	ctx.MustLogin()
	if ctx.User.Role != 0 {
		ctx.Abort500(syserrors.NewError("您没有权限修改文章", nil))
	}
}

// @router /new [get]
func (ctx *NoteController) NewPage() {
	ctx.Data["key"] = ctx.UUID()
	if strings.EqualFold(ctx.User.Editor, "markdown") {
		ctx.TplName = "note_new2.html"
		return
	}
	ctx.TplName = "note_new.html"
}

// @router /edit/:key [get]
func (ctx *NoteController) EditPage() {
	key := ctx.Ctx.Input.Param(":key")
	note, err := ctx.Dao.QueryNoteByKeyAndUserId(key, int(ctx.User.ID))
	if err != nil {
		ctx.Abort500(syserrors.NewError("文章不存在", err))
	}
	ctx.Data["note"] = note
	ctx.Data["key"] = key
	if strings.EqualFold(note.Editor, "markdown") {
		ctx.TplName = "note_new2.html"
		return
	}
	ctx.TplName = "note_new.html"
}

// @router /del/:key [post]
func (ctx *NoteController) Del() {
	key := ctx.Ctx.Input.Param(":key")
	if err := ctx.Dao.DelNoteByKey(key, int(ctx.User.ID)); err != nil {
		ctx.Abort500(syserrors.NewError("删除失败", err))
	}
	ctx.JSONOk("删除成功！", "/")
}

//
//// @router /save/:key [post]
//func (ctx *NoteController) Save() {
//	key := ctx.Ctx.Input.Param(":key")
//	editor := ctx.GetString("editor","default");
//	title := ctx.GetMustString("title", "标题不能为空！")
//	content := ctx.GetMustString("content", "内容不能为空！")
//	files := ctx.GetString("files", "")
//	summary, _ := getSummary(content)
//	note, err := ctx.Dao.QueryNoteByKeyAndUserId(key, int(ctx.User.ID))
//	var n models.Note
//	if err != nil {
//		if err != gorm.ErrRecordNotFound {
//			ctx.Abort500(syserrors.NewError("保存失败！", err))
//		}
//		n = models.Note{
//			Key:     key,
//			Summary: summary,
//			Title:   title,
//			Files:   files,
//			Content: content,
//			UserID:  int(ctx.User.ID),
//		}
//	} else {
//		n = note
//		n.Title = title
//		n.Content = content
//		n.Summary = summary
//		n.Files = files
//		n.UpdatedAt = time.Now()
//	}
//	n.Editor = editor
//	if strings.EqualFold(editor,"markdown"){
//		n.Source = ctx.GetMustString("source","内容不能为空！")
//	}
//
//	if err := ctx.Dao.SaveNote(&n); err != nil {
//		ctx.Abort500(syserrors.NewError("保存失败！", err))
//	}
//	ctx.JSONOk("成功", "/details/"+key)
//}

//func getSummary(content string) (string, error) {
//	var buf bytes.Buffer
//	buf.Write([]byte(content))
//	doc, err := goquery.NewDocumentFromReader(&buf)
//	if err != nil {
//		return "", err
//	}
//	str := doc.Find("body").Text()
//	strRune := []rune(str)
//	if len(strRune) > 400 {
//		strRune = strRune[:400]
//	}
//	return string(strRune) + "...", nil
//}
