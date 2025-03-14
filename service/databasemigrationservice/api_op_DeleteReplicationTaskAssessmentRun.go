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

// Deletes the record of a single premigration assessment run. This operation
// removes all metadata that DMS maintains about this assessment run. However, the
// operation leaves untouched all information about this assessment run that is
// stored in your Amazon S3 bucket.
func (c *Client) DeleteReplicationTaskAssessmentRun(ctx context.Context, params *DeleteReplicationTaskAssessmentRunInput, optFns ...func(*Options)) (*DeleteReplicationTaskAssessmentRunOutput, error) {
	if params == nil {
		params = &DeleteReplicationTaskAssessmentRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteReplicationTaskAssessmentRun", params, optFns, c.addOperationDeleteReplicationTaskAssessmentRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteReplicationTaskAssessmentRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteReplicationTaskAssessmentRunInput struct {

	// Amazon Resource Name (ARN) of the premigration assessment run to be deleted.
	//
	// This member is required.
	ReplicationTaskAssessmentRunArn *string

	noSmithyDocumentSerde
}

type DeleteReplicationTaskAssessmentRunOutput struct {

	// The ReplicationTaskAssessmentRun object for the deleted assessment run.
	ReplicationTaskAssessmentRun *types.ReplicationTaskAssessmentRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteReplicationTaskAssessmentRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteReplicationTaskAssessmentRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteReplicationTaskAssessmentRun{}, middleware.After)
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
	if err = addOpDeleteReplicationTaskAssessmentRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReplicationTaskAssessmentRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteReplicationTaskAssessmentRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DeleteReplicationTaskAssessmentRun",
	}
}
