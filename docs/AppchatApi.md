# \AppchatApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppchat**](AppchatApi.md#CreateAppchat) | **Post** /appchat/create | 创建群聊会话
[**GetAppchat**](AppchatApi.md#GetAppchat) | **Get** /appchat/get | 获取群聊会话
[**UpdateAppchat**](AppchatApi.md#UpdateAppchat) | **Post** /appchat/update | 修改群聊会话


# **CreateAppchat**
> CreateAppchatRsp CreateAppchat(ctx, accessToken, optional)
创建群聊会话

创建群聊会话

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***AppchatApiCreateAppchatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppchatApiCreateAppchatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateAppchatReq**](CreateAppchatReq.md)|  | 

### Return type

[**CreateAppchatRsp**](CreateAppchatRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppchat**
> GetAppchatReq GetAppchat(ctx, accessToken, chatid)
获取群聊会话

获取群聊会话

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **chatid** | **string**|  | 

### Return type

[**GetAppchatReq**](GetAppchatReq.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAppchat**
> BaseResponse UpdateAppchat(ctx, accessToken, optional)
修改群聊会话

修改群聊会话

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***AppchatApiUpdateAppchatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppchatApiUpdateAppchatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateAppchatReq**](UpdateAppchatReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

