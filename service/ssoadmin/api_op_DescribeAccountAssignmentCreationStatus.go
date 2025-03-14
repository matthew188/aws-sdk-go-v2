// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the status of the assignment creation request.
func (c *Client) DescribeAccountAssignmentCreationStatus(ctx context.Context, params *DescribeAccountAssignmentCreationStatusInput, optFns ...func(*Options)) (*DescribeAccountAssignmentCreationStatusOutput, error) {
	if params == nil {
		params = &DescribeAccountAssignmentCreationStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountAssignmentCreationStatus", params, optFns, c.addOperationDescribeAccountAssignmentCreationStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountAssignmentCreationStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountAssignmentCreationStatusInput struct {

	// The identifier that is used to track the request operation progress.
	//
	// This member is required.
	AccountAssignmentCreationRequestId *string

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed. For more information about ARNs, see Amazon Resource Names (ARNs) and
	// AWS Service Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	noSmithyDocumentSerde
}

type DescribeAccountAssignmentCreationStatusOutput struct {

	// The status object for the account assignment creation operation.
	AccountAssignmentCreationStatus *types.AccountAssignmentOperationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccountAssignmentCreationStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAccountAssignmentCreationStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAccountAssignmentCreationStatus{}, middleware.After)
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
	if err = addOpDescribeAccountAssignmentCreationStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountAssignmentCreationStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccountAssignmentCreationStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "DescribeAccountAssignmentCreationStatus",
	}
}
