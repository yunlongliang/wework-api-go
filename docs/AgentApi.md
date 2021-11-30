# \AgentApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgent**](AgentApi.md#GetAgent) | **Get** /agent/get | 获取指定的应用详情
[**GetAgentList**](AgentApi.md#GetAgentList) | **Post** /agent/list | 获取access_token对应的应用列表
[**SetAgent**](AgentApi.md#SetAgent) | **Post** /agent/set | 设置应用


# **GetAgent**
> GetAgentRsp GetAgent(ctx, accessToken, agentid)
获取指定的应用详情

对于互联企业的应用，如果需要获取应用可见范围内其他互联企业的部门与成员，请调用互联企业-获取应用可见范围接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **agentid** | **string**|  | 

### Return type

[**GetAgentRsp**](GetAgentRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgentList**
> GetAgentListRsp GetAgentList(ctx, accessToken)
获取access_token对应的应用列表

获取access_token对应的应用列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 

### Return type

[**GetAgentListRsp**](GetAgentListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAgent**
> SetAgentRsp SetAgent(ctx, accessToken, optional)
设置应用

设置应用

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***AgentApiSetAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentApiSetAgentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SetAgentReq**](SetAgentReq.md)|  | 

### Return type

[**SetAgentRsp**](SetAgentRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

