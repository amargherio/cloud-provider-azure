// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ObjectReplicationPoliciesClient contains the methods for the ObjectReplicationPolicies group.
// Don't use this type directly, use NewObjectReplicationPoliciesClient() instead.
type ObjectReplicationPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewObjectReplicationPoliciesClient creates a new instance of ObjectReplicationPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewObjectReplicationPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ObjectReplicationPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ObjectReplicationPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update the object replication policy of the storage account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - objectReplicationPolicyID - For the destination account, provide the value 'default'. Configure the policy on the destination
//     account first. For the source account, provide the value of the policy ID that is returned when you
//     download the policy that was defined on the destination account. The policy is downloaded as a JSON file.
//   - properties - The object replication policy set to a storage account. A unique policy ID will be created if absent.
//   - options - ObjectReplicationPoliciesClientCreateOrUpdateOptions contains the optional parameters for the ObjectReplicationPoliciesClient.CreateOrUpdate
//     method.
func (client *ObjectReplicationPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, properties ObjectReplicationPolicy, options *ObjectReplicationPoliciesClientCreateOrUpdateOptions) (ObjectReplicationPoliciesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ObjectReplicationPoliciesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, objectReplicationPolicyID, properties, options)
	if err != nil {
		return ObjectReplicationPoliciesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ObjectReplicationPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ObjectReplicationPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ObjectReplicationPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, properties ObjectReplicationPolicy, _ *ObjectReplicationPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies/{objectReplicationPolicyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if objectReplicationPolicyID == "" {
		return nil, errors.New("parameter objectReplicationPolicyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{objectReplicationPolicyId}", url.PathEscape(objectReplicationPolicyID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ObjectReplicationPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (ObjectReplicationPoliciesClientCreateOrUpdateResponse, error) {
	result := ObjectReplicationPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectReplicationPolicy); err != nil {
		return ObjectReplicationPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the object replication policy associated with the specified storage account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - objectReplicationPolicyID - For the destination account, provide the value 'default'. Configure the policy on the destination
//     account first. For the source account, provide the value of the policy ID that is returned when you
//     download the policy that was defined on the destination account. The policy is downloaded as a JSON file.
//   - options - ObjectReplicationPoliciesClientDeleteOptions contains the optional parameters for the ObjectReplicationPoliciesClient.Delete
//     method.
func (client *ObjectReplicationPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, options *ObjectReplicationPoliciesClientDeleteOptions) (ObjectReplicationPoliciesClientDeleteResponse, error) {
	var err error
	const operationName = "ObjectReplicationPoliciesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, objectReplicationPolicyID, options)
	if err != nil {
		return ObjectReplicationPoliciesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ObjectReplicationPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ObjectReplicationPoliciesClientDeleteResponse{}, err
	}
	return ObjectReplicationPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ObjectReplicationPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, _ *ObjectReplicationPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies/{objectReplicationPolicyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if objectReplicationPolicyID == "" {
		return nil, errors.New("parameter objectReplicationPolicyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{objectReplicationPolicyId}", url.PathEscape(objectReplicationPolicyID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the object replication policy of the storage account by policy ID.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - objectReplicationPolicyID - For the destination account, provide the value 'default'. Configure the policy on the destination
//     account first. For the source account, provide the value of the policy ID that is returned when you
//     download the policy that was defined on the destination account. The policy is downloaded as a JSON file.
//   - options - ObjectReplicationPoliciesClientGetOptions contains the optional parameters for the ObjectReplicationPoliciesClient.Get
//     method.
func (client *ObjectReplicationPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, options *ObjectReplicationPoliciesClientGetOptions) (ObjectReplicationPoliciesClientGetResponse, error) {
	var err error
	const operationName = "ObjectReplicationPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, objectReplicationPolicyID, options)
	if err != nil {
		return ObjectReplicationPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ObjectReplicationPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ObjectReplicationPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ObjectReplicationPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, _ *ObjectReplicationPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies/{objectReplicationPolicyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if objectReplicationPolicyID == "" {
		return nil, errors.New("parameter objectReplicationPolicyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{objectReplicationPolicyId}", url.PathEscape(objectReplicationPolicyID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ObjectReplicationPoliciesClient) getHandleResponse(resp *http.Response) (ObjectReplicationPoliciesClientGetResponse, error) {
	result := ObjectReplicationPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectReplicationPolicy); err != nil {
		return ObjectReplicationPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the object replication policies associated with the storage account.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - options - ObjectReplicationPoliciesClientListOptions contains the optional parameters for the ObjectReplicationPoliciesClient.NewListPager
//     method.
func (client *ObjectReplicationPoliciesClient) NewListPager(resourceGroupName string, accountName string, options *ObjectReplicationPoliciesClientListOptions) *runtime.Pager[ObjectReplicationPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ObjectReplicationPoliciesClientListResponse]{
		More: func(page ObjectReplicationPoliciesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ObjectReplicationPoliciesClientListResponse) (ObjectReplicationPoliciesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ObjectReplicationPoliciesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			if err != nil {
				return ObjectReplicationPoliciesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ObjectReplicationPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ObjectReplicationPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ObjectReplicationPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, _ *ObjectReplicationPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ObjectReplicationPoliciesClient) listHandleResponse(resp *http.Response) (ObjectReplicationPoliciesClientListResponse, error) {
	result := ObjectReplicationPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectReplicationPolicies); err != nil {
		return ObjectReplicationPoliciesClientListResponse{}, err
	}
	return result, nil
}
