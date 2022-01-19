package vault_mock

import (
	"errors"
	"fmt"
	"regexp"
)

type Client struct {
	logical *Logical
}

type Config struct {
	Address string
}

type Logical struct {
}

type Secret struct {
	Auth *SecretAuth
}

type SecretAuth struct {
	ClientToken string
	Policies    []string
}

var users = map[string]bool{}

func NewClient(config *Config) (*Client, error) {
	if config.Address == "" {
		return nil, errors.New("address is empty")
	}

	return &Client{logical: &Logical{}}, nil
}

func (c *Client) Logical() *Logical {
	return c.logical
}

func (l *Logical) Write(path string, data map[string]interface{}) (*Secret, error) {
	if path == "" {
		return nil, errors.New("path is empty")
	}

	organisationSlug := fmt.Sprintf("%v", data["BUILDKITE_ORGANIZATION_SLUG"])
	agentId := fmt.Sprintf("%v", data["BUILDKITE_AGENT_ID"])
	pipelineId := fmt.Sprintf("%v", data["BUILDKITE_PIPELINE_ID"])

	re := regexp.MustCompile("^[a-zA-Z0-9-]+$")
	//fmt.Println("The organisation token is %s", fmt.Sprintf("%v", data["BUILDKITE_ORGANIZATION_SLUG"]))
	if !re.MatchString(organisationSlug) {
		return nil, errors.New("organisationSlug is invalid")
	}
	re = regexp.MustCompile("^\"[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}\"$")

	if !re.MatchString(agentId) {
		return nil, errors.New("agentId is invalid")
	}

	if !re.MatchString(pipelineId) {
		return nil, errors.New("pipelineId is invalid")
	}

	//numUsers := len(users)
	token := fmt.Sprintf("abcdefg%vhijklmnopqr", len(users))
	users[token] = true

	secret := &Secret{
		Auth: &SecretAuth{
			ClientToken: token,
			Policies:    []string{"default"},
		},
	}

	return secret, nil
}
