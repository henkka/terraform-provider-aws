// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	redshift_sdkv2 "github.com/aws/aws-sdk-go-v2/service/redshift"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	redshift_sdkv1 "github.com/aws/aws-sdk-go/service/redshift"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceDataShares,
			Name:    "Data Shares",
		},
		{
			Factory: newDataSourceProducerDataShares,
			Name:    "Producer Data Shares",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceDataShareAuthorization,
			Name:    "Data Share Authorization",
		},
		{
			Factory: newResourceDataShareConsumerAssociation,
			Name:    "Data Share Consumer Association",
		},
		{
			Factory: newResourceLogging,
			Name:    "Logging",
		},
		{
			Factory: newResourceSnapshotCopy,
			Name:    "Snapshot Copy",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_redshift_cluster",
			Name:     "Cluster",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceClusterCredentials,
			TypeName: "aws_redshift_cluster_credentials",
			Name:     "Cluster Credentials",
		},
		{
			Factory:  DataSourceOrderableCluster,
			TypeName: "aws_redshift_orderable_cluster",
		},
		{
			Factory:  DataSourceServiceAccount,
			TypeName: "aws_redshift_service_account",
		},
		{
			Factory:  DataSourceSubnetGroup,
			TypeName: "aws_redshift_subnet_group",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAuthenticationProfile,
			TypeName: "aws_redshift_authentication_profile",
			Name:     "Authentication Profile",
		},
		{
			Factory:  resourceCluster,
			TypeName: "aws_redshift_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceClusterIAMRoles,
			TypeName: "aws_redshift_cluster_iam_roles",
			Name:     "Cluster IAM Roles",
		},
		{
			Factory:  resourceClusterSnapshot,
			TypeName: "aws_redshift_cluster_snapshot",
			Name:     "Cluster Snapshot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  resourceEndpointAccess,
			TypeName: "aws_redshift_endpoint_access",
			Name:     "Endpoint Access",
		},
		{
			Factory:  ResourceEndpointAuthorization,
			TypeName: "aws_redshift_endpoint_authorization",
		},
		{
			Factory:  ResourceEventSubscription,
			TypeName: "aws_redshift_event_subscription",
			Name:     "Event Subscription",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceHSMClientCertificate,
			TypeName: "aws_redshift_hsm_client_certificate",
			Name:     "HSM Client Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceHSMConfiguration,
			TypeName: "aws_redshift_hsm_configuration",
			Name:     "HSM Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceParameterGroup,
			TypeName: "aws_redshift_parameter_group",
			Name:     "Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourcePartner,
			TypeName: "aws_redshift_partner",
		},
		{
			Factory:  ResourceResourcePolicy,
			TypeName: "aws_redshift_resource_policy",
		},
		{
			Factory:  ResourceScheduledAction,
			TypeName: "aws_redshift_scheduled_action",
		},
		{
			Factory:  ResourceSnapshotCopyGrant,
			TypeName: "aws_redshift_snapshot_copy_grant",
			Name:     "Snapshot Copy Grant",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceSnapshotSchedule,
			TypeName: "aws_redshift_snapshot_schedule",
			Name:     "Snapshot Schedule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceSnapshotScheduleAssociation,
			TypeName: "aws_redshift_snapshot_schedule_association",
		},
		{
			Factory:  ResourceSubnetGroup,
			TypeName: "aws_redshift_subnet_group",
			Name:     "Subnet Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceUsageLimit,
			TypeName: "aws_redshift_usage_limit",
			Name:     "Usage Limit",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Redshift
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*redshift_sdkv1.Redshift, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return redshift_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*redshift_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return redshift_sdkv2.NewFromConfig(cfg, func(o *redshift_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
