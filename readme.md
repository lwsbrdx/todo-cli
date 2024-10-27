# Todo List CLI

A command-line tool for managing your daily tasks

## Description

Todo List CLI is a command-line tool that allows you to create, modify, and delete tasks. It is designed to be simple and easy to use, while offering advanced features for users who need them.

## Features

* Create tasks with a title, description, and status (to-do, in progress, done)
* Modify existing tasks
* Delete tasks
* Display the list of tasks
* Search for tasks by title or description
* Add projects and assign tasks to projects

## Installation

To install Todo List CLI, you need to have Go installed on your system. You can download the source code and run it with the command `go run main.go` or install it using `go install .` command.

## Commands

* `todo task add`: Create a new task
* `todo task update`: Modify an existing task
* `todo task delete`: Delete a task
* `todo task list`: Display the list of tasks
* `todo task search`: Search for tasks by title or description
* `todo project add`: Add a new project
* `todo project update`: Modify an existing project (WIP)
* `todo project list`: Display the list of projects
* `todo project delete`: Delete a project (WIP)

## Flags

* `--name` (`-n`): Set the title of the task or project
* `--description` (`-d`): Set the description of the task
* `--status` (`-s`): Set the status of the task (to-do, in progress, done)
* `--project` (`-p`): Set the project ID for the task
* `--trashed` (`-t`): Show trashed tasks
* `--empty-trash`: Empty trashed tasks

## Examples

* `todo task add --name "Buy groceries"`: Create a new task with the title "Buy groceries"
* `todo task update 1 --name "Buy groceries tomorrow"`: Modify the task with ID 1 and the title "Buy groceries tomorrow"
* `todo task delete 1`: Delete the task with ID 1
* `todo task list`: Display the list of tasks
* `todo task search "groceries"`: Search for tasks with the title or description "groceries"
* `todo project add --name "My project"`: Add a new project with the title "My project"
* `todo project list`: Display the list of projects

## Contribution

If you want to contribute to Todo List CLI, you can fork the repository and submit pull requests. We are open to all suggestions and improvements.

## License

This README is published under the [MIT License](https://opensource.org/licenses/MIT). You are free to share, modify, and use this content without restrictions.
