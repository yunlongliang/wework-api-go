/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetGroupmsgListV2Rsp struct {
	Errmsg string `json:"errmsg,omitempty"`
	Errcode int32 `json:"errcode,omitempty"`
	NextCursor string `json:"next_cursor,omitempty"`
	Attachments []interface{} `json:"attachments,omitempty"`
}
