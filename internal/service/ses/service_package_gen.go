// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package ses

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	ses_sdkv2 "github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceActiveReceiptRuleSet,
			TypeName: "aws_ses_active_receipt_rule_set",
		},
		{
			Factory:  dataSourceDomainIdentity,
			TypeName: "aws_ses_domain_identity",
			Name:     "Domain Identity",
		},
		{
			Factory:  dataSourceEmailIdentity,
			TypeName: "aws_ses_email_identity",
			Name:     "Email Identity",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceActiveReceiptRuleSet,
			TypeName: "aws_ses_active_receipt_rule_set",
		},
		{
			Factory:  ResourceConfigurationSet,
			TypeName: "aws_ses_configuration_set",
		},
		{
			Factory:  resourceDomainDKIM,
			TypeName: "aws_ses_domain_dkim",
			Name:     "Domain DKIM",
		},
		{
			Factory:  resourceDomainIdentity,
			TypeName: "aws_ses_domain_identity",
			Name:     "Domain Identity",
		},
		{
			Factory:  resourceDomainIdentityVerification,
			TypeName: "aws_ses_domain_identity_verification",
			Name:     "Domain Identity Verification",
		},
		{
			Factory:  resourceDomainMailFrom,
			TypeName: "aws_ses_domain_mail_from",
			Name:     "MAIL FROM Domain",
		},
		{
			Factory:  resourceEmailIdentity,
			TypeName: "aws_ses_email_identity",
			Name:     "Email Identity",
		},
		{
			Factory:  resourceEventDestination,
			TypeName: "aws_ses_event_destination",
			Name:     "Configuration Set Event Destination",
		},
		{
			Factory:  resourceIdentityNotificationTopic,
			TypeName: "aws_ses_identity_notification_topic",
			Name:     "Identity Notification Topic",
		},
		{
			Factory:  resourceIdentityPolicy,
			TypeName: "aws_ses_identity_policy",
			Name:     "Identity Policy",
		},
		{
			Factory:  resourceReceiptFilter,
			TypeName: "aws_ses_receipt_filter",
			Name:     "Receipt Filter",
		},
		{
			Factory:  resourceReceiptRule,
			TypeName: "aws_ses_receipt_rule",
			Name:     "Receipt Rule",
		},
		{
			Factory:  resourceReceiptRuleSet,
			TypeName: "aws_ses_receipt_rule_set",
			Name:     "Receipt Rule Set",
		},
		{
			Factory:  resourceTemplate,
			TypeName: "aws_ses_template",
			Name:     "Template",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SES
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*ses_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return ses_sdkv2.NewFromConfig(cfg,
		ses_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
