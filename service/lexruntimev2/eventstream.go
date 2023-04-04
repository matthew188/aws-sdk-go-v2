// Code generated by smithy-go-codegen DO NOT EDIT.

package lexruntimev2

import (
	"bytes"
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/protocol/eventstream"
	"github.com/matthew188/aws-sdk-go-v2/aws/protocol/eventstream/eventstreamapi"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lexruntimev2/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"sync"
	"time"
)

// StartConversationRequestEventStreamWriter provides the interface for writing
// events to a stream.
//
// The writer's Close method must allow multiple concurrent calls.
type StartConversationRequestEventStreamWriter interface {
	Send(context.Context, types.StartConversationRequestEventStream) error
	Close() error
	Err() error
}

// StartConversationResponseEventStreamReader provides the interface for reading
// events from a stream.
//
// The writer's Close method must allow multiple concurrent calls.
type StartConversationResponseEventStreamReader interface {
	Events() <-chan types.StartConversationResponseEventStream
	Close() error
	Err() error
}

type eventStreamSigner interface {
	GetSignature(ctx context.Context, headers, payload []byte, signingTime time.Time, optFns ...func(*v4.StreamSignerOptions)) ([]byte, error)
}

type asyncStartConversationRequestEventStream struct {
	Event  types.StartConversationRequestEventStream
	Result chan<- error
}

func (e asyncStartConversationRequestEventStream) ReportResult(cancel <-chan struct{}, err error) bool {
	select {
	case e.Result <- err:
		return true

	case <-cancel:
		return false

	}
}

type startConversationRequestEventStreamWriter struct {
	encoder             *eventstream.Encoder
	signer              eventStreamSigner
	stream              chan asyncStartConversationRequestEventStream
	serializationBuffer *bytes.Buffer
	signingBuffer       *bytes.Buffer
	eventStream         io.WriteCloser
	done                chan struct{}
	closeOnce           sync.Once
	err                 *smithysync.OnceErr
}

func newStartConversationRequestEventStreamWriter(stream io.WriteCloser, encoder *eventstream.Encoder, signer eventStreamSigner) *startConversationRequestEventStreamWriter {
	w := &startConversationRequestEventStreamWriter{
		encoder:             encoder,
		signer:              signer,
		stream:              make(chan asyncStartConversationRequestEventStream),
		eventStream:         stream,
		done:                make(chan struct{}),
		err:                 smithysync.NewOnceErr(),
		serializationBuffer: bytes.NewBuffer(nil),
		signingBuffer:       bytes.NewBuffer(nil),
	}

	go w.writeStream()

	return w

}

func (w *startConversationRequestEventStreamWriter) Send(ctx context.Context, event types.StartConversationRequestEventStream) error {
	return w.send(ctx, event)
}

func (w *startConversationRequestEventStreamWriter) send(ctx context.Context, event types.StartConversationRequestEventStream) error {
	if err := w.err.Err(); err != nil {
		return err
	}

	resultCh := make(chan error)

	wrapped := asyncStartConversationRequestEventStream{
		Event:  event,
		Result: resultCh,
	}

	select {
	case w.stream <- wrapped:
	case <-ctx.Done():
		return ctx.Err()
	case <-w.done:
		return fmt.Errorf("stream closed, unable to send event")

	}

	select {
	case err := <-resultCh:
		return err
	case <-ctx.Done():
		return ctx.Err()
	case <-w.done:
		return fmt.Errorf("stream closed, unable to send event")

	}

}

func (w *startConversationRequestEventStreamWriter) writeStream() {
	defer w.Close()

	for {
		select {
		case wrapper := <-w.stream:
			err := w.writeEvent(wrapper.Event)
			wrapper.ReportResult(w.done, err)
			if err != nil {
				w.err.SetError(err)
				return
			}

		case <-w.done:
			if err := w.closeStream(); err != nil {
				w.err.SetError(err)
			}
			return

		}
	}
}

func (w *startConversationRequestEventStreamWriter) writeEvent(event types.StartConversationRequestEventStream) error {
	// serializedEvent returned bytes refers to an underlying byte buffer and must not
	// escape this writeEvent scope without first copying. Any previous bytes stored in
	// the buffer are cleared by this call.
	serializedEvent, err := w.serializeEvent(event)
	if err != nil {
		return err
	}

	// signedEvent returned bytes refers to an underlying byte buffer and must not
	// escape this writeEvent scope without first copying. Any previous bytes stored in
	// the buffer are cleared by this call.
	signedEvent, err := w.signEvent(serializedEvent)
	if err != nil {
		return err
	}

	// bytes are now copied to the underlying stream writer
	_, err = io.Copy(w.eventStream, bytes.NewReader(signedEvent))
	return err
}

func (w *startConversationRequestEventStreamWriter) serializeEvent(event types.StartConversationRequestEventStream) ([]byte, error) {
	w.serializationBuffer.Reset()

	eventMessage := eventstream.Message{}

	if err := awsRestjson1_serializeEventStreamStartConversationRequestEventStream(event, &eventMessage); err != nil {
		return nil, err
	}

	if err := w.encoder.Encode(w.serializationBuffer, eventMessage); err != nil {
		return nil, err
	}

	return w.serializationBuffer.Bytes(), nil
}

func (w *startConversationRequestEventStreamWriter) signEvent(payload []byte) ([]byte, error) {
	w.signingBuffer.Reset()

	date := time.Now().UTC()

	var msg eventstream.Message
	msg.Headers.Set(eventstreamapi.DateHeader, eventstream.TimestampValue(date))
	msg.Payload = payload

	var headers bytes.Buffer
	if err := eventstream.EncodeHeaders(&headers, msg.Headers); err != nil {
		return nil, err
	}

	sig, err := w.signer.GetSignature(context.Background(), headers.Bytes(), msg.Payload, date)
	if err != nil {
		return nil, err
	}

	msg.Headers.Set(eventstreamapi.ChunkSignatureHeader, eventstream.BytesValue(sig))

	if err := w.encoder.Encode(w.signingBuffer, msg); err != nil {
		return nil, err
	}

	return w.signingBuffer.Bytes(), nil
}

func (w *startConversationRequestEventStreamWriter) closeStream() (err error) {
	defer func() {
		if cErr := w.eventStream.Close(); cErr != nil && err == nil {
			err = cErr
		}
	}()

	// Per the protocol, a signed empty message is used to indicate the end of the stream,
	// and that no subsequent events will be sent.
	signedEvent, err := w.signEvent([]byte{})
	if err != nil {
		return err
	}

	_, err = io.Copy(w.eventStream, bytes.NewReader(signedEvent))
	return err
}

func (w *startConversationRequestEventStreamWriter) ErrorSet() <-chan struct{} {
	return w.err.ErrorSet()
}

func (w *startConversationRequestEventStreamWriter) Close() error {
	w.closeOnce.Do(w.safeClose)
	return w.Err()
}

func (w *startConversationRequestEventStreamWriter) safeClose() {
	close(w.done)
}

func (w *startConversationRequestEventStreamWriter) Err() error {
	return w.err.Err()
}

type startConversationResponseEventStreamReader struct {
	stream      chan types.StartConversationResponseEventStream
	decoder     *eventstream.Decoder
	eventStream io.ReadCloser
	err         *smithysync.OnceErr
	payloadBuf  []byte
	done        chan struct{}
	closeOnce   sync.Once
}

func newStartConversationResponseEventStreamReader(readCloser io.ReadCloser, decoder *eventstream.Decoder) *startConversationResponseEventStreamReader {
	w := &startConversationResponseEventStreamReader{
		stream:      make(chan types.StartConversationResponseEventStream),
		decoder:     decoder,
		eventStream: readCloser,
		err:         smithysync.NewOnceErr(),
		done:        make(chan struct{}),
		payloadBuf:  make([]byte, 10*1024),
	}

	go w.readEventStream()

	return w
}

func (r *startConversationResponseEventStreamReader) Events() <-chan types.StartConversationResponseEventStream {
	return r.stream
}

func (r *startConversationResponseEventStreamReader) readEventStream() {
	defer r.Close()
	defer close(r.stream)

	for {
		r.payloadBuf = r.payloadBuf[0:0]
		decodedMessage, err := r.decoder.Decode(r.eventStream, r.payloadBuf)
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case <-r.done:
				return
			default:
				r.err.SetError(err)
				return
			}
		}

		event, err := r.deserializeEventMessage(&decodedMessage)
		if err != nil {
			r.err.SetError(err)
			return
		}

		select {
		case r.stream <- event:
		case <-r.done:
			return
		}

	}
}

func (r *startConversationResponseEventStreamReader) deserializeEventMessage(msg *eventstream.Message) (types.StartConversationResponseEventStream, error) {
	messageType := msg.Headers.Get(eventstreamapi.MessageTypeHeader)
	if messageType == nil {
		return nil, fmt.Errorf("%s event header not present", eventstreamapi.MessageTypeHeader)
	}

	switch messageType.String() {
	case eventstreamapi.EventMessageType:
		var v types.StartConversationResponseEventStream
		if err := awsRestjson1_deserializeEventStreamStartConversationResponseEventStream(&v, msg); err != nil {
			return nil, err
		}
		return v, nil

	case eventstreamapi.ExceptionMessageType:
		return nil, awsRestjson1_deserializeEventStreamExceptionStartConversationResponseEventStream(msg)

	case eventstreamapi.ErrorMessageType:
		errorCode := "UnknownError"
		errorMessage := errorCode
		if header := msg.Headers.Get(eventstreamapi.ErrorCodeHeader); header != nil {
			errorCode = header.String()
		}
		if header := msg.Headers.Get(eventstreamapi.ErrorMessageHeader); header != nil {
			errorMessage = header.String()
		}
		return nil, &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}

	default:
		mc := msg.Clone()
		return nil, &UnknownEventMessageError{
			Type:    messageType.String(),
			Message: &mc,
		}

	}
}

func (r *startConversationResponseEventStreamReader) ErrorSet() <-chan struct{} {
	return r.err.ErrorSet()
}

func (r *startConversationResponseEventStreamReader) Close() error {
	r.closeOnce.Do(r.safeClose)
	return r.Err()
}

func (r *startConversationResponseEventStreamReader) safeClose() {
	close(r.done)
	r.eventStream.Close()

}

func (r *startConversationResponseEventStreamReader) Err() error {
	return r.err.Err()
}

func (r *startConversationResponseEventStreamReader) Closed() <-chan struct{} {
	return r.done
}

type awsRestjson1_deserializeOpEventStreamStartConversation struct {
	LogEventStreamWrites bool
	LogEventStreamReads  bool
}

func (*awsRestjson1_deserializeOpEventStreamStartConversation) ID() string {
	return "OperationEventStreamDeserializer"
}

func (m *awsRestjson1_deserializeOpEventStreamStartConversation) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	defer func() {
		if err == nil {
			return
		}
		m.closeResponseBody(out)
	}()

	logger := middleware.GetLogger(ctx)

	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", in.Request)
	}
	_ = request

	if err := eventstreamapi.ApplyHTTPTransportFixes(request); err != nil {
		return out, metadata, err
	}

	requestSignature, err := v4.GetSignedRequestSignature(request.Request)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to get event stream seed signature: %v", err)
	}

	signer := v4.NewStreamSigner(
		awsmiddleware.GetSigningCredentials(ctx),
		awsmiddleware.GetSigningName(ctx),
		awsmiddleware.GetSigningRegion(ctx),
		requestSignature,
	)

	eventWriter := newStartConversationRequestEventStreamWriter(
		eventstreamapi.GetInputStreamWriter(ctx),
		eventstream.NewEncoder(func(options *eventstream.EncoderOptions) {
			options.Logger = logger
			options.LogMessages = m.LogEventStreamWrites

		}),
		signer,
	)
	defer func() {
		if err == nil {
			return
		}
		_ = eventWriter.Close()
	}()

	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	deserializeOutput, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", out.RawResponse)
	}
	_ = deserializeOutput

	output, ok := out.Result.(*StartConversationOutput)
	if out.Result != nil && !ok {
		return out, metadata, fmt.Errorf("unexpected output result type: %T", out.Result)
	} else if out.Result == nil {
		output = &StartConversationOutput{}
		out.Result = output
	}

	eventReader := newStartConversationResponseEventStreamReader(
		deserializeOutput.Body,
		eventstream.NewDecoder(func(options *eventstream.DecoderOptions) {
			options.Logger = logger
			options.LogMessages = m.LogEventStreamReads

		}),
	)
	defer func() {
		if err == nil {
			return
		}
		_ = eventReader.Close()
	}()

	output.eventStream = NewStartConversationEventStream(func(stream *StartConversationEventStream) {
		stream.Writer = eventWriter
		stream.Reader = eventReader
	})

	go output.eventStream.waitStreamClose()

	return out, metadata, nil
}

func (*awsRestjson1_deserializeOpEventStreamStartConversation) closeResponseBody(out middleware.DeserializeOutput) {
	if resp, ok := out.RawResponse.(*smithyhttp.Response); ok && resp != nil && resp.Body != nil {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}

func addEventStreamStartConversationMiddleware(stack *middleware.Stack, options Options) error {
	if err := stack.Deserialize.Insert(&awsRestjson1_deserializeOpEventStreamStartConversation{
		LogEventStreamWrites: options.ClientLogMode.IsRequestEventMessage(),
		LogEventStreamReads:  options.ClientLogMode.IsResponseEventMessage(),
	}, "OperationDeserializer", middleware.Before); err != nil {
		return err
	}
	return nil

}

// UnknownEventMessageError provides an error when a message is received from the stream,
// but the reader is unable to determine what kind of message it is.
type UnknownEventMessageError struct {
	Type    string
	Message *eventstream.Message
}

// Error retruns the error message string.
func (e *UnknownEventMessageError) Error() string {
	return "unknown event stream message type, " + e.Type
}

func setSafeEventStreamClientLogMode(o *Options, operation string) {
	switch operation {
	case "StartConversation":
		toggleEventStreamClientLogMode(o, true, true)
		return

	default:
		return

	}
}
func toggleEventStreamClientLogMode(o *Options, request, response bool) {
	mode := o.ClientLogMode

	if request && mode.IsRequestWithBody() {
		mode.ClearRequestWithBody()
		mode |= aws.LogRequest
	}

	if response && mode.IsResponseWithBody() {
		mode.ClearResponseWithBody()
		mode |= aws.LogResponse
	}

	o.ClientLogMode = mode

}
