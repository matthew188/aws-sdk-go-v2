// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacemetering

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// ResolveCustomer is called by a SaaS application during the registration process.
// When a buyer visits your website during the registration process, the buyer
// submits a registration token through their browser. The registration token is
// resolved through this API to obtain a CustomerIdentifier along with the
// CustomerAWSAccountId and ProductCode. The API needs to called from the seller
// account id used to publish the SaaS application to successfully resolve the
// token. For an example of using ResolveCustomer, see  ResolveCustomer code
// example
// (https://docs.aws.amazon.com/marketplace/latest/userguide/saas-code-examples.html#saas-resolvecustomer-example)
// in the AWS Marketplace Seller Guide.
func (c *Client) ResolveCustomer(ctx context.Context, params *ResolveCustomerInput, optFns ...func(*Options)) (*ResolveCustomerOutput, error) {
	if params == nil {
		params = &ResolveCustomerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResolveCustomer", params, optFns, c.addOperationResolveCustomerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResolveCustomerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains input to the ResolveCustomer operation.
type ResolveCustomerInput struct {

	// When a buyer visits your website during the registration process, the buyer
	// submits a registration token through the browser. The registration token is
	// resolved to obtain a CustomerIdentifier along with the CustomerAWSAccountId and
	// ProductCode.
	//
	// This member is required.
	RegistrationToken *string

	noSmithyDocumentSerde
}

// The result of the ResolveCustomer operation. Contains the CustomerIdentifier
// along with the CustomerAWSAccountId and ProductCode.
type ResolveCustomerOutput struct {

	// The CustomerAWSAccountId provides the AWS account ID associated with the
	// CustomerIdentifier for the individual customer.
	CustomerAWSAccountId *string

	// The CustomerIdentifier is used to identify an individual customer in your
	// application. Calls to BatchMeterUsage require CustomerIdentifiers for each
	// UsageRecord.
	CustomerIdentifier *string

	// The product code is returned to confirm that the buyer is registering for your
	// product. Subsequent BatchMeterUsage calls should be made using this product
	// code.
	ProductCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResolveCustomerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResolveCustomer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResolveCustomer{}, middleware.After)
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
	if err = addOpResolveCustomerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResolveCustomer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResolveCustomer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "ResolveCustomer",
	}
}
