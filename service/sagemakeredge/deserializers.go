// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakeredge

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/matthew188/aws-sdk-go-v2/service/sagemakeredge/types"
	smithy "github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"strings"
)

type awsRestjson1_deserializeOpGetDeployments struct {
}

func (*awsRestjson1_deserializeOpGetDeployments) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpGetDeployments) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorGetDeployments(response, &metadata)
	}
	output := &GetDeploymentsOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentGetDeploymentsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorGetDeployments(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InternalServiceException", errorCode):
		return awsRestjson1_deserializeErrorInternalServiceException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentGetDeploymentsOutput(v **GetDeploymentsOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *GetDeploymentsOutput
	if *v == nil {
		sv = &GetDeploymentsOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Deployments":
			if err := awsRestjson1_deserializeDocumentEdgeDeployments(&sv.Deployments, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpGetDeviceRegistration struct {
}

func (*awsRestjson1_deserializeOpGetDeviceRegistration) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpGetDeviceRegistration) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorGetDeviceRegistration(response, &metadata)
	}
	output := &GetDeviceRegistrationOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentGetDeviceRegistrationOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorGetDeviceRegistration(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InternalServiceException", errorCode):
		return awsRestjson1_deserializeErrorInternalServiceException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentGetDeviceRegistrationOutput(v **GetDeviceRegistrationOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *GetDeviceRegistrationOutput
	if *v == nil {
		sv = &GetDeviceRegistrationOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "CacheTTL":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected CacheTTLSeconds to be of type string, got %T instead", value)
				}
				sv.CacheTTL = ptr.String(jtv)
			}

		case "DeviceRegistration":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected DeviceRegistration to be of type string, got %T instead", value)
				}
				sv.DeviceRegistration = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpSendHeartbeat struct {
}

func (*awsRestjson1_deserializeOpSendHeartbeat) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpSendHeartbeat) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorSendHeartbeat(response, &metadata)
	}
	output := &SendHeartbeatOutput{}
	out.Result = output

	if _, err = io.Copy(ioutil.Discard, response.Body); err != nil {
		return out, metadata, &smithy.DeserializationError{
			Err: fmt.Errorf("failed to discard response body, %w", err),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorSendHeartbeat(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(headerCode)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 {
		errorCode = restjson.SanitizeErrorCode(jsonCode)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InternalServiceException", errorCode):
		return awsRestjson1_deserializeErrorInternalServiceException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeErrorInternalServiceException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InternalServiceException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentInternalServiceException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeDocumentChecksum(v **types.Checksum, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.Checksum
	if *v == nil {
		sv = &types.Checksum{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Sum":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ChecksumString to be of type string, got %T instead", value)
				}
				sv.Sum = ptr.String(jtv)
			}

		case "Type":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ChecksumType to be of type string, got %T instead", value)
				}
				sv.Type = types.ChecksumType(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentDefinition(v **types.Definition, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.Definition
	if *v == nil {
		sv = &types.Definition{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Checksum":
			if err := awsRestjson1_deserializeDocumentChecksum(&sv.Checksum, value); err != nil {
				return err
			}

		case "ModelHandle":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected EntityName to be of type string, got %T instead", value)
				}
				sv.ModelHandle = ptr.String(jtv)
			}

		case "S3Url":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected S3Uri to be of type string, got %T instead", value)
				}
				sv.S3Url = ptr.String(jtv)
			}

		case "State":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ModelState to be of type string, got %T instead", value)
				}
				sv.State = types.ModelState(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentDefinitions(v *[]types.Definition, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.Definition
	if *v == nil {
		cv = []types.Definition{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.Definition
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentDefinition(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentEdgeDeployment(v **types.EdgeDeployment, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.EdgeDeployment
	if *v == nil {
		sv = &types.EdgeDeployment{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Definitions":
			if err := awsRestjson1_deserializeDocumentDefinitions(&sv.Definitions, value); err != nil {
				return err
			}

		case "DeploymentName":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected EntityName to be of type string, got %T instead", value)
				}
				sv.DeploymentName = ptr.String(jtv)
			}

		case "FailureHandlingPolicy":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected FailureHandlingPolicy to be of type string, got %T instead", value)
				}
				sv.FailureHandlingPolicy = types.FailureHandlingPolicy(jtv)
			}

		case "Type":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected DeploymentType to be of type string, got %T instead", value)
				}
				sv.Type = types.DeploymentType(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentEdgeDeployments(v *[]types.EdgeDeployment, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.EdgeDeployment
	if *v == nil {
		cv = []types.EdgeDeployment{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.EdgeDeployment
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentEdgeDeployment(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentInternalServiceException(v **types.InternalServiceException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InternalServiceException
	if *v == nil {
		sv = &types.InternalServiceException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}
