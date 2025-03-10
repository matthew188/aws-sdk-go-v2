// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Stops a specified import.
func (c *Client) StopImport(ctx context.Context, params *StopImportInput, optFns ...func(*Options)) (*StopImportOutput, error) {
	if params == nil {
		params = &StopImportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopImport", params, optFns, c.addOperationStopImportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopImportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopImportInput struct {

	// The ID of the import.
	//
	// This member is required.
	ImportId *string

	noSmithyDocumentSerde
}

type StopImportOutput struct {

	// The timestamp of the import's creation.
	CreatedTimestamp *time.Time

	// The ARN of the destination event data store.
	Destinations []string

	// Used with StartEventTime to bound a StartImport request, and limit imported
	// trail events to only those events logged within a specified time period.
	EndEventTime *time.Time

	// The ID for the import.
	ImportId *string

	// The source S3 bucket for the import.
	ImportSource *types.ImportSource

	// Returns information on the stopped import.
	ImportStatistics *types.ImportStatistics

	// The status of the import.
	ImportStatus types.ImportStatus

	// Used with EndEventTime to bound a StartImport request, and limit imported trail
	// events to only those events logged within a specified time period.
	StartEventTime *time.Time

	// The timestamp of the import's last update.
	UpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopImportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopImport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopImport{}, middleware.After)
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
	if err = addOpStopImportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopImport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopImport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "StopImport",
	}
}
