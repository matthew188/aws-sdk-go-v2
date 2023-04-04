// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the given custom connector registered in your Amazon Web Services
// account. This API can be used for custom connectors that are registered in your
// account and also for Amazon authored connectors.
func (c *Client) DescribeConnector(ctx context.Context, params *DescribeConnectorInput, optFns ...func(*Options)) (*DescribeConnectorOutput, error) {
	if params == nil {
		params = &DescribeConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnector", params, optFns, c.addOperationDescribeConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectorInput struct {

	// The connector type, such as CUSTOMCONNECTOR, Saleforce, Marketo. Please choose
	// CUSTOMCONNECTOR for Lambda based custom connectors.
	//
	// This member is required.
	ConnectorType types.ConnectorType

	// The label of the connector. The label is unique for each ConnectorRegistration
	// in your Amazon Web Services account. Only needed if calling for CUSTOMCONNECTOR
	// connector type/.
	ConnectorLabel *string

	noSmithyDocumentSerde
}

type DescribeConnectorOutput struct {

	// Configuration info of all the connectors that the user requested.
	ConnectorConfiguration *types.ConnectorConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConnector{}, middleware.After)
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
	if err = addOpDescribeConnectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "DescribeConnector",
	}
}
