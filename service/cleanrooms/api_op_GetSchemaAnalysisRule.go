// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a schema analysis rule.
func (c *Client) GetSchemaAnalysisRule(ctx context.Context, params *GetSchemaAnalysisRuleInput, optFns ...func(*Options)) (*GetSchemaAnalysisRuleOutput, error) {
	if params == nil {
		params = &GetSchemaAnalysisRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSchemaAnalysisRule", params, optFns, c.addOperationGetSchemaAnalysisRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSchemaAnalysisRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSchemaAnalysisRuleInput struct {

	// A unique identifier for the collaboration that the schema belongs to. Currently
	// accepts a collaboration ID.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The name of the schema to retrieve the analysis rule for.
	//
	// This member is required.
	Name *string

	// The type of the schema analysis rule to retrieve. Schema analysis rules are
	// uniquely identified by a combination of the collaboration, the schema name, and
	// their type.
	//
	// This member is required.
	Type types.AnalysisRuleType

	noSmithyDocumentSerde
}

type GetSchemaAnalysisRuleOutput struct {

	// A specification about how data from the configured table can be used.
	//
	// This member is required.
	AnalysisRule *types.AnalysisRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSchemaAnalysisRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSchemaAnalysisRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSchemaAnalysisRule{}, middleware.After)
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
	if err = addOpGetSchemaAnalysisRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSchemaAnalysisRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSchemaAnalysisRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cleanrooms",
		OperationName: "GetSchemaAnalysisRule",
	}
}
