package s3shared

import (
	"context"

	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

const metadataRetrieverID = "S3MetadataRetriever"

// AddMetadataRetrieverMiddleware adds request id, host id retriever middleware
func AddMetadataRetrieverMiddleware(stack *middleware.Stack) error {
	// add metadata retriever middleware before operation deserializers so that it can retrieve metadata such as
	// host id, request id from response header returned by operation deserializers
	return stack.Deserialize.Insert(&metadataRetriever{}, "OperationDeserializer", middleware.Before)
}

type metadataRetriever struct {
}

// ID returns the middleware identifier
func (m *metadataRetriever) ID() string {
	return metadataRetrieverID
}

func (m *metadataRetriever) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		// No raw response to wrap with.
		return out, metadata, err
	}

	// check for header for Request id
	if v := resp.Header.Get("X-Amz-Request-Id"); len(v) != 0 {
		// set reqID on metadata for successful responses.
		awsmiddleware.SetRequestIDMetadata(&metadata, v)
	}

	// look up host-id
	if v := resp.Header.Get("X-Amz-Id-2"); len(v) != 0 {
		// set reqID on metadata for successful responses.
		SetHostIDMetadata(&metadata, v)
	}

	return out, metadata, err
}
