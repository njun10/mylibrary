package lcommon

import (
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"github.com/njun10/mylibrary/common/dao"
	"github.com/njun10/mylibrary/common/model"
)

var CateSrv = new(cateCommonSrv)

type cateCommonSrv struct{}

// 通过ID查询品类信息
func (s *cateCommonSrv) GetChildsByid(id int) ([]*model.Category, error) {
	if id == 1 {
		c, e := dao.Category.Fields("id,name,upid,weight").FindAll("upid=0 and status=1")
		return c, e
	}
	if id > 100 && id < 999 {
		c, e := dao.Category.Fields("id,name,upid,weight").FindAll("upid=? and status=1", id)
		return c, e
	}
	if id > 1000 && id < 9999 {
		c, e := dao.Product.Fields("id,name,upid,weight").FindAll("cid2=? and upid=0 and status=1", id)
		return p_conv(c), e
	}
	if id > 10000 && id < 99999 {
		c, e := dao.Product.Fields("id,name,upid,weight").FindAll("upid=? and upid>0 and status=1", id)
		return p_conv(c), e
	}
	return nil, gerror.New("id error")
}

// 包内函数调用数据库的查询封装
func (s *cateCommonSrv) BatchAllCate() []*model.Category {
	if p, e := dao.Category.Order("weight desc").FindAll(fmt.Sprintf("status=1")); e != nil {
		return nil
	} else {
		return p
	}
}

func p_conv(arr []*model.Product) []*model.Category {
	var res = make([]*model.Category, len(arr))
	for k, m := range arr {
		var r = new(model.Category)
		_ = gconv.Struct(m, r)
		res[k] = r
	}
	return res
}
