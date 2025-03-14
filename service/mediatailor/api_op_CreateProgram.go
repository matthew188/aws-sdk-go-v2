// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a program within a channel. For information about programs, see Working
// with programs
// (https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-programs.html)
// in the MediaTailor User Guide.
func (c *Client) CreateProgram(ctx context.Context, params *CreateProgramInput, optFns ...func(*Options)) (*CreateProgramOutput, error) {
	if params == nil {
		params = &CreateProgramInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProgram", params, optFns, c.addOperationCreateProgramMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProgramOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProgramInput struct {

	// The name of the channel for this Program.
	//
	// This member is required.
	ChannelName *string

	// The name of the Program.
	//
	// This member is required.
	ProgramName *string

	// The schedule configuration settings.
	//
	// This member is required.
	ScheduleConfiguration *types.ScheduleConfiguration

	// The name of the source location.
	//
	// This member is required.
	SourceLocationName *string

	// The ad break configuration settings.
	AdBreaks []types.AdBreak

	// The name of the LiveSource for this Program.
	LiveSourceName *string

	// The name that's used to refer to a VOD source.
	VodSourceName *string

	noSmithyDocumentSerde
}

type CreateProgramOutput struct {

	// The ad break configuration settings.
	AdBreaks []types.AdBreak

	// The ARN to assign to the program.
	Arn *string

	// The name to assign to the channel for this program.
	ChannelName *string

	// The clip range configuration settings.
	ClipRange *types.ClipRange

	// The time the program was created.
	CreationTime *time.Time

	// The duration of the live program in milliseconds.
	DurationMillis int64

	// The name of the LiveSource for this Program.
	LiveSourceName *string

	// The name to assign to this program.
	ProgramName *string

	// The scheduled start time for this Program.
	ScheduledStartTime *time.Time

	// The name to assign to the source location for this program.
	SourceLocationName *string

	// The name that's used to refer to a VOD source.
	VodSourceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProgramMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProgram{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProgram{}, middleware.After)
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
	if err = addOpCreateProgramValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProgram(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProgram(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "CreateProgram",
	}
}
