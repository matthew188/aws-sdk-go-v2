// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a dashboard.
func (c *Client) DeleteDashboard(ctx context.Context, params *DeleteDashboardInput, optFns ...func(*Options)) (*DeleteDashboardOutput, error) {
	if params == nil {
		params = &DeleteDashboardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDashboard", params, optFns, c.addOperationDeleteDashboardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDashboardInput struct {

	// The ID of the Amazon Web Services account that contains the dashboard that
	// you're deleting.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dashboard.
	//
	// This member is required.
	DashboardId *string

	// The version number of the dashboard. If the version number property is provided,
	// only the specified version of the dashboard is deleted.
	VersionNumber *int64

	noSmithyDocumentSerde
}

type DeleteDashboardOutput struct {

	// The Secure Socket Layer (SSL) properties that apply for the resource.
	Arn *string

	// The ID of the dashboard.
	DashboardId *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDashboardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDashboard{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDashboard{}, middleware.After)
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
	if err = addOpDeleteDashboardValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDashboard(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDashboard(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteDashboard",
	}
}
