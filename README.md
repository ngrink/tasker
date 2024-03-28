# tasker
![GitHub Issues or Pull Requests](https://img.shields.io/github/issues/ngrink/tasker)
[![Release](https://img.shields.io/github/release/ngrink/tasker.svg)](https://github.com/ngrink/tasker/releases/latest)
[![Build Status](https://img.shields.io/github/actions/workflow/status/ngrink/tasker/build.yml?branch=main)](https://github.com/ngrink/tasker/actions?workflow=build)
![GitHub License](https://img.shields.io/github/license/ngrink/tasker)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg)](https://github.com/goreleaser)

## Description

Tasker is a simple command-line task manager. It was built using Golang with Cobra framework and follows the [Command Line Interface Guidelines](https://clig.dev)

The motivation for building this project was to learn how to create user-friendly, robust and reliable CLI program using Golang.

## Installation

`go install github.com/ngrink/tasker@latest`

## Available commands
  
| Command | Description |
| --- | --- |
| `tasker add`              | Add task to the list |
| `tasker ls`               | List all tasks |
| `tasker complete`         | Mark task as completed |
| `tasker uncomplete`       | Mark task as uncompleted staged |
| `tasker rm`               | Remove a task |
| `tasker edit`             | Edit a task desription |
| `tasker purge`            | Remove all completed tasks |
| `tasker clean`            | Remove all tasks |
  
For a detailed description of the use of each command, use the --help flag.

## Lisense
This project is licensed under the terms of the **MIT** license.
