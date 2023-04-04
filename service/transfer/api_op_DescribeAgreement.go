// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the agreement that's identified by the AgreementId.
func (c *Client) DescribeAgreement(ctx context.Context, params *DescribeAgreementInput, optFns ...func(*Options)) (*DescribeAgreementOutput, error) {
	if params == nil {
		params = &DescribeAgreementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAgreement", params, optFns, c.addOperationDescribeAgreementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAgreementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAgreementInput struct {

	// A unique identifier for the agreement. This identifier is returned when you
	// create an agreement.
	//
	// This member is required.
	AgreementId *string

	// The server identifier that's associated with the agreement.
	//
	// This member is required.
	ServerId *string

	noSmithyDocumentSerde
}

type DescribeAgreementOutput struct {

	// The details for the specified agreement, returned as a DescribedAgreement
	// object.
	//
	// This member is required.
	Agreement *types.DescribedAgreement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAgreementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAgreement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAgreement{}, middleware.After)
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
	if err = addOpDescribeAgreementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAgreement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAgreement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "DescribeAgreement",
	}
}
