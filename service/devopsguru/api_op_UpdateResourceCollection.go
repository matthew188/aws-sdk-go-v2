// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the collection of resources that DevOps Guru analyzes. The two types of
// Amazon Web Services resource collections supported are Amazon Web Services
// CloudFormation stacks and Amazon Web Services resources that contain the same
// Amazon Web Services tag. DevOps Guru can be configured to analyze the Amazon Web
// Services resources that are defined in the stacks or that are tagged using the
// same tag key. You can specify up to 500 Amazon Web Services CloudFormation
// stacks. This method also creates the IAM role required for you to use DevOps
// Guru.
func (c *Client) UpdateResourceCollection(ctx context.Context, params *UpdateResourceCollectionInput, optFns ...func(*Options)) (*UpdateResourceCollectionOutput, error) {
	if params == nil {
		params = &UpdateResourceCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResourceCollection", params, optFns, c.addOperationUpdateResourceCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResourceCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResourceCollectionInput struct {

	// Specifies if the resource collection in the request is added or deleted to the
	// resource collection.
	//
	// This member is required.
	Action types.UpdateResourceCollectionAction

	// Contains information used to update a collection of Amazon Web Services
	// resources.
	//
	// This member is required.
	ResourceCollection *types.UpdateResourceCollectionFilter

	noSmithyDocumentSerde
}

type UpdateResourceCollectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResourceCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateResourceCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateResourceCollection{}, middleware.After)
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
	if err = addOpUpdateResourceCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResourceCollection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateResourceCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "UpdateResourceCollection",
	}
}
