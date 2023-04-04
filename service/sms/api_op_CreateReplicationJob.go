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

// Creates a replication job. The replication job schedules periodic replication
// runs to replicate your server to Amazon Web Services. Each replication run
// creates an Amazon Machine Image (AMI).
func (c *Client) CreateReplicationJob(ctx context.Context, params *CreateReplicationJobInput, optFns ...func(*Options)) (*CreateReplicationJobOutput, error) {
	if params == nil {
		params = &CreateReplicationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReplicationJob", params, optFns, c.addOperationCreateReplicationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReplicationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReplicationJobInput struct {

	// The seed replication time.
	//
	// This member is required.
	SeedReplicationTime *time.Time

	// The ID of the server.
	//
	// This member is required.
	ServerId *string

	// The description of the replication job.
	Description *string

	// Indicates whether the replication job produces encrypted AMIs.
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
	// is true but a KMS key ID is not specified, the customer's default KMS key for
	// Amazon EBS is used.
	KmsKeyId *string

	// The license type to be used for the AMI created by a successful replication run.
	LicenseType types.LicenseType

	// The maximum number of SMS-created AMIs to retain. The oldest is deleted after
	// the maximum number is reached and a new AMI is created.
	NumberOfRecentAmisToKeep *int32

	// The name of the IAM role to be used by the Server Migration Service.
	RoleName *string

	// Indicates whether to run the replication job one time.
	RunOnce *bool

	noSmithyDocumentSerde
}

type CreateReplicationJobOutput struct {

	// The unique identifier of the replication job.
	ReplicationJobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReplicationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateReplicationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateReplicationJob{}, middleware.After)
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
	if err = addOpCreateReplicationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicationJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReplicationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "CreateReplicationJob",
	}
}
