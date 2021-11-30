# \ServiceApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceAdminList**](ServiceApi.md#GetServiceAdminList) | **Post** /service/get_admin_list | 获取应用的管理员列表
[**GetServiceAuthInfo**](ServiceApi.md#GetServiceAuthInfo) | **Post** /service/get_auth_info | 获取企业授权信息
[**GetServiceCorpToken**](ServiceApi.md#GetServiceCorpToken) | **Post** /service/get_corp_token | 获取企业凭证
[**GetServicePermanentCode**](ServiceApi.md#GetServicePermanentCode) | **Post** /service/get_permanent_code | 获取企业永久授权码
[**GetServicePreAuthCode**](ServiceApi.md#GetServicePreAuthCode) | **Get** /service/get_pre_auth_code | 获取预授权码
[**GetServiceSuiteToken**](ServiceApi.md#GetServiceSuiteToken) | **Post** /service/get_suite_token | 获取第三方应用凭证
[**GetserviceProviderToken**](ServiceApi.md#GetserviceProviderToken) | **Post** /service/get_provider_token | 获取服务商凭证
[**SetServiceSessionInfo**](ServiceApi.md#SetServiceSessionInfo) | **Post** /service/set_session_info | 设置授权配置


# **GetServiceAdminList**
> GetServiceAdminListRsp GetServiceAdminList(ctx, suiteAccessToken, optional)
获取应用的管理员列表

第三方服务商可以用此接口获取授权企业中某个第三方应用的管理员列表(不包括外部管理员)，以便服务商在用户进入应用主页之后根据是否管理员身份做权限的区分。该应用必须与SUITE_ACCESS_TOKEN对应的suiteid对应，否则没权限查看

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **suiteAccessToken** | **string**|  | 
 **optional** | ***ServiceApiGetServiceAdminListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetServiceAdminListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetServiceAdminListReq**](GetServiceAdminListReq.md)|  | 

### Return type

[**GetServiceAdminListRsp**](GetServiceAdminListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceAuthInfo**
> GetServiceAuthInfoRsp GetServiceAuthInfo(ctx, suiteAccessToken, optional)
获取企业授权信息

该API用于通过永久授权码换取企业微信的授权信息。 永久code的获取，是通过临时授权码使用get_permanent_code 接口获取到的permanent_code。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **suiteAccessToken** | **string**|  | 
 **optional** | ***ServiceApiGetServiceAuthInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetServiceAuthInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetServiceAuthInfoReq**](GetServiceAuthInfoReq.md)|  | 

### Return type

[**GetServiceAuthInfoRsp**](GetServiceAuthInfoRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceCorpToken**
> GetServiceCorpTokenRsp GetServiceCorpToken(ctx, suiteAccessToken, optional)
获取企业凭证

第三方服务商在取得企业的永久授权码后，通过此接口可以获取到企业的access_token。获取后可通过通讯录、应用、消息等企业接口来运营这些应用。此处获得的企业access_token与企业获取access_token拿到的token，本质上是一样的，只不过获取方式不同。获取之后，就跟普通企业一样使用token调用API接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **suiteAccessToken** | **string**|  | 
 **optional** | ***ServiceApiGetServiceCorpTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetServiceCorpTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetServiceCorpTokenReq**](GetServiceCorpTokenReq.md)|  | 

### Return type

[**GetServiceCorpTokenRsp**](GetServiceCorpTokenRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServicePermanentCode**
> GetServicePermanentCodeRsp GetServicePermanentCode(ctx, suiteAccessToken, optional)
获取企业永久授权码

该API用于使用临时授权码换取授权方的永久授权码，并换取授权信息、企业access_token，临时授权码一次有效。建议第三方以userid为主键，来建立自己的管理员账号。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **suiteAccessToken** | **string**|  | 
 **optional** | ***ServiceApiGetServicePermanentCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetServicePermanentCodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetServicePermanentCodeReq**](GetServicePermanentCodeReq.md)|  | 

### Return type

[**GetServicePermanentCodeRsp**](GetServicePermanentCodeRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServicePreAuthCode**
> GetServicePreAuthCodeRsp GetServicePreAuthCode(ctx, suiteAccessToken)
获取预授权码

该API用于获取预授权码。预授权码用于企业授权时的第三方服务商安全验证。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **suiteAccessToken** | **string**|  | 

### Return type

[**GetServicePreAuthCodeRsp**](GetServicePreAuthCodeRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceSuiteToken**
> GetServiceSuiteTokenRsp GetServiceSuiteToken(ctx, optional)
获取第三方应用凭证

该API用于获取第三方应用凭证（suite_access_token）。由于第三方服务商可能托管了大量的企业，其安全问题造成的影响会更加严重，故API中除了合法来源IP校验之外，还额外增加了suite_ticket作为安全凭证。获取suite_access_token时，需要suite_ticket参数。suite_ticket由企业微信后台定时推送给“指令回调URL”，每十分钟更新一次，见推送suite_ticket。suite_ticket实际有效期为30分钟，可以容错连续两次获取suite_ticket失败的情况，但是请永远使用最新接收到的suite_ticket。通过本接口获取的suite_access_token有效期为2小时，开发者需要进行缓存，不可频繁获取。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceApiGetServiceSuiteTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetServiceSuiteTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GetServiceSuiteTokenReq**](GetServiceSuiteTokenReq.md)|  | 

### Return type

[**GetServiceSuiteTokenRsp**](GetServiceSuiteTokenRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetserviceProviderToken**
> GetServiceProviderTokenRsp GetserviceProviderToken(ctx, optional)
获取服务商凭证

获取服务商凭证

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceApiGetserviceProviderTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiGetserviceProviderTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GetServiceProviderTokenReq**](GetServiceProviderTokenReq.md)|  | 

### Return type

[**GetServiceProviderTokenRsp**](GetServiceProviderTokenRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetServiceSessionInfo**
> BaseResponse SetServiceSessionInfo(ctx, suiteAccessToken, optional)
设置授权配置

该接口可对某次授权进行配置。可支持测试模式（应用未发布时）。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **suiteAccessToken** | **string**|  | 
 **optional** | ***ServiceApiSetServiceSessionInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiSetServiceSessionInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SetServiceSessionInfoReq**](SetServiceSessionInfoReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

