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

	member, err := c.PostOrgsIDMembers(ctx, orgID).
		AddResourceMemberRequestBody(*api.NewAddResourceMemberRequestBody(params.MemberId)).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to add user %q to org %q: %w", params.MemberId, orgID, err)
	}

	_, err = c.StdIO.Write([]byte(fmt.Sprintf("user %q has been added as a member of org %q\n", *member.Id, orgID)))
	return err
}

type ListMemberParams struct {
	clients.OrgParams
}

const maxConcurrentGetUserRequests = 10

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

	type indexedUser struct {
		user  api.UserResponse
		index int
	}
	userChan := make(chan indexedUser, maxConcurrentGetUserRequests)
	semChan := make(chan struct{}, maxConcurrentGetUserRequests)
	errChan := make(chan error)

	resourceMembers := members.GetUsers()

	// Fetch user details about all members of the org.
	for i, member := range resourceMembers {
		go func(i int, memberId string) {
			semChan <- struct{}{}
			defer func() { <-semChan }()

			user, err := c.GetUsersID(ctx, memberId).Execute()
			if err != nil {
				errChan <- fmt.Errorf("failed to retrieve details for user %q: %w", memberId, err)
				return
			}
			userChan <- indexedUser{user: user, index: i}
		}(i, member.GetId())
	}

	users := make([]api.UserResponse, len(resourceMembers))
	for range resourceMembers {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errChan:
			return err
		case user := <-userChan:
			users[user.index] = user.user
		}
	}

	if c.PrintAsJSON {
		return c.PrintJSON(users)
	}

	rows := make([]map[string]interface{}, len(resourceMembers))
	for i, user := range users {
		rows[i] = map[string]interface{}{
			"ID":        user.GetId(),
			"Name":      user.GetName(),
			"User Type": "member",
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
		return fmt.Errorf("failed to remove member %q from org %q", params.MemberId, orgID)
	}

	_, err = c.StdIO.Write([]byte(fmt.Sprintf("user %q has been removed from org %q\n", params.MemberId, orgID)))
	return err
}
