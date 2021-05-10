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

var ProductSrv = new(productCommonSrv)

type productCommonSrv struct{}

// 保内函数调用数据库的查询封装
func (s *productCommonSrv) getInfoByid(id int) *model.Product {
	if p, e := dao.Product.FindOne("id=? and status=1", id); e != nil {
		return nil
	} else {
		return p
	}

}

// 通过ID查询品类信息
func (s *productCommonSrv) GetInfoByid(id int) (*model.Product, error) {
	if e := s.Checkid(id, ""); e != nil {
		return nil, gerror.New("check id err")
	}

	if res := s.getInfoByid(id); res == nil {
		return nil, gerror.New("no info")
	} else {
		return res, nil
	}
}

// 校验ID是否真实在数据库存在
func (s *productCommonSrv) CheckidExit(id int, t string) error {
	// 格式校验
	if e := s.Checkid(id, t); e != nil {
		return e
	}
	// 查询数据
	if p := s.getInfoByid(id); p == nil {
		return gerror.New("id not exit")
	}
	return nil
}

// ID 格式校验函数
func (s *productCommonSrv) Checkid(id int, t string) error {
	var rule string

	// 校验ID的格式
	switch t {
	case "product":
		rule = "integer|min:10000|max:99999"
		break
	case "breed":
		rule = "integer|min:100000|max:999999"
		break
	default:
		rule = "integer|min:10000|max:999999"
	}
	if e := gvalid.Check(id, rule, nil); e != nil {
		return e
	}

	return nil
}

// 校验productid和breedid的对应关系
func (s *productCommonSrv) CheckidRelate(pid int, bid int) (*model.Product, error) {
	// 校验pid
	if e := s.Checkid(pid, "product"); e != nil {
		return nil, e
	}
	// 校验bid
	if e := s.Checkid(bid, "breed"); e != nil {
		return nil, e
	}
	// 校验数据库的存储关系
	if p, e := dao.Product.FindOne("id=? and upid=? and status=1", bid, pid); e != nil || p == nil {
		return nil, gerror.New("check fail")
	}else {
		return p, nil
	}
}

// 包内函数调用数据库的查询封装
func (s *productCommonSrv) BatchByids(ids []int) []*model.Product {
	if len(ids) == 0 {
		return nil
	}
	strids := gconv.SliceStr(ids)
	if p, e := dao.Product.FindAll(fmt.Sprintf("id in (%s) and status=1", gstr.Implode(",",strids))); e != nil {
		return nil
	} else {
		return p
	}
}
