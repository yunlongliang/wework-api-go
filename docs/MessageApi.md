# \MessageApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendMessage**](MessageApi.md#SendMessage) | **Post** /message/send | 发送消息
[**UpdateMessageTaskcard**](MessageApi.md#UpdateMessageTaskcard) | **Post** /message/update_taskcard | 更新任务卡片消息状态


# **SendMessage**
> SendMessageRsp SendMessage(ctx, accessToken, body)
发送消息

应用支持推送文本、图片、视频、文件、图文等类型

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **body** | [**SendMessageReq**](SendMessageReq.md)|  | 

### Return type

[**SendMessageRsp**](SendMessageRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMessageTaskcard**
> UpdateTaskcardRsp UpdateMessageTaskcard(ctx, accessToken, optional)
更新任务卡片消息状态

应用可以发送任务卡片消息，发送之后可再通过接口更新用户任务卡片消息的替换文案信息。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***MessageApiUpdateMessageTaskcardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessageApiUpdateMessageTaskcardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateTaskcardReq**](UpdateTaskcardReq.md)|  | 

### Return type

[**UpdateTaskcardRsp**](UpdateTaskcardRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

