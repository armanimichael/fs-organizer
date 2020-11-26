# File System Organizer

This project is some sort of learning activity for the Go language. The tool is called FSO (File System Organizer) and provides simple, yet effective, funcionalities to order and organize files.

For now, the CLI tool is able to rename subfolders inside a root directory, following certain rules such as converting them to uppercase or renaming the dirs to follow a numeric order.

## Running the tool

### Building the project

```
go build fso.go
```

### Syntax

```
fso -command
```

For example, to enumerate directories under a "test" directory.

```
fso -enum -root="test"
```

### Commands

- `help`: prints all commands.
- `root dirname`: sets working directory. By default the current directory.
- `enum`: rename dirs (only) to form a list.
- `uppercase`sets dirs (only) names to uppercase
- `lowercase`sets dirs (only) names to lowercase
