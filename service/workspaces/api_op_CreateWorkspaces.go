// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates one or more WorkSpaces. This operation is asynchronous and returns
// before the WorkSpaces are created. The MANUAL running mode value is only
// supported by Amazon WorkSpaces Core. Contact your account team to be
// allow-listed to use this value. For more information, see Amazon WorkSpaces Core
// (http://aws.amazon.com/workspaces/core/).
func (c *Client) CreateWorkspaces(ctx context.Context, params *CreateWorkspacesInput, optFns ...func(*Options)) (*CreateWorkspacesOutput, error) {
	if params == nil {
		params = &CreateWorkspacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkspaces", params, optFns, c.addOperationCreateWorkspacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkspacesInput struct {

	// The WorkSpaces to create. You can specify up to 25 WorkSpaces.
	//
	// This member is required.
	Workspaces []types.WorkspaceRequest

	noSmithyDocumentSerde
}

type CreateWorkspacesOutput struct {

	// Information about the WorkSpaces that could not be created.
	FailedRequests []types.FailedCreateWorkspaceRequest

	// Information about the WorkSpaces that were created. Because this operation is
	// asynchronous, the identifier returned is not immediately available for use with
	// other operations. For example, if you call DescribeWorkspaces before the
	// WorkSpace is created, the information returned can be incomplete.
	PendingRequests []types.Workspace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkspacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWorkspaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWorkspaces{}, middleware.After)
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
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateWorkspacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkspaces(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkspaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "CreateWorkspaces",
	}
}
