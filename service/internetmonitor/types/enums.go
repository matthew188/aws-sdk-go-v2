// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type HealthEventImpactType string

// Enum values for HealthEventImpactType
const (
	HealthEventImpactTypeAvailability HealthEventImpactType = "AVAILABILITY"
	HealthEventImpactTypePerformance  HealthEventImpactType = "PERFORMANCE"
)

// Values returns all known values for HealthEventImpactType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HealthEventImpactType) Values() []HealthEventImpactType {
	return []HealthEventImpactType{
		"AVAILABILITY",
		"PERFORMANCE",
	}
}

type HealthEventStatus string

// Enum values for HealthEventStatus
const (
	HealthEventStatusActive   HealthEventStatus = "ACTIVE"
	HealthEventStatusResolved HealthEventStatus = "RESOLVED"
)

// Values returns all known values for HealthEventStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HealthEventStatus) Values() []HealthEventStatus {
	return []HealthEventStatus{
		"ACTIVE",
		"RESOLVED",
	}
}

type LogDeliveryStatus string

// Enum values for LogDeliveryStatus
const (
	LogDeliveryStatusEnabled  LogDeliveryStatus = "ENABLED"
	LogDeliveryStatusDisabled LogDeliveryStatus = "DISABLED"
)

// Values returns all known values for LogDeliveryStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LogDeliveryStatus) Values() []LogDeliveryStatus {
	return []LogDeliveryStatus{
		"ENABLED",
		"DISABLED",
	}
}

type MonitorConfigState string

// Enum values for MonitorConfigState
const (
	MonitorConfigStatePending  MonitorConfigState = "PENDING"
	MonitorConfigStateActive   MonitorConfigState = "ACTIVE"
	MonitorConfigStateInactive MonitorConfigState = "INACTIVE"
	MonitorConfigStateError    MonitorConfigState = "ERROR"
)

// Values returns all known values for MonitorConfigState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MonitorConfigState) Values() []MonitorConfigState {
	return []MonitorConfigState{
		"PENDING",
		"ACTIVE",
		"INACTIVE",
		"ERROR",
	}
}

type MonitorProcessingStatusCode string

// Enum values for MonitorProcessingStatusCode
const (
	MonitorProcessingStatusCodeOk                    MonitorProcessingStatusCode = "OK"
	MonitorProcessingStatusCodeInactive              MonitorProcessingStatusCode = "INACTIVE"
	MonitorProcessingStatusCodeCollectingData        MonitorProcessingStatusCode = "COLLECTING_DATA"
	MonitorProcessingStatusCodeInsufficientData      MonitorProcessingStatusCode = "INSUFFICIENT_DATA"
	MonitorProcessingStatusCodeFaultService          MonitorProcessingStatusCode = "FAULT_SERVICE"
	MonitorProcessingStatusCodeFaultAccessCloudwatch MonitorProcessingStatusCode = "FAULT_ACCESS_CLOUDWATCH"
)

// Values returns all known values for MonitorProcessingStatusCode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (MonitorProcessingStatusCode) Values() []MonitorProcessingStatusCode {
	return []MonitorProcessingStatusCode{
		"OK",
		"INACTIVE",
		"COLLECTING_DATA",
		"INSUFFICIENT_DATA",
		"FAULT_SERVICE",
		"FAULT_ACCESS_CLOUDWATCH",
	}
}

type TriangulationEventType string

// Enum values for TriangulationEventType
const (
	TriangulationEventTypeAws      TriangulationEventType = "AWS"
	TriangulationEventTypeInternet TriangulationEventType = "Internet"
)

// Values returns all known values for TriangulationEventType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TriangulationEventType) Values() []TriangulationEventType {
	return []TriangulationEventType{
		"AWS",
		"Internet",
	}
}
