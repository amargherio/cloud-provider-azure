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
	"strconv"
	"strings"
)

// TaskAssignmentInstancesReportClient contains the methods for the StorageTaskAssignmentInstancesReport group.
// Don't use this type directly, use NewTaskAssignmentInstancesReportClient() instead.
type TaskAssignmentInstancesReportClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTaskAssignmentInstancesReportClient creates a new instance of TaskAssignmentInstancesReportClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTaskAssignmentInstancesReportClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TaskAssignmentInstancesReportClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TaskAssignmentInstancesReportClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Fetch the report summary of a single storage task assignment's instances
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - storageTaskAssignmentName - The name of the storage task assignment within the specified resource group. Storage task assignment
//     names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
//   - options - TaskAssignmentInstancesReportClientListOptions contains the optional parameters for the TaskAssignmentInstancesReportClient.NewListPager
//     method.
func (client *TaskAssignmentInstancesReportClient) NewListPager(resourceGroupName string, accountName string, storageTaskAssignmentName string, options *TaskAssignmentInstancesReportClientListOptions) *runtime.Pager[TaskAssignmentInstancesReportClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TaskAssignmentInstancesReportClientListResponse]{
		More: func(page TaskAssignmentInstancesReportClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TaskAssignmentInstancesReportClientListResponse) (TaskAssignmentInstancesReportClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TaskAssignmentInstancesReportClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, accountName, storageTaskAssignmentName, options)
			}, nil)
			if err != nil {
				return TaskAssignmentInstancesReportClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TaskAssignmentInstancesReportClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, options *TaskAssignmentInstancesReportClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/storageTaskAssignments/{storageTaskAssignmentName}/reports"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageTaskAssignmentName == "" {
		return nil, errors.New("parameter storageTaskAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTaskAssignmentName}", url.PathEscape(storageTaskAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("$maxpagesize", strconv.FormatInt(int64(*options.Maxpagesize), 10))
	}
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TaskAssignmentInstancesReportClient) listHandleResponse(resp *http.Response) (TaskAssignmentInstancesReportClientListResponse, error) {
	result := TaskAssignmentInstancesReportClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskReportSummary); err != nil {
		return TaskAssignmentInstancesReportClientListResponse{}, err
	}
	return result, nil
}
