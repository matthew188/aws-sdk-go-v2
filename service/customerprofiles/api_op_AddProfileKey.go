// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a new key value with a specific profile, such as a Contact Record
// ContactId. A profile object can have a single unique key and any number of
// additional keys that can be used to identify the profile that it belongs to.
func (c *Client) AddProfileKey(ctx context.Context, params *AddProfileKeyInput, optFns ...func(*Options)) (*AddProfileKeyOutput, error) {
	if params == nil {
		params = &AddProfileKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddProfileKey", params, optFns, c.addOperationAddProfileKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddProfileKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddProfileKeyInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// A searchable identifier of a customer profile. The predefined keys you can use
	// include: _account, _profileId, _assetId, _caseId, _orderId, _fullName, _phone,
	// _email, _ctrContactId, _marketoLeadId, _salesforceAccountId,
	// _salesforceContactId, _salesforceAssetId, _zendeskUserId, _zendeskExternalId,
	// _zendeskTicketId, _serviceNowSystemId, _serviceNowIncidentId, _segmentUserId,
	// _shopifyCustomerId, _shopifyOrderId.
	//
	// This member is required.
	KeyName *string

	// The unique identifier of a customer profile.
	//
	// This member is required.
	ProfileId *string

	// A list of key values.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

type AddProfileKeyOutput struct {

	// A searchable identifier of a customer profile.
	KeyName *string

	// A list of key values.
	Values []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddProfileKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAddProfileKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAddProfileKey{}, middleware.After)
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
	if err = addOpAddProfileKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddProfileKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddProfileKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "AddProfileKey",
	}
}
