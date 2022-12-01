// Package connector 数据连接器，基于gorm，可支持多种类型的数据库
package connector

import (
	"fmt"
	"math"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"transformer_mz/libs/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// isDigit字符串是否纯数字
func isDigit(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsDigit(x) && x != rune('-') && x != rune('.') {
			return false
		}
	}
	return true
}

// ConditionSupport 添加数据库查询的条件支持
func ConditionSupport(c *gin.Context, db *gorm.DB) (*gorm.DB, error) {
	query, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		return nil, errors.New(err, "请求参数不正确，请联系管理员进行处理")
	}
	for name, value := range query {
		if strings.HasPrefix(name, "field_") && len(name) > 9 {
			fieldName := "`" + name[9:] + "`"
			if strings.Index(name, "JSON_EXTRACT") > 0 {
				fieldName = name[9:]
			} else {
				fieldName = strings.Replace(fieldName, ".", "`.`", -1)
			}
			var fieldValue []interface{}
			for _, v := range value {
				if isDigit(v) {
					numV, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return nil, errors.New(err, "请求参数不正确，请联系管理员进行处理")
					}
					if math.Mod(numV, 1) == 0 {
						fieldValue = append(fieldValue, int64(numV))
						continue
					}
					fieldValue = append(fieldValue, numV)
				} else {
					fieldValue = append(fieldValue, v)
				}
			}
			opt := name[6:8]
			switch opt {
			case "eq":
				if len(value) == 1 {
					db = db.Where(fieldName+"=?", value[0])
				} else {
					db = db.Where(fieldName+" IN (?)", fieldValue)
				}
			case "ne":
				db = db.Where(fieldName+"<>?", fieldValue[0])
			case "lt":
				db = db.Where(fieldName+"<?", fieldValue[0])
			case "le":
				db = db.Where(fieldName+"<=?", fieldValue[0])
			case "gt":
				db = db.Where(fieldName+">?", fieldValue[0])
			case "ge":
				db = db.Where(fieldName+">=?", fieldValue[0])
			case "lk":
				db = db.Where(fieldName+" LIKE ?", "%"+fmt.Sprintf("%v", fieldValue[0])+"%")
			case "re":
				db = db.Where(fieldName+" REGEXP ?", fieldValue[0])
			default:
				return nil, errors.New("不支持的条件查询操作", "条件查询参数不正确，请联系开发人员进行处理")
			}
		}
	}
	return db, nil
}

// Sort 添加数据库查询的排序支持
func Sort(c *gin.Context, db *gorm.DB) *gorm.DB {
	orderBy := c.Request.FormValue("order_by")
	if orderBy != "" {
		order := c.Request.FormValue("order")
		if order != "" {
			db = db.Order(orderBy + " " + order)
		} else {
			db = db.Order(orderBy + " DESC")
		}
	}
	return db
}

// Paging 添加数据库查询的分页支持
func Paging(c *gin.Context, db *gorm.DB) (*gorm.DB, error) {
	limitStr := c.Request.FormValue("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, errors.New(err, "分页参数不正确，请联系开发人员进行处理")
		}
		pageStr := c.Request.FormValue("page")
		if pageStr == "" {
			db = db.Limit(limit)
		} else {
			page, err := strconv.Atoi(pageStr)
			if err != nil {
				return nil, errors.New(err, "分页参数不正确，请联系开发人员进行处理")
			}
			db = db.Offset((page - 1) * limit).Limit(limit)
		}
	}
	return db, nil
}

// GetRecords 添加数据库复杂查询的支持
func GetRecords(c *gin.Context, db *gorm.DB, records interface{}) (int64, error) {
	var count int64
	err := c.Request.ParseForm()
	if err != nil {
		return 0, errors.New(err, "请求参数不正确，请联系开发人员进行处理")
	}
	db, err = ConditionSupport(c, db)
	if err != nil {
		return 0, errors.New(err, "查询条件有误，请联系开发人员进行处理")
	}
	err = db.Model(reflect.New(reflect.ValueOf(records).Type().Elem()).Interface()).Count(&count).Error
	if err != nil {
		return 0, errors.New(err, "数据库记录统计失败，请稍后再试")
	}
	db, err = Paging(c, db)
	if err != nil {
		return 0, errors.New(err, "分页参数不正确，请联系开发人员进行处理")
	}
	err = Sort(c, db).Find(records).Error
	return count, errors.New(err, "数据库记录查询失败，请稍后再试")
}

// DataSource 全局数据库源
var DataSource *gorm.DB
