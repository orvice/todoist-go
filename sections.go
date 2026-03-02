package todoist

import (
	"context"
	"fmt"
)

// GetSectionsOptions contains query parameters for listing sections.
type GetSectionsOptions struct {
	ProjectID string `url:"project_id,omitempty"`
}

// GetSections returns sections, optionally filtered by project.
func (c *Client) GetSections(ctx context.Context, opts *GetSectionsOptions) ([]Section, error) {
	return getList[Section](c, ctx, "/sections", opts)
}

// GetSection returns a single section by ID.
func (c *Client) GetSection(ctx context.Context, id string) (*Section, error) {
	var section Section
	err := c.get(ctx, fmt.Sprintf("/sections/%s", id), nil, &section)
	if err != nil {
		return nil, err
	}
	return &section, nil
}

// CreateSection creates a new section.
func (c *Client) CreateSection(ctx context.Context, req CreateSectionRequest) (*Section, error) {
	var section Section
	err := c.post(ctx, "/sections", req, &section)
	if err != nil {
		return nil, err
	}
	return &section, nil
}

// UpdateSection updates an existing section.
func (c *Client) UpdateSection(ctx context.Context, id string, req UpdateSectionRequest) (*Section, error) {
	var section Section
	err := c.post(ctx, fmt.Sprintf("/sections/%s", id), req, &section)
	if err != nil {
		return nil, err
	}
	return &section, nil
}

// DeleteSection deletes a section.
func (c *Client) DeleteSection(ctx context.Context, id string) error {
	return c.delete(ctx, fmt.Sprintf("/sections/%s", id))
}
