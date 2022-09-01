# VMerge

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/vmerge.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/vmerge/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/vmerge/actions/workflows/go.yml) [![CodeQL](https://github.com/gowizzard/vmerge/actions/workflows/codeql.yml/badge.svg)](https://github.com/gowizzard/vmerge/actions/workflows/codeql.yml) [![VMerge](https://github.com/gowizzard/vmerge/actions/workflows/vmerge.yml/badge.svg)](https://github.com/gowizzard/vmerge/actions/workflows/vmerge.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/vmerge)](https://goreportcard.com/report/github.com/gowizzard/vmerge) [![GitHub issues](https://img.shields.io/github/issues/gowizzard/vmerge)](https://github.com/gowizzard/vmerge/issues) [![GitHub forks](https://img.shields.io/github/forks/gowizzard/vmerge)](https://github.com/gowizzard/vmerge/network) [![GitHub stars](https://img.shields.io/github/stars/gowizzard/vmerge)](https://github.com/gowizzard/vmerge/stargazers) [![GitHub license](https://img.shields.io/github/license/gowizzard/vmerge)](https://github.com/gowizzard/vmerge/blob/master/LICENSE)

This small solution is intended to automatically merge a branch, with each new tag / release, the data from the default branch. 

This is especially important for solutions like Go libraries, where you have to maintain an associated `v2` branch from version 2 on. So you can use this action exactly for this and with every new version the versions branch gets a merge to the default branch. Thus, this is always supplied with the correct data.

For this to work correctly, the corresponding branch must be created. You can do this directly at [GitHub](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-and-deleting-branches-within-your-repository). You can decide how this should be called. But you have to consider this when creating the action. [Here](https://github.com/gowizzard/vmerge#using-the-github-action) you can find an example of how you can use the action. In this example, because it is convenient for golang, we use a suffix of `v`, so that when the new tag is checked, the major version number is output with the `v` prepended.

That means we would currently create a `v1` branch, which is only maintained as long as there are releases / tags in the major version 1. After that we create a branch for `v2`, so that this can also be maintained automatically.

## Using the GitHub Action

Here you can find an example for the action, so you can use it in your project. How to create a github action you can find [here](https://docs.github.com/en/actions/quickstart).

If you use this example, please pay attention to the version number below the dot `Get the major version`.

```yaml
name: VMerge

on:
  push:
    tags:
        - "v*.*.*"

env:
  USER_NAME: "GitHub Action"
  USER_EMAIL: "actions@github.com"
  DEFAULT_BRANCH: ${{ github.event.repository.default_branch }}
  COMMIT_MESSAGE: "ci: The data of the master branch was merged automatically."

jobs:
  version:
    runs-on: ubuntu-latest
    steps:

      - name: Clone repository
        uses: actions/checkout@v3
        with:
            fetch-depth: 0

      - name: Set repository tag information
        id: information
        run: echo ::set-output name=tag::${GITHUB_REF#refs/*/}

      - name: Get the major version
        id: vmerge
        uses: gowizzard/vmerge@v1.0.0
        env:
            VERSION: ${{ steps.information.outputs.tag }}
            SUFFIX: "v"

      - name: Set git config
        run: |
            git config --local user.name "$USER_NAME"
            git config --local user.email "$USER_EMAIL"

      - name: Merge data from master branch
        run: |
            git fetch
            git checkout ${{ steps.vmerge.outputs.branch_name }}
            git pull
            git merge --no-ff "origin/$DEFAULT_BRANCH" -m "COMMIT_MESSAGE"
            git push
```

## Special thanks

Thanks to [JetBrains](https://github.com/JetBrains) for supporting me with this and other [open source projects](https://www.jetbrains.com/community/opensource/#support).