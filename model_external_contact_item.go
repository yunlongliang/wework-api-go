/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ExternalContactItem struct {
	FollowInfo *FollowInfo `json:"follow_info,omitempty"`
	ExternalContact interface{} `json:"external_contact,omitempty"`
}