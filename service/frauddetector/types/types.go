// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// The log odds metric details. Account Takeover Insights (ATI) model uses event
// variables from the login data you provide to continuously calculate a set of
// variables (aggregated variables) based on historical events. For example, your
// ATI model might calculate the number of times an user has logged in using the
// same IP address. In this case, event variables used to derive the aggregated
// variables are IP address and user.
type AggregatedLogOddsMetric struct {

	// The relative importance of the variables in the list to the other event
	// variable.
	//
	// This member is required.
	AggregatedVariablesImportance *float32

	// The names of all the variables.
	//
	// This member is required.
	VariableNames []string

	noSmithyDocumentSerde
}

// The details of the impact of aggregated variables on the prediction score.
// Account Takeover Insights (ATI) model uses the login data you provide to
// continuously calculate a set of variables (aggregated variables) based on
// historical events. For example, the model might calculate the number of times an
// user has logged in using the same IP address. In this case, event variables used
// to derive the aggregated variables are IP address and user.
type AggregatedVariablesImpactExplanation struct {

	// The names of all the event variables that were used to derive the aggregated
	// variables.
	EventVariableNames []string

	// The raw, uninterpreted value represented as log-odds of the fraud. These values
	// are usually between -10 to +10, but range from -infinity to +infinity.
	//
	// * A
	// positive value indicates that the variables drove the risk score up.
	//
	// * A
	// negative value indicates that the variables drove the risk score down.
	LogOddsImpact *float32

	// The relative impact of the aggregated variables in terms of magnitude on the
	// prediction scores.
	RelativeImpact *string

	noSmithyDocumentSerde
}

// The details of the relative importance of the aggregated variables. Account
// Takeover Insights (ATI) model uses event variables from the login data you
// provide to continuously calculate a set of variables (aggregated variables)
// based on historical events. For example, your ATI model might calculate the
// number of times an user has logged in using the same IP address. In this case,
// event variables used to derive the aggregated variables are IP address and user.
type AggregatedVariablesImportanceMetrics struct {

	// List of variables' metrics.
	LogOddsMetrics []AggregatedLogOddsMetric

	noSmithyDocumentSerde
}

// The metadata of a list.
type AllowDenyList struct {

	// The name of the list.
	//
	// This member is required.
	Name *string

	// The ARN of the list.
	Arn *string

	// The time the list was created.
	CreatedTime *string

	// The description of the list.
	Description *string

	// The time the list was last updated.
	UpdatedTime *string

	// The variable type of the list.
	VariableType *string

	noSmithyDocumentSerde
}

// The Account Takeover Insights (ATI) model performance metrics data points.
type ATIMetricDataPoint struct {

	// The anomaly discovery rate. This metric quantifies the percentage of anomalies
	// that can be detected by the model at the selected score threshold. A lower score
	// threshold increases the percentage of anomalies captured by the model, but would
	// also require challenging a larger percentage of login events, leading to a
	// higher customer friction.
	Adr *float32

	// The account takeover discovery rate. This metric quantifies the percentage of
	// account compromise events that can be detected by the model at the selected
	// score threshold. This metric is only available if 50 or more entities with
	// at-least one labeled account takeover event is present in the ingested dataset.
	Atodr *float32

	// The challenge rate. This indicates the percentage of login events that the model
	// recommends to challenge such as one-time password, multi-factor authentication,
	// and investigations.
	Cr *float32

	// The model's threshold that specifies an acceptable fraud capture rate. For
	// example, a threshold of 500 means any model score 500 or above is labeled as
	// fraud.
	Threshold *float32

	noSmithyDocumentSerde
}

// The Account Takeover Insights (ATI) model performance score.
type ATIModelPerformance struct {

	// The anomaly separation index (ASI) score. This metric summarizes the overall
	// ability of the model to separate anomalous activities from the normal behavior.
	// Depending on the business, a large fraction of these anomalous activities can be
	// malicious and correspond to the account takeover attacks. A model with no
	// separability power will have the lowest possible ASI score of 0.5, whereas the a
	// model with a high separability power will have the highest possible ASI score of
	// 1.0
	Asi *float32

	noSmithyDocumentSerde
}

// The Account Takeover Insights (ATI) model training metric details.
type ATITrainingMetricsValue struct {

	// The model's performance metrics data points.
	MetricDataPoints []ATIMetricDataPoint

	// The model's overall performance scores.
	ModelPerformance *ATIModelPerformance

	noSmithyDocumentSerde
}

// Provides the error of the batch create variable API.
type BatchCreateVariableError struct {

	// The error code.
	Code int32

	// The error message.
	Message *string

	// The name.
	Name *string

	noSmithyDocumentSerde
}

// Provides the error of the batch get variable API.
type BatchGetVariableError struct {

	// The error code.
	Code int32

	// The error message.
	Message *string

	// The error name.
	Name *string

	noSmithyDocumentSerde
}

// The batch import job details.
type BatchImport struct {

	// The ARN of the batch import job.
	Arn *string

	// Timestamp of when batch import job completed.
	CompletionTime *string

	// The name of the event type.
	EventTypeName *string

	// The number of records that failed to import.
	FailedRecordsCount *int32

	// The reason batch import job failed.
	FailureReason *string

	// The ARN of the IAM role to use for this job request.
	IamRoleArn *string

	// The Amazon S3 location of your data file for batch import.
	InputPath *string

	// The ID of the batch import job.
	JobId *string

	// The Amazon S3 location of your output file.
	OutputPath *string

	// The number of records processed by batch import job.
	ProcessedRecordsCount *int32

	// Timestamp of when the batch import job started.
	StartTime *string

	// The status of the batch import job.
	Status AsyncJobStatus

	// The total number of records in the batch import job.
	TotalRecordsCount *int32

	noSmithyDocumentSerde
}

// The batch prediction details.
type BatchPrediction struct {

	// The ARN of batch prediction job.
	Arn *string

	// Timestamp of when the batch prediction job completed.
	CompletionTime *string

	// The name of the detector.
	DetectorName *string

	// The detector version.
	DetectorVersion *string

	// The name of the event type.
	EventTypeName *string

	// The reason a batch prediction job failed.
	FailureReason *string

	// The ARN of the IAM role to use for this job request.
	IamRoleArn *string

	// The Amazon S3 location of your training file.
	InputPath *string

	// The job ID for the batch prediction.
	JobId *string

	// Timestamp of most recent heartbeat indicating the batch prediction job was
	// making progress.
	LastHeartbeatTime *string

	// The Amazon S3 location of your output file.
	OutputPath *string

	// The number of records processed by the batch prediction job.
	ProcessedRecordsCount *int32

	// Timestamp of when the batch prediction job started.
	StartTime *string

	// The batch prediction status.
	Status AsyncJobStatus

	// The total number of records in the batch prediction job.
	TotalRecordsCount *int32

	noSmithyDocumentSerde
}

// The model training data validation metrics.
type DataValidationMetrics struct {

	// The field-specific model training validation messages.
	FieldLevelMessages []FieldValidationMessage

	// The file-specific model training data validation messages.
	FileLevelMessages []FileValidationMessage

	noSmithyDocumentSerde
}

// The detector.
type Detector struct {

	// The detector ARN.
	Arn *string

	// Timestamp of when the detector was created.
	CreatedTime *string

	// The detector description.
	Description *string

	// The detector ID.
	DetectorId *string

	// The name of the event type.
	EventTypeName *string

	// Timestamp of when the detector was last updated.
	LastUpdatedTime *string

	noSmithyDocumentSerde
}

// The summary of the detector version.
type DetectorVersionSummary struct {

	// The detector version description.
	Description *string

	// The detector version ID.
	DetectorVersionId *string

	// Timestamp of when the detector version was last updated.
	LastUpdatedTime *string

	// The detector version status.
	Status DetectorVersionStatus

	noSmithyDocumentSerde
}

// The entity details.
type Entity struct {

	// The entity ID. If you do not know the entityId, you can pass unknown, which is
	// areserved string literal.
	//
	// This member is required.
	EntityId *string

	// The entity type.
	//
	// This member is required.
	EntityType *string

	noSmithyDocumentSerde
}

// The entity type details.
type EntityType struct {

	// The entity type ARN.
	Arn *string

	// Timestamp of when the entity type was created.
	CreatedTime *string

	// The entity type description.
	Description *string

	// Timestamp of when the entity type was last updated.
	LastUpdatedTime *string

	// The entity type name.
	Name *string

	noSmithyDocumentSerde
}

// The details of the external (Amazon Sagemaker) model evaluated for generating
// predictions.
type EvaluatedExternalModel struct {

	// Input variables use for generating predictions.
	InputVariables map[string]string

	// The endpoint of the external (Amazon Sagemaker) model.
	ModelEndpoint *string

	// Output variables.
	OutputVariables map[string]string

	// Indicates whether event variables were used to generate predictions.
	UseEventVariables *bool

	noSmithyDocumentSerde
}

// The model version evaluated for generating prediction.
type EvaluatedModelVersion struct {

	// Evaluations generated for the model version.
	Evaluations []ModelVersionEvaluation

	// The model ID.
	ModelId *string

	// The model type. Valid values: ONLINE_FRAUD_INSIGHTS | TRANSACTION_FRAUD_INSIGHTS
	ModelType *string

	// The model version.
	ModelVersion *string

	noSmithyDocumentSerde
}

// The details of the rule used for evaluating variable values.
type EvaluatedRule struct {

	// Indicates whether the rule was evaluated.
	Evaluated *bool

	// The rule expression.
	Expression *string

	// The rule expression value.
	ExpressionWithValues *string

	// Indicates whether the rule matched.
	Matched *bool

	// The rule outcome.
	Outcomes []string

	// The rule ID.
	RuleId *string

	// The rule version.
	RuleVersion *string

	noSmithyDocumentSerde
}

// The event details.
type Event struct {

	// The label associated with the event.
	CurrentLabel *string

	// The event entities.
	Entities []Entity

	// The event ID.
	EventId *string

	// The timestamp that defines when the event under evaluation occurred. The
	// timestamp must be specified using ISO 8601 standard in UTC.
	EventTimestamp *string

	// The event type.
	EventTypeName *string

	// Names of the event type's variables you defined in Amazon Fraud Detector to
	// represent data elements and their corresponding values for the event you are
	// sending for evaluation.
	EventVariables map[string]string

	// The timestamp associated with the label to update. The timestamp must be
	// specified using ISO 8601 standard in UTC.
	LabelTimestamp *string

	noSmithyDocumentSerde
}

// Information about the summary of an event prediction.
type EventPredictionSummary struct {

	// The detector ID.
	DetectorId *string

	// The detector version ID.
	DetectorVersionId *string

	// The event ID.
	EventId *string

	// The timestamp of the event.
	EventTimestamp *string

	// The event type.
	EventTypeName *string

	// The timestamp when the prediction was generated.
	PredictionTimestamp *string

	noSmithyDocumentSerde
}

// The event type details.
type EventType struct {

	// The entity type ARN.
	Arn *string

	// Timestamp of when the event type was created.
	CreatedTime *string

	// The event type description.
	Description *string

	// The event type entity types.
	EntityTypes []string

	// If Enabled, Amazon Fraud Detector stores event data when you generate a
	// prediction and uses that data to update calculated variables in near real-time.
	// Amazon Fraud Detector uses this data, known as INGESTED_EVENTS, to train your
	// model and improve fraud predictions.
	EventIngestion EventIngestion

	// The event type event variables.
	EventVariables []string

	// Data about the stored events.
	IngestedEventStatistics *IngestedEventStatistics

	// The event type labels.
	Labels []string

	// Timestamp of when the event type was last updated.
	LastUpdatedTime *string

	// The event type name.
	Name *string

	noSmithyDocumentSerde
}

// Information about the summary of an event variable that was evaluated for
// generating prediction.
type EventVariableSummary struct {

	// The event variable name.
	Name *string

	// The event variable source.
	Source *string

	// The value of the event variable.
	Value *string

	noSmithyDocumentSerde
}

// Details for the external events data used for model version training.
type ExternalEventsDetail struct {

	// The ARN of the role that provides Amazon Fraud Detector access to the data
	// location.
	//
	// This member is required.
	DataAccessRoleArn *string

	// The Amazon S3 bucket location for the data.
	//
	// This member is required.
	DataLocation *string

	noSmithyDocumentSerde
}

// The Amazon SageMaker model.
type ExternalModel struct {

	// The model ARN.
	Arn *string

	// Timestamp of when the model was last created.
	CreatedTime *string

	// The input configuration.
	InputConfiguration *ModelInputConfiguration

	// The role used to invoke the model.
	InvokeModelEndpointRoleArn *string

	// Timestamp of when the model was last updated.
	LastUpdatedTime *string

	// The Amazon SageMaker model endpoints.
	ModelEndpoint *string

	// The Amazon Fraud Detector status for the external model endpoint
	ModelEndpointStatus ModelEndpointStatus

	// The source of the model.
	ModelSource ModelSource

	// The output configuration.
	OutputConfiguration *ModelOutputConfiguration

	noSmithyDocumentSerde
}

// The fraud prediction scores from Amazon SageMaker model.
type ExternalModelOutputs struct {

	// The Amazon SageMaker model.
	ExternalModel *ExternalModelSummary

	// The fraud prediction scores from Amazon SageMaker model.
	Outputs map[string]string

	noSmithyDocumentSerde
}

// The Amazon SageMaker model.
type ExternalModelSummary struct {

	// The endpoint of the Amazon SageMaker model.
	ModelEndpoint *string

	// The source of the model.
	ModelSource ModelSource

	noSmithyDocumentSerde
}

// The message details.
type FieldValidationMessage struct {

	// The message content.
	Content *string

	// The field name.
	FieldName *string

	// The message ID.
	Identifier *string

	// The message title.
	Title *string

	// The message type.
	Type *string

	noSmithyDocumentSerde
}

// The message details.
type FileValidationMessage struct {

	// The message content.
	Content *string

	// The message title.
	Title *string

	// The message type.
	Type *string

	noSmithyDocumentSerde
}

// A conditional statement for filtering a list of past predictions.
type FilterCondition struct {

	// A statement containing a resource property and a value to specify filter
	// condition.
	Value *string

	noSmithyDocumentSerde
}

// The details of the ingested event.
type IngestedEventsDetail struct {

	// The start and stop time of the ingested events.
	//
	// This member is required.
	IngestedEventsTimeWindow *IngestedEventsTimeWindow

	noSmithyDocumentSerde
}

// Data about the stored events.
type IngestedEventStatistics struct {

	// The total size of the stored events.
	EventDataSizeInBytes *int64

	// Timestamp of when the stored event was last updated.
	LastUpdatedTime *string

	// The oldest stored event.
	LeastRecentEvent *string

	// The newest stored event.
	MostRecentEvent *string

	// The number of stored events.
	NumberOfEvents *int64

	noSmithyDocumentSerde
}

// The start and stop time of the ingested events.
type IngestedEventsTimeWindow struct {

	// Timestamp of the final ingested event.
	//
	// This member is required.
	EndTime *string

	// Timestamp of the first ingensted event.
	//
	// This member is required.
	StartTime *string

	noSmithyDocumentSerde
}

// The KMS key details.
type KMSKey struct {

	// The encryption key ARN.
	KmsEncryptionKeyArn *string

	noSmithyDocumentSerde
}

// The label details.
type Label struct {

	// The label ARN.
	Arn *string

	// Timestamp of when the event type was created.
	CreatedTime *string

	// The label description.
	Description *string

	// Timestamp of when the label was last updated.
	LastUpdatedTime *string

	// The label name.
	Name *string

	noSmithyDocumentSerde
}

// The label schema.
type LabelSchema struct {

	// The label mapper maps the Amazon Fraud Detector supported model classification
	// labels (FRAUD, LEGIT) to the appropriate event type labels. For example, if
	// "FRAUD" and "LEGIT" are Amazon Fraud Detector supported labels, this mapper
	// could be: {"FRAUD" => ["0"], "LEGIT" => ["1"]} or {"FRAUD" => ["false"], "LEGIT"
	// => ["true"]} or {"FRAUD" => ["fraud", "abuse"], "LEGIT" => ["legit", "safe"]}.
	// The value part of the mapper is a list, because you may have multiple label
	// variants from your event type for a single Amazon Fraud Detector label.
	LabelMapper map[string][]string

	// The action to take for unlabeled events.
	//
	// * Use IGNORE if you want the unlabeled
	// events to be ignored. This is recommended when the majority of the events in the
	// dataset are labeled.
	//
	// * Use FRAUD if you want to categorize all unlabeled events
	// as “Fraud”. This is recommended when most of the events in your dataset are
	// fraudulent.
	//
	// * Use LEGIT f you want to categorize all unlabeled events as
	// “Legit”. This is recommended when most of the events in your dataset are
	// legitimate.
	//
	// * Use AUTO if you want Amazon Fraud Detector to decide how to use
	// the unlabeled data. This is recommended when there is significant unlabeled
	// events in the dataset.
	//
	// By default, Amazon Fraud Detector ignores the unlabeled
	// data.
	UnlabeledEventsTreatment UnlabeledEventsTreatment

	noSmithyDocumentSerde
}

// The log odds metric details.
type LogOddsMetric struct {

	// The relative importance of the variable. For more information, see Model
	// variable importance
	// (https://docs.aws.amazon.com/frauddetector/latest/ug/model-variable-importance.html).
	//
	// This member is required.
	VariableImportance *float32

	// The name of the variable.
	//
	// This member is required.
	VariableName *string

	// The type of variable.
	//
	// This member is required.
	VariableType *string

	noSmithyDocumentSerde
}

// Model performance metrics data points.
type MetricDataPoint struct {

	// The false positive rate. This is the percentage of total legitimate events that
	// are incorrectly predicted as fraud.
	Fpr *float32

	// The percentage of fraud events correctly predicted as fraudulent as compared to
	// all events predicted as fraudulent.
	Precision *float32

	// The model threshold that specifies an acceptable fraud capture rate. For
	// example, a threshold of 500 means any model score 500 or above is labeled as
	// fraud.
	Threshold *float32

	// The true positive rate. This is the percentage of total fraud the model detects.
	// Also known as capture rate.
	Tpr *float32

	noSmithyDocumentSerde
}

// The model.
type Model struct {

	// The ARN of the model.
	Arn *string

	// Timestamp of when the model was created.
	CreatedTime *string

	// The model description.
	Description *string

	// The name of the event type.
	EventTypeName *string

	// Timestamp of last time the model was updated.
	LastUpdatedTime *string

	// The model ID.
	ModelId *string

	// The model type.
	ModelType ModelTypeEnum

	noSmithyDocumentSerde
}

// A pre-formed Amazon SageMaker model input you can include if your detector
// version includes an imported Amazon SageMaker model endpoint with pass-through
// input configuration.
type ModelEndpointDataBlob struct {

	// The byte buffer of the Amazon SageMaker model endpoint input data blob.
	ByteBuffer []byte

	// The content type of the Amazon SageMaker model endpoint input data blob.
	ContentType *string

	noSmithyDocumentSerde
}

// The Amazon SageMaker model input configuration.
type ModelInputConfiguration struct {

	// The event variables.
	//
	// This member is required.
	UseEventVariables *bool

	// Template for constructing the CSV input-data sent to SageMaker. At
	// event-evaluation, the placeholders for variable-names in the template will be
	// replaced with the variable values before being sent to SageMaker.
	CsvInputTemplate *string

	// The event type name.
	EventTypeName *string

	// The format of the model input configuration. The format differs depending on if
	// it is passed through to SageMaker or constructed by Amazon Fraud Detector.
	Format ModelInputDataFormat

	// Template for constructing the JSON input-data sent to SageMaker. At
	// event-evaluation, the placeholders for variable names in the template will be
	// replaced with the variable values before being sent to SageMaker.
	JsonInputTemplate *string

	noSmithyDocumentSerde
}

// Provides the Amazon Sagemaker model output configuration.
type ModelOutputConfiguration struct {

	// The format of the model output configuration.
	//
	// This member is required.
	Format ModelOutputDataFormat

	// A map of CSV index values in the SageMaker response to the Amazon Fraud Detector
	// variables.
	CsvIndexToVariableMap map[string]string

	// A map of JSON keys in response from SageMaker to the Amazon Fraud Detector
	// variables.
	JsonKeyToVariableMap map[string]string

	noSmithyDocumentSerde
}

// The fraud prediction scores.
type ModelScores struct {

	// The model version.
	ModelVersion *ModelVersion

	// The model's fraud prediction scores.
	Scores map[string]float32

	noSmithyDocumentSerde
}

// The model version.
type ModelVersion struct {

	// The model ID.
	//
	// This member is required.
	ModelId *string

	// The model type.
	//
	// This member is required.
	ModelType ModelTypeEnum

	// The model version number.
	//
	// This member is required.
	ModelVersionNumber *string

	// The model version ARN.
	Arn *string

	noSmithyDocumentSerde
}

// The details of the model version.
type ModelVersionDetail struct {

	// The model version ARN.
	Arn *string

	// The timestamp when the model was created.
	CreatedTime *string

	// The external events data details. This will be populated if the
	// trainingDataSource for the model version is specified as EXTERNAL_EVENTS.
	ExternalEventsDetail *ExternalEventsDetail

	// The ingested events data details. This will be populated if the
	// trainingDataSource for the model version is specified as INGESTED_EVENTS.
	IngestedEventsDetail *IngestedEventsDetail

	// The timestamp when the model was last updated.
	LastUpdatedTime *string

	// The model ID.
	ModelId *string

	// The model type.
	ModelType ModelTypeEnum

	// The model version number.
	ModelVersionNumber *string

	// The status of the model version.
	Status *string

	// The training data schema.
	TrainingDataSchema *TrainingDataSchema

	// The model version training data source.
	TrainingDataSource TrainingDataSourceEnum

	// The training results.
	TrainingResult *TrainingResult

	// The training result details. The details include the relative importance of the
	// variables.
	TrainingResultV2 *TrainingResultV2

	noSmithyDocumentSerde
}

// The model version evalutions.
type ModelVersionEvaluation struct {

	// The evaluation score generated for the model version.
	EvaluationScore *string

	// The output variable name.
	OutputVariableName *string

	// The prediction explanations generated for the model version.
	PredictionExplanations *PredictionExplanations

	noSmithyDocumentSerde
}

// The Online Fraud Insights (OFI) model performance metrics data points.
type OFIMetricDataPoint struct {

	// The false positive rate. This is the percentage of total legitimate events that
	// are incorrectly predicted as fraud.
	Fpr *float32

	// The percentage of fraud events correctly predicted as fraudulent as compared to
	// all events predicted as fraudulent.
	Precision *float32

	// The model threshold that specifies an acceptable fraud capture rate. For
	// example, a threshold of 500 means any model score 500 or above is labeled as
	// fraud.
	Threshold *float32

	// The true positive rate. This is the percentage of total fraud the model detects.
	// Also known as capture rate.
	Tpr *float32

	noSmithyDocumentSerde
}

// The Online Fraud Insights (OFI) model performance score.
type OFIModelPerformance struct {

	// The area under the curve (auc). This summarizes the total positive rate (tpr)
	// and false positive rate (FPR) across all possible model score thresholds.
	Auc *float32

	// Indicates the range of area under curve (auc) expected from the OFI model. A
	// range greater than 0.1 indicates higher model uncertainity.
	UncertaintyRange *UncertaintyRange

	noSmithyDocumentSerde
}

// The Online Fraud Insights (OFI) model training metric details.
type OFITrainingMetricsValue struct {

	// The model's performance metrics data points.
	MetricDataPoints []OFIMetricDataPoint

	// The model's overall performance score.
	ModelPerformance *OFIModelPerformance

	noSmithyDocumentSerde
}

// The outcome.
type Outcome struct {

	// The outcome ARN.
	Arn *string

	// The timestamp when the outcome was created.
	CreatedTime *string

	// The outcome description.
	Description *string

	// The timestamp when the outcome was last updated.
	LastUpdatedTime *string

	// The outcome name.
	Name *string

	noSmithyDocumentSerde
}

// The prediction explanations that provide insight into how each event variable
// impacted the model version's fraud prediction score.
type PredictionExplanations struct {

	// The details of the aggregated variables impact on the prediction score. Account
	// Takeover Insights (ATI) model uses event variables from the login data you
	// provide to continuously calculate a set of variables (aggregated variables)
	// based on historical events. For example, your ATI model might calculate the
	// number of times an user has logged in using the same IP address. In this case,
	// event variables used to derive the aggregated variables are IP address and user.
	AggregatedVariablesImpactExplanations []AggregatedVariablesImpactExplanation

	// The details of the event variable's impact on the prediction score.
	VariableImpactExplanations []VariableImpactExplanation

	noSmithyDocumentSerde
}

// The time period for when the predictions were generated.
type PredictionTimeRange struct {

	// The end time of the time period for when the predictions were generated.
	//
	// This member is required.
	EndTime *string

	// The start time of the time period for when the predictions were generated.
	//
	// This member is required.
	StartTime *string

	noSmithyDocumentSerde
}

// A rule.
type Rule struct {

	// The detector for which the rule is associated.
	//
	// This member is required.
	DetectorId *string

	// The rule ID.
	//
	// This member is required.
	RuleId *string

	// The rule version.
	//
	// This member is required.
	RuleVersion *string

	noSmithyDocumentSerde
}

// The details of the rule.
type RuleDetail struct {

	// The rule ARN.
	Arn *string

	// The timestamp of when the rule was created.
	CreatedTime *string

	// The rule description.
	Description *string

	// The detector for which the rule is associated.
	DetectorId *string

	// The rule expression.
	Expression *string

	// The rule language.
	Language Language

	// Timestamp of the last time the rule was updated.
	LastUpdatedTime *string

	// The rule outcomes.
	Outcomes []string

	// The rule ID.
	RuleId *string

	// The rule version.
	RuleVersion *string

	noSmithyDocumentSerde
}

// The rule results.
type RuleResult struct {

	// The outcomes of the matched rule, based on the rule execution mode.
	Outcomes []string

	// The rule ID that was matched, based on the rule execution mode.
	RuleId *string

	noSmithyDocumentSerde
}

// A key and value pair.
type Tag struct {

	// A tag key.
	//
	// This member is required.
	Key *string

	// A value assigned to a tag key.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The performance metrics data points for Transaction Fraud Insights (TFI) model.
type TFIMetricDataPoint struct {

	// The false positive rate. This is the percentage of total legitimate events that
	// are incorrectly predicted as fraud.
	Fpr *float32

	// The percentage of fraud events correctly predicted as fraudulent as compared to
	// all events predicted as fraudulent.
	Precision *float32

	// The model threshold that specifies an acceptable fraud capture rate. For
	// example, a threshold of 500 means any model score 500 or above is labeled as
	// fraud.
	Threshold *float32

	// The true positive rate. This is the percentage of total fraud the model detects.
	// Also known as capture rate.
	Tpr *float32

	noSmithyDocumentSerde
}

// The Transaction Fraud Insights (TFI) model performance score.
type TFIModelPerformance struct {

	// The area under the curve (auc). This summarizes the total positive rate (tpr)
	// and false positive rate (FPR) across all possible model score thresholds.
	Auc *float32

	// Indicates the range of area under curve (auc) expected from the TFI model. A
	// range greater than 0.1 indicates higher model uncertainity.
	UncertaintyRange *UncertaintyRange

	noSmithyDocumentSerde
}

// The Transaction Fraud Insights (TFI) model training metric details.
type TFITrainingMetricsValue struct {

	// The model's performance metrics data points.
	MetricDataPoints []TFIMetricDataPoint

	// The model performance score.
	ModelPerformance *TFIModelPerformance

	noSmithyDocumentSerde
}

// The training data schema.
type TrainingDataSchema struct {

	// The training data schema variables.
	//
	// This member is required.
	ModelVariables []string

	// The label schema.
	LabelSchema *LabelSchema

	noSmithyDocumentSerde
}

// The training metric details.
type TrainingMetrics struct {

	// The area under the curve. This summarizes true positive rate (TPR) and false
	// positive rate (FPR) across all possible model score thresholds. A model with no
	// predictive power has an AUC of 0.5, whereas a perfect model has a score of 1.0.
	Auc *float32

	// The data points details.
	MetricDataPoints []MetricDataPoint

	noSmithyDocumentSerde
}

// The training metrics details.
type TrainingMetricsV2 struct {

	// The Account Takeover Insights (ATI) model training metric details.
	Ati *ATITrainingMetricsValue

	// The Online Fraud Insights (OFI) model training metric details.
	Ofi *OFITrainingMetricsValue

	// The Transaction Fraud Insights (TFI) model training metric details.
	Tfi *TFITrainingMetricsValue

	noSmithyDocumentSerde
}

// The training result details.
type TrainingResult struct {

	// The validation metrics.
	DataValidationMetrics *DataValidationMetrics

	// The training metric details.
	TrainingMetrics *TrainingMetrics

	// The variable importance metrics.
	VariableImportanceMetrics *VariableImportanceMetrics

	noSmithyDocumentSerde
}

// The training result details.
type TrainingResultV2 struct {

	// The variable importance metrics of the aggregated variables. Account Takeover
	// Insights (ATI) model uses event variables from the login data you provide to
	// continuously calculate a set of variables (aggregated variables) based on
	// historical events. For example, your ATI model might calculate the number of
	// times an user has logged in using the same IP address. In this case, event
	// variables used to derive the aggregated variables are IP address and user.
	AggregatedVariablesImportanceMetrics *AggregatedVariablesImportanceMetrics

	// The model training data validation metrics.
	DataValidationMetrics *DataValidationMetrics

	// The training metric details.
	TrainingMetricsV2 *TrainingMetricsV2

	// The variable importance metrics details.
	VariableImportanceMetrics *VariableImportanceMetrics

	noSmithyDocumentSerde
}

// Range of area under curve (auc) expected from the model. A range greater than
// 0.1 indicates higher model uncertainity. A range is the difference between upper
// and lower bound of auc.
type UncertaintyRange struct {

	// The lower bound value of the area under curve (auc).
	//
	// This member is required.
	LowerBoundValue *float32

	// The lower bound value of the area under curve (auc).
	//
	// This member is required.
	UpperBoundValue *float32

	noSmithyDocumentSerde
}

// The variable.
type Variable struct {

	// The ARN of the variable.
	Arn *string

	// The time when the variable was created.
	CreatedTime *string

	// The data source of the variable.
	DataSource DataSource

	// The data type of the variable. For more information see Variable types
	// (https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types).
	DataType DataType

	// The default value of the variable.
	DefaultValue *string

	// The description of the variable.
	Description *string

	// The time when variable was last updated.
	LastUpdatedTime *string

	// The name of the variable.
	Name *string

	// The variable type of the variable. Valid Values: AUTH_CODE | AVS |
	// BILLING_ADDRESS_L1 | BILLING_ADDRESS_L2 | BILLING_CITY | BILLING_COUNTRY |
	// BILLING_NAME | BILLING_PHONE | BILLING_STATE | BILLING_ZIP | CARD_BIN |
	// CATEGORICAL | CURRENCY_CODE | EMAIL_ADDRESS | FINGERPRINT | FRAUD_LABEL |
	// FREE_FORM_TEXT | IP_ADDRESS | NUMERIC | ORDER_ID | PAYMENT_TYPE | PHONE_NUMBER |
	// PRICE | PRODUCT_CATEGORY | SHIPPING_ADDRESS_L1 | SHIPPING_ADDRESS_L2 |
	// SHIPPING_CITY | SHIPPING_COUNTRY | SHIPPING_NAME | SHIPPING_PHONE |
	// SHIPPING_STATE | SHIPPING_ZIP | USERAGENT
	VariableType *string

	noSmithyDocumentSerde
}

// A variable in the list of variables for the batch create variable request.
type VariableEntry struct {

	// The data source of the variable.
	DataSource *string

	// The data type of the variable.
	DataType *string

	// The default value of the variable.
	DefaultValue *string

	// The description of the variable.
	Description *string

	// The name of the variable.
	Name *string

	// The type of the variable. For more information see Variable types
	// (https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types).
	// Valid Values: AUTH_CODE | AVS | BILLING_ADDRESS_L1 | BILLING_ADDRESS_L2 |
	// BILLING_CITY | BILLING_COUNTRY | BILLING_NAME | BILLING_PHONE | BILLING_STATE |
	// BILLING_ZIP | CARD_BIN | CATEGORICAL | CURRENCY_CODE | EMAIL_ADDRESS |
	// FINGERPRINT | FRAUD_LABEL | FREE_FORM_TEXT | IP_ADDRESS | NUMERIC | ORDER_ID |
	// PAYMENT_TYPE | PHONE_NUMBER | PRICE | PRODUCT_CATEGORY | SHIPPING_ADDRESS_L1 |
	// SHIPPING_ADDRESS_L2 | SHIPPING_CITY | SHIPPING_COUNTRY | SHIPPING_NAME |
	// SHIPPING_PHONE | SHIPPING_STATE | SHIPPING_ZIP | USERAGENT
	VariableType *string

	noSmithyDocumentSerde
}

// The details of the event variable's impact on the prediction score.
type VariableImpactExplanation struct {

	// The event variable name.
	EventVariableName *string

	// The raw, uninterpreted value represented as log-odds of the fraud. These values
	// are usually between -10 to +10, but range from - infinity to + infinity.
	//
	// * A
	// positive value indicates that the variable drove the risk score up.
	//
	// * A
	// negative value indicates that the variable drove the risk score down.
	LogOddsImpact *float32

	// The event variable's relative impact in terms of magnitude on the prediction
	// scores. The relative impact values consist of a numerical rating (0-5, 5 being
	// the highest) and direction (increased/decreased) impact of the fraud risk.
	RelativeImpact *string

	noSmithyDocumentSerde
}

// The variable importance metrics details.
type VariableImportanceMetrics struct {

	// List of variable metrics.
	LogOddsMetrics []LogOddsMetric

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
