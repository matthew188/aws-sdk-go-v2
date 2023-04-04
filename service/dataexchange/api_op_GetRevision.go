// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This operation returns information about a revision.
func (c *Client) GetRevision(ctx context.Context, params *GetRevisionInput, optFns ...func(*Options)) (*GetRevisionOutput, error) {
	if params == nil {
		params = &GetRevisionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRevision", params, optFns, c.addOperationGetRevisionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRevisionInput struct {

	// The unique identifier for a data set.
	//
	// This member is required.
	DataSetId *string

	// The unique identifier for a revision.
	//
	// This member is required.
	RevisionId *string

	noSmithyDocumentSerde
}

type GetRevisionOutput struct {

	// The ARN for the revision.
	Arn *string

	// An optional comment about the revision.
	Comment *string

	// The date and time that the revision was created, in ISO 8601 format.
	CreatedAt *time.Time

	// The unique identifier for the data set associated with the data set revision.
	DataSetId *string

	// To publish a revision to a data set in a product, the revision must first be
	// finalized. Finalizing a revision tells AWS Data Exchange that your changes to
	// the assets in the revision are complete. After it's in this read-only state, you
	// can publish the revision to your products. Finalized revisions can be published
	// through the AWS Data Exchange console or the AWS Marketplace Catalog API, using
	// the StartChangeSet AWS Marketplace Catalog API action. When using the API,
	// revisions are uniquely identified by their ARN.
	Finalized bool

	// The unique identifier for the revision.
	Id *string

	// A required comment to inform subscribers of the reason their access to the
	// revision was revoked.
	RevocationComment *string

	// A status indicating that subscribers' access to the revision was revoked.
	Revoked bool

	// The date and time that the revision was revoked, in ISO 8601 format.
	RevokedAt *time.Time

	// The revision ID of the owned revision corresponding to the entitled revision
	// being viewed. This parameter is returned when a revision owner is viewing the
	// entitled copy of its owned revision.
	SourceId *string

	// The tags for the revision.
	Tags map[string]string

	// The date and time that the revision was last updated, in ISO 8601 format.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRevisionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRevision{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRevision{}, middleware.After)
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
	if err = addOpGetRevisionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRevision(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRevision(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "GetRevision",
	}
}
