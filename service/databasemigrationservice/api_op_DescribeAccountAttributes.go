// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the DMS attributes for a customer account. These attributes include
// DMS quotas for the account and a unique account identifier in a particular DMS
// region. DMS quotas include a list of resource quotas supported by the account,
// such as the number of replication instances allowed. The description for each
// resource quota, includes the quota name, current usage toward that quota, and
// the quota's maximum value. DMS uses the unique account identifier to name each
// artifact used by DMS in the given region. This command does not take any
// parameters.
func (c *Client) DescribeAccountAttributes(ctx context.Context, params *DescribeAccountAttributesInput, optFns ...func(*Options)) (*DescribeAccountAttributesOutput, error) {
	if params == nil {
		params = &DescribeAccountAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountAttributes", params, optFns, c.addOperationDescribeAccountAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountAttributesInput struct {
	noSmithyDocumentSerde
}

type DescribeAccountAttributesOutput struct {

	// Account quota information.
	AccountQuotas []types.AccountQuota

	// A unique DMS identifier for an account in a particular Amazon Web Services
	// Region. The value of this identifier has the following format: c99999999999. DMS
	// uses this identifier to name artifacts. For example, DMS uses this identifier to
	// name the default Amazon S3 bucket for storing task assessment reports in a given
	// Amazon Web Services Region. The format of this S3 bucket name is the following:
	// dms-AccountNumber-UniqueAccountIdentifier. Here is an example name for this
	// default S3 bucket: dms-111122223333-c44445555666. DMS supports the
	// UniqueAccountIdentifier parameter in versions 3.1.4 and later.
	UniqueAccountIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccountAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAccountAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAccountAttributes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccountAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeAccountAttributes",
	}
}
