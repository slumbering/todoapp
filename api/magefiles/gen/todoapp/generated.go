// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package todoapp

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/dagger/cloak/sdk/go/dagger"
)

// DeployResponse is returned by Deploy on success.
type DeployResponse struct {
	Todoapp DeployTodoappTodoApp `json:"todoapp"`
}

// GetTodoapp returns DeployResponse.Todoapp, and is useful for accessing the field via an interface.
func (v *DeployResponse) GetTodoapp() DeployTodoappTodoApp { return v.Todoapp }

// DeployTodoappTodoApp includes the requested fields of the GraphQL type TodoApp.
type DeployTodoappTodoApp struct {
	Deploy string `json:"deploy"`
}

// GetDeploy returns DeployTodoappTodoApp.Deploy, and is useful for accessing the field via an interface.
func (v *DeployTodoappTodoApp) GetDeploy() string { return v.Deploy }

// __DeployInput is used internally by genqlient
type __DeployInput struct {
	Source   dagger.FSID     `json:"source"`
	SiteName string          `json:"siteName"`
	Token    dagger.SecretID `json:"token"`
}

// GetSource returns __DeployInput.Source, and is useful for accessing the field via an interface.
func (v *__DeployInput) GetSource() dagger.FSID { return v.Source }

// GetSiteName returns __DeployInput.SiteName, and is useful for accessing the field via an interface.
func (v *__DeployInput) GetSiteName() string { return v.SiteName }

// GetToken returns __DeployInput.Token, and is useful for accessing the field via an interface.
func (v *__DeployInput) GetToken() dagger.SecretID { return v.Token }

func Deploy(
	ctx context.Context,
	source dagger.FSID,
	siteName string,
	token dagger.SecretID,
) (*DeployResponse, error) {
	req := &graphql.Request{
		OpName: "Deploy",
		Query: `
query Deploy ($source: FSID!, $siteName: String!, $token: SecretID!) {
	todoapp {
		deploy(source: $source, siteName: $siteName, token: $token)
	}
}
`,
		Variables: &__DeployInput{
			Source:   source,
			SiteName: siteName,
			Token:    token,
		},
	}
	var err error
	var client graphql.Client

	client, err = dagger.Client(ctx)
	if err != nil {
		return nil, err
	}

	var data DeployResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
