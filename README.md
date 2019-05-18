[![GoodFirstIssue](docs/goodfirstissue.png)](https://github.com/rajatjindal/goodfirstissue) 

[![OpenFaaS](https://img.shields.io/badge/openfaas-cloud-blue.svg)](https://www.openfaas.com)  [![Twitter URL](https://img.shields.io/twitter/follow/goodfirstissue.svg?label=Follow&style=social)](https://twitter.com/goodfirstissue) [![good first issues](https://img.shields.io/github/issues/rajatjindal/goodfirstissue/good%20first%20issue.svg
)](https://github.com/rajatjindal/goodfirstissue/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) 

This is a [openfaas](https://www.openfaas.com) function, deployed on [openfaas-cloud](https://github.com/openfaas/openfaas-cloud) running on [Kubernetes](https://github.com/kubernetes/kubernetes), listening to webhook events from [goodfirstissue](https://github.com/apps/goodfirstissue) github app which configures a webhook for listening to `issues` github-event. 

It tweets the link to issue through [@goodfirstissue](https://twitter.com/goodfirstissue) `twitter` handle if:

- The issue has `good first issue` or `good-first-issue` label AND
- if action is one of `opened`, `reopened`, `labeled` or `unassigned`.

The target audience (or followers) for this twitter account are users who are looking forward for starting their journey in `open source contributions` and target audience for this app are github users/org who encourage `first time contributors` to make contributions to their repo(s).

# How to Install

- Go to https://github.com/apps/goodfirstissue
- Click on Configure
- Select the User/Org which owns the repo where you plan to install this app.
- Confirm Password (required by `github`). App don't get access to this password.
- Refer that only `read` access is required to `issues` and `metadata`.
- From `Repository Access` box, select the repositories where you want to enable it. You can enable for `all` or `only selected` repositories.
- Click Save and you are all set.

# Permissions required

The github app needs `read` access to `issues` and `metadata` of the repository. Refer to the screenshot below:

![Permissions](docs/permissions.png)

# Current users

Orgs/Users who have enabled `goodfirstissue` for all/selected repositories (generated using [github-app-installations](https://github.com/rajatjindal/github-app-installations))

| Org/User | Repository |
| ------ | ------ |
| [asyncapi](https://github.com/asyncapi) | [All](https://github.com/asyncapi) |
| [zuzakistan](https://github.com/zuzakistan) | [All](https://github.com/zuzakistan) |
| [tektoncd](https://github.com/tektoncd) | [All](https://github.com/tektoncd) |
| [helm](https://github.com/helm) | [All](https://github.com/helm) |
| [storyscript](https://github.com/storyscript) | [All](https://github.com/storyscript) |
| [openfaas](https://github.com/openfaas) | [All](https://github.com/openfaas) |
| [rajatjindal](https://github.com/rajatjindal) | - [github-app-installations](https://github.com/rajatjindal/github-app-installations)<br/>- [goodfirstissue](https://github.com/rajatjindal/goodfirstissue) |
| [citrusframework](https://github.com/citrusframework) | - [citrus](https://github.com/citrusframework/citrus)<br/>- [citrus-db](https://github.com/citrusframework/citrus-db)<br/>- [citrus-simulator](https://github.com/citrusframework/citrus-simulator) |
| [reactiverse](https://github.com/reactiverse) | [es4x](https://github.com/reactiverse/es4x) |
| [pmlopes](https://github.com/pmlopes) | [vertx-starter](https://github.com/pmlopes/vertx-starter) |
| [Ewocker](https://github.com/Ewocker) | [vue-lodash](https://github.com/Ewocker/vue-lodash) |
| [google](https://github.com/google) | [go-github](https://github.com/google/go-github) |
| [jetstack](https://github.com/jetstack) | [cert-manager](https://github.com/jetstack/cert-manager) |
| [alexellis](https://github.com/alexellis) | - [derek](https://github.com/alexellis/derek)<br/>- [github-exporter](https://github.com/alexellis/github-exporter)<br/>- [inlets](https://github.com/alexellis/inlets)<br/>- [minikube-operator](https://github.com/alexellis/minikube-operator)<br/>- [ubiquitous-octo-guacamole](https://github.com/alexellis/ubiquitous-octo-guacamole) |


# Acknowledgements

Many thanks to [Alex Ellis](https://twitter.com/alexellisuk) for helping me write, build and deploy this openfaas-function to [openfaas-cloud](https://github.com/openfaas/openfaas-cloud).
