// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a natively supported Amazon Web Service as an Amazon Security Lake source.
// Enables source types for member accounts in required Amazon Web Services
// Regions, based on the parameters you specify. You can choose any source type in
// any Region for either accounts that are part of a trusted organization or
// standalone accounts. At least one of the three dimensions is a mandatory input
// to this API. However, you can supply any combination of the three dimensions to
// this API. By default, a dimension refers to the entire set. When you don't
// provide a dimension, Security Lake assumes that the missing dimension refers to
// the entire set. This is overridden when you supply any one of the inputs. For
// instance, when you do not specify members, the API enables all Security Lake
// member accounts for all sources. Similarly, when you do not specify Regions,
// Security Lake is enabled for all the Regions where Security Lake is available as
// a service. You can use this API only to enable natively supported Amazon Web
// Services as a source. Use CreateCustomLogSource to enable data collection from a
// custom source.
func (c *Client) CreateAwsLogSource(ctx context.Context, params *CreateAwsLogSourceInput, optFns ...func(*Options)) (*CreateAwsLogSourceOutput, error) {
	if params == nil {
		params = &CreateAwsLogSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAwsLogSource", params, optFns, c.addOperationCreateAwsLogSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAwsLogSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAwsLogSourceInput struct {

	// Specifies the input order to enable dimensions in Security Lake, namely Region,
	// source type, and member account.
	//
	// This member is required.
	InputOrder []types.Dimension

	// Enables data collection from specific Amazon Web Services sources in all
	// specific accounts and specific Regions.
	EnableAllDimensions map[string]map[string][]string

	// Enables data collection from all Amazon Web Services sources in specific
	// accounts or Regions.
	EnableSingleDimension []string

	// Enables data collection from specific Amazon Web Services sources in specific
	// accounts or Regions.
	EnableTwoDimensions map[string][]string

	noSmithyDocumentSerde
}

type CreateAwsLogSourceOutput struct {

	// Lists all accounts in which enabling a natively supported Amazon Web Service as
	// a Security Lake source failed. The failure occurred as these accounts are not
	// part of an organization.
	Failed []string

	// Lists the accounts that are in the process of enabling a natively supported
	// Amazon Web Service as a Security Lake source.
	Processing []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAwsLogSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAwsLogSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAwsLogSource{}, middleware.After)
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
	if err = addOpCreateAwsLogSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAwsLogSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAwsLogSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateAwsLogSource",
	}
}
