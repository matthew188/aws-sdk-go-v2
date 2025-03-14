// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/kinesisvideoarchivedmedia/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpGetClip struct {
}

func (*validateOpGetClip) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetClip) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetClipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetClipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetImages struct {
}

func (*validateOpGetImages) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetImages) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetImagesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetImagesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMediaForFragmentList struct {
}

func (*validateOpGetMediaForFragmentList) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMediaForFragmentList) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMediaForFragmentListInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMediaForFragmentListInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListFragments struct {
}

func (*validateOpListFragments) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListFragments) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListFragmentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListFragmentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpGetClipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetClip{}, middleware.After)
}

func addOpGetImagesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetImages{}, middleware.After)
}

func addOpGetMediaForFragmentListValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMediaForFragmentList{}, middleware.After)
}

func addOpListFragmentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListFragments{}, middleware.After)
}

func validateClipFragmentSelector(v *types.ClipFragmentSelector) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ClipFragmentSelector"}
	if len(v.FragmentSelectorType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("FragmentSelectorType"))
	}
	if v.TimestampRange == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TimestampRange"))
	} else if v.TimestampRange != nil {
		if err := validateClipTimestampRange(v.TimestampRange); err != nil {
			invalidParams.AddNested("TimestampRange", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateClipTimestampRange(v *types.ClipTimestampRange) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ClipTimestampRange"}
	if v.StartTimestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTimestamp"))
	}
	if v.EndTimestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTimestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFragmentSelector(v *types.FragmentSelector) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FragmentSelector"}
	if len(v.FragmentSelectorType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("FragmentSelectorType"))
	}
	if v.TimestampRange == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TimestampRange"))
	} else if v.TimestampRange != nil {
		if err := validateTimestampRange(v.TimestampRange); err != nil {
			invalidParams.AddNested("TimestampRange", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTimestampRange(v *types.TimestampRange) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TimestampRange"}
	if v.StartTimestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTimestamp"))
	}
	if v.EndTimestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTimestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetClipInput(v *GetClipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetClipInput"}
	if v.ClipFragmentSelector == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClipFragmentSelector"))
	} else if v.ClipFragmentSelector != nil {
		if err := validateClipFragmentSelector(v.ClipFragmentSelector); err != nil {
			invalidParams.AddNested("ClipFragmentSelector", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetImagesInput(v *GetImagesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetImagesInput"}
	if len(v.ImageSelectorType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ImageSelectorType"))
	}
	if v.StartTimestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTimestamp"))
	}
	if v.EndTimestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTimestamp"))
	}
	if v.SamplingInterval == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SamplingInterval"))
	}
	if len(v.Format) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Format"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMediaForFragmentListInput(v *GetMediaForFragmentListInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMediaForFragmentListInput"}
	if v.Fragments == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Fragments"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListFragmentsInput(v *ListFragmentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListFragmentsInput"}
	if v.FragmentSelector != nil {
		if err := validateFragmentSelector(v.FragmentSelector); err != nil {
			invalidParams.AddNested("FragmentSelector", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
