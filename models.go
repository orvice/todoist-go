package todoist

// PaginatedResponse is the wrapper for list endpoints.
type PaginatedResponse[T any] struct {
	Results    []T     `json:"results"`
	NextCursor *string `json:"next_cursor"`
}

// Project represents a Todoist project.
type Project struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Color          string  `json:"color"`
	ParentID       *string `json:"parent_id"`
	ChildOrder     int     `json:"child_order"`
	IsShared       bool    `json:"is_shared"`
	IsFavorite     bool    `json:"is_favorite"`
	InboxProject   bool    `json:"inbox_project"`
	IsArchived     bool    `json:"is_archived"`
	IsDeleted      bool    `json:"is_deleted"`
	IsFrozen       bool    `json:"is_frozen"`
	IsCollapsed    bool    `json:"is_collapsed"`
	ViewStyle      string  `json:"view_style"`
	Description    string  `json:"description"`
	CreatorUID     string  `json:"creator_uid"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	CanAssignTasks bool    `json:"can_assign_tasks"`
}

// CreateProjectRequest contains parameters for creating a project.
type CreateProjectRequest struct {
	Name       string  `json:"name"`
	ParentID   *string `json:"parent_id,omitempty"`
	Color      string  `json:"color,omitempty"`
	IsFavorite *bool   `json:"is_favorite,omitempty"`
	ViewStyle  string  `json:"view_style,omitempty"`
}

// UpdateProjectRequest contains parameters for updating a project.
type UpdateProjectRequest struct {
	Name       string  `json:"name,omitempty"`
	ParentID   *string `json:"parent_id,omitempty"`
	Color      string  `json:"color,omitempty"`
	IsFavorite *bool   `json:"is_favorite,omitempty"`
	ViewStyle  string  `json:"view_style,omitempty"`
}

// Section represents a Todoist section.
type Section struct {
	ID           string  `json:"id"`
	ProjectID    string  `json:"project_id"`
	Name         string  `json:"name"`
	SectionOrder int     `json:"section_order"`
	UserID       string  `json:"user_id"`
	AddedAt      string  `json:"added_at"`
	UpdatedAt    string  `json:"updated_at"`
	ArchivedAt   *string `json:"archived_at"`
	IsArchived   bool    `json:"is_archived"`
	IsDeleted    bool    `json:"is_deleted"`
	IsCollapsed  bool    `json:"is_collapsed"`
}

// CreateSectionRequest contains parameters for creating a section.
type CreateSectionRequest struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
	Order     int    `json:"order,omitempty"`
}

// UpdateSectionRequest contains parameters for updating a section.
type UpdateSectionRequest struct {
	Name string `json:"name"`
}

// Due represents the due date of a task.
type Due struct {
	String      string `json:"string"`
	Date        string `json:"date"`
	IsRecurring bool   `json:"is_recurring"`
	Datetime    string `json:"datetime,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	Lang        string `json:"lang,omitempty"`
}

// Deadline represents a task deadline.
type Deadline struct {
	Date string `json:"date"`
}

// Duration represents a task time estimate.
type Duration struct {
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}

// Task represents a Todoist task.
type Task struct {
	ID             string    `json:"id"`
	Content        string    `json:"content"`
	Description    string    `json:"description"`
	ProjectID      string    `json:"project_id"`
	SectionID      *string   `json:"section_id"`
	ParentID       *string   `json:"parent_id"`
	ChildOrder     int       `json:"child_order"`
	Checked        bool      `json:"checked"`
	Labels         []string  `json:"labels"`
	Priority       int       `json:"priority"`
	NoteCount      int       `json:"note_count"`
	UserID         string    `json:"user_id"`
	AddedByUID     string    `json:"added_by_uid"`
	AddedAt        string    `json:"added_at"`
	UpdatedAt      string    `json:"updated_at"`
	CompletedAt    *string   `json:"completed_at"`
	AssignedByUID  *string   `json:"assigned_by_uid"`
	ResponsibleUID *string   `json:"responsible_uid"`
	Due            *Due      `json:"due"`
	Deadline       *Deadline `json:"deadline"`
	Duration       *Duration `json:"duration"`
	IsDeleted      bool      `json:"is_deleted"`
	IsCollapsed    bool      `json:"is_collapsed"`
	DayOrder       int       `json:"day_order"`
}

// CreateTaskRequest contains parameters for creating a task.
type CreateTaskRequest struct {
	Content        string   `json:"content"`
	Description    string   `json:"description,omitempty"`
	ProjectID      string   `json:"project_id,omitempty"`
	SectionID      string   `json:"section_id,omitempty"`
	ParentID       string   `json:"parent_id,omitempty"`
	Order          int      `json:"order,omitempty"`
	Labels         []string `json:"labels,omitempty"`
	Priority       int      `json:"priority,omitempty"`
	DueString      string   `json:"due_string,omitempty"`
	DueDate        string   `json:"due_date,omitempty"`
	DueDatetime    string   `json:"due_datetime,omitempty"`
	DueLang        string   `json:"due_lang,omitempty"`
	ResponsibleUID string   `json:"responsible_uid,omitempty"`
	Duration       int      `json:"duration,omitempty"`
	DurationUnit   string   `json:"duration_unit,omitempty"`
	DeadlineDate   string   `json:"deadline_date,omitempty"`
}

// UpdateTaskRequest contains parameters for updating a task.
type UpdateTaskRequest struct {
	Content        *string  `json:"content,omitempty"`
	Description    *string  `json:"description,omitempty"`
	Labels         []string `json:"labels,omitempty"`
	Priority       *int     `json:"priority,omitempty"`
	DueString      *string  `json:"due_string,omitempty"`
	DueDate        *string  `json:"due_date,omitempty"`
	DueDatetime    *string  `json:"due_datetime,omitempty"`
	DueLang        *string  `json:"due_lang,omitempty"`
	ResponsibleUID *string  `json:"responsible_uid,omitempty"`
	Duration       *int     `json:"duration,omitempty"`
	DurationUnit   *string  `json:"duration_unit,omitempty"`
	DeadlineDate   *string  `json:"deadline_date,omitempty"`
}

// GetTasksOptions contains query parameters for listing tasks.
type GetTasksOptions struct {
	ProjectID string `url:"project_id,omitempty"`
	SectionID string `url:"section_id,omitempty"`
	Label     string `url:"label,omitempty"`
	Filter    string `url:"filter,omitempty"`
	Lang      string `url:"lang,omitempty"`
	IDs       string `url:"ids,omitempty"`
}

// Comment represents a Todoist comment.
type Comment struct {
	ID         string      `json:"id"`
	TaskID     *string     `json:"task_id"`
	ProjectID  *string     `json:"project_id"`
	PostedAt   string      `json:"posted_at"`
	Content    string      `json:"content"`
	Attachment *Attachment `json:"attachment"`
}

// Attachment represents a file attachment on a comment.
type Attachment struct {
	FileName     string `json:"file_name"`
	FileType     string `json:"file_type"`
	FileURL      string `json:"file_url"`
	ResourceType string `json:"resource_type"`
}

// CreateCommentRequest contains parameters for creating a comment.
type CreateCommentRequest struct {
	TaskID     string      `json:"task_id,omitempty"`
	ProjectID  string      `json:"project_id,omitempty"`
	Content    string      `json:"content"`
	Attachment *Attachment `json:"attachment,omitempty"`
}

// UpdateCommentRequest contains parameters for updating a comment.
type UpdateCommentRequest struct {
	Content string `json:"content"`
}

// Label represents a personal Todoist label.
type Label struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	Order      int    `json:"order"`
	IsFavorite bool   `json:"is_favorite"`
}

// CreateLabelRequest contains parameters for creating a label.
type CreateLabelRequest struct {
	Name       string `json:"name"`
	Color      string `json:"color,omitempty"`
	Order      int    `json:"order,omitempty"`
	IsFavorite *bool  `json:"is_favorite,omitempty"`
}

// UpdateLabelRequest contains parameters for updating a label.
type UpdateLabelRequest struct {
	Name       string `json:"name,omitempty"`
	Color      string `json:"color,omitempty"`
	Order      *int   `json:"order,omitempty"`
	IsFavorite *bool  `json:"is_favorite,omitempty"`
}

// Collaborator represents a project collaborator.
type Collaborator struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
