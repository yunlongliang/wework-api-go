/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetMenuRsp struct {
	Errmsg string `json:"errmsg,omitempty"`
	Errcode int32 `json:"errcode,omitempty"`
	Button []interface{} `json:"button,omitempty"`
}
