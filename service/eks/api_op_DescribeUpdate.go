// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns descriptive information about an update against your Amazon EKS cluster
// or associated managed node group or Amazon EKS add-on. When the status of the
// update is Succeeded, the update is complete. If an update fails, the status is
// Failed, and an error detail explains the reason for the failure.
func (c *Client) DescribeUpdate(ctx context.Context, params *DescribeUpdateInput, optFns ...func(*Options)) (*DescribeUpdateOutput, error) {
	if params == nil {
		params = &DescribeUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUpdate", params, optFns, c.addOperationDescribeUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUpdateInput struct {

	// The name of the Amazon EKS cluster associated with the update.
	//
	// This member is required.
	Name *string

	// The ID of the update to describe.
	//
	// This member is required.
	UpdateId *string

	// The name of the add-on. The name must match one of the names returned by
	// ListAddons
	// (https://docs.aws.amazon.com/eks/latest/APIReference/API_ListAddons.html). This
	// parameter is required if the update is an add-on update.
	AddonName *string

	// The name of the Amazon EKS node group associated with the update. This parameter
	// is required if the update is a node group update.
	NodegroupName *string

	noSmithyDocumentSerde
}

type DescribeUpdateOutput struct {

	// The full description of the specified update.
	Update *types.Update

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeUpdate{}, middleware.After)
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
	if err = addOpDescribeUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUpdate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "DescribeUpdate",
	}
}
