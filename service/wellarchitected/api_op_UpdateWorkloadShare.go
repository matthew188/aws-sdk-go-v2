// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a workload share.
func (c *Client) UpdateWorkloadShare(ctx context.Context, params *UpdateWorkloadShareInput, optFns ...func(*Options)) (*UpdateWorkloadShareOutput, error) {
	if params == nil {
		params = &UpdateWorkloadShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkloadShare", params, optFns, c.addOperationUpdateWorkloadShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkloadShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for Update Workload Share
type UpdateWorkloadShareInput struct {

	// Permission granted on a workload share.
	//
	// This member is required.
	PermissionType types.PermissionType

	// The ID associated with the workload share.
	//
	// This member is required.
	ShareId *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web Services
	// Region.
	//
	// This member is required.
	WorkloadId *string

	noSmithyDocumentSerde
}

// Input for Update Workload Share
type UpdateWorkloadShareOutput struct {

	// The ID assigned to the workload. This ID is unique within an Amazon Web Services
	// Region.
	WorkloadId *string

	// A workload share return object.
	WorkloadShare *types.WorkloadShare

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWorkloadShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateWorkloadShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateWorkloadShare{}, middleware.After)
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
	if err = addOpUpdateWorkloadShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkloadShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWorkloadShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "UpdateWorkloadShare",
	}
}
