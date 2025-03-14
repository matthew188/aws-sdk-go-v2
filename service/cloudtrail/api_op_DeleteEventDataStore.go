// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables the event data store specified by EventDataStore, which accepts an
// event data store ARN. After you run DeleteEventDataStore, the event data store
// enters a PENDING_DELETION state, and is automatically deleted after a wait
// period of seven days. TerminationProtectionEnabled must be set to False on the
// event data store; this operation cannot work if TerminationProtectionEnabled is
// True. After you run DeleteEventDataStore on an event data store, you cannot run
// ListQueries, DescribeQuery, or GetQueryResults on queries that are using an
// event data store in a PENDING_DELETION state. An event data store in the
// PENDING_DELETION state does not incur costs.
func (c *Client) DeleteEventDataStore(ctx context.Context, params *DeleteEventDataStoreInput, optFns ...func(*Options)) (*DeleteEventDataStoreOutput, error) {
	if params == nil {
		params = &DeleteEventDataStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEventDataStore", params, optFns, c.addOperationDeleteEventDataStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEventDataStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEventDataStoreInput struct {

	// The ARN (or the ID suffix of the ARN) of the event data store to delete.
	//
	// This member is required.
	EventDataStore *string

	noSmithyDocumentSerde
}

type DeleteEventDataStoreOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteEventDataStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteEventDataStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteEventDataStore{}, middleware.After)
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
	if err = addOpDeleteEventDataStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEventDataStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteEventDataStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "DeleteEventDataStore",
	}
}
