// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of available log streams for a specific database in Amazon
// Lightsail.
func (c *Client) GetRelationalDatabaseLogStreams(ctx context.Context, params *GetRelationalDatabaseLogStreamsInput, optFns ...func(*Options)) (*GetRelationalDatabaseLogStreamsOutput, error) {
	if params == nil {
		params = &GetRelationalDatabaseLogStreamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabaseLogStreams", params, optFns, c.addOperationGetRelationalDatabaseLogStreamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabaseLogStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseLogStreamsInput struct {

	// The name of your database for which to get log streams.
	//
	// This member is required.
	RelationalDatabaseName *string

	noSmithyDocumentSerde
}

type GetRelationalDatabaseLogStreamsOutput struct {

	// An object describing the result of your get relational database log streams
	// request.
	LogStreams []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRelationalDatabaseLogStreamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseLogStreams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseLogStreams{}, middleware.After)
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
	if err = addOpGetRelationalDatabaseLogStreamsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseLogStreams(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRelationalDatabaseLogStreams(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseLogStreams",
	}
}
