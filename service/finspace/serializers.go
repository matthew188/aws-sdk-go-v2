// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"bytes"
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/finspace/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCreateEnvironment struct {
}

func (*awsRestjson1_serializeOpCreateEnvironment) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateEnvironmentInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/environment")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentCreateEnvironmentInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsCreateEnvironmentInput(v *CreateEnvironmentInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateEnvironmentInput(v *CreateEnvironmentInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DataBundles != nil {
		ok := object.Key("dataBundles")
		if err := awsRestjson1_serializeDocumentDataBundleArns(v.DataBundles, ok); err != nil {
			return err
		}
	}

	if v.Description != nil {
		ok := object.Key("description")
		ok.String(*v.Description)
	}

	if len(v.FederationMode) > 0 {
		ok := object.Key("federationMode")
		ok.String(string(v.FederationMode))
	}

	if v.FederationParameters != nil {
		ok := object.Key("federationParameters")
		if err := awsRestjson1_serializeDocumentFederationParameters(v.FederationParameters, ok); err != nil {
			return err
		}
	}

	if v.KmsKeyId != nil {
		ok := object.Key("kmsKeyId")
		ok.String(*v.KmsKeyId)
	}

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	if v.SuperuserParameters != nil {
		ok := object.Key("superuserParameters")
		if err := awsRestjson1_serializeDocumentSuperuserParameters(v.SuperuserParameters, ok); err != nil {
			return err
		}
	}

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDeleteEnvironment struct {
}

func (*awsRestjson1_serializeOpDeleteEnvironment) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteEnvironmentInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/environment/{environmentId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDeleteEnvironmentInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeleteEnvironmentInput(v *DeleteEnvironmentInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.EnvironmentId == nil || len(*v.EnvironmentId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member environmentId must not be empty")}
	}
	if v.EnvironmentId != nil {
		if err := encoder.SetURI("environmentId").String(*v.EnvironmentId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetEnvironment struct {
}

func (*awsRestjson1_serializeOpGetEnvironment) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetEnvironmentInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/environment/{environmentId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetEnvironmentInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetEnvironmentInput(v *GetEnvironmentInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.EnvironmentId == nil || len(*v.EnvironmentId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member environmentId must not be empty")}
	}
	if v.EnvironmentId != nil {
		if err := encoder.SetURI("environmentId").String(*v.EnvironmentId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListEnvironments struct {
}

func (*awsRestjson1_serializeOpListEnvironments) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListEnvironments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListEnvironmentsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/environment")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListEnvironmentsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListEnvironmentsInput(v *ListEnvironmentsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != 0 {
		encoder.SetQuery("maxResults").Integer(v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpListTagsForResource struct {
}

func (*awsRestjson1_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTagsForResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{resourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("resourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpTagResource struct {
}

func (*awsRestjson1_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{resourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsTagResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("resourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpUntagResource struct {
}

func (*awsRestjson1_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{resourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUntagResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("resourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	if v.TagKeys != nil {
		for i := range v.TagKeys {
			encoder.AddQuery("tagKeys").String(v.TagKeys[i])
		}
	}

	return nil
}

type awsRestjson1_serializeOpUpdateEnvironment struct {
}

func (*awsRestjson1_serializeOpUpdateEnvironment) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateEnvironment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateEnvironmentInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/environment/{environmentId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PUT"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUpdateEnvironmentInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdateEnvironmentInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsUpdateEnvironmentInput(v *UpdateEnvironmentInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.EnvironmentId == nil || len(*v.EnvironmentId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member environmentId must not be empty")}
	}
	if v.EnvironmentId != nil {
		if err := encoder.SetURI("environmentId").String(*v.EnvironmentId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdateEnvironmentInput(v *UpdateEnvironmentInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Description != nil {
		ok := object.Key("description")
		ok.String(*v.Description)
	}

	if len(v.FederationMode) > 0 {
		ok := object.Key("federationMode")
		ok.String(string(v.FederationMode))
	}

	if v.FederationParameters != nil {
		ok := object.Key("federationParameters")
		if err := awsRestjson1_serializeDocumentFederationParameters(v.FederationParameters, ok); err != nil {
			return err
		}
	}

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	return nil
}

func awsRestjson1_serializeDocumentAttributeMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsRestjson1_serializeDocumentDataBundleArns(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentFederationParameters(v *types.FederationParameters, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ApplicationCallBackURL != nil {
		ok := object.Key("applicationCallBackURL")
		ok.String(*v.ApplicationCallBackURL)
	}

	if v.AttributeMap != nil {
		ok := object.Key("attributeMap")
		if err := awsRestjson1_serializeDocumentAttributeMap(v.AttributeMap, ok); err != nil {
			return err
		}
	}

	if v.FederationProviderName != nil {
		ok := object.Key("federationProviderName")
		ok.String(*v.FederationProviderName)
	}

	if v.FederationURN != nil {
		ok := object.Key("federationURN")
		ok.String(*v.FederationURN)
	}

	if v.SamlMetadataDocument != nil {
		ok := object.Key("samlMetadataDocument")
		ok.String(*v.SamlMetadataDocument)
	}

	if v.SamlMetadataURL != nil {
		ok := object.Key("samlMetadataURL")
		ok.String(*v.SamlMetadataURL)
	}

	return nil
}

func awsRestjson1_serializeDocumentSuperuserParameters(v *types.SuperuserParameters, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EmailAddress != nil {
		ok := object.Key("emailAddress")
		ok.String(*v.EmailAddress)
	}

	if v.FirstName != nil {
		ok := object.Key("firstName")
		ok.String(*v.FirstName)
	}

	if v.LastName != nil {
		ok := object.Key("lastName")
		ok.String(*v.LastName)
	}

	return nil
}

func awsRestjson1_serializeDocumentTagMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}
