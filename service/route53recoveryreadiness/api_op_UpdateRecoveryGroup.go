// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a recovery group.
func (c *Client) UpdateRecoveryGroup(ctx context.Context, params *UpdateRecoveryGroupInput, optFns ...func(*Options)) (*UpdateRecoveryGroupOutput, error) {
	if params == nil {
		params = &UpdateRecoveryGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRecoveryGroup", params, optFns, c.addOperationUpdateRecoveryGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRecoveryGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Name of a recovery group.
type UpdateRecoveryGroupInput struct {

	// A list of cell Amazon Resource Names (ARNs). This list completely replaces the
	// previous list.
	//
	// This member is required.
	Cells []string

	// The name of a recovery group.
	//
	// This member is required.
	RecoveryGroupName *string

	noSmithyDocumentSerde
}

type UpdateRecoveryGroupOutput struct {

	// A list of a cell's Amazon Resource Names (ARNs).
	Cells []string

	// The Amazon Resource Name (ARN) for the recovery group.
	RecoveryGroupArn *string

	// The name of the recovery group.
	RecoveryGroupName *string

	// The tags associated with the recovery group.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRecoveryGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRecoveryGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRecoveryGroup{}, middleware.After)
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
	if err = addOpUpdateRecoveryGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRecoveryGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRecoveryGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-readiness",
		OperationName: "UpdateRecoveryGroup",
	}
}
