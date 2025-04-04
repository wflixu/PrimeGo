package hello

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// 定义本地请求响应结构
type HelloReq struct {
	Name string `json:"name"`
}

type HelloResp struct {
	Message string `json:"message"`
}

type OtherReq struct {
	Data string `json:"data"`
}

type OtherResp struct {
	Result string `json:"result"`
}

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 返回模拟数据
	resp := &HelloResp{
		Message: "Hello,111 " + req.Name,
	}

	c.JSON(consts.StatusOK, resp)
}

// OtherMethod .
// @router /other [POST]
func OtherMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req OtherReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 返回模拟数据
	resp := &OtherResp{
		Result: "Processed: " + req.Data,
	}

	c.JSON(consts.StatusOK, resp)
}
