// Package 系统维护信息
package maintenance

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"transformer_mz/databases/model"
	"transformer_mz/libs/monitor"
	"transformer_mz/libs/permission"
	"transformer_mz/libs/permission/action"
	"transformer_mz/libs/permission/resource"
)

// UpdateDesignTaskStatus 审核设计任务的结果，状态修改只能为审核通过或者不接受
/**
系统基本信息
*/
func SystemBasicInfoStatus(c *gin.Context) {
	var (
		systembasic interface{}
		err         error
	)
	err = permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	systembasic, err = monitor.SysHostAdvancedInfo()
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: systembasic})
	return

}

/**
系统内存
*/
func SystemMemInfoStatus(c *gin.Context) {
	var (
		systemInfo interface{}
		err        error
	)
	err = permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	systemInfo, err = monitor.SysMemInfo()
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: systemInfo})
	return

}

/**
系统cpu
*/
func SystemCpuInfoInfoStatus(c *gin.Context) {
	var (
		systemInfo interface{}
		err        error
	)
	err = permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	systemInfo, err = monitor.SysCpuInfo()
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: systemInfo})
	return

}

/**
系统host
*/
func SystemHostsInfoStatus(c *gin.Context) {
	var (
		systemInfo interface{}
		err        error
	)
	err = permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	systemInfo, err = monitor.SysHostStatInfo()
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: systemInfo})
	return

}

/**
系统硬盘存储信息
*/
func SystemDisksInfoStatus(c *gin.Context) {
	var (
		systemInfo interface{}
		err        error
	)
	err = permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	//todo
	//systemInfo,err = monitor.SysAdvancedDiskInfo()
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: systemInfo})
	return
}

/**
系统网络信息
*/
func SystemNetInfoStatus(c *gin.Context) {
	var (
		systemInfo interface{}
		err        error
	)
	err = permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	systemInfo, err = monitor.SysAdvancedNetInfo(true)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: systemInfo})
	return
}

/**
系统进程信息
*/
func SystemProcessInfoStatus(c *gin.Context) {
	var (
		systemInfo interface{}
		err        error
	)
	err = permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	systemInfo, err = monitor.SysAdvancedProcessInfo()
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: systemInfo})
	return
}

/**
系统负载
*/
func SystemLoadInfoStatus(c *gin.Context) {
	var (
		systemInfo interface{}
		err        error
	)
	err = permission.Check(c, resource.Libraries, action.Read)
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusUnauthorized, Err: err})
		return
	}
	systemInfo, err = monitor.SysAdvancedLoadInfo()
	if err != nil {
		c.Set("response", &model.Response{Code: http.StatusInternalServerError, Err: err})
		return
	}
	c.Set("response", &model.Response{Spec: systemInfo})
	return
}
