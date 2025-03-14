// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Amazon Kendra data source connector.
func (c *Client) UpdateDataSource(ctx context.Context, params *UpdateDataSourceInput, optFns ...func(*Options)) (*UpdateDataSourceOutput, error) {
	if params == nil {
		params = &UpdateDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataSource", params, optFns, c.addOperationUpdateDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSourceInput struct {

	// The identifier of the data source connector you want to update.
	//
	// This member is required.
	Id *string

	// The identifier of the index used with the data source connector.
	//
	// This member is required.
	IndexId *string

	// Configuration information you want to update for the data source connector.
	Configuration *types.DataSourceConfiguration

	// Configuration information you want to update for altering document metadata and
	// content during the document ingestion process. For more information on how to
	// create, modify and delete document metadata, or make other content alterations
	// when you ingest documents into Amazon Kendra, see Customizing document metadata
	// during the ingestion process
	// (https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html).
	CustomDocumentEnrichmentConfiguration *types.CustomDocumentEnrichmentConfiguration

	// A new description for the data source connector.
	Description *string

	// The code for a language you want to update for the data source connector. This
	// allows you to support a language for all documents when updating the data
	// source. English is supported by default. For more information on supported
	// languages, including their codes, see Adding documents in languages other than
	// English (https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html).
	LanguageCode *string

	// A new name for the data source connector.
	Name *string

	// The Amazon Resource Name (ARN) of a role with permission to access the data
	// source and required resources. For more information, see IAM roles for Amazon
	// Kendra (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
	RoleArn *string

	// The sync schedule you want to update for the data source connector.
	Schedule *string

	// Configuration information for an Amazon Virtual Private Cloud to connect to your
	// data source. For more information, see Configuring a VPC
	// (https://docs.aws.amazon.com/kendra/latest/dg/vpc-configuration.html).
	VpcConfiguration *types.DataSourceVpcConfiguration

	noSmithyDocumentSerde
}

type UpdateDataSourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDataSource{}, middleware.After)
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
	if err = addOpUpdateDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "UpdateDataSource",
	}
}
