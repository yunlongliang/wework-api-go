/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SendMessageRsp struct {
	Errmsg string `json:"errmsg,omitempty"`
	Errcode int32 `json:"errcode,omitempty"`
	Invaliduser string `json:"invaliduser,omitempty"`
	Invalidparty string `json:"invalidparty,omitempty"`
	Invalidtag string `json:"invalidtag,omitempty"`
}
