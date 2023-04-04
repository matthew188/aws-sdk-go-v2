// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This example serializes intEnums as top level properties, in lists, sets, and
// maps.
func (c *Client) JsonIntEnums(ctx context.Context, params *JsonIntEnumsInput, optFns ...func(*Options)) (*JsonIntEnumsOutput, error) {
	if params == nil {
		params = &JsonIntEnumsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "JsonIntEnums", params, optFns, c.addOperationJsonIntEnumsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*JsonIntEnumsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type JsonIntEnumsInput struct {
	IntegerEnum1 types.IntegerEnum

	IntegerEnum2 types.IntegerEnum

	IntegerEnum3 types.IntegerEnum

	IntegerEnumList []types.IntegerEnum

	IntegerEnumMap map[string]types.IntegerEnum

	IntegerEnumSet []types.IntegerEnum

	noSmithyDocumentSerde
}

type JsonIntEnumsOutput struct {
	IntegerEnum1 types.IntegerEnum

	IntegerEnum2 types.IntegerEnum

	IntegerEnum3 types.IntegerEnum

	IntegerEnumList []types.IntegerEnum

	IntegerEnumMap map[string]types.IntegerEnum

	IntegerEnumSet []types.IntegerEnum

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationJsonIntEnumsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpJsonIntEnums{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpJsonIntEnums{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opJsonIntEnums(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opJsonIntEnums(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "JsonIntEnums",
	}
}
