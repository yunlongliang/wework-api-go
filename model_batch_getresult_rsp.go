/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BatchGetresultRsp struct {
	Errmsg string `json:"errmsg,omitempty"`
	Errcode int32 `json:"errcode,omitempty"`
	Status int32 `json:"status,omitempty"`
	Type_ string `json:"type,omitempty"`
	Total int32 `json:"total,omitempty"`
	Percentage int32 `json:"percentage,omitempty"`
	Result []interface{} `json:"result,omitempty"`
}
