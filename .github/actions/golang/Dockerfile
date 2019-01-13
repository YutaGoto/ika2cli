FROM golang:1.11.4

LABEL "name"="Go lang workflow" \
      "maintainer"="YutaGoto <you.goto.510@gmail.com>" \
      "com.github.actions.icon"="code" \
      "com.github.actions.color"="green-dark" \
      "com.github.actions.name"="golang workflow" \
      "com.github.actions.description"="This is an Action to run go and golang ci-lint commands."

ENV LINT_VERSION="v1.12.5"
COPY entrypoint.sh /entrypoint.sh
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b "$(go env GOPATH)"/bin ${LINT_VERSION} \
  && chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
