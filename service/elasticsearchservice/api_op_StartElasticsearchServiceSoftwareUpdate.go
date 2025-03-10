// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Schedules a service software update for an Amazon ES domain.
func (c *Client) StartElasticsearchServiceSoftwareUpdate(ctx context.Context, params *StartElasticsearchServiceSoftwareUpdateInput, optFns ...func(*Options)) (*StartElasticsearchServiceSoftwareUpdateOutput, error) {
	if params == nil {
		params = &StartElasticsearchServiceSoftwareUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartElasticsearchServiceSoftwareUpdate", params, optFns, c.addOperationStartElasticsearchServiceSoftwareUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartElasticsearchServiceSoftwareUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the StartElasticsearchServiceSoftwareUpdate
// operation. Specifies the name of the Elasticsearch domain that you wish to
// schedule a service software update on.
type StartElasticsearchServiceSoftwareUpdateInput struct {

	// The name of the domain that you want to update to the latest service software.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

// The result of a StartElasticsearchServiceSoftwareUpdate operation. Contains the
// status of the update.
type StartElasticsearchServiceSoftwareUpdateOutput struct {

	// The current status of the Elasticsearch service software update.
	ServiceSoftwareOptions *types.ServiceSoftwareOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartElasticsearchServiceSoftwareUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartElasticsearchServiceSoftwareUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartElasticsearchServiceSoftwareUpdate{}, middleware.After)
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
	if err = addOpStartElasticsearchServiceSoftwareUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartElasticsearchServiceSoftwareUpdate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartElasticsearchServiceSoftwareUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "StartElasticsearchServiceSoftwareUpdate",
	}
}
