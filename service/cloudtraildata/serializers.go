// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtraildata

import (
	"bytes"
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudtraildata/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpPutAuditEvents struct {
}

func (*awsRestjson1_serializeOpPutAuditEvents) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutAuditEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutAuditEventsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/PutAuditEvents")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsPutAuditEventsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentPutAuditEventsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsPutAuditEventsInput(v *PutAuditEventsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ChannelArn != nil {
		encoder.SetQuery("channelArn").String(*v.ChannelArn)
	}

	if v.ExternalId != nil {
		encoder.SetQuery("externalId").String(*v.ExternalId)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentPutAuditEventsInput(v *PutAuditEventsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AuditEvents != nil {
		ok := object.Key("auditEvents")
		if err := awsRestjson1_serializeDocumentAuditEvents(v.AuditEvents, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentAuditEvent(v *types.AuditEvent, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EventData != nil {
		ok := object.Key("eventData")
		ok.String(*v.EventData)
	}

	if v.EventDataChecksum != nil {
		ok := object.Key("eventDataChecksum")
		ok.String(*v.EventDataChecksum)
	}

	if v.Id != nil {
		ok := object.Key("id")
		ok.String(*v.Id)
	}

	return nil
}

func awsRestjson1_serializeDocumentAuditEvents(v []types.AuditEvent, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentAuditEvent(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}
