/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CorpTag struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	CreateTime int32 `json:"create_time,omitempty"`
	Order int32 `json:"order,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
}