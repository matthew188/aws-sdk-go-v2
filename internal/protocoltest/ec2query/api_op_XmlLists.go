// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This test case serializes XML lists for the following cases for both input and
// output:
//
// * Normal XML lists.
//
// * Normal XML sets.
//
// * XML lists of lists.
//
// * XML
// lists with @xmlName on its members
//
// * Flattened XML lists.
//
// * Flattened XML
// lists with @xmlName.
//
// * Flattened XML lists with @xmlNamespace.
//
// * Lists of
// structures.
func (c *Client) XmlLists(ctx context.Context, params *XmlListsInput, optFns ...func(*Options)) (*XmlListsOutput, error) {
	if params == nil {
		params = &XmlListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "XmlLists", params, optFns, c.addOperationXmlListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*XmlListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlListsInput struct {
	noSmithyDocumentSerde
}

type XmlListsOutput struct {
	BooleanList []bool

	EnumList []types.FooEnum

	FlattenedList []string

	FlattenedList2 []string

	FlattenedListWithMemberNamespace []string

	FlattenedListWithNamespace []string

	IntEnumList []types.IntegerEnum

	IntegerList []int32

	// A list of lists of strings.
	NestedStringList [][]string

	RenamedListMembers []string

	StringList []string

	StringSet []string

	StructureList []types.StructureListMember

	TimestampList []time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationXmlListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpXmlLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpXmlLists{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opXmlLists(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opXmlLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "XmlLists",
	}
}
