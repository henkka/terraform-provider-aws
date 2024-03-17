// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrpublic

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/sweep"
)

func RegisterSweepers() {
	resource.AddTestSweepers("aws_ecrpublic_repository", &resource.Sweeper{
		Name: "aws_ecrpublic_repository",
		F:    sweepRepositories,
	})
}

func sweepRepositories(region string) error {
	ctx := sweep.Context(region)
	client, err := sweep.SharedRegionalSweepClient(ctx, region)
	if err != nil {
		return fmt.Errorf("error getting client: %s", err)
	}
	conn := client.ECRPublicClient(ctx)
	input := &ecrpublic.DescribeRepositoriesInput{}
	sweepResources := make([]sweep.Sweepable, 0)

	paginator := ecrpublic.NewDescribeRepositoriesPaginator(conn, input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("error describing repositories: %w", err)
		}
		for _, repository := range page.Repositories {
			r := ResourceRepository()
			d := r.Data(nil)
			d.SetId(aws.ToString(repository.RepositoryName))
			d.Set("registry_id", repository.RegistryId)
			d.Set("force_destroy", true)

			sweepResources = append(sweepResources, sweep.NewSweepResource(r, d, client))
		}
	}

	err = sweep.SweepOrchestrator(ctx, sweepResources)

	if err != nil {
		return fmt.Errorf("error sweeping ECR Public Repositories (%s): %w", region, err)
	}

	return nil
}
