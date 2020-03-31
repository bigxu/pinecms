package taglibs

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"reflect"
)

func Channel(args jet.Arguments) reflect.Value {
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Error("Channel Failed", err)
		}
	}()
	row := getInt(args.Get(2))
	if row <= 0 {
		row = 10
	}
	catid := getInt(args.Get(0))
	typeInfo := args.Get(1).String()
	sess := pine.Make("*xorm.Engine").(*xorm.Engine).Table(&tables.Category{})
	sess.Limit(row)
	switch typeInfo {
	case "son":
		sess.Where("parentid = ?", catid)
	case "self":
		s := &tables.Category{}
		exist,_ := pine.Make("*xorm.Engine").(*xorm.Engine).Table(s).ID(catid).Cols("parentid").Get(s)
		if exist {
			sess.Where("parentid = ?", catid)
		} else {
			return reflect.ValueOf([]tables.Category{})
		}
	}
	var data = []tables.Category{}
	sess.Find(&data)
	m := models.NewCategoryModel()
	for k, v := range data {
		if v.Type == 0 {
			data[k].Url = fmt.Sprintf("/%s/", m.GetUrlPrefix(v.Catid))
		} else if v.Type == 1 {
			data[k].Url = fmt.Sprintf("/%s/", m.GetUrlPrefix(v.Catid))
		}
	}
	return reflect.ValueOf(data)
}