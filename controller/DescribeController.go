package controller

import (
	"Hest/model/web"
	"Hest/rpc"
	"Hest/service"
	"Hest/util"
	"github.com/kataras/iris/v12"
)

type DescribeControllerStr struct {
	service *service.DescribeServiceStr
	util    *util.OperatorUtil
}

// GetBaseSync 同步
func (contr *DescribeControllerStr) GetBaseSync(context iris.Context) web.HestResp {
	baseDescribe, err := rpc.ReadBaseDescribe()
	if err != nil {
		return web.ERROR(err.Error())
	}
	length := len(*baseDescribe)
	for i := 0; i < length; i++ {
		base := (*baseDescribe)[i]
		base.Id = contr.util.GetId()
	}
	//service.
	return web.SUCCESS()
}
