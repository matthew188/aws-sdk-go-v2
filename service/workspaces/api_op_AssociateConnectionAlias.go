// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates the specified connection alias with the specified directory to enable
// cross-Region redirection. For more information, see  Cross-Region Redirection
// for Amazon WorkSpaces
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html).
// Before performing this operation, call  DescribeConnectionAliases
// (https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeConnectionAliases.html)
// to make sure that the current state of the connection alias is CREATED.
func (c *Client) AssociateConnectionAlias(ctx context.Context, params *AssociateConnectionAliasInput, optFns ...func(*Options)) (*AssociateConnectionAliasOutput, error) {
	if params == nil {
		params = &AssociateConnectionAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateConnectionAlias", params, optFns, c.addOperationAssociateConnectionAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateConnectionAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateConnectionAliasInput struct {

	// The identifier of the connection alias.
	//
	// This member is required.
	AliasId *string

	// The identifier of the directory to associate the connection alias with.
	//
	// This member is required.
	ResourceId *string

	noSmithyDocumentSerde
}

type AssociateConnectionAliasOutput struct {

	// The identifier of the connection alias association. You use the connection
	// identifier in the DNS TXT record when you're configuring your DNS routing
	// policies.
	ConnectionIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateConnectionAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateConnectionAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateConnectionAlias{}, middleware.After)
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
	if err = addOpAssociateConnectionAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateConnectionAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateConnectionAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "AssociateConnectionAlias",
	}
}
