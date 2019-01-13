workflow "main" {
  on = "push"
  resolves = ["test"]
}

action "test" {
  needs = ["lint"]
  uses = "./.github/actions/golang"
  args = "test"
}

action "lint" {
  uses = "./.github/actions/golang"
  args = "lint"
}
