// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns backup job details for the specified BackupJobId.
func (c *Client) DescribeBackupJob(ctx context.Context, params *DescribeBackupJobInput, optFns ...func(*Options)) (*DescribeBackupJobOutput, error) {
	if params == nil {
		params = &DescribeBackupJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBackupJob", params, optFns, c.addOperationDescribeBackupJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBackupJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBackupJobInput struct {

	// Uniquely identifies a request to Backup to back up a resource.
	//
	// This member is required.
	BackupJobId *string

	noSmithyDocumentSerde
}

type DescribeBackupJobOutput struct {

	// Returns the account ID that owns the backup job.
	AccountId *string

	// Uniquely identifies a request to Backup to back up a resource.
	BackupJobId *string

	// Represents the options specified as part of backup plan or on-demand backup job.
	BackupOptions map[string]string

	// The size, in bytes, of a backup.
	BackupSizeInBytes *int64

	// Represents the actual backup type selected for a backup job. For example, if a
	// successful Windows Volume Shadow Copy Service (VSS) backup was taken, BackupType
	// returns "WindowsVSS". If BackupType is empty, then the backup type was a regular
	// backup.
	BackupType *string

	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Amazon Web Services Region where they are created. They consist of lowercase
	// letters, numbers, and hyphens.
	BackupVaultName *string

	// The size in bytes transferred to a backup vault at the time that the job status
	// was queried.
	BytesTransferred *int64

	// This returns the statistics of the included child (nested) backup jobs.
	ChildJobsInState map[string]int64

	// The date and time that a job to create a backup job is completed, in Unix format
	// and Coordinated Universal Time (UTC). The value of CompletionDate is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time

	// Contains identifying information about the creation of a backup job, including
	// the BackupPlanArn, BackupPlanId, BackupPlanVersion, and BackupRuleId of the
	// backup plan that is used to create it.
	CreatedBy *types.RecoveryPointCreator

	// The date and time that a backup job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// The date and time that a job to back up resources is expected to be completed,
	// in Unix format and Coordinated Universal Time (UTC). The value of
	// ExpectedCompletionDate is accurate to milliseconds. For example, the value
	// 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.
	ExpectedCompletionDate *time.Time

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string

	// This returns the boolean value that a backup job is a parent (composite) job.
	IsParent bool

	// This returns the number of child (nested) backup jobs.
	NumberOfChildJobs *int64

	// This returns the parent (composite) resource backup job ID.
	ParentJobId *string

	// Contains an estimated percentage that is complete of a job at the time the job
	// status was queried.
	PercentDone *string

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string

	// An ARN that uniquely identifies a saved resource. The format of the ARN depends
	// on the resource type.
	ResourceArn *string

	// This is the non-unique name of the resource that belongs to the specified
	// backup.
	ResourceName *string

	// The type of Amazon Web Services resource to be backed up; for example, an Amazon
	// Elastic Block Store (Amazon EBS) volume or an Amazon Relational Database Service
	// (Amazon RDS) database.
	ResourceType *string

	// Specifies the time in Unix format and Coordinated Universal Time (UTC) when a
	// backup job must be started before it is canceled. The value is calculated by
	// adding the start window to the scheduled time. So if the scheduled time were
	// 6:00 PM and the start window is 2 hours, the StartBy time would be 8:00 PM on
	// the date specified. The value of StartBy is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	StartBy *time.Time

	// The current state of a resource recovery point.
	State types.BackupJobState

	// A detailed message explaining the status of the job to back up a resource.
	StatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBackupJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBackupJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBackupJob{}, middleware.After)
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
	if err = addOpDescribeBackupJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBackupJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBackupJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeBackupJob",
	}
}
