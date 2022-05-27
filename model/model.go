package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"qiu/blog/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID         uint `gorm:"primary_key" uri:"id" `
	CreatedOn  time.Time
	ModifiedOn time.Time
	DeletedAt  *time.Time `sql:"index"`
}

func Setup() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec := setting.DatabaseSetting
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Type
	dbName = sec.Name
	user = sec.User
	password = sec.Password
	host = sec.Host
	tablePrefix = sec.TablePrefix
	// log.Println("数据库初始化：", dbType, dbName, user, password, host, tablePrefix)

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Article{})
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		// nowTime := time.Now().Unix()
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("State"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(1)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now())
	}
}

// func deleteCallback(scope *gorm.Scope) {
// 	if !scope.HasError() {
// 		var extraOption string
// 		if str, ok := scope.Get("gorm:delete_option"); ok {
// 			extraOption = fmt.Sprint(str)
// 		}

// 		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

// 		if !scope.Search.Unscoped && hasDeletedOnField {
// 			scope.Raw(fmt.Sprintf(
// 				"UPDATE %v SET %v=%v%v%v",
// 				scope.QuotedTableName(),
// 				scope.Quote(deletedOnField.DBName),
// 				scope.AddToVars(time.Now()),
// 				addExtraSpaceIfExist(scope.CombinedConditionSql()),
// 				addExtraSpaceIfExist(extraOption),
// 			)).Exec()
// 		} else {
// 			scope.Raw(fmt.Sprintf(
// 				"DELETE FROM %v%v%v",
// 				scope.QuotedTableName(),
// 				addExtraSpaceIfExist(scope.CombinedConditionSql()),
// 				addExtraSpaceIfExist(extraOption),
// 			)).Exec()
// 		}
// 	}
// }

// func addExtraSpaceIfExist(str string) string {
// 	if str != "" {
// 		return " " + str
// 	}
// 	return ""
// }

func CloseDB() {
	defer db.Close()
}
