/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetServicePreAuthCodeReq struct {
	Corpid string `json:"corpid,omitempty"`
	ProviderSecret string `json:"provider_secret,omitempty"`
}
