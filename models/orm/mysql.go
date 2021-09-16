package orm

import (
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "")

	orm.RegisterModel(new(IPFS_saved_file))

	o = orm.NewOrm() //new一个Orm，默认使用名为default的数据库
	//o.Using("default") // 你可以使用Using函数指定其他数据库
}

/*
func readAllData(table interface{}, field string, value, dstStruct interface{}) bool {

	qs := o.QueryTable(table).Filter(field, value)

	if qs.Exist() {
		_, err := qs.All(dstStruct)
		if err != nil {
			log.Println(err)
			return false
		}
		//log.Println(dstStruct)
	} else {
		log.Println("qs is not Exist")
		return false
	}

	return true
}
*/

func readAllOrderBy(table interface{}, field string, dstStruct interface{}) bool {

	qs := o.QueryTable(table).OrderBy(field)

	if qs.Exist() {
		_, err := qs.All(dstStruct)
		if err != nil {
			log.Println(err)
			return false
		}
		//log.Println(dstStruct)
	} else {
		log.Println("qs is not Exist")
		return false
	}

	return true
}

func saveSingleData(data interface{}) {
	_, err := o.InsertOrUpdate(data)
	if err != nil {
		log.Println(err)
	}
}

func updateSingleData(data interface{}, cols []string) error {
	_, err := o.Update(data, cols...)
	if err != nil {
		log.Println(err)
	}
	return err
}
