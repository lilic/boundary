// Code generated by "make api"; DO NOT EDIT.
package targets

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/scopes"
)

type Target struct {
	Id                     string                 `json:"id,omitempty"`
	ScopeId                string                 `json:"scope_id,omitempty"`
	Scope                  *scopes.ScopeInfo      `json:"scope,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Description            string                 `json:"description,omitempty"`
	CreatedTime            time.Time              `json:"created_time,omitempty"`
	UpdatedTime            time.Time              `json:"updated_time,omitempty"`
	Version                uint32                 `json:"version,omitempty"`
	Type                   string                 `json:"type,omitempty"`
	HostSetIds             []string               `json:"host_set_ids,omitempty"`
	HostSets               []*HostSet             `json:"host_sets,omitempty"`
	SessionMaxSeconds      uint32                 `json:"session_max_seconds,omitempty"`
	SessionConnectionLimit int32                  `json:"session_connection_limit,omitempty"`
	WorkerFilter           string                 `json:"worker_filter,omitempty"`
	CredentialLibraryIds   []string               `json:"credential_library_ids,omitempty"`
	CredentialLibraries    []*CredentialLibrary   `json:"credential_libraries,omitempty"`
	Attributes             map[string]interface{} `json:"attributes,omitempty"`
	AuthorizedActions      []string               `json:"authorized_actions,omitempty"`

	response *api.Response
}

type TargetReadResult struct {
	Item     *Target
	response *api.Response
}

func (n TargetReadResult) GetItem() interface{} {
	return n.Item
}

func (n TargetReadResult) GetResponse() *api.Response {
	return n.response
}

type (
	TargetCreateResult = TargetReadResult
	TargetUpdateResult = TargetReadResult
)

type TargetDeleteResult struct {
	response *api.Response
}

// GetItem will always be nil for TargetDeleteResult
func (n TargetDeleteResult) GetItem() interface{} {
	return nil
}

func (n TargetDeleteResult) GetResponse() *api.Response {
	return n.response
}

type TargetListResult struct {
	Items    []*Target
	response *api.Response
}

func (n TargetListResult) GetItems() interface{} {
	return n.Items
}

func (n TargetListResult) GetResponse() *api.Response {
	return n.response
}

// Client is a client for this collection
type Client struct {
	client *api.Client
}

// Creates a new client for this collection. The submitted API client is cloned;
// modifications to it after generating this client will not have effect. If you
// need to make changes to the underlying API client, use ApiClient() to access
// it.
func NewClient(c *api.Client) *Client {
	return &Client{client: c.Clone()}
}

// ApiClient returns the underlying API client
func (c *Client) ApiClient() *api.Client {
	return c.client
}

func (c *Client) Create(ctx context.Context, resourceType string, scopeId string, opt ...Option) (*TargetCreateResult, error) {
	if scopeId == "" {
		return nil, fmt.Errorf("empty scopeId value passed into Create request")
	}

	opts, apiOpts := getOpts(opt...)

	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}
	if resourceType == "" {
		return nil, fmt.Errorf("empty resourceType value passed into Create request")
	} else {
		opts.postMap["type"] = resourceType
	}

	opts.postMap["scope_id"] = scopeId

	req, err := c.client.NewRequest(ctx, "POST", "targets", opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Create request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Create call: %w", err)
	}

	target := new(TargetCreateResult)
	target.Item = new(Target)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Create response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Read(ctx context.Context, targetId string, opt ...Option) (*TargetReadResult, error) {
	if targetId == "" {
		return nil, fmt.Errorf("empty targetId value passed into Read request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("targets/%s", targetId), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Read request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(TargetReadResult)
	target.Item = new(Target)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Update(ctx context.Context, targetId string, version uint32, opt ...Option) (*TargetUpdateResult, error) {
	if targetId == "" {
		return nil, fmt.Errorf("empty targetId value passed into Update request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into Update request and automatic versioning not specified")
		}
		existingTarget, existingErr := c.Read(ctx, targetId, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	req, err := c.client.NewRequest(ctx, "PATCH", fmt.Sprintf("targets/%s", targetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Update request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Update call: %w", err)
	}

	target := new(TargetUpdateResult)
	target.Item = new(Target)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Update response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Delete(ctx context.Context, targetId string, opt ...Option) (*TargetDeleteResult, error) {
	if targetId == "" {
		return nil, fmt.Errorf("empty targetId value passed into Delete request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("targets/%s", targetId), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Delete request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Delete call: %w", err)
	}

	apiErr, err := resp.Decode(nil)
	if err != nil {
		return nil, fmt.Errorf("error decoding Delete response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}

	target := &TargetDeleteResult{
		response: resp,
	}
	return target, nil
}

func (c *Client) List(ctx context.Context, scopeId string, opt ...Option) (*TargetListResult, error) {
	if scopeId == "" {
		return nil, fmt.Errorf("empty scopeId value passed into List request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	opts.queryMap["scope_id"] = scopeId

	req, err := c.client.NewRequest(ctx, "GET", "targets", nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating List request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	target := new(TargetListResult)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) AddCredentialLibraries(ctx context.Context, targetId string, version uint32, credentialLibraryIds []string, opt ...Option) (*TargetUpdateResult, error) {
	if targetId == "" {
		return nil, fmt.Errorf("empty targetId value passed into AddCredentialLibraries request")
	}
	if len(credentialLibraryIds) == 0 {
		return nil, errors.New("empty credentialLibraryIds passed into AddCredentialLibraries request")
	}
	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into AddCredentialLibraries request")
		}
		existingTarget, existingErr := c.Read(ctx, targetId, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version
	opts.postMap["credential_library_ids"] = credentialLibraryIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("targets/%s:add-credential-libraries", targetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating AddCredentialLibraries request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during AddCredentialLibraries call: %w", err)
	}

	target := new(TargetUpdateResult)
	target.Item = new(Target)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding AddCredentialLibraries response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) AddHostSets(ctx context.Context, targetId string, version uint32, hostSetIds []string, opt ...Option) (*TargetUpdateResult, error) {
	if targetId == "" {
		return nil, fmt.Errorf("empty targetId value passed into AddHostSets request")
	}
	if len(hostSetIds) == 0 {
		return nil, errors.New("empty hostSetIds passed into AddHostSets request")
	}
	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into AddHostSets request")
		}
		existingTarget, existingErr := c.Read(ctx, targetId, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version
	opts.postMap["host_set_ids"] = hostSetIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("targets/%s:add-host-sets", targetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating AddHostSets request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during AddHostSets call: %w", err)
	}

	target := new(TargetUpdateResult)
	target.Item = new(Target)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding AddHostSets response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) SetCredentialLibraries(ctx context.Context, targetId string, version uint32, credentialLibraryIds []string, opt ...Option) (*TargetUpdateResult, error) {
	if targetId == "" {
		return nil, fmt.Errorf("empty targetId value passed into SetCredentialLibraries request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into SetCredentialLibraries request")
		}
		existingTarget, existingErr := c.Read(ctx, targetId, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version
	opts.postMap["credential_library_ids"] = credentialLibraryIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("targets/%s:set-credential-libraries", targetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating SetCredentialLibraries request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during SetCredentialLibraries call: %w", err)
	}

	target := new(TargetUpdateResult)
	target.Item = new(Target)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding SetCredentialLibraries response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) SetHostSets(ctx context.Context, targetId string, version uint32, hostSetIds []string, opt ...Option) (*TargetUpdateResult, error) {
	if targetId == "" {
		return nil, fmt.Errorf("empty targetId value passed into SetHostSets request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into SetHostSets request")
		}
		existingTarget, existingErr := c.Read(ctx, targetId, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version
	opts.postMap["host_set_ids"] = hostSetIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("targets/%s:set-host-sets", targetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating SetHostSets request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during SetHostSets call: %w", err)
	}

	target := new(TargetUpdateResult)
	target.Item = new(Target)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding SetHostSets response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) RemoveCredentialLibraries(ctx context.Context, targetId string, version uint32, credentialLibraryIds []string, opt ...Option) (*TargetUpdateResult, error) {
	if targetId == "" {
		return nil, fmt.Errorf("empty targetId value passed into RemoveCredentialLibraries request")
	}
	if len(credentialLibraryIds) == 0 {
		return nil, errors.New("empty credentialLibraryIds passed into RemoveCredentialLibraries request")
	}
	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into RemoveCredentialLibraries request")
		}
		existingTarget, existingErr := c.Read(ctx, targetId, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version
	opts.postMap["credential_library_ids"] = credentialLibraryIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("targets/%s:remove-credential-libraries", targetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating RemoveCredentialLibraries request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during RemoveCredentialLibraries call: %w", err)
	}

	target := new(TargetUpdateResult)
	target.Item = new(Target)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding RemoveCredentialLibraries response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) RemoveHostSets(ctx context.Context, targetId string, version uint32, hostSetIds []string, opt ...Option) (*TargetUpdateResult, error) {
	if targetId == "" {
		return nil, fmt.Errorf("empty targetId value passed into RemoveHostSets request")
	}
	if len(hostSetIds) == 0 {
		return nil, errors.New("empty hostSetIds passed into RemoveHostSets request")
	}
	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into RemoveHostSets request")
		}
		existingTarget, existingErr := c.Read(ctx, targetId, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version
	opts.postMap["host_set_ids"] = hostSetIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("targets/%s:remove-host-sets", targetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating RemoveHostSets request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during RemoveHostSets call: %w", err)
	}

	target := new(TargetUpdateResult)
	target.Item = new(Target)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding RemoveHostSets response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}
