### Security
  - [dependabot-alerts](https://docs.github.com/en/code-security/dependabot/dependabot-alerts/configuring-dependabot-alerts)

### Versioning & Release
  - [semver.org](https://semver.org/)
  - [semantic-release](https://github.com/semantic-release/semantic-release) ([discuss](https://github.com/semantic-release/semantic-release/discussions))  ([gitbook](https://semantic-release.gitbook.io/semantic-release/))([example](https://github.com/GeekHomeInside/semantic-release-example))

### Commit
  - [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)
  - [commit message format](https://github.com/angular/angular/blob/main/CONTRIBUTING.md#-commit-message-format)
  - [the perfect commit](https://simonwillison.net/2022/Oct/29/the-perfect-commit/)
  - [commit lint](https://github.com/conventional-changelog/commitlint)
  - [commitizen CLI](https://github.com/commitizen/cz-cli)

### Changelog
  - [keepachangelog.com](https://keepachangelog.com/en/1.1.0/)
  - [changelog.com](https://changelog.com/)
  - [CHANGELOG.md](https://gist.github.com/juampynr/4c18214a8eb554084e21d6e288a18a2c)
  - [generate-changelog](https://github.com/lob/generate-changelog)
  - [git-chglog](https://github.com/git-chglog/git-chglog)

### Lint
  - [golangci-lint](https://github.com/golangci/golangci-lint-action) ([doc](https://golangci-lint.run/)) ([git-action](https://github.com/marketplace/actions/run-golangci-lint))

```bash
# Check
> golangci-lint run
> golangci-lint run --enable-all

# Fix linting
> go fmt ./...
> gofmt -s -w -l .
> gofumpt -l -w .
```
