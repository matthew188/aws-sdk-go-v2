// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metadata about a backup vault specified by its name.
func (c *Client) DescribeBackupVault(ctx context.Context, params *DescribeBackupVaultInput, optFns ...func(*Options)) (*DescribeBackupVaultOutput, error) {
	if params == nil {
		params = &DescribeBackupVaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBackupVault", params, optFns, c.addOperationDescribeBackupVaultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBackupVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBackupVaultInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Amazon Web Services Region where they are created. They consist of lowercase
	// letters, numbers, and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	noSmithyDocumentSerde
}

type DescribeBackupVaultOutput struct {

	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Region where they are created. They consist of lowercase letters, numbers, and
	// hyphens.
	BackupVaultName *string

	// The date and time that a backup vault is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of running the operation twice.
	CreatorRequestId *string

	// The server-side encryption key that is used to protect your backups; for
	// example,
	// arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	EncryptionKeyArn *string

	// The date and time when Backup Vault Lock configuration cannot be changed or
	// deleted. If you applied Vault Lock to your vault without specifying a lock date,
	// you can change any of your Vault Lock settings, or delete Vault Lock from the
	// vault entirely, at any time. This value is in Unix format, Coordinated Universal
	// Time (UTC), and accurate to milliseconds. For example, the value 1516925490.087
	// represents Friday, January 26, 2018 12:11:30.087 AM.
	LockDate *time.Time

	// A Boolean that indicates whether Backup Vault Lock is currently protecting the
	// backup vault. True means that Vault Lock causes delete or update operations on
	// the recovery points stored in the vault to fail.
	Locked *bool

	// The Backup Vault Lock setting that specifies the maximum retention period that
	// the vault retains its recovery points. If this parameter is not specified, Vault
	// Lock does not enforce a maximum retention period on the recovery points in the
	// vault (allowing indefinite storage). If specified, any backup or copy job to the
	// vault must have a lifecycle policy with a retention period equal to or shorter
	// than the maximum retention period. If the job's retention period is longer than
	// that maximum retention period, then the vault fails the backup or copy job, and
	// you should either modify your lifecycle settings or use a different vault.
	// Recovery points already stored in the vault prior to Vault Lock are not
	// affected.
	MaxRetentionDays *int64

	// The Backup Vault Lock setting that specifies the minimum retention period that
	// the vault retains its recovery points. If this parameter is not specified, Vault
	// Lock does not enforce a minimum retention period. If specified, any backup or
	// copy job to the vault must have a lifecycle policy with a retention period equal
	// to or longer than the minimum retention period. If the job's retention period is
	// shorter than that minimum retention period, then the vault fails the backup or
	// copy job, and you should either modify your lifecycle settings or use a
	// different vault. Recovery points already stored in the vault prior to Vault Lock
	// are not affected.
	MinRetentionDays *int64

	// The number of recovery points that are stored in a backup vault.
	NumberOfRecoveryPoints int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBackupVaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBackupVault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBackupVault{}, middleware.After)
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
	if err = addOpDescribeBackupVaultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBackupVault(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBackupVault(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeBackupVault",
	}
}
