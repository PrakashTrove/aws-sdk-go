package ses

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// DeleteIdentityRequest generates a request for the DeleteIdentity operation.
func (c *SES) DeleteIdentityRequest(input *DeleteIdentityInput) (req *aws.Request, output *DeleteIdentityOutput) {
	if opDeleteIdentity == nil {
		opDeleteIdentity = &aws.Operation{
			Name:       "DeleteIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteIdentity, input, output)
	output = &DeleteIdentityOutput{}
	req.Data = output
	return
}

func (c *SES) DeleteIdentity(input *DeleteIdentityInput) (output *DeleteIdentityOutput, err error) {
	req, out := c.DeleteIdentityRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteIdentity *aws.Operation

// DeleteVerifiedEmailAddressRequest generates a request for the DeleteVerifiedEmailAddress operation.
func (c *SES) DeleteVerifiedEmailAddressRequest(input *DeleteVerifiedEmailAddressInput) (req *aws.Request, output *DeleteVerifiedEmailAddressOutput) {
	if opDeleteVerifiedEmailAddress == nil {
		opDeleteVerifiedEmailAddress = &aws.Operation{
			Name:       "DeleteVerifiedEmailAddress",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteVerifiedEmailAddress, input, output)
	output = &DeleteVerifiedEmailAddressOutput{}
	req.Data = output
	return
}

func (c *SES) DeleteVerifiedEmailAddress(input *DeleteVerifiedEmailAddressInput) (output *DeleteVerifiedEmailAddressOutput, err error) {
	req, out := c.DeleteVerifiedEmailAddressRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteVerifiedEmailAddress *aws.Operation

// GetIdentityDKIMAttributesRequest generates a request for the GetIdentityDKIMAttributes operation.
func (c *SES) GetIdentityDKIMAttributesRequest(input *GetIdentityDKIMAttributesInput) (req *aws.Request, output *GetIdentityDKIMAttributesOutput) {
	if opGetIdentityDKIMAttributes == nil {
		opGetIdentityDKIMAttributes = &aws.Operation{
			Name:       "GetIdentityDkimAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetIdentityDKIMAttributes, input, output)
	output = &GetIdentityDKIMAttributesOutput{}
	req.Data = output
	return
}

func (c *SES) GetIdentityDKIMAttributes(input *GetIdentityDKIMAttributesInput) (output *GetIdentityDKIMAttributesOutput, err error) {
	req, out := c.GetIdentityDKIMAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetIdentityDKIMAttributes *aws.Operation

// GetIdentityNotificationAttributesRequest generates a request for the GetIdentityNotificationAttributes operation.
func (c *SES) GetIdentityNotificationAttributesRequest(input *GetIdentityNotificationAttributesInput) (req *aws.Request, output *GetIdentityNotificationAttributesOutput) {
	if opGetIdentityNotificationAttributes == nil {
		opGetIdentityNotificationAttributes = &aws.Operation{
			Name:       "GetIdentityNotificationAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetIdentityNotificationAttributes, input, output)
	output = &GetIdentityNotificationAttributesOutput{}
	req.Data = output
	return
}

func (c *SES) GetIdentityNotificationAttributes(input *GetIdentityNotificationAttributesInput) (output *GetIdentityNotificationAttributesOutput, err error) {
	req, out := c.GetIdentityNotificationAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetIdentityNotificationAttributes *aws.Operation

// GetIdentityVerificationAttributesRequest generates a request for the GetIdentityVerificationAttributes operation.
func (c *SES) GetIdentityVerificationAttributesRequest(input *GetIdentityVerificationAttributesInput) (req *aws.Request, output *GetIdentityVerificationAttributesOutput) {
	if opGetIdentityVerificationAttributes == nil {
		opGetIdentityVerificationAttributes = &aws.Operation{
			Name:       "GetIdentityVerificationAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetIdentityVerificationAttributes, input, output)
	output = &GetIdentityVerificationAttributesOutput{}
	req.Data = output
	return
}

func (c *SES) GetIdentityVerificationAttributes(input *GetIdentityVerificationAttributesInput) (output *GetIdentityVerificationAttributesOutput, err error) {
	req, out := c.GetIdentityVerificationAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetIdentityVerificationAttributes *aws.Operation

// GetSendQuotaRequest generates a request for the GetSendQuota operation.
func (c *SES) GetSendQuotaRequest(input *GetSendQuotaInput) (req *aws.Request, output *GetSendQuotaOutput) {
	if opGetSendQuota == nil {
		opGetSendQuota = &aws.Operation{
			Name:       "GetSendQuota",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetSendQuota, input, output)
	output = &GetSendQuotaOutput{}
	req.Data = output
	return
}

func (c *SES) GetSendQuota(input *GetSendQuotaInput) (output *GetSendQuotaOutput, err error) {
	req, out := c.GetSendQuotaRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetSendQuota *aws.Operation

// GetSendStatisticsRequest generates a request for the GetSendStatistics operation.
func (c *SES) GetSendStatisticsRequest(input *GetSendStatisticsInput) (req *aws.Request, output *GetSendStatisticsOutput) {
	if opGetSendStatistics == nil {
		opGetSendStatistics = &aws.Operation{
			Name:       "GetSendStatistics",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetSendStatistics, input, output)
	output = &GetSendStatisticsOutput{}
	req.Data = output
	return
}

func (c *SES) GetSendStatistics(input *GetSendStatisticsInput) (output *GetSendStatisticsOutput, err error) {
	req, out := c.GetSendStatisticsRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetSendStatistics *aws.Operation

// ListIdentitiesRequest generates a request for the ListIdentities operation.
func (c *SES) ListIdentitiesRequest(input *ListIdentitiesInput) (req *aws.Request, output *ListIdentitiesOutput) {
	if opListIdentities == nil {
		opListIdentities = &aws.Operation{
			Name:       "ListIdentities",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListIdentities, input, output)
	output = &ListIdentitiesOutput{}
	req.Data = output
	return
}

func (c *SES) ListIdentities(input *ListIdentitiesInput) (output *ListIdentitiesOutput, err error) {
	req, out := c.ListIdentitiesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListIdentities *aws.Operation

// ListVerifiedEmailAddressesRequest generates a request for the ListVerifiedEmailAddresses operation.
func (c *SES) ListVerifiedEmailAddressesRequest(input *ListVerifiedEmailAddressesInput) (req *aws.Request, output *ListVerifiedEmailAddressesOutput) {
	if opListVerifiedEmailAddresses == nil {
		opListVerifiedEmailAddresses = &aws.Operation{
			Name:       "ListVerifiedEmailAddresses",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListVerifiedEmailAddresses, input, output)
	output = &ListVerifiedEmailAddressesOutput{}
	req.Data = output
	return
}

func (c *SES) ListVerifiedEmailAddresses(input *ListVerifiedEmailAddressesInput) (output *ListVerifiedEmailAddressesOutput, err error) {
	req, out := c.ListVerifiedEmailAddressesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListVerifiedEmailAddresses *aws.Operation

// SendEmailRequest generates a request for the SendEmail operation.
func (c *SES) SendEmailRequest(input *SendEmailInput) (req *aws.Request, output *SendEmailOutput) {
	if opSendEmail == nil {
		opSendEmail = &aws.Operation{
			Name:       "SendEmail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSendEmail, input, output)
	output = &SendEmailOutput{}
	req.Data = output
	return
}

func (c *SES) SendEmail(input *SendEmailInput) (output *SendEmailOutput, err error) {
	req, out := c.SendEmailRequest(input)
	output = out
	err = req.Send()
	return
}

var opSendEmail *aws.Operation

// SendRawEmailRequest generates a request for the SendRawEmail operation.
func (c *SES) SendRawEmailRequest(input *SendRawEmailInput) (req *aws.Request, output *SendRawEmailOutput) {
	if opSendRawEmail == nil {
		opSendRawEmail = &aws.Operation{
			Name:       "SendRawEmail",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSendRawEmail, input, output)
	output = &SendRawEmailOutput{}
	req.Data = output
	return
}

func (c *SES) SendRawEmail(input *SendRawEmailInput) (output *SendRawEmailOutput, err error) {
	req, out := c.SendRawEmailRequest(input)
	output = out
	err = req.Send()
	return
}

var opSendRawEmail *aws.Operation

// SetIdentityDKIMEnabledRequest generates a request for the SetIdentityDKIMEnabled operation.
func (c *SES) SetIdentityDKIMEnabledRequest(input *SetIdentityDKIMEnabledInput) (req *aws.Request, output *SetIdentityDKIMEnabledOutput) {
	if opSetIdentityDKIMEnabled == nil {
		opSetIdentityDKIMEnabled = &aws.Operation{
			Name:       "SetIdentityDkimEnabled",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetIdentityDKIMEnabled, input, output)
	output = &SetIdentityDKIMEnabledOutput{}
	req.Data = output
	return
}

func (c *SES) SetIdentityDKIMEnabled(input *SetIdentityDKIMEnabledInput) (output *SetIdentityDKIMEnabledOutput, err error) {
	req, out := c.SetIdentityDKIMEnabledRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetIdentityDKIMEnabled *aws.Operation

// SetIdentityFeedbackForwardingEnabledRequest generates a request for the SetIdentityFeedbackForwardingEnabled operation.
func (c *SES) SetIdentityFeedbackForwardingEnabledRequest(input *SetIdentityFeedbackForwardingEnabledInput) (req *aws.Request, output *SetIdentityFeedbackForwardingEnabledOutput) {
	if opSetIdentityFeedbackForwardingEnabled == nil {
		opSetIdentityFeedbackForwardingEnabled = &aws.Operation{
			Name:       "SetIdentityFeedbackForwardingEnabled",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetIdentityFeedbackForwardingEnabled, input, output)
	output = &SetIdentityFeedbackForwardingEnabledOutput{}
	req.Data = output
	return
}

func (c *SES) SetIdentityFeedbackForwardingEnabled(input *SetIdentityFeedbackForwardingEnabledInput) (output *SetIdentityFeedbackForwardingEnabledOutput, err error) {
	req, out := c.SetIdentityFeedbackForwardingEnabledRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetIdentityFeedbackForwardingEnabled *aws.Operation

// SetIdentityNotificationTopicRequest generates a request for the SetIdentityNotificationTopic operation.
func (c *SES) SetIdentityNotificationTopicRequest(input *SetIdentityNotificationTopicInput) (req *aws.Request, output *SetIdentityNotificationTopicOutput) {
	if opSetIdentityNotificationTopic == nil {
		opSetIdentityNotificationTopic = &aws.Operation{
			Name:       "SetIdentityNotificationTopic",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetIdentityNotificationTopic, input, output)
	output = &SetIdentityNotificationTopicOutput{}
	req.Data = output
	return
}

func (c *SES) SetIdentityNotificationTopic(input *SetIdentityNotificationTopicInput) (output *SetIdentityNotificationTopicOutput, err error) {
	req, out := c.SetIdentityNotificationTopicRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetIdentityNotificationTopic *aws.Operation

// VerifyDomainDKIMRequest generates a request for the VerifyDomainDKIM operation.
func (c *SES) VerifyDomainDKIMRequest(input *VerifyDomainDKIMInput) (req *aws.Request, output *VerifyDomainDKIMOutput) {
	if opVerifyDomainDKIM == nil {
		opVerifyDomainDKIM = &aws.Operation{
			Name:       "VerifyDomainDkim",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opVerifyDomainDKIM, input, output)
	output = &VerifyDomainDKIMOutput{}
	req.Data = output
	return
}

func (c *SES) VerifyDomainDKIM(input *VerifyDomainDKIMInput) (output *VerifyDomainDKIMOutput, err error) {
	req, out := c.VerifyDomainDKIMRequest(input)
	output = out
	err = req.Send()
	return
}

var opVerifyDomainDKIM *aws.Operation

// VerifyDomainIdentityRequest generates a request for the VerifyDomainIdentity operation.
func (c *SES) VerifyDomainIdentityRequest(input *VerifyDomainIdentityInput) (req *aws.Request, output *VerifyDomainIdentityOutput) {
	if opVerifyDomainIdentity == nil {
		opVerifyDomainIdentity = &aws.Operation{
			Name:       "VerifyDomainIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opVerifyDomainIdentity, input, output)
	output = &VerifyDomainIdentityOutput{}
	req.Data = output
	return
}

func (c *SES) VerifyDomainIdentity(input *VerifyDomainIdentityInput) (output *VerifyDomainIdentityOutput, err error) {
	req, out := c.VerifyDomainIdentityRequest(input)
	output = out
	err = req.Send()
	return
}

var opVerifyDomainIdentity *aws.Operation

// VerifyEmailAddressRequest generates a request for the VerifyEmailAddress operation.
func (c *SES) VerifyEmailAddressRequest(input *VerifyEmailAddressInput) (req *aws.Request, output *VerifyEmailAddressOutput) {
	if opVerifyEmailAddress == nil {
		opVerifyEmailAddress = &aws.Operation{
			Name:       "VerifyEmailAddress",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opVerifyEmailAddress, input, output)
	output = &VerifyEmailAddressOutput{}
	req.Data = output
	return
}

func (c *SES) VerifyEmailAddress(input *VerifyEmailAddressInput) (output *VerifyEmailAddressOutput, err error) {
	req, out := c.VerifyEmailAddressRequest(input)
	output = out
	err = req.Send()
	return
}

var opVerifyEmailAddress *aws.Operation

// VerifyEmailIdentityRequest generates a request for the VerifyEmailIdentity operation.
func (c *SES) VerifyEmailIdentityRequest(input *VerifyEmailIdentityInput) (req *aws.Request, output *VerifyEmailIdentityOutput) {
	if opVerifyEmailIdentity == nil {
		opVerifyEmailIdentity = &aws.Operation{
			Name:       "VerifyEmailIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opVerifyEmailIdentity, input, output)
	output = &VerifyEmailIdentityOutput{}
	req.Data = output
	return
}

func (c *SES) VerifyEmailIdentity(input *VerifyEmailIdentityInput) (output *VerifyEmailIdentityOutput, err error) {
	req, out := c.VerifyEmailIdentityRequest(input)
	output = out
	err = req.Send()
	return
}

var opVerifyEmailIdentity *aws.Operation

type Body struct {
	HTML *Content `locationName:"Html" type:"structure"`
	Text *Content `type:"structure"`

	metadataBody `json:"-", xml:"-"`
}

type metadataBody struct {
	SDKShapeTraits bool `type:"structure"`
}

type Content struct {
	Charset *string `type:"string"`
	Data    *string `type:"string" required:"true"`

	metadataContent `json:"-", xml:"-"`
}

type metadataContent struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteIdentityInput struct {
	Identity *string `type:"string" required:"true"`

	metadataDeleteIdentityInput `json:"-", xml:"-"`
}

type metadataDeleteIdentityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteIdentityOutput struct {
	metadataDeleteIdentityOutput `json:"-", xml:"-"`
}

type metadataDeleteIdentityOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteVerifiedEmailAddressInput struct {
	EmailAddress *string `type:"string" required:"true"`

	metadataDeleteVerifiedEmailAddressInput `json:"-", xml:"-"`
}

type metadataDeleteVerifiedEmailAddressInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteVerifiedEmailAddressOutput struct {
	metadataDeleteVerifiedEmailAddressOutput `json:"-", xml:"-"`
}

type metadataDeleteVerifiedEmailAddressOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Destination struct {
	BCCAddresses []*string `locationName:"BccAddresses" type:"list"`
	CCAddresses  []*string `locationName:"CcAddresses" type:"list"`
	ToAddresses  []*string `type:"list"`

	metadataDestination `json:"-", xml:"-"`
}

type metadataDestination struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetIdentityDKIMAttributesInput struct {
	Identities []*string `type:"list" required:"true"`

	metadataGetIdentityDKIMAttributesInput `json:"-", xml:"-"`
}

type metadataGetIdentityDKIMAttributesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetIdentityDKIMAttributesOutput struct {
	DKIMAttributes *map[string]*IdentityDKIMAttributes `locationName:"DkimAttributes" type:"map" required:"true"`

	metadataGetIdentityDKIMAttributesOutput `json:"-", xml:"-"`
}

type metadataGetIdentityDKIMAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetIdentityNotificationAttributesInput struct {
	Identities []*string `type:"list" required:"true"`

	metadataGetIdentityNotificationAttributesInput `json:"-", xml:"-"`
}

type metadataGetIdentityNotificationAttributesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetIdentityNotificationAttributesOutput struct {
	NotificationAttributes *map[string]*IdentityNotificationAttributes `type:"map" required:"true"`

	metadataGetIdentityNotificationAttributesOutput `json:"-", xml:"-"`
}

type metadataGetIdentityNotificationAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetIdentityVerificationAttributesInput struct {
	Identities []*string `type:"list" required:"true"`

	metadataGetIdentityVerificationAttributesInput `json:"-", xml:"-"`
}

type metadataGetIdentityVerificationAttributesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetIdentityVerificationAttributesOutput struct {
	VerificationAttributes *map[string]*IdentityVerificationAttributes `type:"map" required:"true"`

	metadataGetIdentityVerificationAttributesOutput `json:"-", xml:"-"`
}

type metadataGetIdentityVerificationAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetSendQuotaInput struct {
	metadataGetSendQuotaInput `json:"-", xml:"-"`
}

type metadataGetSendQuotaInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetSendQuotaOutput struct {
	Max24HourSend   *float64 `type:"double"`
	MaxSendRate     *float64 `type:"double"`
	SentLast24Hours *float64 `type:"double"`

	metadataGetSendQuotaOutput `json:"-", xml:"-"`
}

type metadataGetSendQuotaOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetSendStatisticsInput struct {
	metadataGetSendStatisticsInput `json:"-", xml:"-"`
}

type metadataGetSendStatisticsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetSendStatisticsOutput struct {
	SendDataPoints []*SendDataPoint `type:"list"`

	metadataGetSendStatisticsOutput `json:"-", xml:"-"`
}

type metadataGetSendStatisticsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type IdentityDKIMAttributes struct {
	DKIMEnabled            *bool     `locationName:"DkimEnabled" type:"boolean" required:"true"`
	DKIMTokens             []*string `locationName:"DkimTokens" type:"list"`
	DKIMVerificationStatus *string   `locationName:"DkimVerificationStatus" type:"string" required:"true"`

	metadataIdentityDKIMAttributes `json:"-", xml:"-"`
}

type metadataIdentityDKIMAttributes struct {
	SDKShapeTraits bool `type:"structure"`
}

type IdentityNotificationAttributes struct {
	BounceTopic       *string `type:"string" required:"true"`
	ComplaintTopic    *string `type:"string" required:"true"`
	DeliveryTopic     *string `type:"string" required:"true"`
	ForwardingEnabled *bool   `type:"boolean" required:"true"`

	metadataIdentityNotificationAttributes `json:"-", xml:"-"`
}

type metadataIdentityNotificationAttributes struct {
	SDKShapeTraits bool `type:"structure"`
}

type IdentityVerificationAttributes struct {
	VerificationStatus *string `type:"string" required:"true"`
	VerificationToken  *string `type:"string"`

	metadataIdentityVerificationAttributes `json:"-", xml:"-"`
}

type metadataIdentityVerificationAttributes struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListIdentitiesInput struct {
	IdentityType *string `type:"string"`
	MaxItems     *int    `type:"integer"`
	NextToken    *string `type:"string"`

	metadataListIdentitiesInput `json:"-", xml:"-"`
}

type metadataListIdentitiesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListIdentitiesOutput struct {
	Identities []*string `type:"list" required:"true"`
	NextToken  *string   `type:"string"`

	metadataListIdentitiesOutput `json:"-", xml:"-"`
}

type metadataListIdentitiesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListVerifiedEmailAddressesInput struct {
	metadataListVerifiedEmailAddressesInput `json:"-", xml:"-"`
}

type metadataListVerifiedEmailAddressesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListVerifiedEmailAddressesOutput struct {
	VerifiedEmailAddresses []*string `type:"list"`

	metadataListVerifiedEmailAddressesOutput `json:"-", xml:"-"`
}

type metadataListVerifiedEmailAddressesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Message struct {
	Body    *Body    `type:"structure" required:"true"`
	Subject *Content `type:"structure" required:"true"`

	metadataMessage `json:"-", xml:"-"`
}

type metadataMessage struct {
	SDKShapeTraits bool `type:"structure"`
}

type RawMessage struct {
	Data []byte `type:"blob" required:"true"`

	metadataRawMessage `json:"-", xml:"-"`
}

type metadataRawMessage struct {
	SDKShapeTraits bool `type:"structure"`
}

type SendDataPoint struct {
	Bounces          *int64     `type:"long"`
	Complaints       *int64     `type:"long"`
	DeliveryAttempts *int64     `type:"long"`
	Rejects          *int64     `type:"long"`
	Timestamp        *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	metadataSendDataPoint `json:"-", xml:"-"`
}

type metadataSendDataPoint struct {
	SDKShapeTraits bool `type:"structure"`
}

type SendEmailInput struct {
	Destination      *Destination `type:"structure" required:"true"`
	Message          *Message     `type:"structure" required:"true"`
	ReplyToAddresses []*string    `type:"list"`
	ReturnPath       *string      `type:"string"`
	Source           *string      `type:"string" required:"true"`

	metadataSendEmailInput `json:"-", xml:"-"`
}

type metadataSendEmailInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SendEmailOutput struct {
	MessageID *string `locationName:"MessageId" type:"string" required:"true"`

	metadataSendEmailOutput `json:"-", xml:"-"`
}

type metadataSendEmailOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SendRawEmailInput struct {
	Destinations []*string   `type:"list"`
	RawMessage   *RawMessage `type:"structure" required:"true"`
	Source       *string     `type:"string"`

	metadataSendRawEmailInput `json:"-", xml:"-"`
}

type metadataSendRawEmailInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SendRawEmailOutput struct {
	MessageID *string `locationName:"MessageId" type:"string" required:"true"`

	metadataSendRawEmailOutput `json:"-", xml:"-"`
}

type metadataSendRawEmailOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetIdentityDKIMEnabledInput struct {
	DKIMEnabled *bool   `locationName:"DkimEnabled" type:"boolean" required:"true"`
	Identity    *string `type:"string" required:"true"`

	metadataSetIdentityDKIMEnabledInput `json:"-", xml:"-"`
}

type metadataSetIdentityDKIMEnabledInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetIdentityDKIMEnabledOutput struct {
	metadataSetIdentityDKIMEnabledOutput `json:"-", xml:"-"`
}

type metadataSetIdentityDKIMEnabledOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetIdentityFeedbackForwardingEnabledInput struct {
	ForwardingEnabled *bool   `type:"boolean" required:"true"`
	Identity          *string `type:"string" required:"true"`

	metadataSetIdentityFeedbackForwardingEnabledInput `json:"-", xml:"-"`
}

type metadataSetIdentityFeedbackForwardingEnabledInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetIdentityFeedbackForwardingEnabledOutput struct {
	metadataSetIdentityFeedbackForwardingEnabledOutput `json:"-", xml:"-"`
}

type metadataSetIdentityFeedbackForwardingEnabledOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetIdentityNotificationTopicInput struct {
	Identity         *string `type:"string" required:"true"`
	NotificationType *string `type:"string" required:"true"`
	SNSTopic         *string `locationName:"SnsTopic" type:"string"`

	metadataSetIdentityNotificationTopicInput `json:"-", xml:"-"`
}

type metadataSetIdentityNotificationTopicInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetIdentityNotificationTopicOutput struct {
	metadataSetIdentityNotificationTopicOutput `json:"-", xml:"-"`
}

type metadataSetIdentityNotificationTopicOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type VerifyDomainDKIMInput struct {
	Domain *string `type:"string" required:"true"`

	metadataVerifyDomainDKIMInput `json:"-", xml:"-"`
}

type metadataVerifyDomainDKIMInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type VerifyDomainDKIMOutput struct {
	DKIMTokens []*string `locationName:"DkimTokens" type:"list" required:"true"`

	metadataVerifyDomainDKIMOutput `json:"-", xml:"-"`
}

type metadataVerifyDomainDKIMOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type VerifyDomainIdentityInput struct {
	Domain *string `type:"string" required:"true"`

	metadataVerifyDomainIdentityInput `json:"-", xml:"-"`
}

type metadataVerifyDomainIdentityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type VerifyDomainIdentityOutput struct {
	VerificationToken *string `type:"string" required:"true"`

	metadataVerifyDomainIdentityOutput `json:"-", xml:"-"`
}

type metadataVerifyDomainIdentityOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type VerifyEmailAddressInput struct {
	EmailAddress *string `type:"string" required:"true"`

	metadataVerifyEmailAddressInput `json:"-", xml:"-"`
}

type metadataVerifyEmailAddressInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type VerifyEmailAddressOutput struct {
	metadataVerifyEmailAddressOutput `json:"-", xml:"-"`
}

type metadataVerifyEmailAddressOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type VerifyEmailIdentityInput struct {
	EmailAddress *string `type:"string" required:"true"`

	metadataVerifyEmailIdentityInput `json:"-", xml:"-"`
}

type metadataVerifyEmailIdentityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type VerifyEmailIdentityOutput struct {
	metadataVerifyEmailIdentityOutput `json:"-", xml:"-"`
}

type metadataVerifyEmailIdentityOutput struct {
	SDKShapeTraits bool `type:"structure"`
}