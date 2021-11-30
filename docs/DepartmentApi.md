# \DepartmentApi

All URIs are relative to *https://qyapi.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDepartment**](DepartmentApi.md#CreateDepartment) | **Post** /department/create | 创建部门
[**DeleteDepartment**](DepartmentApi.md#DeleteDepartment) | **Get** /department/delete | 删除部门
[**GetBatchDepartmentResult**](DepartmentApi.md#GetBatchDepartmentResult) | **Get** /batch/getresult | 获取异步任务结果
[**GetDepartmentList**](DepartmentApi.md#GetDepartmentList) | **Get** /department/list | 获取部门列表
[**ReplaceBatchDepartment**](DepartmentApi.md#ReplaceBatchDepartment) | **Post** /batch/replaceparty | 全量覆盖部门
[**UpdateDepartment**](DepartmentApi.md#UpdateDepartment) | **Post** /department/update | 更新部门


# **CreateDepartment**
> CreateDepartmentRsp CreateDepartment(ctx, accessToken, optional)
创建部门

创建部门

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***DepartmentApiCreateDepartmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DepartmentApiCreateDepartmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateDepartmentReq**](CreateDepartmentReq.md)|  | 

### Return type

[**CreateDepartmentRsp**](CreateDepartmentRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDepartment**
> BaseResponse DeleteDepartment(ctx, accessToken, id)
删除部门

删除部门

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **id** | **int32**|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBatchDepartmentResult**
> BatchGetresultRsp GetBatchDepartmentResult(ctx, accessToken, jonid)
获取异步任务结果

获取异步任务结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **jonid** | **string**|  | 

### Return type

[**BatchGetresultRsp**](BatchGetresultRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDepartmentList**
> GetDepartmentListRsp GetDepartmentList(ctx, accessToken, optional)
获取部门列表

获取部门列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***DepartmentApiGetDepartmentListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DepartmentApiGetDepartmentListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int32**|  | 

### Return type

[**GetDepartmentListRsp**](GetDepartmentListRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceBatchDepartment**
> BatchReplacepartyRsp ReplaceBatchDepartment(ctx, accessToken, optional)
全量覆盖部门

本接口以partyid为键，全量覆盖企业的通讯录组织架构，任务完成后企业的通讯录组织架构与提交的文件完全保持一致。请先下载CSV文件(下载全量覆盖部门模版)，根据需求填写文件内容。注意事项：1.文件中存在、通讯录中也存在的部门，执行修改操作;2.文件中存在、通讯录中不存在的部门，执行添加操作;3.文件中不存在、通讯录中存在的部门，当部门下没有任何成员或子部门时，执行删除操作;4.文件中不存在、通讯录中存在的部门，当部门下仍有成员或子部门时，暂时不会删除，当下次导入成员把人从部门移出后自动删除;5.CSV文件中，部门名称、部门ID、父部门ID为必填字段，部门ID必须为数字，根部门的部门id默认为1；排序为可选字段，置空或填0不修改排序, order值大的排序靠前。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***DepartmentApiReplaceBatchDepartmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DepartmentApiReplaceBatchDepartmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BatchReplacepartyReq**](BatchReplacepartyReq.md)|  | 

### Return type

[**BatchReplacepartyRsp**](BatchReplacepartyRsp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDepartment**
> BaseResponse UpdateDepartment(ctx, accessToken, optional)
更新部门

更新部门

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***DepartmentApiUpdateDepartmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DepartmentApiUpdateDepartmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateDepartmentReq**](UpdateDepartmentReq.md)|  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

