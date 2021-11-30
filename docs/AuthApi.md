# \AuthApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiDomainIp**](AuthApi.md#GetApiDomainIp) | **Get** /get_api_domain_ip | 获取企业微信API域名IP段
[**GetAuthToken**](AuthApi.md#GetAuthToken) | **Get** /gettoken | 每个应用有独立的secret，获取到的access_token只能本应用使用，所以每个应用的access_token应该分开来获取


# **GetApiDomainIp**
> GetApiDomainIpRsp GetApiDomainIp(ctx, accessToken)
获取企业微信API域名IP段



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 

### Return type

[**GetApiDomainIpRsp**](GetApiDomainIpRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthToken**
> TokenRsp GetAuthToken(ctx, corpid, corpsecret)
每个应用有独立的secret，获取到的access_token只能本应用使用，所以每个应用的access_token应该分开来获取



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **corpid** | **string**| 企业ID | 
  **corpsecret** | **string**| 应用的凭证密钥 | 

### Return type

[**TokenRsp**](TokenRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

