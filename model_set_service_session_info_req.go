/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SetServiceSessionInfoReq struct {
	PreAuthCode string `json:"pre_auth_code,omitempty"`
	SessionInfo interface{} `json:"session_info,omitempty"`
}
