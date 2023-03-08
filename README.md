[![GoodFirstIssue](docs/goodfirstissue.png)](https://github.com/rajatjindal/goodfirstissue) 

[![Twitter URL](https://img.shields.io/twitter/follow/goodfirstissue.svg?label=Follow&style=social)](https://twitter.com/goodfirstissue) [![good first issues](https://img.shields.io/github/issues/rajatjindal/goodfirstissue/good%20first%20issue.svg
)](https://github.com/rajatjindal/goodfirstissue/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) 

This is a tool listening to webhook events from [goodfirstissue](https://github.com/apps/goodfirstissue) github app which configures a webhook for listening to `issues` github-event. 

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

Orgs/Users who have enabled `goodfirstissue` for atleast one repository (generated using [github-app-installations](https://github.com/rajatjindal/github-app-installations))


<a href="https://github.com/apache"><img src="https://github.com/apache.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/google"><img src="https://github.com/google.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/layer5io"><img src="https://github.com/layer5io.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/openfaas"><img src="https://github.com/openfaas.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/openfaas-incubator"><img src="https://github.com/openfaas-incubator.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/alexellis"><img src="https://github.com/alexellis.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/tektoncd"><img src="https://github.com/tektoncd.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/linkerd"><img src="https://github.com/linkerd.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/adobe"><img src="https://github.com/adobe.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/inlets"><img src="https://github.com/inlets.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/fastify"><img src="https://github.com/fastify.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/open-sauced"><img src="https://github.com/open-sauced.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/helm"><img src="https://github.com/helm.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/parthpvaghani"><img src="https://github.com/parthpvaghani.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/melsorrells23"><img src="https://github.com/melsorrells23.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/ntedgi"><img src="https://github.com/ntedgi.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/HectorGalindoPedraza"><img src="https://github.com/HectorGalindoPedraza.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/naman-tiwari"><img src="https://github.com/naman-tiwari.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/PropaneC3H8"><img src="https://github.com/PropaneC3H8.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/nifty-toolbox"><img src="https://github.com/nifty-toolbox.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/Project-Universe"><img src="https://github.com/Project-Universe.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/j-dogcoder"><img src="https://github.com/j-dogcoder.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/Heptagram-Project"><img src="https://github.com/Heptagram-Project.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/VictoryWekwa"><img src="https://github.com/VictoryWekwa.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/service-mesh-patterns"><img src="https://github.com/service-mesh-patterns.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/service-mesh-performance"><img src="https://github.com/service-mesh-performance.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/mui"><img src="https://github.com/mui.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/hantsy"><img src="https://github.com/hantsy.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/akulsr0"><img src="https://github.com/akulsr0.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/Retenodus"><img src="https://github.com/Retenodus.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/yanyu95"><img src="https://github.com/yanyu95.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/HospitalRun"><img src="https://github.com/HospitalRun.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/mayonaze2g"><img src="https://github.com/mayonaze2g.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/kbudde"><img src="https://github.com/kbudde.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/Satendra124"><img src="https://github.com/Satendra124.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/amitech"><img src="https://github.com/amitech.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/meshery"><img src="https://github.com/meshery.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/light-bringer"><img src="https://github.com/light-bringer.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/Sambalicious"><img src="https://github.com/Sambalicious.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/fairlearn"><img src="https://github.com/fairlearn.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/developerfred"><img src="https://github.com/developerfred.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/React95"><img src="https://github.com/React95.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/asyncapi"><img src="https://github.com/asyncapi.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/zuzakistan"><img src="https://github.com/zuzakistan.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/RustScan"><img src="https://github.com/RustScan.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/govdirectory"><img src="https://github.com/govdirectory.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/eps1lon"><img src="https://github.com/eps1lon.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/kgashok"><img src="https://github.com/kgashok.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/pmlopes"><img src="https://github.com/pmlopes.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/KenanBek"><img src="https://github.com/KenanBek.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/reactiverse"><img src="https://github.com/reactiverse.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/wyattowalsh"><img src="https://github.com/wyattowalsh.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/iusehooks"><img src="https://github.com/iusehooks.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/sakuli"><img src="https://github.com/sakuli.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/nut-tree"><img src="https://github.com/nut-tree.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/carsonoid"><img src="https://github.com/carsonoid.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/citrusframework"><img src="https://github.com/citrusframework.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/milvus-io"><img src="https://github.com/milvus-io.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/rajatjindal"><img src="https://github.com/rajatjindal.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/lucbpz"><img src="https://github.com/lucbpz.png" width="100"></a><span width="10px">&nbsp;</span>
<a href="https://github.com/Ewocker"><img src="https://github.com/Ewocker.png" width="100"></a><span width="10px">&nbsp;</span>

# Acknowledgements

Many thanks to [Alex Ellis](https://twitter.com/alexellisuk) for helping me write, build, and host this as openfaas-function on [openfaas-cloud](https://github.com/openfaas/openfaas-cloud) for more than four years without costing a dime to me. I will be thankful forever for that.
