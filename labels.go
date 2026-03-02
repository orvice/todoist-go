package todoist

import (
	"context"
	"fmt"
)

// GetLabels returns all personal labels.
func (c *Client) GetLabels(ctx context.Context) ([]Label, error) {
	return getList[Label](c, ctx, "/labels", nil)
}

// GetLabel returns a single personal label by ID.
func (c *Client) GetLabel(ctx context.Context, id string) (*Label, error) {
	var label Label
	err := c.get(ctx, fmt.Sprintf("/labels/%s", id), nil, &label)
	if err != nil {
		return nil, err
	}
	return &label, nil
}

// CreateLabel creates a new personal label.
func (c *Client) CreateLabel(ctx context.Context, req CreateLabelRequest) (*Label, error) {
	var label Label
	err := c.post(ctx, "/labels", req, &label)
	if err != nil {
		return nil, err
	}
	return &label, nil
}

// UpdateLabel updates an existing personal label.
func (c *Client) UpdateLabel(ctx context.Context, id string, req UpdateLabelRequest) (*Label, error) {
	var label Label
	err := c.post(ctx, fmt.Sprintf("/labels/%s", id), req, &label)
	if err != nil {
		return nil, err
	}
	return &label, nil
}

// DeleteLabel deletes a personal label.
func (c *Client) DeleteLabel(ctx context.Context, id string) error {
	return c.delete(ctx, fmt.Sprintf("/labels/%s", id))
}
