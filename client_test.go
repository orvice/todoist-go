package todoist_test

import (
	"context"
	"os"
	"testing"

	todoist "github.com/orvice/todoist-go"
)

func newTestClient(t *testing.T) *todoist.Client {
	t.Helper()
	token := os.Getenv("TEST_TODOIST_TOKEN")
	if token == "" {
		t.Skip("TEST_TODOIST_TOKEN not set")
	}
	return todoist.New(token)
}

func TestProjects(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	// Create
	project, err := client.CreateProject(ctx, todoist.CreateProjectRequest{
		Name: "SDK Test Project",
	})
	if err != nil {
		t.Fatalf("CreateProject: %v", err)
	}
	if project.Name != "SDK Test Project" {
		t.Fatalf("expected name 'SDK Test Project', got %q", project.Name)
	}
	t.Logf("created project %s", project.ID)

	// Cleanup at end
	defer func() {
		if err := client.DeleteProject(ctx, project.ID); err != nil {
			t.Errorf("DeleteProject: %v", err)
		}
	}()

	// Get
	got, err := client.GetProject(ctx, project.ID)
	if err != nil {
		t.Fatalf("GetProject: %v", err)
	}
	if got.ID != project.ID {
		t.Fatalf("expected id %s, got %s", project.ID, got.ID)
	}

	// List
	projects, err := client.GetProjects(ctx)
	if err != nil {
		t.Fatalf("GetProjects: %v", err)
	}
	found := false
	for _, p := range projects {
		if p.ID == project.ID {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("created project not found in list")
	}

	// Update
	updated, err := client.UpdateProject(ctx, project.ID, todoist.UpdateProjectRequest{
		Name: "SDK Test Project Updated",
	})
	if err != nil {
		t.Fatalf("UpdateProject: %v", err)
	}
	if updated.Name != "SDK Test Project Updated" {
		t.Fatalf("expected updated name, got %q", updated.Name)
	}
}

func TestTasks(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	// Create a project for the task
	project, err := client.CreateProject(ctx, todoist.CreateProjectRequest{
		Name: "SDK Test Tasks Project",
	})
	if err != nil {
		t.Fatalf("CreateProject: %v", err)
	}
	defer func() {
		_ = client.DeleteProject(ctx, project.ID)
	}()

	// Create task
	task, err := client.CreateTask(ctx, todoist.CreateTaskRequest{
		Content:   "SDK Test Task",
		ProjectID: project.ID,
		Priority:  2,
	})
	if err != nil {
		t.Fatalf("CreateTask: %v", err)
	}
	t.Logf("created task %s", task.ID)

	defer func() {
		_ = client.DeleteTask(ctx, task.ID)
	}()

	// Get
	got, err := client.GetTask(ctx, task.ID)
	if err != nil {
		t.Fatalf("GetTask: %v", err)
	}
	if got.Content != "SDK Test Task" {
		t.Fatalf("expected content 'SDK Test Task', got %q", got.Content)
	}
	if got.Priority != 2 {
		t.Fatalf("expected priority 2, got %d", got.Priority)
	}

	// List by project
	tasks, err := client.GetTasks(ctx, &todoist.GetTasksOptions{
		ProjectID: project.ID,
	})
	if err != nil {
		t.Fatalf("GetTasks: %v", err)
	}
	if len(tasks) == 0 {
		t.Fatal("expected at least 1 task")
	}

	// Update
	newContent := "SDK Test Task Updated"
	updated, err := client.UpdateTask(ctx, task.ID, todoist.UpdateTaskRequest{
		Content: &newContent,
	})
	if err != nil {
		t.Fatalf("UpdateTask: %v", err)
	}
	if updated.Content != newContent {
		t.Fatalf("expected updated content, got %q", updated.Content)
	}

	// Close
	if err := client.CloseTask(ctx, task.ID); err != nil {
		t.Fatalf("CloseTask: %v", err)
	}

	// Reopen
	if err := client.ReopenTask(ctx, task.ID); err != nil {
		t.Fatalf("ReopenTask: %v", err)
	}
}

func TestSections(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	// Create a project
	project, err := client.CreateProject(ctx, todoist.CreateProjectRequest{
		Name: "SDK Test Sections Project",
	})
	if err != nil {
		t.Fatalf("CreateProject: %v", err)
	}
	defer func() {
		_ = client.DeleteProject(ctx, project.ID)
	}()

	// Create section
	section, err := client.CreateSection(ctx, todoist.CreateSectionRequest{
		Name:      "SDK Test Section",
		ProjectID: project.ID,
	})
	if err != nil {
		t.Fatalf("CreateSection: %v", err)
	}
	t.Logf("created section %s", section.ID)

	defer func() {
		_ = client.DeleteSection(ctx, section.ID)
	}()

	// Get
	got, err := client.GetSection(ctx, section.ID)
	if err != nil {
		t.Fatalf("GetSection: %v", err)
	}
	if got.Name != "SDK Test Section" {
		t.Fatalf("expected name 'SDK Test Section', got %q", got.Name)
	}

	// List by project
	sections, err := client.GetSections(ctx, &todoist.GetSectionsOptions{
		ProjectID: project.ID,
	})
	if err != nil {
		t.Fatalf("GetSections: %v", err)
	}
	if len(sections) == 0 {
		t.Fatal("expected at least 1 section")
	}

	// Update
	updated, err := client.UpdateSection(ctx, section.ID, todoist.UpdateSectionRequest{
		Name: "SDK Test Section Updated",
	})
	if err != nil {
		t.Fatalf("UpdateSection: %v", err)
	}
	if updated.Name != "SDK Test Section Updated" {
		t.Fatalf("expected updated name, got %q", updated.Name)
	}
}

func TestComments(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	// Create a project and task
	project, err := client.CreateProject(ctx, todoist.CreateProjectRequest{
		Name: "SDK Test Comments Project",
	})
	if err != nil {
		t.Fatalf("CreateProject: %v", err)
	}
	defer func() {
		_ = client.DeleteProject(ctx, project.ID)
	}()

	task, err := client.CreateTask(ctx, todoist.CreateTaskRequest{
		Content:   "SDK Test Comment Task",
		ProjectID: project.ID,
	})
	if err != nil {
		t.Fatalf("CreateTask: %v", err)
	}
	defer func() {
		_ = client.DeleteTask(ctx, task.ID)
	}()

	// Create comment on task
	comment, err := client.CreateComment(ctx, todoist.CreateCommentRequest{
		TaskID:  task.ID,
		Content: "SDK test comment",
	})
	if err != nil {
		t.Fatalf("CreateComment: %v", err)
	}
	t.Logf("created comment %s", comment.ID)

	defer func() {
		_ = client.DeleteComment(ctx, comment.ID)
	}()

	// Get
	got, err := client.GetComment(ctx, comment.ID)
	if err != nil {
		t.Fatalf("GetComment: %v", err)
	}
	if got.Content != "SDK test comment" {
		t.Fatalf("expected content 'SDK test comment', got %q", got.Content)
	}

	// List by task
	comments, err := client.GetComments(ctx, &todoist.GetCommentsOptions{
		TaskID: task.ID,
	})
	if err != nil {
		t.Fatalf("GetComments: %v", err)
	}
	if len(comments) == 0 {
		t.Fatal("expected at least 1 comment")
	}

	// Update
	updated, err := client.UpdateComment(ctx, comment.ID, todoist.UpdateCommentRequest{
		Content: "SDK test comment updated",
	})
	if err != nil {
		t.Fatalf("UpdateComment: %v", err)
	}
	if updated.Content != "SDK test comment updated" {
		t.Fatalf("expected updated content, got %q", updated.Content)
	}
}

func TestLabels(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	// Create
	label, err := client.CreateLabel(ctx, todoist.CreateLabelRequest{
		Name:  "sdk-test-label",
		Color: "red",
	})
	if err != nil {
		t.Fatalf("CreateLabel: %v", err)
	}
	t.Logf("created label %s", label.ID)

	defer func() {
		_ = client.DeleteLabel(ctx, label.ID)
	}()

	// Get
	got, err := client.GetLabel(ctx, label.ID)
	if err != nil {
		t.Fatalf("GetLabel: %v", err)
	}
	if got.Name != "sdk-test-label" {
		t.Fatalf("expected name 'sdk-test-label', got %q", got.Name)
	}

	// List
	labels, err := client.GetLabels(ctx)
	if err != nil {
		t.Fatalf("GetLabels: %v", err)
	}
	found := false
	for _, l := range labels {
		if l.ID == label.ID {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("created label not found in list")
	}

	// Update
	updated, err := client.UpdateLabel(ctx, label.ID, todoist.UpdateLabelRequest{
		Name: "sdk-test-label-updated",
	})
	if err != nil {
		t.Fatalf("UpdateLabel: %v", err)
	}
	if updated.Name != "sdk-test-label-updated" {
		t.Fatalf("expected updated name, got %q", updated.Name)
	}
}
