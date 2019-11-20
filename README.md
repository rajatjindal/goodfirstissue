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
<a href="https://github.com/developerfred"><img src="https://github.com/developerfred.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/helm"><img src="https://github.com/helm.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/asyncapi"><img src="https://github.com/asyncapi.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/inlets"><img src="https://github.com/inlets.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/openfaas-incubator"><img src="https://github.com/openfaas-incubator.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/fastify"><img src="https://github.com/fastify.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/storyscript"><img src="https://github.com/storyscript.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/tektoncd"><img src="https://github.com/tektoncd.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/openfaas"><img src="https://github.com/openfaas.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/zuzakistan"><img src="https://github.com/zuzakistan.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/apache"><img src="https://github.com/apache.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/sakuli"><img src="https://github.com/sakuli.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/rajatjindal"><img src="https://github.com/rajatjindal.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/google"><img src="https://github.com/google.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/jetstack"><img src="https://github.com/jetstack.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/pmlopes"><img src="https://github.com/pmlopes.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/reactiverse"><img src="https://github.com/reactiverse.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/citrusframework"><img src="https://github.com/citrusframework.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/Ewocker"><img src="https://github.com/Ewocker.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/nut-tree"><img src="https://github.com/nut-tree.png" width="100"></a><span width="10px">&nbsp;</span>
* Connection #0 to host rajatjindal.o6s.io left intact
<a href="https://github.com/alexellis"><img src="https://github.com/alexellis.png" width="100"></a><span width="10px">&nbsp;</span>%


# Acknowledgements

Many thanks to [Alex Ellis](https://twitter.com/alexellisuk) for helping me write, build and deploy this openfaas-function to [openfaas-cloud](https://github.com/openfaas/openfaas-cloud).
