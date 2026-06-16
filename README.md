# jx-helpers

[![Documentation](https://godoc.org/github.com/spring-financial-group/jx-helpers?status.svg)](https://pkg.go.dev/mod/github.com/spring-financial-group/jx-helpers)
[![Go Report Card](https://goreportcard.com/badge/github.com/spring-financial-group/jx-helpers)](https://goreportcard.com/report/github.com/spring-financial-group/jx-helpers)
[![Releases](https://img.shields.io/github/release-pre/spring-financial-group/jx-helpers.svg)](https://github.com/spring-financial-group/jx-helpers/releases)
[![LICENSE](https://img.shields.io/github/license/spring-financial-group/jx-helpers.svg)](https://github.com/spring-financial-group/jx-helpers/blob/master/LICENSE)
[![Slack Status](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://slack.k8s.io/)

`jx-helpers` is a small library of helper functions for working with the commands, git, scm and kubernetes

## Refactoring notes

If you are refactoring code from the v2 branch of jenkins-x/jx out into a separate library/microservice here's some tips on switching code over to `jx-helpers`

The `util` package has been split up into separate packages:

Files:

* `util.Copy*` => `files.Copy*`
* `util.Dir*` => `files.Dir*`
* `util.File*` => `files.File*`
* `util.IO*` => `files.IO*`
* `util.Un*` => `files.Un*`

Strings:

* `util.String*` => `stringhelpers.String*`
* `util.Url*` => `stringhelpers.Url*`

Git:

* `gits.Gitter` => `gitclient.Interface`  then the git commands are simple CLI arguments like:

```go 
results, err := gitClient.Command(dir, "commit", "-a", "-m", "my message")
```

Cobra:

* `cmd/*` => `cobras/*`

