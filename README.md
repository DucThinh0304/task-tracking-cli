# Task Tracker CLI

A lightweight command-line interface (CLI) tool written in **Go** to help you track and manage your to-do list. This project is built as a learning project following the specification from [roadmap.sh Task Tracker Project](https://roadmap.sh/projects/task-tracker).

---

## 📌 Features

- **Task Management**: Add, update, and delete tasks.
- **Status Tracking**: Mark tasks as `todo`, `in-progress`, or `done`.
- **Filtering**: View all tasks or filter tasks by their current status.
- **Persistent Storage**: Automatically saves and loads tasks from a JSON file in your home directory (`~/.task-tracker-cli/tasks.json`).
- **Zero External Dependencies**: Built entirely using Go's standard library.

---

## 🛠️ Task Data Structure

Each task stored in `tasks.json` has the following schema:

```json
{
  "id": 1,
  "description": "Buy groceries",
  "status": "todo",
  "createdAt": "2026-08-07T17:00:00Z",
  "updatedAt": "2026-08-07T17:00:00Z"
}
```

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) (version 1.22 or higher recommended)

### Build & Installation

1. **Clone or navigate to the project directory:**
   ```bash
   cd task-tracker-cli
   ```

2. **Build the CLI executable:**
   ```bash
   go build -o task-cli .
   ```

---

## 💻 Usage

Once built, you can execute commands using `./task-cli` (or run directly using `go run main.go -- <command>`):

### 1. Add a New Task
```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### 2. Update a Task
```bash
task-cli update 1 "Buy groceries and cook dinner"
# Output: Task updated successfully
```

### 3. Delete a Task
```bash
task-cli delete 1
# Output: Task deleted successfully
```

### 4. Update Task Status
```bash
# Mark as todo
task-cli mark-todo 1

# Mark as in-progress
task-cli mark-in-progress 1

# Mark as done
task-cli mark-done 1
# Output: Task marked as <status> successfully
```

### 5. List Tasks
```bash
# List all tasks
task-cli list
# Output format: Task ID: 1, Name: Buy groceries, Status: todo

# List tasks by status
task-cli list todo
task-cli list in-progress
task-cli list done
```

---

## 📁 File Storage

Tasks are stored in `~/.task-tracker-cli/tasks.json` in the user's home directory (or `%USERPROFILE%\.task-tracker-cli\tasks.json` on Windows). If the home directory cannot be resolved, it falls back to `tasks.json` in the local execution directory.

If the directory or file does not exist when a command is executed, the application creates them automatically.

---

## 🔗 Reference

- Project Specification: [roadmap.sh Task Tracker](https://roadmap.sh/projects/task-tracker)

