package entities

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/natefinch/lumberjack"
)

/*
* set mysql master and slave service
* master as a writer
* slave as reader
 */
var MasterEngine *xorm.Engine
var SlaveEngine *xorm.Engine

func init() {
	var err error
	// Just for test
	// MasterEngine, err = xorm.NewEngine("mysql", "root:master@tcp(localhost:3307)/ccrsystem?charset=utf8&parseTime=true")
	// True connection for mysql
	MasterEngine, err = xorm.NewEngine("mysql", "root:master@tcp(dbmaster:3306)/ccrsystem?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}

	err = MasterEngine.Sync2(new(StudentInfo), new(AdminInfo), new(ApproverInfo), new(ClassroomInfo),
		new(ReservationInfo), new(DepartmentInfo))
	if err != nil {
		panic(err)
	}

	// Just for test
	// SlaveEngine, err = xorm.NewEngine("mysql", "root:slave@tcp(localhost:3308)/ccrsystem?charset=utf8&parseTime=true")
	// True connection for mysql
	SlaveEngine, err = xorm.NewEngine("mysql", "root:slave@tcp(dbslave:3306)/ccrsystem?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}

	// Open log function of xorm, and write the logs to file
	f := &lumberjack.Logger{
		Filename:   "./data/logs/sqllog/" + time.Now().Format("2006-01-02") + "-sql.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

	MasterEngine.ShowSQL(true)
	MasterEngine.Logger().SetLevel(core.LOG_DEBUG)
	MasterEngine.SetLogger(xorm.NewSimpleLogger(f))
	SlaveEngine.ShowSQL(true)
	SlaveEngine.Logger().SetLevel(core.LOG_DEBUG)
	SlaveEngine.SetLogger(xorm.NewSimpleLogger(f))
}
