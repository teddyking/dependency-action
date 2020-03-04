# dependency-action

![](https://github.com/teddyking/dependency-action/workflows/pipeline/badge.svg)

A GitHub action that can be used to download and extract dependencies during workflow jobs.

## Usage

Simply add a `step` to a job in your workflow, as follows:

```
name: my-example-workflow
on: [push]
jobs:
  test:
    name: unit-test
    runs-on: ubuntu-latest
    steps:
    - name: Checkout master
      uses: actions/checkout@master
    - name: Download test dependencies
      uses: teddyking/dependency-action@master
      with:
        deps: https://dep1.tar.gz,https://dep2.txz
    - name: Run tests
      ...
```

The `deps` input is a comma-separated list of URLs that point to dependency files. Supported filetypes = `.tgz`, `.tar.gz`, `.txz` and `.tar.xz`.

The action will download and extract the files to `$HOME` (`/github/home`), so that they can be used during any following steps.
