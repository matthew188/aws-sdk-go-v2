// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Instructs the specified agents or connectors to stop collecting data.
func (c *Client) StopDataCollectionByAgentIds(ctx context.Context, params *StopDataCollectionByAgentIdsInput, optFns ...func(*Options)) (*StopDataCollectionByAgentIdsOutput, error) {
	if params == nil {
		params = &StopDataCollectionByAgentIdsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDataCollectionByAgentIds", params, optFns, c.addOperationStopDataCollectionByAgentIdsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDataCollectionByAgentIdsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDataCollectionByAgentIdsInput struct {

	// The IDs of the agents or connectors from which to stop collecting data.
	//
	// This member is required.
	AgentIds []string

	noSmithyDocumentSerde
}

type StopDataCollectionByAgentIdsOutput struct {

	// Information about the agents or connector that were instructed to stop
	// collecting data. Information includes the agent/connector ID, a description of
	// the operation performed, and whether the agent/connector configuration was
	// updated.
	AgentsConfigurationStatus []types.AgentConfigurationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopDataCollectionByAgentIdsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopDataCollectionByAgentIds{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopDataCollectionByAgentIds{}, middleware.After)
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
	if err = addOpStopDataCollectionByAgentIdsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopDataCollectionByAgentIds(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopDataCollectionByAgentIds(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "StopDataCollectionByAgentIds",
	}
}
