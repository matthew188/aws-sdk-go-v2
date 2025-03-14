// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/dataexchange/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This operation updates a data set.
func (c *Client) UpdateDataSet(ctx context.Context, params *UpdateDataSetInput, optFns ...func(*Options)) (*UpdateDataSetOutput, error) {
	if params == nil {
		params = &UpdateDataSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataSet", params, optFns, c.addOperationUpdateDataSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSetInput struct {

	// The unique identifier for a data set.
	//
	// This member is required.
	DataSetId *string

	// The description for the data set.
	Description *string

	// The name of the data set.
	Name *string

	noSmithyDocumentSerde
}

type UpdateDataSetOutput struct {

	// The ARN for the data set.
	Arn *string

	// The type of asset that is added to a data set.
	AssetType types.AssetType

	// The date and time that the data set was created, in ISO 8601 format.
	CreatedAt *time.Time

	// The description for the data set.
	Description *string

	// The unique identifier for the data set.
	Id *string

	// The name of the data set.
	Name *string

	// A property that defines the data set as OWNED by the account (for providers) or
	// ENTITLED to the account (for subscribers).
	Origin types.Origin

	// If the origin of this data set is ENTITLED, includes the details for the product
	// on AWS Marketplace.
	OriginDetails *types.OriginDetails

	// The data set ID of the owned data set corresponding to the entitled data set
	// being viewed. This parameter is returned when a data set owner is viewing the
	// entitled copy of its owned data set.
	SourceId *string

	// The date and time that the data set was last updated, in ISO 8601 format.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDataSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataSet{}, middleware.After)
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
	if err = addOpUpdateDataSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "UpdateDataSet",
	}
}
