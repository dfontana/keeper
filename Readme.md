# Keeper

Keeper is a CLI tool that manages your codebase. Keeper is designed to ease the burden of working on multiple, large mono-repos at once that need to stay in-sync and lean. As a result, a lot of what you'll find keeper has to offer thrives under the idea that keeper runs out of your "GitHub" folder.

## Getting started

Since no binaries are currently available, the best option you have is to get the package and install it:

1. `go get -u github.com/dfontana/keeper`
2. `cd $GOPATH/src/github.com/dfontana/keeper`
3. `go install`

Note you'll want to make sure your `GOPATH` is set and `$GOPATH/bin` is in your `PATH`.

#### Side note:

If you'd like to clone via SSH rather than HTTPS, a helpful setting: `git config --global url.git@github.com:.insteadOf https://github.com/`

## Config

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
