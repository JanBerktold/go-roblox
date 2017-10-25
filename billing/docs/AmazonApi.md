# \AmazonApi

All URIs are relative to *https://billing.roblox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTV1AmazonPurchase**](AmazonApi.md#POSTV1AmazonPurchase) | **Post** /v1/amazon/purchase | Perform a purchase and grant desired product.
[**POSTV1AmazonValidate**](AmazonApi.md#POSTV1AmazonValidate) | **Post** /v1/amazon/validate | Validate a ProductId before making a purchase.


# **POSTV1AmazonPurchase**
> ApiEmptyResponseModel POSTV1AmazonPurchase($amazonStorePurchaseModel)

Perform a purchase and grant desired product.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amazonStorePurchaseModel** | [**AmazonStorePurchaseModel**](AmazonStorePurchaseModel.md)| Amazon Store Purchase Model | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTV1AmazonValidate**
> ApiEmptyResponseModel POSTV1AmazonValidate($validateModel)

Validate a ProductId before making a purchase.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateModel** | [**ValidateModel**](ValidateModel.md)| Product Id for example: com.roblox.client.bc1month, com.roblox.client.tbc1month, ... | 

### Return type

[**ApiEmptyResponseModel**](ApiEmptyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

