/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetServiceSuiteTokenReq struct {
	SuiteId string `json:"suite_id,omitempty"`
	SuiteSecret string `json:"suite_secret,omitempty"`
	SuiteTicket string `json:"suite_ticket,omitempty"`
}
