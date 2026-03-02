package todoist

import (
	"context"
	"fmt"
)

// GetTasks returns active tasks, optionally filtered.
func (c *Client) GetTasks(ctx context.Context, opts *GetTasksOptions) ([]Task, error) {
	return getList[Task](c, ctx, "/tasks", opts)
}

// GetTask returns a single active task by ID.
func (c *Client) GetTask(ctx context.Context, id string) (*Task, error) {
	var task Task
	err := c.get(ctx, fmt.Sprintf("/tasks/%s", id), nil, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// CreateTask creates a new task.
func (c *Client) CreateTask(ctx context.Context, req CreateTaskRequest) (*Task, error) {
	var task Task
	err := c.post(ctx, "/tasks", req, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// UpdateTask updates an existing task.
func (c *Client) UpdateTask(ctx context.Context, id string, req UpdateTaskRequest) (*Task, error) {
	var task Task
	err := c.post(ctx, fmt.Sprintf("/tasks/%s", id), req, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// CloseTask marks a task as complete.
func (c *Client) CloseTask(ctx context.Context, id string) error {
	return c.post(ctx, fmt.Sprintf("/tasks/%s/close", id), nil, nil)
}

// ReopenTask reopens a completed task.
func (c *Client) ReopenTask(ctx context.Context, id string) error {
	return c.post(ctx, fmt.Sprintf("/tasks/%s/reopen", id), nil, nil)
}

// DeleteTask deletes a task.
func (c *Client) DeleteTask(ctx context.Context, id string) error {
	return c.delete(ctx, fmt.Sprintf("/tasks/%s", id))
}
