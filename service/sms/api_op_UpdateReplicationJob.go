// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified settings for the specified replication job.
func (c *Client) UpdateReplicationJob(ctx context.Context, params *UpdateReplicationJobInput, optFns ...func(*Options)) (*UpdateReplicationJobOutput, error) {
	if params == nil {
		params = &UpdateReplicationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateReplicationJob", params, optFns, c.addOperationUpdateReplicationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateReplicationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateReplicationJobInput struct {

	// The ID of the replication job.
	//
	// This member is required.
	ReplicationJobId *string

	// The description of the replication job.
	Description *string

	// When true, the replication job produces encrypted AMIs. For more information,
	// KmsKeyId.
	Encrypted *bool

	// The time between consecutive replication runs, in hours.
	Frequency *int32

	// The ID of the KMS key for replication jobs that produce encrypted AMIs. This
	// value can be any of the following:
	//
	// * KMS key ID
	//
	// * KMS key alias
	//
	// * ARN
	// referring to the KMS key ID
	//
	// * ARN referring to the KMS key alias
	//
	// If encrypted
	// is enabled but a KMS key ID is not specified, the customer's default KMS key for
	// Amazon EBS is used.
	KmsKeyId *string

	// The license type to be used for the AMI created by a successful replication run.
	LicenseType types.LicenseType

	// The start time of the next replication run.
	NextReplicationRunStartTime *time.Time

	// The maximum number of SMS-created AMIs to retain. The oldest is deleted after
	// the maximum number is reached and a new AMI is created.
	NumberOfRecentAmisToKeep *int32

	// The name of the IAM role to be used by Server Migration Service.
	RoleName *string

	noSmithyDocumentSerde
}

type UpdateReplicationJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateReplicationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateReplicationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateReplicationJob{}, middleware.After)
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
	if err = addOpUpdateReplicationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateReplicationJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateReplicationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "UpdateReplicationJob",
	}
}
