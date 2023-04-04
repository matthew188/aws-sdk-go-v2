// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubstrategy

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/migrationhubstrategy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves details about an application component.
func (c *Client) GetApplicationComponentDetails(ctx context.Context, params *GetApplicationComponentDetailsInput, optFns ...func(*Options)) (*GetApplicationComponentDetailsOutput, error) {
	if params == nil {
		params = &GetApplicationComponentDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApplicationComponentDetails", params, optFns, c.addOperationGetApplicationComponentDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApplicationComponentDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApplicationComponentDetailsInput struct {

	// The ID of the application component. The ID is unique within an AWS account.
	//
	// This member is required.
	ApplicationComponentId *string

	noSmithyDocumentSerde
}

type GetApplicationComponentDetailsOutput struct {

	// Detailed information about an application component.
	ApplicationComponentDetail *types.ApplicationComponentDetail

	// The associated application group as defined in AWS Application Discovery
	// Service.
	AssociatedApplications []types.AssociatedApplication

	// A list of the IDs of the servers on which the application component is running.
	AssociatedServerIds []string

	// Set to true if the application component belongs to more than one application
	// group.
	MoreApplicationResource *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetApplicationComponentDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetApplicationComponentDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApplicationComponentDetails{}, middleware.After)
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
	if err = addOpGetApplicationComponentDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetApplicationComponentDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetApplicationComponentDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-strategy",
		OperationName: "GetApplicationComponentDetails",
	}
}
