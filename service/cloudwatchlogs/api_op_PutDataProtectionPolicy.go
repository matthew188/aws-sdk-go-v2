// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a data protection policy for the specified log group. A data protection
// policy can help safeguard sensitive data that's ingested by the log group by
// auditing and masking the sensitive log data. Sensitive data is detected and
// masked when it is ingested into the log group. When you set a data protection
// policy, log events ingested into the log group before that time are not masked.
// By default, when a user views a log event that includes masked data, the
// sensitive data is replaced by asterisks. A user who has the logs:Unmask
// permission can use a GetLogEvents
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogEvents.html)
// or FilterLogEvents
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_FilterLogEvents.html)
// operation with the unmask parameter set to true to view the unmasked log events.
// Users with the logs:Unmask can also view unmasked data in the CloudWatch Logs
// console by running a CloudWatch Logs Insights query with the unmask query
// command. For more information, including a list of types of data that can be
// audited and masked, see Protect sensitive log data with masking
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html).
func (c *Client) PutDataProtectionPolicy(ctx context.Context, params *PutDataProtectionPolicyInput, optFns ...func(*Options)) (*PutDataProtectionPolicyOutput, error) {
	if params == nil {
		params = &PutDataProtectionPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDataProtectionPolicy", params, optFns, c.addOperationPutDataProtectionPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDataProtectionPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutDataProtectionPolicyInput struct {

	// Specify either the log group name or log group ARN.
	//
	// This member is required.
	LogGroupIdentifier *string

	// Specify the data protection policy, in JSON. This policy must include two JSON
	// blocks:
	//
	// * The first block must include both a DataIdentifer array and an
	// Operation property with an Audit action. The DataIdentifer array lists the types
	// of sensitive data that you want to mask. For more information about the
	// available options, see Types of data that you can mask
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-types.html).
	// The Operation property with an Audit action is required to find the sensitive
	// data terms. This Audit action must contain a FindingsDestination object. You can
	// optionally use that FindingsDestination object to list one or more destinations
	// to send audit findings to. If you specify destinations such as log groups,
	// Kinesis Data Firehose streams, and S3 buckets, they must already exist.
	//
	// * The
	// second block must include both a DataIdentifer array and an Operation property
	// with an Deidentify action. The DataIdentifer array must exactly match the
	// DataIdentifer array in the first block of the policy. The Operation property
	// with the Deidentify action is what actually masks the data, and it must contain
	// the  "MaskConfig": {} object. The  "MaskConfig": {} object must be empty.
	//
	// For
	// an example data protection policy, see the Examples section on this page. The
	// contents of two DataIdentifer arrays must match exactly.
	//
	// This member is required.
	PolicyDocument *string

	noSmithyDocumentSerde
}

type PutDataProtectionPolicyOutput struct {

	// The date and time that this policy was most recently updated.
	LastUpdatedTime *int64

	// The log group name or ARN that you specified in your request.
	LogGroupIdentifier *string

	// The data protection policy used for this log group.
	PolicyDocument *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutDataProtectionPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutDataProtectionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutDataProtectionPolicy{}, middleware.After)
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
	if err = addOpPutDataProtectionPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutDataProtectionPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutDataProtectionPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutDataProtectionPolicy",
	}
}
