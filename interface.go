package contextmanager

import "context"

// ContextManager defines the interface for managing contexts within the system.
type ContextManager interface {
	// WriteContext writes a new context into the system. It returns an error if
	// the operation fails, for example, if the context already exists for the TaskID.
	//
	WriteContext(ctx context.Context, taskID string, parentTaskID string, content string) error

	// UpdateContext updates an existing context identified by TaskID. It may use
	// the version number to handle concurrency and ensure consistency. It returns
	// an error if the update fails.
	// This is mainly used to correct the context during manual review.
	// If task dependencies change, it is recommended to create a new version instead of updating it directly.
	UpdateContext(ctx context.Context, taskID string, newContent string, version string) error

	// ReadContext retrieves the context based on the TaskID. It returns the corresponding
	// context and an error, if any. The error could indicate that the context does not exist.
	ReadContext(ctx context.Context, taskID string) (string, error)

	// SetVersionLatest set the specified version as the latest version.
	// It usually used to rollback the changes to the history version.
	SetVersionLatest(ctx context.Context, taskID string, version string) error

	// DeleteContext removes a context from the system based on the TaskID. It returns
	// an error if the delete operation fails, for example, if the context does not exist.
	// We may not need to manually delete the context.
	// This method is used to preserve design integrity and can also be used for some rollback operations.
	DeleteContext(ctx context.Context, taskID string) error

	// DeleteContextByVersion removes a specific version context from the system based on the TaskID. It returns
	// an error if the delete operation fails, for example, if the context does not exist.
	// We may not need to manually delete the context.
	// This method is used to preserve design integrity and can also be used for some rollback operations.
	DeleteContextByVersion(ctx context.Context, taskID string) error
}
