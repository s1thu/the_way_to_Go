todo-cli/
├── cmd/
│ └── root.go # Entry point CLI command handlers
├── todo/
│ ├── task.go # Task struct and helpers
│ ├── manager.go # Add, List, Done, Delete logic (pure functions)
│ ├── storage.go # Load/Save to JSON file
│ ├── task_test.go # Unit tests for task logic
│ ├── manager_test.go # Unit tests for business logic
│ └── storage_test.go # Unit tests for file I/O (can be mocked)
├── tasks.json # File to store tasks (auto-created)
├── go.mod
├── main.go # Main entry — delegates to cmd/root.go
└── README.md # Project instructions
