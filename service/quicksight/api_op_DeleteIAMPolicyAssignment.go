// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an existing IAM policy assignment.
func (c *Client) DeleteIAMPolicyAssignment(ctx context.Context, params *DeleteIAMPolicyAssignmentInput, optFns ...func(*Options)) (*DeleteIAMPolicyAssignmentOutput, error) {
	if params == nil {
		params = &DeleteIAMPolicyAssignmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIAMPolicyAssignment", params, optFns, c.addOperationDeleteIAMPolicyAssignmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIAMPolicyAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteIAMPolicyAssignmentInput struct {

	// The name of the assignment.
	//
	// This member is required.
	AssignmentName *string

	// The Amazon Web Services account ID where you want to delete the IAM policy
	// assignment.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace that contains the assignment.
	//
	// This member is required.
	Namespace *string

	noSmithyDocumentSerde
}

type DeleteIAMPolicyAssignmentOutput struct {

	// The name of the assignment.
	AssignmentName *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteIAMPolicyAssignmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteIAMPolicyAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteIAMPolicyAssignment{}, middleware.After)
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
	if err = addOpDeleteIAMPolicyAssignmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIAMPolicyAssignment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteIAMPolicyAssignment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteIAMPolicyAssignment",
	}
}
