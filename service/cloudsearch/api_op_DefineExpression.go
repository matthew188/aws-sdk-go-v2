// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures an Expression for the search domain. Used to create new expressions
// and modify existing ones. If the expression exists, the new configuration
// replaces the old one. For more information, see Configuring Expressions
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-expressions.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DefineExpression(ctx context.Context, params *DefineExpressionInput, optFns ...func(*Options)) (*DefineExpressionOutput, error) {
	if params == nil {
		params = &DefineExpressionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DefineExpression", params, optFns, c.addOperationDefineExpressionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DefineExpressionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DefineExpression operation. Specifies the
// name of the domain you want to update and the expression you want to configure.
type DefineExpressionInput struct {

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string

	// A named expression that can be evaluated at search time. Can be used to sort the
	// search results, define other expressions, or return computed information in the
	// search results.
	//
	// This member is required.
	Expression *types.Expression

	noSmithyDocumentSerde
}

// The result of a DefineExpression request. Contains the status of the
// newly-configured expression.
type DefineExpressionOutput struct {

	// The value of an Expression and its current status.
	//
	// This member is required.
	Expression *types.ExpressionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDefineExpressionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDefineExpression{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDefineExpression{}, middleware.After)
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
	if err = addOpDefineExpressionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDefineExpression(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDefineExpression(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DefineExpression",
	}
}
