// Package model 后端响应的数据结构
package model

// Response 后端响应的数据结构
type Response struct {
	Code  int         // 错误码
	Err   error       // 错误信息
	Total int64       // 符合条件的记录数目
	Spec  interface{} // 响应的实际数据
}
