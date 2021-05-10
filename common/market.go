package lcommon

import (
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/njun10/mylibrary/common/dao"
	"github.com/njun10/mylibrary/common/model"
)

var MarketSrv = new(marketCommonSrv)

type marketCommonSrv struct{}

// 保内函数调用数据库的查询封装
func (s *marketCommonSrv) getInfoByid(id int) *model.MarketInfo {
	if p, e := dao.MarketInfo.FindOne("id=? and status=1", id); e != nil {
		return nil
	} else {
		return p
	}

}

// 通过ID查询市场详细信息
func (s *marketCommonSrv) GetInfoByid(id int) (*model.MarketInfo, error) {
	if res := s.getInfoByid(id); res == nil {
		return nil, gerror.New("no info")
	} else {
		return res, nil
	}
}

// 校验ID是否真实在数据库存在
func (s *marketCommonSrv) CheckidExit(id int, t string) error {
	// 查询数据
	if p := s.getInfoByid(id); p == nil {
		return gerror.New("id not exit")
	}
	return nil
}

// 包内函数调用数据库的查询封装
func (s *marketCommonSrv) BatchByids(ids []int) []*model.MarketInfo {
	if len(ids) == 0 {
		return nil
	}
	strids := gconv.SliceStr(ids)
	if p, e := dao.MarketInfo.FindAll(fmt.Sprintf("id in (%s) and status=1", gstr.Implode(",",strids))); e != nil {
		return nil
	} else {
		return p
	}
}
