## Keeper

Keeper is a CLI tool that manages your codebase. Keeper is designed to ease the burden of working on multiple, large mono-repos at once that need to stay in-sync and lean. As a result, a lot of what you'll find keeper has to offer thrives under the idea that keeper runs out of your "GitHub" folder.

### Getting started

Since no binaries are currently available, the best option you have is to:

1. Clone this repo
2. `go get` it's dependencies (cobra)
3. `go install`
4. Ensure your `$GOPATH` is set and that `$GOPATH/bin` is in your `$PATH`

### Config

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

This specifies the folder (codebase) where all your dirs live. Each dir entry is a key value where the key is the name of the dir in that folder and the value is the shorthand flag you'd like to use to specify it. Thus, you can do `--example` or `-e` to operate on the`/path/to/where/dirs/live/example` folder.
