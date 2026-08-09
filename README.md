# Task Tracker CLI

A lightweight command-line interface (CLI) tool written in **Go** to help you track and manage your to-do list. This project is built as a learning project following the specification from [roadmap.sh Task Tracker Project](https://roadmap.sh/projects/task-tracker).

---

## 📌 Features

- **Task Management**: Add, update, and delete tasks.
- **Status Tracking**: Mark tasks as `todo`, `in-progress`, or `done`.
- **Filtering**: View all tasks or filter tasks by their current status.
- **Persistent Storage**: Automatically saves and loads tasks from a local JSON file (`tasks.json`).
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
# Mark as in-progress
task-cli mark-in-progress 1

# Mark as done
task-cli mark-done 1
```

### 5. List Tasks
```bash
# List all tasks
task-cli list

# List tasks by status
task-cli list todo
task-cli list in-progress
task-cli list done
```

---

## 📁 File Storage

Tasks are stored locally in a `tasks.json` file in the execution directory. If the file does not exist when a command is executed, the application creates it automatically.

---

## 🔗 Reference

- Project Specification: [roadmap.sh Task Tracker](https://roadmap.sh/projects/task-tracker)
