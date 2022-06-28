package org

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type AddMemberParams struct {
	clients.OrgParams
	MemberId string
	IsOwner  bool
}

func (c Client) AddMember(ctx context.Context, params *AddMemberParams) (err error) {
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}
	if params.IsOwner {
		owner, err := c.PostOrgsIDOwners(ctx, orgID).
			AddResourceMemberRequestBody(*api.NewAddResourceMemberRequestBody(params.MemberId)).
			Execute()
		if err != nil {
			return fmt.Errorf("failed to add user %q as owner of org %q: %w", params.MemberId, orgID, err)
		}
		_, err = c.StdIO.Write([]byte(fmt.Sprintf("user %q has been added as an owner of org %q\n", *owner.Id, orgID)))
		return err
	} else {
		member, err := c.PostOrgsIDMembers(ctx, orgID).
			AddResourceMemberRequestBody(*api.NewAddResourceMemberRequestBody(params.MemberId)).
			Execute()
		if err != nil {
			return fmt.Errorf("failed to add user %q as member of org %q: %w", params.MemberId, orgID, err)
		}
		_, err = c.StdIO.Write([]byte(fmt.Sprintf("user %q has been added as a member of org %q\n", *member.Id, orgID)))
		return err
	}
}

type ListMemberParams struct {
	clients.OrgParams
}

func (c Client) ListMembers(ctx context.Context, params *ListMemberParams) (err error) {
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	members, err := c.GetOrgsIDMembers(ctx, orgID).Execute()
	if err != nil {
		return fmt.Errorf("failed to find members of org %q: %w", orgID, err)
	}
	owners, err := c.GetOrgsIDOwners(ctx, orgID).Execute()
	if err != nil {
		return fmt.Errorf("failed to find owners of org %q: %w", orgID, err)
	}
	allMembers := members.GetUsers()
	resourceOwners := owners.GetUsers()
	resourceOwnersAsMembers := make([]api.ResourceMember, len(resourceOwners))
	for i, owner := range resourceOwners {
		resourceOwnersAsMembers[i] = api.ResourceMember(owner)
	}
	allMembers = append(resourceOwnersAsMembers, allMembers...)

	if c.PrintAsJSON {
		return c.PrintJSON(allMembers)
	}

	rows := make([]map[string]interface{}, len(allMembers))
	for i, user := range allMembers {
		rows[i] = map[string]interface{}{
			"ID":        user.GetId(),
			"Name":      user.GetName(),
			"User Type": user.GetRole(),
			"Status":    user.GetStatus(),
		}
	}
	return c.PrintTable([]string{"ID", "Name", "User Type", "Status"}, rows...)
}

type RemoveMemberParams struct {
	clients.OrgParams
	MemberId string
}

func (c Client) RemoveMember(ctx context.Context, params *RemoveMemberParams) (err error) {
	orgID, err := params.GetOrgID(ctx, c.ActiveConfig, c.OrganizationsApi)
	if err != nil {
		return err
	}

	if err = c.DeleteOrgsIDMembersID(ctx, params.MemberId, orgID).Execute(); err != nil {
		if err = c.DeleteOrgsIDOwnersID(ctx, params.MemberId, orgID).Execute(); err != nil {
			return fmt.Errorf("failed to remove user %q from org %q", params.MemberId, orgID)
		}
		_, err = c.StdIO.Write([]byte(fmt.Sprintf("owner %q has been removed from org %q\n", params.MemberId, orgID)))
		return err
	}
	_, err = c.StdIO.Write([]byte(fmt.Sprintf("user %q has been removed from org %q\n", params.MemberId, orgID)))
	return err
}
