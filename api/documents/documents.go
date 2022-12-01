// Package 文件管理，包括标准库文件等
/**
todo  处理标准库文件上传、下载、
todo 文件上传后生成文件存储路径，uuid 等信息后将文件保存至文件系统，
todo 文件下载时跟进文件uuid等信息拼接文件下载路径进行文件下载。

*/
package Documents

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"transformer_mz/databases/connector"
	"transformer_mz/databases/model"
	"transformer_mz/libs/config"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
	"transformer_mz/utils"
)

var BasePath string

// MysqlConfigWrap mysql连接设置
type DocsConfigWrap struct {
	Uploadpath string `yaml:"uploadpath"`
}

// MysqlConfig mysql连接设置
type DocsConfig struct {
	DocsConfigWrap `yaml:"docs"`
}

// UploadDocuments 添加文档信息
func UploadDocuments(c *gin.Context) {

	err := permission.Check(c, resource.Documents, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	fileForm := struct {
		Category string `form:"category"`
	}{}
	if err = errors.New(c.Bind(&fileForm), "请求数据有误，请联系开发人员进行处理"); err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	filePath := BasePath + filepath.Base(file.Filename)
	//// 判断目录是否存在，不存在则新建目录
	_, err = os.Stat(BasePath)
	if os.IsNotExist(err) {
		err := os.Mkdir(BasePath, os.ModePerm)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
			return
		}
	}
	//// 保存文件
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	docsType := utils.Eext(file.Filename)
	////docsType := path.Ext(file.Filename)
	////split := strings.Split(file.Filename, `.`)
	////docsType := split[len(split)-1]
	id := uuid.NewV4().String()
	size := float64(file.Size) / 1024 / 1024
	newDocuments := &model.Documents{}
	newDocuments.UUID = id
	newDocuments.Name = file.Filename
	newDocuments.DocumentType = docsType
	newDocuments.FileSize, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", size), 64)
	newDocuments.Location = filePath
	newDocuments.DocsCreator = self.ID
	newDocuments.Category = model.DocumentCategory(fileForm.Category)
	err = errors.New(connector.DataSource.Create(newDocuments).Error, "文档信息添加失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "文档信息添加成功"})
}

// DownloadDocuments 处理文件下载
func DownloadDocuments(c *gin.Context) {
	var documents model.Documents
	err := permission.Check(c, resource.Documents, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&documents).Error,
		"文档信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	filePath := BasePath + documents.Location
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: errors.New(err, "文件不存在")})
		return
	}
	//打开文件
	fileTmp, errByOpenFile := os.Open(filePath)

	//非空处理
	if errByOpenFile != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	defer fileTmp.Close()
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	file, err := os.Create(filePath) //创建文件
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	exportName := strings.TrimSpace(documents.Name) + documents.DocumentType
	c.Header("Content-Type", "application/octet-stream")
	//强制浏览器下载
	c.Header("Content-Disposition", "attachment; filename="+exportName)
	//浏览器下载或预览
	c.Header("Content-Disposition", "inline;filename="+exportName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")

	buf := bufio.NewWriter(file) //创建新的 Writer 对象
	buf.WriteString(string(f))
	buf.Flush()
	defer file.Close()
	//返回文件流
	c.File(file.Name())
	return
}

// UpdateDocuments 更新文档信息
func UpdateDocuments(c *gin.Context) {
	var Documents model.Documents
	err := permission.Check(c, resource.Documents, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(c.BindJSON(&Documents), "请求数据有误，请联系开发人员进行处理")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusBadRequest, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Model(&model.Documents{}).Where("id=?", id).Updates(&Documents).Error,
		"文档信息更新失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "文档信息更新成功"})
}

// DeleteDocuments 删除文档信息
func DeleteDocuments(c *gin.Context) {
	err := permission.Check(c, resource.Documents, action.Write)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).Delete(&model.Documents{}).Error,
		"文档信息删除失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: "文档信息删除成功"})
}

// GetSpecDocuments 查找特定文档信息
func GetSpecDocuments(c *gin.Context) {
	var Documents model.Documents
	self, err := utils.GetUserFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	id, err := utils.GetIdFromRequest(c)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	err = errors.New(connector.DataSource.Where("id=?", id).First(&Documents).Error,
		"文档信息查找失败，请稍后再试")
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	if self.ID != Documents.ID {
		err = permission.Check(c, resource.Documents, action.Read)
		if err != nil {
			c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
			return
		}
	}
	c.Set("response", &model.Response{Spec: Documents})
}

// FindDocumentss 查找文档信息
func FindDocuments(c *gin.Context) {
	var Documents []model.Documents
	err := permission.Check(c, resource.Documents, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	total, err := connector.GetRecords(c, connector.DataSource.Preload("CreatorUser").Preload("DesignProject"), &Documents)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Total: total, Spec: Documents})
}

// InitDocuments 初始化
func InitDocuments() error {
	docsConfig := &DocsConfig{}
	err := config.Load(docsConfig)
	if err != nil {
		return err
	}
	BasePath = docsConfig.Uploadpath
	return nil
}
