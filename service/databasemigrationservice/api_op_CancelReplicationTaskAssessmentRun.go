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

// Cancels a single premigration assessment run. This operation prevents any
// individual assessments from running if they haven't started running. It also
// attempts to cancel any individual assessments that are currently running.
func (c *Client) CancelReplicationTaskAssessmentRun(ctx context.Context, params *CancelReplicationTaskAssessmentRunInput, optFns ...func(*Options)) (*CancelReplicationTaskAssessmentRunOutput, error) {
	if params == nil {
		params = &CancelReplicationTaskAssessmentRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelReplicationTaskAssessmentRun", params, optFns, c.addOperationCancelReplicationTaskAssessmentRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelReplicationTaskAssessmentRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelReplicationTaskAssessmentRunInput struct {

	// Amazon Resource Name (ARN) of the premigration assessment run to be canceled.
	//
	// This member is required.
	ReplicationTaskAssessmentRunArn *string

	noSmithyDocumentSerde
}

type CancelReplicationTaskAssessmentRunOutput struct {

	// The ReplicationTaskAssessmentRun object for the canceled assessment run.
	ReplicationTaskAssessmentRun *types.ReplicationTaskAssessmentRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelReplicationTaskAssessmentRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelReplicationTaskAssessmentRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelReplicationTaskAssessmentRun{}, middleware.After)
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
	if err = addOpCancelReplicationTaskAssessmentRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelReplicationTaskAssessmentRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelReplicationTaskAssessmentRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "CancelReplicationTaskAssessmentRun",
	}
}
