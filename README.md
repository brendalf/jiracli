# jiracli

Jira CLI written in Go that allows developers to work seamlessly with Jira inside their favorite terminal.

This is a side project to help me be productive with Jira and also sharp my Golang skills.

## Install

1. Clone this repository.
2. Edit the jira configuration constants inside api/api.go file.
3. Run `$ go run . --help`.

## Usage

There are currently two commands available:

1. List tickets:
```
$ jiracli list
```

2. View ticket:
```
$ jiracli view <ticket_id>
```
