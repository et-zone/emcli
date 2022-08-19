package service

import (
	"context"
	"fmt"
	servCli "github.com/et-zone/serv_api/gen_go"
	"gmcli/client"
)

func Ping(ctx context.Context) string{

	// Make request

	rsp, err := client.SrvClient.Ping(ctx, &servCli.Request{
		Name: "John",
		Common: &servCli.Common{
			City: "12",
			Contry: "900"},
	})
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(rsp.Msg)
	return rsp.Msg
}

