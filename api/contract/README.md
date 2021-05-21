# API Contract

This directory contains the source YMLs used to drive code generation of HTTP clients in [`api`](../).

## YML Structure

Most YMLs used here are pulled from the source-of-truth [`openapi`](https://github.com/influxdata/openapi) repo via a git
submodule. In rare cases, the full description of an API is too complex for our codegen tooling to handle;
the [`overrides/`](./overrides) directory contains alternate definitions for paths/schemas that work around these cases.
[`cli.yml`](./cli.yml) ties together all the pieces by linking all routes and schemas used by the CLI.

## Updating the API contract

To extend/modify the API contract used by the CLI, first make sure the `openapi` submodule is cloned and up-to-date:
```shell
# Run from the project root.
git submodule update --init --recursive
```

Then create a new branch to track your work:
```shell
git checkout <new-branch-name>
```

Next, decide if any modifications are needed in the source-of-truth `openapi` repo. If so, create a branch in the
submodule to track changes there:
```shell
cd api/contract/openapi && git checkout -b <new-branch-name>
```

Edit/add to the files under `api-contract/` to describe the new API contract. Run the following from the project
root test your changes and see the outputs in Go code:
```shell
make openapi
# Use `git status` to see new/modified files under `api`
```

Once you're happy with the new API contract, submit your changes for review & merge.
If you added/edited files within `openapi`, you'll first need to:
1. Push your submodule branch to GitHub
   ```shell
   cd api/contract/openapi && git push <your-branch-name>
   ```
2. Create a PR in `openapi`, eventually merge to `master` there
3. Update your submodule to point at the merge result:
   ```shell
   cd api/contract/openapi && git fetch && git checkout master && git pull origin master
   ```
4. Update the submodule reference from the main repo:
   ```shell
   git add api/contract/openapi
   git commit
   ```
