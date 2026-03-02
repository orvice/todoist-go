package todoist

import (
	"context"
	"fmt"
)

// GetCommentsOptions contains query parameters for listing comments.
type GetCommentsOptions struct {
	TaskID    string `url:"task_id,omitempty"`
	ProjectID string `url:"project_id,omitempty"`
}

// GetComments returns comments for a task or project.
func (c *Client) GetComments(ctx context.Context, opts *GetCommentsOptions) ([]Comment, error) {
	return getList[Comment](c, ctx, "/comments", opts)
}

// GetComment returns a single comment by ID.
func (c *Client) GetComment(ctx context.Context, id string) (*Comment, error) {
	var comment Comment
	err := c.get(ctx, fmt.Sprintf("/comments/%s", id), nil, &comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// CreateComment creates a new comment on a task or project.
func (c *Client) CreateComment(ctx context.Context, req CreateCommentRequest) (*Comment, error) {
	var comment Comment
	err := c.post(ctx, "/comments", req, &comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// UpdateComment updates an existing comment.
func (c *Client) UpdateComment(ctx context.Context, id string, req UpdateCommentRequest) (*Comment, error) {
	var comment Comment
	err := c.post(ctx, fmt.Sprintf("/comments/%s", id), req, &comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// DeleteComment deletes a comment.
func (c *Client) DeleteComment(ctx context.Context, id string) error {
	return c.delete(ctx, fmt.Sprintf("/comments/%s", id))
}
