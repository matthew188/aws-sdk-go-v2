// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes an existing opt-out list. All opted out phone numbers in the opt-out
// list are deleted. If the specified opt-out list name doesn't exist or is in-use
// by an origination phone number or pool, an Error is returned.
func (c *Client) DeleteOptOutList(ctx context.Context, params *DeleteOptOutListInput, optFns ...func(*Options)) (*DeleteOptOutListOutput, error) {
	if params == nil {
		params = &DeleteOptOutListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOptOutList", params, optFns, c.addOperationDeleteOptOutListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOptOutListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOptOutListInput struct {

	// The OptOutListName or OptOutListArn of the OptOutList to delete. You can use
	// DescribeOptOutLists to find the values for OptOutListName and OptOutListArn.
	//
	// This member is required.
	OptOutListName *string

	noSmithyDocumentSerde
}

type DeleteOptOutListOutput struct {

	// The time when the OptOutList was created, in UNIX epoch time
	// (https://www.epochconverter.com/) format.
	CreatedTimestamp *time.Time

	// The Amazon Resource Name (ARN) of the OptOutList that was removed.
	OptOutListArn *string

	// The name of the OptOutList that was removed.
	OptOutListName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteOptOutListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteOptOutList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteOptOutList{}, middleware.After)
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
	if err = addOpDeleteOptOutListValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOptOutList(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteOptOutList(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DeleteOptOutList",
	}
}
