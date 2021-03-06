/*
 * 企业微信服务端API
 *
 * 企业微信服务端API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SendMessageReq struct {
	Touser string `json:"touser,omitempty"`
	Toparty string `json:"toparty,omitempty"`
	Totag string `json:"totag,omitempty"`
	Msgtype string `json:"msgtype,omitempty"`
	Agentid int32 `json:"agentid,omitempty"`
	Text *TextMsg `json:"text,omitempty"`
	Image *ImageMsg `json:"image,omitempty"`
	Voice *VoiceMsg `json:"voice,omitempty"`
	Video *VideoMsg `json:"video,omitempty"`
	Textcard *TextcardMsg `json:"textcard,omitempty"`
	News *NewsMsg `json:"news,omitempty"`
	Mpnews *MpNewsMsg `json:"mpnews,omitempty"`
	Markdown *MarkdownMsg `json:"markdown,omitempty"`
	MiniprogramNotice *MiniProgramMsg `json:"miniprogram_notice,omitempty"`
	Taskcard *TaskcardMsg `json:"taskcard,omitempty"`
}
