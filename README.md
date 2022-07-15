# Banner

A simple, lightweight, easy to use, cross-platform, server-intended banner generator

## Installation

Download the binary from release page.

## How it works

The printed banner is templated (see: `main.go`). The title can be optionally set via the `title` flag. The default title value is `"Wrecking Ball"`. Commands are described in a text file, (see: `resources/commands.txt` for an example). Path to the text command file can be set via the `path` flag which is required.

## Usage:

```
$ banner --path /some/absolute/path/commands.txt
$ banner --title Server 123 --path commands.txt
```