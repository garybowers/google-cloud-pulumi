// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectState struct {
	pulumi.ResourceState
}

type ProjectArgs struct {
	OrganizationId    pulumi.StringInput `pulumi:"organizationid"`
	FolderId          pulumi.StringInput `pulumi:"folderid"`
	Name              pulumi.StringInput `pulumi:"name"`
	ProjectId         pulumi.StringInput `pulumi:"projectid"`
	BillingAccount    pulumi.StringInput `pulumi:"billingaccount"`
	AutoCreateNetwork pulumi.Bool        `pulumi:"autocreatenetwork"`
	OSLogin           struct {
		Enabled pulumi.Bool        `pulumi:"enabled"`
		Admins  pulumi.StringInput `pulumi:"admins"`
		Users   pulumi.StringInput `pulumi:"users"`
	}
}

func NewProject(ctx *pulumi.Context, name string, args ProjectArgs, opts pulumi.ResourceOption) (*ProjectState, error) {
	project := &ProjectState{}
	err := ctx.RegisterComponentResource("pkg:google:project", name, project, opts)

	project, err := organizations.NewProject(ctx, name, &organizations.ProjectArgs{
		ProjectId:         args.ProjectId,
		BillingAccount:    args.BillingAccount,
		AutoCreateNetwork: args.AutoCreateNetwork,
	})
	if err != nil {
		return err
	}
}
