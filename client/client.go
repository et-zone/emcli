package client

import (
	"errors"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/kubernetes/v3"
	"github.com/asim/go-micro/v3"
	servCli "github.com/et-zone/serv_api/gen_go"
)
const(
	EnvOnline = "online"
	EnvPre = "pre"
	EnvDev = "dev"
)

// client
var	SrvClient servCli.SrvInterfaceService



func NewService(env ,servName string,port int)(micro.Service,error){

	var service micro.Service
	switch env {
	case EnvOnline:
		var options []micro.Option
		options = append(options, micro.Name(servName))
		options = append(options, micro.Address(fmt.Sprintf(":%v",port)))
		options = append(options, micro.Registry(kubernetes.NewRegistry()))
		service = micro.NewService(options...)

	case EnvPre:
		var options []micro.Option
		options = append(options, micro.Name(servName))
		options = append(options, micro.Address(fmt.Sprintf(":%v",port)))
		options = append(options, micro.Registry(kubernetes.NewRegistry()))
		service = micro.NewService(options...)

	case EnvDev:
		service = micro.NewService(
			micro.Name(servName),
			micro.Address(fmt.Sprintf(":%v",port)),
		)
	default:
		return nil,errors.New("env err, not fond env information ")
	}
	service.Init()
	return service,nil
}


//register client
func Register(mic micro.Service){
	SrvClient=servCli.NewSrvInterfaceService("carsome.serv.com",mic.Client())
}