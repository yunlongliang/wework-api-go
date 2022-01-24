# \ExternalcontactApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContactWay**](ExternalcontactApi.md#AddContactWay) | **Post** /externalcontact/add_contact_way | 配置客户联系「联系我」方式
[**AddCustomerTag**](ExternalcontactApi.md#AddCustomerTag) | **Post** /externalcontact/add_corp_tag | 添加企业客户标签
[**AddGroupWelcomeTemplate**](ExternalcontactApi.md#AddGroupWelcomeTemplate) | **Post** /externalcontact/group_welcome_template/add | 添加入群欢迎语素材
[**AddMsgTemplate**](ExternalcontactApi.md#AddMsgTemplate) | **Post** /externalcontact/add_msg_template | 创建企业群发
[**CloseTempChat**](ExternalcontactApi.md#CloseTempChat) | **Post** /externalcontact/close_temp_chat | 结束临时会话
[**ConvertOpengidToChatid**](ExternalcontactApi.md#ConvertOpengidToChatid) | **Post** /externalcontact/opengid_to_chatid | 客户群opengid转换
[**DelContactWay**](ExternalcontactApi.md#DelContactWay) | **Post** /externalcontact/del_contact_way | 删除企业已配置的「联系我」方式
[**DelCustomerTag**](ExternalcontactApi.md#DelCustomerTag) | **Post** /externalcontact/del_corp_tag | 删除企业客户标签
[**EditCustomerTag**](ExternalcontactApi.md#EditCustomerTag) | **Post** /externalcontact/edit_corp_tag | 编辑企业客户标签
[**GetBatchCustomerUser**](ExternalcontactApi.md#GetBatchCustomerUser) | **Post** /externalcontact/batch/get_by_user | 批量获取客户详情
[**GetContactWay**](ExternalcontactApi.md#GetContactWay) | **Post** /externalcontact/get_contact_way | 获取企业已配置的「联系我」方式
[**GetCorpTagList**](ExternalcontactApi.md#GetCorpTagList) | **Post** /externalcontact/get_corp_tag_list | 获取企业标签库
[**GetCustomerDetail**](ExternalcontactApi.md#GetCustomerDetail) | **Get** /externalcontact/get | 获取客户详情
[**GetCustomerList**](ExternalcontactApi.md#GetCustomerList) | **Get** /externalcontact/list | 获取客户列表
[**GetFollowUserList**](ExternalcontactApi.md#GetFollowUserList) | **Get** /externalcontact/get_follow_user_list | 获取配置了客户联系功能的成员列表
[**GetGroupmsgListV2**](ExternalcontactApi.md#GetGroupmsgListV2) | **Post** /externalcontact/get_groupmsg_list_v2 | 获取群发记录列表
[**GetGroupmsgSendResult**](ExternalcontactApi.md#GetGroupmsgSendResult) | **Post** /externalcontact/get_groupmsg_send_result | 获取企业群发成员执行结果
[**GetGroupmsgTask**](ExternalcontactApi.md#GetGroupmsgTask) | **Post** /externalcontact/get_groupmsg_task | 获取群发成员发送任务列表
[**GetMomentComments**](ExternalcontactApi.md#GetMomentComments) | **Post** /externalcontact/get_moment_comments | 获取客户朋友圈的互动数据
[**GetMomentCustomerList**](ExternalcontactApi.md#GetMomentCustomerList) | **Post** /externalcontact/get_moment_customer_list | 获取客户朋友圈发表时选择的可见范围
[**GetMomentList**](ExternalcontactApi.md#GetMomentList) | **Post** /externalcontact/get_moment_list | 获取企业全部的发表列表
[**GetMomentSendResult**](ExternalcontactApi.md#GetMomentSendResult) | **Post** /externalcontact/get_moment_send_result | 获取客户朋友圈发表后的可见客户列表
[**GetMomentTask**](ExternalcontactApi.md#GetMomentTask) | **Post** /externalcontact/get_moment_task | 获取客户朋友圈企业发表的列表
[**GetTransferCustomerResult**](ExternalcontactApi.md#GetTransferCustomerResult) | **Post** /externalcontact/resigned/transfer_result | 查询客户接替状态
[**GetUnassignedList**](ExternalcontactApi.md#GetUnassignedList) | **Post** /externalcontact/get_unassigned_list | 获取待分配的离职成员列表
[**GetUserBehaviorData**](ExternalcontactApi.md#GetUserBehaviorData) | **Post** /externalcontact/get_user_behavior_data | 获取「联系客户统计」数据
[**MarkCustomerTag**](ExternalcontactApi.md#MarkCustomerTag) | **Post** /externalcontact/mark_tag | 编辑客户企业标签
[**RemarkCustomer**](ExternalcontactApi.md#RemarkCustomer) | **Post** /externalcontact/remark | 修改客户备注信息
[**SendWelcomeMsg**](ExternalcontactApi.md#SendWelcomeMsg) | **Post** /externalcontact/send_welcome_msg | 发送新客户欢迎语
[**TransferCustomer**](ExternalcontactApi.md#TransferCustomer) | **Post** /externalcontact/transfer_customer | 分配在职成员的客户
[**TransferCustomerResult**](ExternalcontactApi.md#TransferCustomerResult) | **Post** /externalcontact/transfer_result | 查询客户接替状态
[**TransferResignedCustomer**](ExternalcontactApi.md#TransferResignedCustomer) | **Post** /externalcontact/resigned/transfer_customer | 分配离职成员的客户
[**UpdateContactWay**](ExternalcontactApi.md#UpdateContactWay) | **Post** /externalcontact/update_contact_way | 更新企业已配置的「联系我」方式


# **AddContactWay**
> AddContactWayRsp AddContactWay(ctx, accessToken, body)
配置客户联系「联系我」方式

企业可以在管理后台-客户联系-加客户中配置成员的「联系我」的二维码或者小程序按钮，客户通过扫描二维码或点击小程序上的按钮，即可获取成员联系方式,主动联系到成员。企业可通过此接口为具有客户联系功能的成员生成专属的「联系我」二维码或者「联系我」按钮。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**AddContactWayReq**](AddContactWayReq.md)|  | 

### Return type

[**AddContactWayRsp**](AddContactWayRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCustomerTag**
> AddCorpTagRsp AddCustomerTag(ctx, accessToken, optional)
添加企业客户标签

企业可通过此接口向客户标签库中添加新的标签组和标签，每个企业最多可配置3000个企业标签。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiAddCustomerTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiAddCustomerTagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AddCorpTagReq**](AddCorpTagReq.md)|  | 

### Return type

[**AddCorpTagRsp**](AddCorpTagRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddGroupWelcomeTemplate**
> AddGroupWelcomeTemplateRsp AddGroupWelcomeTemplate(ctx, accessToken, optional)
添加入群欢迎语素材

企业可通过此API向企业的入群欢迎语素材库中添加素材。每个企业的入群欢迎语素材库中，最多容纳100个素材。text中支持配置多个%NICKNAME%(大小写敏感)形式的欢迎语，当配置了欢迎语占位符后，发送给客户时会自动替换为客户的昵称;text、image、link和miniprogram四者不能同时为空；text与另外三者可以同时发送，此时将会以两条消息的形式触达客户;image、link和miniprogram只能有一个，如果三者同时填，则按image、link、miniprogram的优先顺序取参，也就是说，如果image与link同时传值，则只有image生效。media_id和pic_url只需填写一个，两者同时填写时使用media_id，二者不可同时为空。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiAddGroupWelcomeTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiAddGroupWelcomeTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AddGroupWelcomeTemplateReq**](AddGroupWelcomeTemplateReq.md)|  | 

### Return type

[**AddGroupWelcomeTemplateRsp**](AddGroupWelcomeTemplateRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddMsgTemplate**
> AddMsgTemplateRsp AddMsgTemplate(ctx, accessToken, optional)
创建企业群发

企业跟第三方应用可通过此接口添加企业群发消息的任务并通知成员发送给相关客户或客户群。（注：企业微信终端需升级到2.7.5版本及以上）注意：调用该接口并不会直接发送消息给客户/客户群，需要成员确认后才会执行发送（客服人员的企业微信需要升级到2.7.5及以上版本）旧接口创建企业群发已经废弃，接口升级后支持发送视频文件，并且支持最多同时发送9个附件。同一个企业每个自然月内仅可针对一个客户/客户群发送4条消息，超过接收上限的客户将无法再收到群发消息。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiAddMsgTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiAddMsgTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AddMsgTemplateReq**](AddMsgTemplateReq.md)|  | 

### Return type

[**AddMsgTemplateRsp**](AddMsgTemplateRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloseTempChat**
> BaseResponse CloseTempChat(ctx, accessToken, body)
结束临时会话

将指定的企业成员和客户之前的临时会话断开，断开前会自动下发已配置的结束语。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**CloseTempChatReq**](CloseTempChatReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertOpengidToChatid**
> ConvertOpengidToChatidRsp ConvertOpengidToChatid(ctx, accessToken, optional)
客户群opengid转换

用户在微信里的客户群里打开小程序时，某些场景下可以获取到群的opengid，如果该群是企业微信的客户群，则企业或第三方可以调用此接口将一个opengid转换为客户群chat_id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiConvertOpengidToChatidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiConvertOpengidToChatidOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ConvertOpengidToChatidReq**](ConvertOpengidToChatidReq.md)|  | 

### Return type

[**ConvertOpengidToChatidRsp**](ConvertOpengidToChatidRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DelContactWay**
> BaseResponse DelContactWay(ctx, accessToken, body)
删除企业已配置的「联系我」方式

删除一个已配置的「联系我」二维码或者「联系我」小程序按钮。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**DelContactWayReq**](DelContactWayReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DelCustomerTag**
> BaseResponse DelCustomerTag(ctx, accessToken, optional)
删除企业客户标签

企业可通过此接口删除客户标签库中的标签，或删除整个标签组。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiDelCustomerTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiDelCustomerTagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DelCorpTagReq**](DelCorpTagReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditCustomerTag**
> BaseResponse EditCustomerTag(ctx, accessToken, optional)
编辑企业客户标签

企业可通过此接口编辑客户标签/标签组的名称或次序值。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiEditCustomerTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiEditCustomerTagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EditCorpTagReq**](EditCorpTagReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBatchCustomerUser**
> GetExternalContactBatchUserRsp GetBatchCustomerUser(ctx, accessToken, optional)
批量获取客户详情

企业/第三方可通过此接口获取指定成员添加的客户信息列表。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetBatchCustomerUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetBatchCustomerUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetExternalContactBatchUserReq**](GetExternalContactBatchUserReq.md)|  | 

### Return type

[**GetExternalContactBatchUserRsp**](GetExternalContactBatchUserRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactWay**
> GetContactWayRsp GetContactWay(ctx, accessToken, body)
获取企业已配置的「联系我」方式

获取企业配置的「联系我」二维码和「联系我」小程序按钮。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**GetContactWayReq**](GetContactWayReq.md)|  | 

### Return type

[**GetContactWayRsp**](GetContactWayRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorpTagList**
> GetCorpTagListRsp GetCorpTagList(ctx, accessToken, optional)
获取企业标签库

企业可通过此接口获取企业客户标签详情。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetCorpTagListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetCorpTagListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetCorpTagListReq**](GetCorpTagListReq.md)|  | 

### Return type

[**GetCorpTagListRsp**](GetCorpTagListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerDetail**
> GetExternalContactDetailRsp GetCustomerDetail(ctx, accessToken, externalUserid, optional)
获取客户详情

企业可通过此接口，根据外部联系人的userid（如何获取?），拉取客户详情。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **externalUserid** | **string**|  | 
 **optional** | ***ExternalcontactApiGetCustomerDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetCustomerDetailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**|  | 

### Return type

[**GetExternalContactDetailRsp**](GetExternalContactDetailRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerList**
> GetExternalContactListRsp GetCustomerList(ctx, accessToken, userid)
获取客户列表

企业可通过此接口获取指定成员添加的客户列表。客户是指配置了客户联系功能的成员所添加的外部联系人。没有配置客户联系功能的成员，所添加的外部联系人将不会作为客户返回。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **userid** | **string**|  | 

### Return type

[**GetExternalContactListRsp**](GetExternalContactListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFollowUserList**
> GetFollowUserListRsp GetFollowUserList(ctx, accessToken)
获取配置了客户联系功能的成员列表

企业和第三方服务商可通过此接口获取配置了客户联系功能的成员列表。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 

### Return type

[**GetFollowUserListRsp**](GetFollowUserListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupmsgListV2**
> GetGroupmsgListV2Rsp GetGroupmsgListV2(ctx, accessToken, optional)
获取群发记录列表

企业和第三方应用可通过此接口获取企业与成员的群发记录。群发任务记录的起止时间间隔不能超过1个月

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetGroupmsgListV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetGroupmsgListV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetGroupmsgListV2Req**](GetGroupmsgListV2Req.md)|  | 

### Return type

[**GetGroupmsgListV2Rsp**](GetGroupmsgListV2Rsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupmsgSendResult**
> GetGroupmsgSendResultRsp GetGroupmsgSendResult(ctx, accessToken, optional)
获取企业群发成员执行结果

获取企业群发成员执行结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetGroupmsgSendResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetGroupmsgSendResultOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetGroupmsgSendResultReq**](GetGroupmsgSendResultReq.md)|  | 

### Return type

[**GetGroupmsgSendResultRsp**](GetGroupmsgSendResultRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupmsgTask**
> GetGroupmsgTaskRsp GetGroupmsgTask(ctx, accessToken, optional)
获取群发成员发送任务列表

获取群发成员发送任务列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetGroupmsgTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetGroupmsgTaskOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetGroupmsgTaskReq**](GetGroupmsgTaskReq.md)|  | 

### Return type

[**GetGroupmsgTaskRsp**](GetGroupmsgTaskRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMomentComments**
> GetMomentCommentsRsp GetMomentComments(ctx, accessToken, optional)
获取客户朋友圈的互动数据

企业和第三方应用可通过此接口获取客户朋友圈的互动数据。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetMomentCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetMomentCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetMomentCommentsReq**](GetMomentCommentsReq.md)|  | 

### Return type

[**GetMomentCommentsRsp**](GetMomentCommentsRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMomentCustomerList**
> GetMomentCustomerListRsp GetMomentCustomerList(ctx, accessToken, optional)
获取客户朋友圈发表时选择的可见范围

企业和第三方应用可通过该接口获取客户朋友圈创建时，选择的客户可见范围

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetMomentCustomerListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetMomentCustomerListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetMomentCustomerListReq**](GetMomentCustomerListReq.md)|  | 

### Return type

[**GetMomentCustomerListRsp**](GetMomentCustomerListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMomentList**
> GetMomentListRsp GetMomentList(ctx, accessToken, optional)
获取企业全部的发表列表

企业和第三方应用可通过该接口获取企业全部的发表内容。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetMomentListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetMomentListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetMomentListReq**](GetMomentListReq.md)|  | 

### Return type

[**GetMomentListRsp**](GetMomentListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMomentSendResult**
> GetMomentSendResultRsp GetMomentSendResult(ctx, accessToken, optional)
获取客户朋友圈发表后的可见客户列表

企业和第三方应用可通过该接口获取客户朋友圈发表后，可在微信朋友圈中查看的客户列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetMomentSendResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetMomentSendResultOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetMomentSendResultReq**](GetMomentSendResultReq.md)|  | 

### Return type

[**GetMomentSendResultRsp**](GetMomentSendResultRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMomentTask**
> GetMomentTaskRsp GetMomentTask(ctx, accessToken, optional)
获取客户朋友圈企业发表的列表

企业和第三方应用可通过该接口获取企业发表的朋友圈成员执行情况.第三方应用调用需要企业授权客户朋友圈下获取企业全部的发表记录的权限

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetMomentTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetMomentTaskOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetMomentTaskReq**](GetMomentTaskReq.md)|  | 

### Return type

[**GetMomentTaskRsp**](GetMomentTaskRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransferCustomerResult**
> ResignedTransferResultRsp GetTransferCustomerResult(ctx, accessToken, optional)
查询客户接替状态

企业和第三方可通过此接口查询离职成员的客户分配情况。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetTransferCustomerResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetTransferCustomerResultOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ResignedTransferResultReq**](ResignedTransferResultReq.md)|  | 

### Return type

[**ResignedTransferResultRsp**](ResignedTransferResultRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnassignedList**
> GetUnassignedListRsp GetUnassignedList(ctx, accessToken, optional)
获取待分配的离职成员列表

企业和第三方可通过此接口，获取所有离职成员的客户列表，并可进一步调用分配离职成员的客户接口将这些客户重新分配给其他企业成员。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetUnassignedListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetUnassignedListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetUnassignedListReq**](GetUnassignedListReq.md)|  | 

### Return type

[**GetUnassignedListRsp**](GetUnassignedListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserBehaviorData**
> GetUserBehaviorDataRsp GetUserBehaviorData(ctx, accessToken, optional)
获取「联系客户统计」数据

企业可通过此接口获取成员联系客户的数据，包括发起申请数、新增客户数、聊天数、发送消息数和删除/拉黑成员的客户数等指标。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiGetUserBehaviorDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiGetUserBehaviorDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetUserBehaviorDataReq**](GetUserBehaviorDataReq.md)|  | 

### Return type

[**GetUserBehaviorDataRsp**](GetUserBehaviorDataRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkCustomerTag**
> BaseResponse MarkCustomerTag(ctx, accessToken, optional)
编辑客户企业标签

企业可通过此接口为指定成员的客户添加上由企业统一配置的标签。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiMarkCustomerTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiMarkCustomerTagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MarkCorpContactTagReq**](MarkCorpContactTagReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemarkCustomer**
> BaseResponse RemarkCustomer(ctx, accessToken, body)
修改客户备注信息

企业可通过此接口修改指定用户添加的客户的备注信息。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**UpdateExternalContactRemarkReq**](UpdateExternalContactRemarkReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendWelcomeMsg**
> SendWelcomeMsgRsp SendWelcomeMsg(ctx, accessToken, optional)
发送新客户欢迎语

企业微信在向企业推送添加外部联系人事件时，会额外返回一个welcome_code，企业以此为凭据调用接口，即可通过成员向新添加的客户发送个性化的欢迎语。为了保证用户体验以及避免滥用，企业仅可在收到相关事件后20秒内调用，且只可调用一次。如果企业已经在管理端为相关成员配置了可用的欢迎语，则推送添加外部联系人事件时不会返回welcome_code。每次添加新客户时可能有多个企业自建应用/第三方应用收到带有welcome_code的回调事件，但仅有最先调用的可以发送成功。后续调用将返回41051（externaluser has started chatting）错误，请用户根据实际使用需求，合理设置应用可见范围，避免冲突。旧接口发送新客户欢迎语已经废弃，接口升级后支持发送视频文件，并且最多支持同时发送9个附件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiSendWelcomeMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiSendWelcomeMsgOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SendWelcomeMsgReq**](SendWelcomeMsgReq.md)|  | 

### Return type

[**SendWelcomeMsgRsp**](SendWelcomeMsgRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferCustomer**
> TransferCustomerRsp TransferCustomer(ctx, accessToken, optional)
分配在职成员的客户

企业可通过此接口，转接在职成员的客户给其他成员。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiTransferCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiTransferCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TransferCustomerReq**](TransferCustomerReq.md)|  | 

### Return type

[**TransferCustomerRsp**](TransferCustomerRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferCustomerResult**
> GetTransferResultRsp TransferCustomerResult(ctx, accessToken, optional)
查询客户接替状态

企业和第三方可通过此接口查询在职成员的客户转接情况。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiTransferCustomerResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiTransferCustomerResultOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetTransferResultReq**](GetTransferResultReq.md)|  | 

### Return type

[**GetTransferResultRsp**](GetTransferResultRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferResignedCustomer**
> ResignedTransferCustomerRsp TransferResignedCustomer(ctx, accessToken, optional)
分配离职成员的客户

企业可通过此接口，分配离职成员的客户给其他成员。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ExternalcontactApiTransferResignedCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalcontactApiTransferResignedCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ResignedTransferCustomerReq**](ResignedTransferCustomerReq.md)|  | 

### Return type

[**ResignedTransferCustomerRsp**](ResignedTransferCustomerRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContactWay**
> BaseResponse UpdateContactWay(ctx, accessToken, body)
更新企业已配置的「联系我」方式

更新企业配置的「联系我」二维码和「联系我」小程序按钮中的信息，如使用人员和备注等。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**UpdateContactWayReq**](UpdateContactWayReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

