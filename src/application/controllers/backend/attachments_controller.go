package backend

import (
	"html/template"

	"github.com/xiusin/pine"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
)

type AttachmentController struct {
	pine.Controller
}

func (c *AttachmentController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/attachments/list", "List")
	b.POST("/attachments/delete", "Delete")
}

func (c *AttachmentController) List(orm *xorm.Engine) {
	page, _ := c.Ctx().URLParamInt64("page")
	rows, _ := c.Ctx().URLParamInt64("rows")

	if page > 0 {
		keywords := c.Ctx().GetString("keywords")
		list, total := models.NewAttachmentsModel().GetList(keywords, page, rows)
		c.Ctx().Render().JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}

	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Datagrid("attachments_list_datagrid", "/b/attachments/list", helper.EasyuiOptions{
		"title":   models.NewMenuModel().CurrentPos(menuid),
		"toolbar": "attachments_list_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"原始名称":   {"field": "original", "width": "20", "index": "0"},
		"资源名称":   {"field": "name", "width": "20", "index": "1"},
		"图片":   {"field": "url", "width": "10", "index": "2", "formatter": "attachmentsListUrlFormatter"},
		"大小":   {"field": "size", "width": "5", "index": "3", "formatter": "attachmentsListSizeFormatter"},
		"上传时间": {"field": "upload_time", "width": "15", "index": "4"},
		"类型":   {"field": "type", "width": "10", "index": "5"},
		"操作":   {"field": "id", "formatter": "attachmentsListSizeOptFormatter", "index": "6"},
	})

	c.Ctx().Render().ViewData("dataGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/attachments_list.html")
}

func (c *AttachmentController) Delete(orm *xorm.Engine) {
	id, _ := c.Ctx().PostInt64("id")
	if id < 1 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	if models.NewAttachmentsModel().Delete(id) {
		helper.Ajax("删除附件成功", 0, c.Ctx())
	} else {
		helper.Ajax("删除附件失败", 1, c.Ctx())
	}
}
