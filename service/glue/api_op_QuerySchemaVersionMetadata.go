// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Queries for the schema version metadata information.
func (c *Client) QuerySchemaVersionMetadata(ctx context.Context, params *QuerySchemaVersionMetadataInput, optFns ...func(*Options)) (*QuerySchemaVersionMetadataOutput, error) {
	if params == nil {
		params = &QuerySchemaVersionMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "QuerySchemaVersionMetadata", params, optFns, c.addOperationQuerySchemaVersionMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QuerySchemaVersionMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QuerySchemaVersionMetadataInput struct {

	// Maximum number of results required per page. If the value is not supplied, this
	// will be defaulted to 25 per page.
	MaxResults int32

	// Search key-value pairs for metadata, if they are not provided all the metadata
	// information will be fetched.
	MetadataList []types.MetadataKeyValuePair

	// A continuation token, if this is a continuation call.
	NextToken *string

	// A wrapper structure that may contain the schema name and Amazon Resource Name
	// (ARN).
	SchemaId *types.SchemaId

	// The unique version ID of the schema version.
	SchemaVersionId *string

	// The version number of the schema.
	SchemaVersionNumber *types.SchemaVersionNumber

	noSmithyDocumentSerde
}

type QuerySchemaVersionMetadataOutput struct {

	// A map of a metadata key and associated values.
	MetadataInfoMap map[string]types.MetadataInfo

	// A continuation token for paginating the returned list of tokens, returned if the
	// current segment of the list is not the last.
	NextToken *string

	// The unique version ID of the schema version.
	SchemaVersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationQuerySchemaVersionMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpQuerySchemaVersionMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpQuerySchemaVersionMetadata{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQuerySchemaVersionMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opQuerySchemaVersionMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "QuerySchemaVersionMetadata",
	}
}
