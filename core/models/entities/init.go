package entities

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

// set mysql master and slave service
// master as a writer
// slave as reader
var MasterEngine *xorm.Engine
var SlaveEngine *xorm.Engine

func init() {
	var err error
	MasterEngine, err = xorm.NewEngine("mysql", "root:master@tcp(dbmaster:3306)/ressystem?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	err = MasterEngine.Sync2(new(StudentInfo))
	if err != nil {
		panic(err)
	}
	err = MasterEngine.Sync2(new(AdminInfo))
	if err != nil {
		panic(err)
	}
	err = MasterEngine.Sync2(new(ApproverInfo))
	if err != nil {
		panic(err)
	}
	err = MasterEngine.Sync2(new(ClassroomInfo))
	if err != nil {
		panic(err)
	}
	err = MasterEngine.Sync2(new(ReservationInfo))
	if err != nil {
		panic(err)
	}

	SlaveEngine, err = xorm.NewEngine("mysql", "root:slave@tcp(dbslave:3306)/ressystem?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}

	if os.Getenv("DEBUG") == "TRUE" {
		MasterEngine.ShowSQL(true)
		MasterEngine.Logger().SetLevel(core.LOG_DEBUG)
		SlaveEngine.ShowSQL(true)
		SlaveEngine.Logger().SetLevel(core.LOG_DEBUG)
	}
}
