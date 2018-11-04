# Keeper

Keeper is a CLI tool that manages your codebase. Keeper is designed to ease the burden of working on multiple, large mono-repos at once that need to stay in-sync and lean. As a result, a lot of what you'll find keeper has to offer thrives under the idea that keeper runs out of your "GitHub" folder.

## Table of Contents

- [Getting Started](#getting-started)
- [Basic Config](#basic-config)
- [Commands](#commands)
  - [list](#list)
  - [del](#del)
  - [run](#run)
  - [start](#start)
- [TODO](#todo)

## Getting started

Since no binaries are currently available, the best option you have is to get the package and install it:

1. `go get -u github.com/dfontana/keeper`
2. `cd $GOPATH/src/github.com/dfontana/keeper`
3. `go install`

Note you'll want to make sure your `GOPATH` is set and `$GOPATH/bin` is in your `PATH`.

#### Helpful (?) Hints:

- If prefer SSH to HTTPS, make sure this is set: `git config --global url.git@github.com:.insteadOf https://github.com/`
- If you don't have `go` installed, `brew install go` does the trick :)

## Basic Config

To work correctly you must specify a config in your `$HOME` called `.keeper`. Inside should be, at minimum, a JSON with keys:

```
{
  "dirs" : {
    "dir_name" : "abbreviation_letter",
    "example" : "e"
  },
  "codebase" : "/path/to/where/dirs/live"
}
```

This specifies the folder (codebase) where all your dirs live. Each dir entry is a key value where the key is the name of the dir in that folder and the value is the shorthand flag you'd like to use to specify it. Thus, you can do `--example` or `-e` to operate on the`/path/to/where/dirs/live/example` folder with keeper.

## Commands

This section gives a brief overview of what each command does. For details on how to utilize and run, see `--help`.

### list

Displays branches with their author. By providing a search string, you can filter the output to only those containing the filter. By default, if no filter is provided the `git config user.name` is utilized as the filter instead. Omitting directory flags will run the command in the current directory. Running with multiple directory flags will run the command in each, printing the directory of the branch in addition to the author and name.

`-i` will make the search insensitive.

### del

Deletes the given list of branches, from both the local and remote. `<remote>/` need not be prepended to the branch name. If a directory flag is omitted, the current directory is used. Providing multiple directory flags will invalidate the command. Before deletion, you will be prompted to confirm each directory to be deleted, _after which the deletion will occur_ allowing the user time to cancel if they made a mistake.

### run

Runs the given command against each directory, returning once all commands have finished. Must provide at least one directory flag.

### start

An each east way to start a new feature by creating a branch off `origin/master` in each of the provided directory flags. The command will attempt to checkout into these branches, but is not guaranteed to succeed if any directory contains uncommitted changes. If your `.keeper` config contains a template string for the start command, then the provided name will be substituted into the template for each instance of `${NAME}` before naming each branch.

```
{
  ...,
  "start": "feature_${NAME}"
  ...
}
```
