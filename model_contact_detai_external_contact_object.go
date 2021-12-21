/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ContactDetaiExternalContactObject struct {
	ExternalUserid string `json:"external_userid,omitempty"`
	Name string `json:"name,omitempty"`
	Position string `json:"position,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	CorpName string `json:"corp_name,omitempty"`
	CorpFullName string `json:"corp_full_name,omitempty"`
	Type_ int32 `json:"type,omitempty"`
	Gender int32 `json:"gender,omitempty"`
	Unionid string `json:"unionid,omitempty"`
	ExternalProfile interface{} `json:"external_profile,omitempty"`
}