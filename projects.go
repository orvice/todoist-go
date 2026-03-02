package todoist

import (
	"context"
	"fmt"
)

// GetProjects returns all user projects.
func (c *Client) GetProjects(ctx context.Context) ([]Project, error) {
	return getList[Project](c, ctx, "/projects", nil)
}

// GetProject returns a single project by ID.
func (c *Client) GetProject(ctx context.Context, id string) (*Project, error) {
	var project Project
	err := c.get(ctx, fmt.Sprintf("/projects/%s", id), nil, &project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// CreateProject creates a new project.
func (c *Client) CreateProject(ctx context.Context, req CreateProjectRequest) (*Project, error) {
	var project Project
	err := c.post(ctx, "/projects", req, &project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// UpdateProject updates an existing project.
func (c *Client) UpdateProject(ctx context.Context, id string, req UpdateProjectRequest) (*Project, error) {
	var project Project
	err := c.post(ctx, fmt.Sprintf("/projects/%s", id), req, &project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// ArchiveProject archives a project.
func (c *Client) ArchiveProject(ctx context.Context, id string) error {
	return c.post(ctx, fmt.Sprintf("/projects/%s/archive", id), nil, nil)
}

// UnarchiveProject unarchives a project.
func (c *Client) UnarchiveProject(ctx context.Context, id string) error {
	return c.post(ctx, fmt.Sprintf("/projects/%s/unarchive", id), nil, nil)
}

// DeleteProject deletes a project.
func (c *Client) DeleteProject(ctx context.Context, id string) error {
	return c.delete(ctx, fmt.Sprintf("/projects/%s", id))
}

// GetProjectCollaborators returns collaborators for a shared project.
func (c *Client) GetProjectCollaborators(ctx context.Context, projectID string) ([]Collaborator, error) {
	return getList[Collaborator](c, ctx, fmt.Sprintf("/projects/%s/collaborators", projectID), nil)
}
