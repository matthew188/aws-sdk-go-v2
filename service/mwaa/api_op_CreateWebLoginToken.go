// Code generated by smithy-go-codegen DO NOT EDIT.

package mwaa

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a web login token for the Airflow Web UI. To learn more, see Creating an
// Apache Airflow web login token
// (https://docs.aws.amazon.com/mwaa/latest/userguide/call-mwaa-apis-web.html).
func (c *Client) CreateWebLoginToken(ctx context.Context, params *CreateWebLoginTokenInput, optFns ...func(*Options)) (*CreateWebLoginTokenOutput, error) {
	if params == nil {
		params = &CreateWebLoginTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWebLoginToken", params, optFns, c.addOperationCreateWebLoginTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWebLoginTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWebLoginTokenInput struct {

	// The name of the Amazon MWAA environment. For example, MyMWAAEnvironment.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type CreateWebLoginTokenOutput struct {

	// The Airflow web server hostname for the environment.
	WebServerHostname *string

	// An Airflow web server login token.
	WebToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWebLoginTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWebLoginToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWebLoginToken{}, middleware.After)
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
	if err = addEndpointPrefix_opCreateWebLoginTokenMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateWebLoginTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWebLoginToken(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateWebLoginTokenMiddleware struct {
}

func (*endpointPrefix_opCreateWebLoginTokenMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateWebLoginTokenMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "env." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateWebLoginTokenMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateWebLoginTokenMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCreateWebLoginToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "airflow",
		OperationName: "CreateWebLoginToken",
	}
}
