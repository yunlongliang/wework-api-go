package main

import (
	"fmt"
	weworkApi "github.com/yunlongliang/wework-api-go"
	"context"
)

func main() {
	cfg := weworkApi.NewConfiguration()
	ApiClient := weworkApi.NewAPIClient(cfg)
	corpID := "wwe99c32aded3d14c2"
	secret := "6cq42U5RfuHseIwPC68O-4ZoET4mibYX5BEIuCzJyQ4"
	tokenRsp, _, _ := ApiClient.AuthApi.GetAuthToken(context.Background(), corpID, secret)
	if tokenRsp.Errcode != 0 {
		fmt.Println(tokenRsp.Errmsg)
		return
	}
	fmt.Println(tokenRsp.AccessToken)
}
