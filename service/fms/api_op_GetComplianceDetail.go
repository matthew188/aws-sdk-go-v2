// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns detailed compliance information about the specified member account.
// Details include resources that are in and out of compliance with the specified
// policy.
//
// * Resources are considered noncompliant for WAF and Shield Advanced
// policies if the specified policy has not been applied to them.
//
// * Resources are
// considered noncompliant for security group policies if they are in scope of the
// policy, they violate one or more of the policy rules, and remediation is
// disabled or not possible.
//
// * Resources are considered noncompliant for Network
// Firewall policies if a firewall is missing in the VPC, if the firewall endpoint
// isn't set up in an expected Availability Zone and subnet, if a subnet created by
// the Firewall Manager doesn't have the expected route table, and for
// modifications to a firewall policy that violate the Firewall Manager policy's
// rules.
//
// * Resources are considered noncompliant for DNS Firewall policies if a
// DNS Firewall rule group is missing from the rule group associations for the VPC.
func (c *Client) GetComplianceDetail(ctx context.Context, params *GetComplianceDetailInput, optFns ...func(*Options)) (*GetComplianceDetailOutput, error) {
	if params == nil {
		params = &GetComplianceDetailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComplianceDetail", params, optFns, c.addOperationGetComplianceDetailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComplianceDetailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetComplianceDetailInput struct {

	// The Amazon Web Services account that owns the resources that you want to get the
	// details for.
	//
	// This member is required.
	MemberAccount *string

	// The ID of the policy that you want to get the details for. PolicyId is returned
	// by PutPolicy and by ListPolicies.
	//
	// This member is required.
	PolicyId *string

	noSmithyDocumentSerde
}

type GetComplianceDetailOutput struct {

	// Information about the resources and the policy that you specified in the
	// GetComplianceDetail request.
	PolicyComplianceDetail *types.PolicyComplianceDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetComplianceDetailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComplianceDetail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComplianceDetail{}, middleware.After)
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
	if err = addOpGetComplianceDetailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComplianceDetail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetComplianceDetail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "GetComplianceDetail",
	}
}
