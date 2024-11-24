// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorydb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/memorydb"
	awstypes "github.com/aws/aws-sdk-go-v2/service/memorydb/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func FindSubnetGroupByName(ctx context.Context, conn *memorydb.Client, name string) (*awstypes.SubnetGroup, error) {
	input := memorydb.DescribeSubnetGroupsInput{
		SubnetGroupName: aws.String(name),
	}

	output, err := conn.DescribeSubnetGroups(ctx, &input)

	if errs.IsA[*awstypes.SubnetGroupNotFoundFault](err) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	return tfresource.AssertSingleValueResult(output.SubnetGroups)
}

func FindUserByName(ctx context.Context, conn *memorydb.Client, name string) (*awstypes.User, error) {
	input := memorydb.DescribeUsersInput{
		UserName: aws.String(name),
	}

	output, err := conn.DescribeUsers(ctx, &input)

	if errs.IsA[*awstypes.UserNotFoundFault](err) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	return tfresource.AssertSingleValueResult(output.Users)
}
