// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the data available for the group.
func (c *Client) DescribeGroup(ctx context.Context, params *DescribeGroupInput, optFns ...func(*Options)) (*DescribeGroupOutput, error) {
	if params == nil {
		params = &DescribeGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGroup", params, optFns, c.addOperationDescribeGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGroupInput struct {

	// The identifier for the group to be described.
	//
	// This member is required.
	GroupId *string

	// The identifier for the organization under which the group exists.
	//
	// This member is required.
	OrganizationId *string

	noSmithyDocumentSerde
}

type DescribeGroupOutput struct {

	// The date and time when a user was deregistered from WorkMail, in UNIX epoch time
	// format.
	DisabledDate *time.Time

	// The email of the described group.
	Email *string

	// The date and time when a user was registered to WorkMail, in UNIX epoch time
	// format.
	EnabledDate *time.Time

	// The identifier of the described group.
	GroupId *string

	// The name of the described group.
	Name *string

	// The state of the user: enabled (registered to WorkMail) or disabled
	// (deregistered or never registered to WorkMail).
	State types.EntityState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGroup{}, middleware.After)
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
	if err = addOpDescribeGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "DescribeGroup",
	}
}
