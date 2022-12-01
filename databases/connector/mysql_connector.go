// Package connector mysql数据库连接器
package connector

import (
	"reflect"
	"time"

	"transformer_mz/databases/model"
	"transformer_mz/libs/config"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// MysqlConfigWrap mysql连接设置
type MysqlConfigWrap struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Debug    bool   `yaml:"debug"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxOpen  int    `yaml:"max_open"`
	MaxLife  int    `yaml:"max_life"`
}

// MysqlConfig mysql连接设置
type MysqlConfig struct {
	MysqlConfigWrap `yaml:"database"`
}

// excludeCallback 添加查询记录时的回调处理
func excludeCallback(db *gorm.DB) {
	if db.Error == nil {
		if result, ok := db.Get("skipCallback"); ok && result.(bool) {
			return
		}
		if deletedAtField, ok := db.Statement.Schema.FieldsByName["DeletedAt"]; ok {
			db.Statement.AddClause(clause.Where{Exprs: []clause.Expression{
				clause.Eq{Column: clause.Column{Name: deletedAtField.DBName}, Value: nil}}})
		}
	}
}

// createCallback 添加创建记录时的回调处理
func createCallback(db *gorm.DB) {
	if db.Error == nil {
		now := model.Timestamp{Time: time.Now()}
		reflectValue := reflect.ValueOf(db.Statement.Dest)
		if createdAtField, ok := db.Statement.Schema.FieldsByName["CreatedAt"]; ok {
			switch reflectValue.Kind() {
			case reflect.Slice, reflect.Array:
				for i := 0; i < reflectValue.Len(); i++ {
					db.Error = createdAtField.Set(reflectValue.Index(i), now)
					if db.Error != nil {
						return
					}
				}
			case reflect.Ptr, reflect.Struct:
				db.Error = createdAtField.Set(reflect.ValueOf(db.Statement.Dest), now)
				if db.Error != nil {
					return
				}
			}
		}
		if updatedAtField, ok := db.Statement.Schema.FieldsByName["UpdatedAt"]; ok {
			switch reflectValue.Kind() {
			case reflect.Slice, reflect.Array:
				for i := 0; i < reflectValue.Len(); i++ {
					db.Error = updatedAtField.Set(reflectValue.Index(i), now)
					if db.Error != nil {
						return
					}
				}
			case reflect.Ptr, reflect.Struct:
				db.Error = updatedAtField.Set(reflect.ValueOf(db.Statement.Dest), now)
				if db.Error != nil {
					return
				}
			}
		}
		if operatedAtField, ok := db.Statement.Schema.FieldsByName["OperatedAt"]; ok {
			switch reflectValue.Kind() {
			case reflect.Slice, reflect.Array:
				for i := 0; i < reflectValue.Len(); i++ {
					db.Error = operatedAtField.Set(reflectValue.Index(i), now)
					if db.Error != nil {
						return
					}
				}
			case reflect.Ptr, reflect.Struct:
				db.Error = operatedAtField.Set(reflect.ValueOf(db.Statement.Dest), now)
				if db.Error != nil {
					return
				}
			}
		}
	}
}

// updateCallback 添加更新记录时的回调处理
func updateCallback(db *gorm.DB) {
	if db.Error == nil {
		if updatedAtField, ok := db.Statement.Schema.FieldsByName["UpdatedAt"]; ok {
			updatedAtField.AutoUpdateTime = schema.TimeType(0)
			newClause := clause.Set{{Column: clause.Column{Name: updatedAtField.DBName}, Value: model.Timestamp{Time: time.Now()}}}
			if db.Statement.SQL.String() == "" {
				db.Statement.SQL.Grow(180)
				db.Statement.AddClauseIfNotExists(clause.Update{})
				if set := callbacks.ConvertToAssignments(db.Statement); len(set) != 0 {
					db.Statement.AddClause(append(set, newClause...))
				} else {
					db.Statement.AddClause(newClause)
				}
				db.Statement.Build(db.Statement.BuildClauses...)
			}
		}
	}
}

// deleteCallback 添加软删除记录时的回调处理
func deleteCallback(db *gorm.DB) {
	if db.Error == nil {
		if deletedAtField, ok := db.Statement.Schema.FieldsByName["DeletedAt"]; ok {
			stmt := db.Statement
			if stmt.SQL.String() == "" {
				stmt.AddClause(clause.Set{{Column: clause.Column{Name: deletedAtField.DBName}, Value: model.Timestamp{Time: time.Now()}}})
				if stmt.Schema != nil {
					_, queryValues := schema.GetIdentityFieldValuesMap(stmt.ReflectValue, stmt.Schema.PrimaryFields)
					column, values := schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)
					if len(values) > 0 {
						stmt.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
					}
					if stmt.ReflectValue.CanAddr() && stmt.Dest != stmt.Model && stmt.Model != nil {
						_, queryValues = schema.GetIdentityFieldValuesMap(reflect.ValueOf(stmt.Model), stmt.Schema.PrimaryFields)
						column, values = schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)
						if len(values) > 0 {
							stmt.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
						}
					}
				}
				stmt.AddClauseIfNotExists(clause.Update{})
				stmt.Build("UPDATE", "SET", "WHERE")
			}
		}
	}
}

// InitMysqlConnector 初始化数据库连接器
func InitMysqlConnector() error {
	mysqlConfig := &MysqlConfig{}
	err := config.Load(mysqlConfig)
	if err != nil {
		return err
	}
	DataSource, err = gorm.Open(mysql.New(mysql.Config{
		DSN: mysqlConfig.Username + ":" + mysqlConfig.Password + "@tcp" + mysqlConfig.Address,
	}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return errors.New(err)
	}
	// 设置数据库操作日志
	loggerConfig := logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Error,
		Colorful:      false,
	}
	// 开启数据库操作调试信息
	if mysqlConfig.Debug {
		loggerConfig.LogLevel = logger.Info
	}
	DataSource.Logger = logger.New(log.Logger, loggerConfig)
	DataSource.AutoMigrate(&model.DesignResults{})
	// 设置数据库连接池
	sqlDB, err := DataSource.DB()
	if err != nil {
		return errors.New(err)
	}
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConfig.MaxLife) * time.Minute)
	_ = DataSource.Callback().Query().Before("gorm:query").Register("exclude_soft_deleted", excludeCallback)
	_ = DataSource.Callback().Create().Before("gorm:create").Register("update_timestamp", createCallback)
	_ = DataSource.Callback().Update().Before("gorm:update").Register("update_timestamp", updateCallback)
	_ = DataSource.Callback().Delete().Before("gorm:delete").Register("update_timestamp", deleteCallback)
	return nil
}
