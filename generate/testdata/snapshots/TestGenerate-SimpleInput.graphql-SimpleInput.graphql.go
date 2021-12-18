// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// SimpleInputQueryResponse is returned by SimpleInputQuery on success.
type SimpleInputQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User SimpleInputQueryUser `json:"user"`
}

// GetUser returns SimpleInputQueryResponse.User, and is useful for accessing the field via an interface.
func (v *SimpleInputQueryResponse) GetUser() SimpleInputQueryUser { return v.User }

// SimpleInputQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type SimpleInputQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// GetId returns SimpleInputQueryUser.Id, and is useful for accessing the field via an interface.
func (v *SimpleInputQueryUser) GetId() testutil.ID { return v.Id }

// __SimpleInputQueryInput is used internally by genqlient
type __SimpleInputQueryInput struct {
	Name string `json:"name"`
}

// GetName returns __SimpleInputQueryInput.Name, and is useful for accessing the field via an interface.
func (v *__SimpleInputQueryInput) GetName() string { return v.Name }

func SimpleInputQuery(
	client graphql.Client,
	name string,
) (*SimpleInputQueryResponse, error) {
	__input := __SimpleInputQueryInput{
		Name: name,
	}
	var err error

	var retval SimpleInputQueryResponse
	err = client.MakeRequest(
		nil,
		"SimpleInputQuery",
		`
query SimpleInputQuery ($name: String!) {
	user(query: {name:$name}) {
		id
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

