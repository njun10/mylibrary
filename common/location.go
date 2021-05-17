package lcommon

import (
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"github.com/njun10/mylibrary/common/dao"
	"github.com/njun10/mylibrary/common/model"
)

var LocSrv = new(locCommonSrv)

type locCommonSrv struct{}

// 通过ID查询区域关联信息
func (s *locCommonSrv) GetChildsByid(id int) ([]*model.Location, error) {
	if id == 1 {
		c, e := dao.Location.FindAll("level=1 and status=1")
		return c, e
	}
	if id > 100 && id < 999 {
		c, e := dao.Location.FindAll("upid=? and level in (2,3) and status=1", id)
		return c, e
	}
	if id > 1000 && id < 9999 {
		c, e := dao.Location.FindAll("upid=? and level=3 and status=1", id)
		return c, e
	}
	if id > 10000 && id < 99999 {
		c, e := dao.Location.FindAll("upid=? and level=4 and status=1", id)
		return c, e
	}
	if id > 100000 && id < 999999 {
		c, e := dao.Location.FindAll("upid=? and level=5 and status=1", id)
		return c, e
	}
	return nil, gerror.New("id error")
}

// 通过ID查询位置信息
func (s *locCommonSrv) GetInfoByid(id int) (*model.Location, error) {
	if e := s.Checkid(id, 0); e != nil {
		return nil, gerror.New("check id err")
	}

	if res := s.getInfoByid(id); res == nil {
		return nil, gerror.New("no info")
	} else {
		return res, nil
	}
}

// ID 格式校验函数
func (s *locCommonSrv) Checkid(id int, l int) error {
	var rule string

	// 校验ID的格式
	switch l {
	case 1:
		rule = "integer|min:100|max:999"
		break
	case 2:
		rule = "integer|min:1000|max:9999"
		break
	case 3:
		rule = "integer|min:10000|max:99999"
		break
	case 4:
		rule = "integer|min:100000|max:999999"
		break
	case 5:
		rule = "integer|min:1000000|max:9999999"
		break
	default:
		rule = "integer|min:100|max:9999999"
	}
	if e := gvalid.Check(id, rule, nil); e != nil {
		return e
	}

	return nil
}

// 根据ID获取全路径
func (s *locCommonSrv) PathAllByid(id int) ([]*model.Location, error) {
	if e := s.Checkid(id, 0); e != nil {
		return nil, gerror.New("check id err")
	}

	if res := s.getInfoByid(id); res == nil {
		return nil, gerror.New("no info")
	} else {
		path := gstr.Explode(".", res.Path)
		ids := gconv.SliceInt(path)
		if ids == nil {
			return nil, gerror.New("path all get ids err")
		}
		if all := s.BatchByids(ids); all==nil {
			return nil, gerror.New("batch info err")
		}else{
			return all, nil
		}
	}
}

// 包内函数调用数据库的查询封装
func (s *locCommonSrv) getInfoByid(id int) *model.Location {
	if p, e := dao.Location.FindOne("id=? and status=1", id); e != nil {
		return nil
	} else {
		return p
	}

}

// 包内函数调用数据库的查询封装
func (s *locCommonSrv) BatchByids(ids []int) []*model.Location {
	if len(ids) == 0 {
		return nil
	}
	strids := gconv.SliceStr(ids)
	if p, e := dao.Location.FindAll(fmt.Sprintf("id in (%s) and status=1", gstr.Implode(",",strids))); e != nil {
		return nil
	} else {
		return p
	}
}

// 包内函数调用数据库的查询封装
func (s *locCommonSrv) BatchByNames(names []string) []*model.Location {
	if len(names) == 0 {
		return nil
	}
	if p, e := dao.Location.FindAll(fmt.Sprintf("name in ('%s') and status=1", gstr.Implode("','",names))); e != nil {
		return nil
	} else {
		return p
	}
}

// 包内函数调用数据库的查询封装
func (s *locCommonSrv) BatchAll() []*model.Location {
	if p, e := dao.Location.FindAll(fmt.Sprintf("status=1")); e != nil {
		return nil
	} else {
		return p
	}
}
