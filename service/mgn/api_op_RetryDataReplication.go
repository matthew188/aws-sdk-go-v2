// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Causes the data replication initiation sequence to begin immediately upon next
// Handshake for specified SourceServer IDs, regardless of when the previous
// initiation started. This command will not work if the SourceServer is not
// stalled or is in a DISCONNECTED or STOPPED state.
func (c *Client) RetryDataReplication(ctx context.Context, params *RetryDataReplicationInput, optFns ...func(*Options)) (*RetryDataReplicationOutput, error) {
	if params == nil {
		params = &RetryDataReplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RetryDataReplication", params, optFns, c.addOperationRetryDataReplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetryDataReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RetryDataReplicationInput struct {

	// Retry data replication for Source Server ID.
	//
	// This member is required.
	SourceServerID *string

	noSmithyDocumentSerde
}

type RetryDataReplicationOutput struct {

	// Source server application ID.
	ApplicationID *string

	// Source server ARN.
	Arn *string

	// Source server data replication info.
	DataReplicationInfo *types.DataReplicationInfo

	// Source server fqdn for action framework.
	FqdnForActionFramework *string

	// Source server archived status.
	IsArchived *bool

	// Source server launched instance.
	LaunchedInstance *types.LaunchedInstance

	// Source server lifecycle state.
	LifeCycle *types.LifeCycle

	// Source server replication type.
	ReplicationType types.ReplicationType

	// Source server properties.
	SourceProperties *types.SourceProperties

	// Source server ID.
	SourceServerID *string

	// Source server Tags.
	Tags map[string]string

	// Source server user provided ID.
	UserProvidedID *string

	// Source server vCenter client id.
	VcenterClientID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRetryDataReplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRetryDataReplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRetryDataReplication{}, middleware.After)
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
	if err = addOpRetryDataReplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRetryDataReplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRetryDataReplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "RetryDataReplication",
	}
}
