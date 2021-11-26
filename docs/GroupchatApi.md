# \GroupchatApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupchat**](GroupchatApi.md#GetGroupchat) | **Post** /externalcontact/groupchat/get | 获取客户群详情
[**GetGroupchatStatistic**](GroupchatApi.md#GetGroupchatStatistic) | **Post** /externalcontact/groupchat/statistic | 获取「群聊数据统计」数据 - 按群主聚合的方式
[**GetGroupchatStatisticGroupByDay**](GroupchatApi.md#GetGroupchatStatisticGroupByDay) | **Post** /externalcontact/groupchat/statistic_group_by_day | 获取「群聊数据统计」数据 - 按自然日聚合的方式
[**GetgroupchatList**](GroupchatApi.md#GetgroupchatList) | **Post** /externalcontact/groupchat/list | 获取客户群列表
[**TransferGroupchat**](GroupchatApi.md#TransferGroupchat) | **Post** /externalcontact/groupchat/transfer | 分配离职成员的客户群


# **GetGroupchat**
> GetGroupchatRsp GetGroupchat(ctx, accessToken, optional)
获取客户群详情

通过客户群ID，获取详情。包括群名、群成员列表、群成员入群时间、入群方式。（客户群是由具有客户群使用权限的成员创建的外部群）

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***GroupchatApiGetGroupchatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupchatApiGetGroupchatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetGroupchatReq**](GetGroupchatReq.md)|  | 

### Return type

[**GetGroupchatRsp**](GetGroupchatRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupchatStatistic**
> GetGroupchatStatisticRsp GetGroupchatStatistic(ctx, accessToken, optional)
获取「群聊数据统计」数据 - 按群主聚合的方式

获取指定日期的统计数据。注意，企业微信仅存储180天的数据。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***GroupchatApiGetGroupchatStatisticOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupchatApiGetGroupchatStatisticOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetGroupchatStatisticReq**](GetGroupchatStatisticReq.md)|  | 

### Return type

[**GetGroupchatStatisticRsp**](GetGroupchatStatisticRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupchatStatisticGroupByDay**
> GetGroupchatStatisticGroupByDayRsp GetGroupchatStatisticGroupByDay(ctx, accessToken, optional)
获取「群聊数据统计」数据 - 按自然日聚合的方式

获取「群聊数据统计」数据 - 按自然日聚合的方式

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***GroupchatApiGetGroupchatStatisticGroupByDayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupchatApiGetGroupchatStatisticGroupByDayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetGroupchatStatisticGroupByDayReq**](GetGroupchatStatisticGroupByDayReq.md)|  | 

### Return type

[**GetGroupchatStatisticGroupByDayRsp**](GetGroupchatStatisticGroupByDayRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetgroupchatList**
> GroupchatListRsp GetgroupchatList(ctx, accessToken, optional)
获取客户群列表

该接口用于获取配置过客户群管理的客户群列表。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***GroupchatApiGetgroupchatListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupchatApiGetgroupchatListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GroupchatListReq**](GroupchatListReq.md)|  | 

### Return type

[**GroupchatListRsp**](GroupchatListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferGroupchat**
> GroupchatTransferRsp TransferGroupchat(ctx, accessToken, optional)
分配离职成员的客户群

企业可通过此接口，将已离职成员为群主的群，分配给另一个客服成员。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***GroupchatApiTransferGroupchatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupchatApiTransferGroupchatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GroupchatTransferReq**](GroupchatTransferReq.md)|  | 

### Return type

[**GroupchatTransferRsp**](GroupchatTransferRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

