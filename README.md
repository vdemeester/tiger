# 🐯 tiger [![Build Status](https://travis-ci.org/vdemeester/tiger.svg?branch=master)](https://travis-ci.org/vdemeester/tiger)

`tiger` is a a command line tool that wraps git in order to extend it with extra features and commands to make
your life wonderful. In a way it's like [hub](https://github.com/vdemeester/hub) but not tied to GitHub.

```bash
λ tiger st
# forward to git st
λ tiger rebase foo bar
# […]
λ tiger pr 1234
# […]
```

## Tiger hooks ⛳

One of the feature of `tiger` is hooks being run on any commands. It can allows you to do some pretty nifty things.
Several hooks are built-in `tiger` but as it's almost just commands run at certain point of time, it's pretty composable
and easy to defines yours.

An example is something I use with `nixos` and `direnv`. I like to have an `.envrc` and a `default.nix` that sets up
my development environment on a specific folder. But sometimes, these files are not present in an upstream project I
don't control. I tend to ignore them globally but I want to be able to keep them updated across my different working
station.

```bash
λ ls
λ tiger st
🐅 nix-direnv detected for this repo, files not present, updating the files
λ ls
# […]
# Later, the file has been update (by some other means)
λ tiger fetch
🐅 nix-direnv detected for this repo, files out of date, updating the files

```