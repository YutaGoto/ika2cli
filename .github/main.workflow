workflow "main" {
  on = "push"
  resolves = ["lint"]
}

action "lint" {
  uses = "./.github/actions/golang"
  args = "lint"
}
