package cloudtrail

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// CreateTrailRequest generates a request for the CreateTrail operation.
func (c *CloudTrail) CreateTrailRequest(input *CreateTrailInput) (req *aws.Request, output *CreateTrailOutput) {
	if opCreateTrail == nil {
		opCreateTrail = &aws.Operation{
			Name:       "CreateTrail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreateTrail, input, output)
	output = &CreateTrailOutput{}
	req.Data = output
	return
}

func (c *CloudTrail) CreateTrail(input *CreateTrailInput) (output *CreateTrailOutput, err error) {
	req, out := c.CreateTrailRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateTrail *aws.Operation

// DeleteTrailRequest generates a request for the DeleteTrail operation.
func (c *CloudTrail) DeleteTrailRequest(input *DeleteTrailInput) (req *aws.Request, output *DeleteTrailOutput) {
	if opDeleteTrail == nil {
		opDeleteTrail = &aws.Operation{
			Name:       "DeleteTrail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteTrail, input, output)
	output = &DeleteTrailOutput{}
	req.Data = output
	return
}

func (c *CloudTrail) DeleteTrail(input *DeleteTrailInput) (output *DeleteTrailOutput, err error) {
	req, out := c.DeleteTrailRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteTrail *aws.Operation

// DescribeTrailsRequest generates a request for the DescribeTrails operation.
func (c *CloudTrail) DescribeTrailsRequest(input *DescribeTrailsInput) (req *aws.Request, output *DescribeTrailsOutput) {
	if opDescribeTrails == nil {
		opDescribeTrails = &aws.Operation{
			Name:       "DescribeTrails",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeTrails, input, output)
	output = &DescribeTrailsOutput{}
	req.Data = output
	return
}

func (c *CloudTrail) DescribeTrails(input *DescribeTrailsInput) (output *DescribeTrailsOutput, err error) {
	req, out := c.DescribeTrailsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeTrails *aws.Operation

// GetTrailStatusRequest generates a request for the GetTrailStatus operation.
func (c *CloudTrail) GetTrailStatusRequest(input *GetTrailStatusInput) (req *aws.Request, output *GetTrailStatusOutput) {
	if opGetTrailStatus == nil {
		opGetTrailStatus = &aws.Operation{
			Name:       "GetTrailStatus",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetTrailStatus, input, output)
	output = &GetTrailStatusOutput{}
	req.Data = output
	return
}

func (c *CloudTrail) GetTrailStatus(input *GetTrailStatusInput) (output *GetTrailStatusOutput, err error) {
	req, out := c.GetTrailStatusRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetTrailStatus *aws.Operation

// StartLoggingRequest generates a request for the StartLogging operation.
func (c *CloudTrail) StartLoggingRequest(input *StartLoggingInput) (req *aws.Request, output *StartLoggingOutput) {
	if opStartLogging == nil {
		opStartLogging = &aws.Operation{
			Name:       "StartLogging",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opStartLogging, input, output)
	output = &StartLoggingOutput{}
	req.Data = output
	return
}

func (c *CloudTrail) StartLogging(input *StartLoggingInput) (output *StartLoggingOutput, err error) {
	req, out := c.StartLoggingRequest(input)
	output = out
	err = req.Send()
	return
}

var opStartLogging *aws.Operation

// StopLoggingRequest generates a request for the StopLogging operation.
func (c *CloudTrail) StopLoggingRequest(input *StopLoggingInput) (req *aws.Request, output *StopLoggingOutput) {
	if opStopLogging == nil {
		opStopLogging = &aws.Operation{
			Name:       "StopLogging",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opStopLogging, input, output)
	output = &StopLoggingOutput{}
	req.Data = output
	return
}

func (c *CloudTrail) StopLogging(input *StopLoggingInput) (output *StopLoggingOutput, err error) {
	req, out := c.StopLoggingRequest(input)
	output = out
	err = req.Send()
	return
}

var opStopLogging *aws.Operation

// UpdateTrailRequest generates a request for the UpdateTrail operation.
func (c *CloudTrail) UpdateTrailRequest(input *UpdateTrailInput) (req *aws.Request, output *UpdateTrailOutput) {
	if opUpdateTrail == nil {
		opUpdateTrail = &aws.Operation{
			Name:       "UpdateTrail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateTrail, input, output)
	output = &UpdateTrailOutput{}
	req.Data = output
	return
}

func (c *CloudTrail) UpdateTrail(input *UpdateTrailInput) (output *UpdateTrailOutput, err error) {
	req, out := c.UpdateTrailRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateTrail *aws.Operation

type CreateTrailInput struct {
	CloudWatchLogsLogGroupARN  *string `locationName:"CloudWatchLogsLogGroupArn" type:"string" json:"CloudWatchLogsLogGroupArn,omitempty"`
	CloudWatchLogsRoleARN      *string `locationName:"CloudWatchLogsRoleArn" type:"string" json:"CloudWatchLogsRoleArn,omitempty"`
	IncludeGlobalServiceEvents *bool   `type:"boolean" json:",omitempty"`
	Name                       *string `type:"string" required:"true"json:",omitempty"`
	S3BucketName               *string `type:"string" required:"true"json:",omitempty"`
	S3KeyPrefix                *string `type:"string" json:",omitempty"`
	SNSTopicName               *string `locationName:"SnsTopicName" type:"string" json:"SnsTopicName,omitempty"`

	metadataCreateTrailInput `json:"-", xml:"-"`
}

type metadataCreateTrailInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type CreateTrailOutput struct {
	CloudWatchLogsLogGroupARN  *string `locationName:"CloudWatchLogsLogGroupArn" type:"string" json:"CloudWatchLogsLogGroupArn,omitempty"`
	CloudWatchLogsRoleARN      *string `locationName:"CloudWatchLogsRoleArn" type:"string" json:"CloudWatchLogsRoleArn,omitempty"`
	IncludeGlobalServiceEvents *bool   `type:"boolean" json:",omitempty"`
	Name                       *string `type:"string" json:",omitempty"`
	S3BucketName               *string `type:"string" json:",omitempty"`
	S3KeyPrefix                *string `type:"string" json:",omitempty"`
	SNSTopicName               *string `locationName:"SnsTopicName" type:"string" json:"SnsTopicName,omitempty"`

	metadataCreateTrailOutput `json:"-", xml:"-"`
}

type metadataCreateTrailOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteTrailInput struct {
	Name *string `type:"string" required:"true"json:",omitempty"`

	metadataDeleteTrailInput `json:"-", xml:"-"`
}

type metadataDeleteTrailInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteTrailOutput struct {
	metadataDeleteTrailOutput `json:"-", xml:"-"`
}

type metadataDeleteTrailOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeTrailsInput struct {
	TrailNameList []*string `locationName:"trailNameList" type:"list" json:"trailNameList,omitempty"`

	metadataDescribeTrailsInput `json:"-", xml:"-"`
}

type metadataDescribeTrailsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeTrailsOutput struct {
	TrailList []*Trail `locationName:"trailList" type:"list" json:"trailList,omitempty"`

	metadataDescribeTrailsOutput `json:"-", xml:"-"`
}

type metadataDescribeTrailsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetTrailStatusInput struct {
	Name *string `type:"string" required:"true"json:",omitempty"`

	metadataGetTrailStatusInput `json:"-", xml:"-"`
}

type metadataGetTrailStatusInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetTrailStatusOutput struct {
	IsLogging                         *bool      `type:"boolean" json:",omitempty"`
	LatestCloudWatchLogsDeliveryError *string    `type:"string" json:",omitempty"`
	LatestCloudWatchLogsDeliveryTime  *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	LatestDeliveryError               *string    `type:"string" json:",omitempty"`
	LatestDeliveryTime                *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	LatestNotificationError           *string    `type:"string" json:",omitempty"`
	LatestNotificationTime            *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	StartLoggingTime                  *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	StopLoggingTime                   *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`

	metadataGetTrailStatusOutput `json:"-", xml:"-"`
}

type metadataGetTrailStatusOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type StartLoggingInput struct {
	Name *string `type:"string" required:"true"json:",omitempty"`

	metadataStartLoggingInput `json:"-", xml:"-"`
}

type metadataStartLoggingInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type StartLoggingOutput struct {
	metadataStartLoggingOutput `json:"-", xml:"-"`
}

type metadataStartLoggingOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type StopLoggingInput struct {
	Name *string `type:"string" required:"true"json:",omitempty"`

	metadataStopLoggingInput `json:"-", xml:"-"`
}

type metadataStopLoggingInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type StopLoggingOutput struct {
	metadataStopLoggingOutput `json:"-", xml:"-"`
}

type metadataStopLoggingOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type Trail struct {
	CloudWatchLogsLogGroupARN  *string `locationName:"CloudWatchLogsLogGroupArn" type:"string" json:"CloudWatchLogsLogGroupArn,omitempty"`
	CloudWatchLogsRoleARN      *string `locationName:"CloudWatchLogsRoleArn" type:"string" json:"CloudWatchLogsRoleArn,omitempty"`
	IncludeGlobalServiceEvents *bool   `type:"boolean" json:",omitempty"`
	Name                       *string `type:"string" json:",omitempty"`
	S3BucketName               *string `type:"string" json:",omitempty"`
	S3KeyPrefix                *string `type:"string" json:",omitempty"`
	SNSTopicName               *string `locationName:"SnsTopicName" type:"string" json:"SnsTopicName,omitempty"`

	metadataTrail `json:"-", xml:"-"`
}

type metadataTrail struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UpdateTrailInput struct {
	CloudWatchLogsLogGroupARN  *string `locationName:"CloudWatchLogsLogGroupArn" type:"string" json:"CloudWatchLogsLogGroupArn,omitempty"`
	CloudWatchLogsRoleARN      *string `locationName:"CloudWatchLogsRoleArn" type:"string" json:"CloudWatchLogsRoleArn,omitempty"`
	IncludeGlobalServiceEvents *bool   `type:"boolean" json:",omitempty"`
	Name                       *string `type:"string" required:"true"json:",omitempty"`
	S3BucketName               *string `type:"string" json:",omitempty"`
	S3KeyPrefix                *string `type:"string" json:",omitempty"`
	SNSTopicName               *string `locationName:"SnsTopicName" type:"string" json:"SnsTopicName,omitempty"`

	metadataUpdateTrailInput `json:"-", xml:"-"`
}

type metadataUpdateTrailInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UpdateTrailOutput struct {
	CloudWatchLogsLogGroupARN  *string `locationName:"CloudWatchLogsLogGroupArn" type:"string" json:"CloudWatchLogsLogGroupArn,omitempty"`
	CloudWatchLogsRoleARN      *string `locationName:"CloudWatchLogsRoleArn" type:"string" json:"CloudWatchLogsRoleArn,omitempty"`
	IncludeGlobalServiceEvents *bool   `type:"boolean" json:",omitempty"`
	Name                       *string `type:"string" json:",omitempty"`
	S3BucketName               *string `type:"string" json:",omitempty"`
	S3KeyPrefix                *string `type:"string" json:",omitempty"`
	SNSTopicName               *string `locationName:"SnsTopicName" type:"string" json:"SnsTopicName,omitempty"`

	metadataUpdateTrailOutput `json:"-", xml:"-"`
}

type metadataUpdateTrailOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}