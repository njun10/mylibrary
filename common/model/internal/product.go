// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// Product is the golang structure for table product.
type Product struct {
	Id          uint   `orm:"id"           json:"id"`          //
	Name        string `orm:"name"         json:"name"`        // 类目名称
	Cid1        int    `orm:"cid1"         json:"cid1"`        // 一级分类id
	Cid2        int    `orm:"cid2"         json:"cid2"`        // 二级分类id
	Weight      int    `orm:"weight"       json:"weight"`      //
	PinyinS     string `orm:"pinyin_s"     json:"pinyinS"`     // 拼音简写
	PinyinA     string `orm:"pinyin_a"     json:"pinyinA"`     // 拼音全拼
	Status      int    `orm:"status"       json:"status"`      //
	Upid        int    `orm:"upid"         json:"upid"`        // 上级类目ID
	CreatedTime int    `orm:"created_time" json:"createdTime"` //
	UpdatedTime int    `orm:"updated_time" json:"updatedTime"` //
}
