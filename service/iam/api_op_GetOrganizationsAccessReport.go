// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the service last accessed data report for Organizations that was
// previously generated using the GenerateOrganizationsAccessReport operation. This
// operation retrieves the status of your report job and the report contents.
// Depending on the parameters that you passed when you generated the report, the
// data returned could include different information. For details, see
// GenerateOrganizationsAccessReport. To call this operation, you must be signed in
// to the management account in your organization. SCPs must be enabled for your
// organization root. You must have permissions to perform this operation. For more
// information, see Refining permissions using service last accessed data
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html)
// in the IAM User Guide. For each service that principals in an account (root
// user, IAM users, or IAM roles) could access using SCPs, the operation returns
// details about the most recent access attempt. If there was no attempt, the
// service is listed without details about the most recent attempt to access the
// service. If the operation fails, it returns the reason that it failed. By
// default, the list is sorted by service namespace.
func (c *Client) GetOrganizationsAccessReport(ctx context.Context, params *GetOrganizationsAccessReportInput, optFns ...func(*Options)) (*GetOrganizationsAccessReportOutput, error) {
	if params == nil {
		params = &GetOrganizationsAccessReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOrganizationsAccessReport", params, optFns, c.addOperationGetOrganizationsAccessReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOrganizationsAccessReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOrganizationsAccessReportInput struct {

	// The identifier of the request generated by the GenerateOrganizationsAccessReport
	// operation.
	//
	// This member is required.
	JobId *string

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32

	// The key that is used to sort the results. If you choose the namespace key, the
	// results are returned in alphabetical order. If you choose the time key, the
	// results are sorted numerically by the date and time.
	SortKey types.SortKeyType

	noSmithyDocumentSerde
}

type GetOrganizationsAccessReportOutput struct {

	// The date and time, in ISO 8601 date-time format
	// (http://www.iso.org/iso/iso8601), when the report job was created.
	//
	// This member is required.
	JobCreationDate *time.Time

	// The status of the job.
	//
	// This member is required.
	JobStatus types.JobStatusType

	// An object that contains details about the most recent attempt to access the
	// service.
	AccessDetails []types.AccessDetail

	// Contains information about the reason that the operation failed. This data type
	// is used as a response element in the GetOrganizationsAccessReport,
	// GetServiceLastAccessedDetails, and GetServiceLastAccessedDetailsWithEntities
	// operations.
	ErrorDetails *types.ErrorDetails

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you receive
	// all your results.
	IsTruncated bool

	// The date and time, in ISO 8601 date-time format
	// (http://www.iso.org/iso/iso8601), when the generated report job was completed or
	// failed. This field is null if the job is still in progress, as indicated by a
	// job status value of IN_PROGRESS.
	JobCompletionDate *time.Time

	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// The number of services that the applicable SCPs allow account principals to
	// access.
	NumberOfServicesAccessible *int32

	// The number of services that account principals are allowed but did not attempt
	// to access.
	NumberOfServicesNotAccessed *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOrganizationsAccessReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetOrganizationsAccessReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetOrganizationsAccessReport{}, middleware.After)
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
	if err = addOpGetOrganizationsAccessReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOrganizationsAccessReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetOrganizationsAccessReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetOrganizationsAccessReport",
	}
}
