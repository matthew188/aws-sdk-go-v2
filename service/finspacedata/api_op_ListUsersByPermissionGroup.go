// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists details of all the users in a specific permission group.
func (c *Client) ListUsersByPermissionGroup(ctx context.Context, params *ListUsersByPermissionGroupInput, optFns ...func(*Options)) (*ListUsersByPermissionGroupOutput, error) {
	if params == nil {
		params = &ListUsersByPermissionGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUsersByPermissionGroup", params, optFns, c.addOperationListUsersByPermissionGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUsersByPermissionGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUsersByPermissionGroupInput struct {

	// The maximum number of results per page.
	//
	// This member is required.
	MaxResults int32

	// The unique identifier for the permission group.
	//
	// This member is required.
	PermissionGroupId *string

	// A token that indicates where a results page should begin.
	NextToken *string

	noSmithyDocumentSerde
}

type ListUsersByPermissionGroupOutput struct {

	// A token that indicates where a results page should begin.
	NextToken *string

	// Lists details of all users in a specific permission group.
	Users []types.UserByPermissionGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUsersByPermissionGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUsersByPermissionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUsersByPermissionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListUsersByPermissionGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUsersByPermissionGroup(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListUsersByPermissionGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "ListUsersByPermissionGroup",
	}
}
