# README

**BEFORE YOU START**, please be aware that there are more ways to integrate with your service that don't require creating a service from this template, see https://keptn.sh/docs/0.10.x/integrations/how_integrate/ for more details.

Examples:

* Webhooks: https://keptn.sh/docs/0.10.x/integrations/webhooks/
* Job-Executor-Service: https://github.com/keptn-sandbox/job-executor-service

---

This is a Keptn Service Template written in GoLang. Follow the instructions below for writing your own Keptn integration.

Quick start:

1. In case you want to contribute your service to keptn-sandbox or keptn-contrib, make sure you have read and understood the [Contributing Guidelines](https://github.com/keptn-sandbox/contributing).
1. Click [Use this template](https://github.com/raffy23/automatic-octo-telegram/generate) on top of the repository, or download the repo as a zip-file, extract it into a new folder named after the service you want to create (e.g., simple-service) 
1. Run GitHub workflow `One-time repository initialization` to tailor deployment files and go modules to the new instance of the keptn service template. This will create a Pull Request containing the necessary changes, review it, adjust if necessary and merge it.
1. Figure out whether your Kubernetes Deployment requires [any RBAC rules or a different service-account](https://github.com/keptn-sandbox/contributing#rbac-guidelines), and adapt [chart/templates/serviceaccount.yaml](chart/templates/serviceaccount.yaml) accordingly for the roles.
1. Last but not least: Remove this intro within the README file and make sure the README file properly states what this repository is about

---

# automatic-octo-telegram
![GitHub release (latest by date)](https://img.shields.io/github/v/release/raffy23/automatic-octo-telegram)
[![Go Report Card](https://goreportcard.com/badge/github.com/raffy23/automatic-octo-telegram)](https://goreportcard.com/report/github.com/raffy23/automatic-octo-telegram)

This implements a automatic-octo-telegram for Keptn. If you want to learn more about Keptn visit us on [keptn.sh](https://keptn.sh)

## Compatibility Matrix

*Please fill in your versions accordingly*

| Keptn Version    | [Keptn-Service-Template-Go Docker Image](https://hub.docker.com/r/raffy23/automatic-octo-telegram/tags) |
|:----------------:|:----------------------------------------:|
|       0.6.1      | raffy23/automatic-octo-telegram:0.1.0 |
|       0.7.1      | raffy23/automatic-octo-telegram:0.1.1 |
|       0.7.2      | raffy23/automatic-octo-telegram:0.1.2 |

## Installation

The *automatic-octo-telegram* can be installed as a part of [Keptn's uniform](https://keptn.sh).

### Deploy in your Kubernetes cluster

To deploy the current version of the *automatic-octo-telegram* in your Keptn Kubernetes cluster use the [`helm chart`](chart/Chart.yaml) file,
for example:

```console
helm install -n keptn automatic-octo-telegram chart/
```

This should install the `automatic-octo-telegram` together with a Keptn `distributor` into the `keptn` namespace, which you can verify using

```console
kubectl -n keptn get deployment automatic-octo-telegram -o wide
kubectl -n keptn get pods -l run=automatic-octo-telegram
```

### Up- or Downgrading

Adapt and use the following command in case you want to up- or downgrade your installed version (specified by the `$VERSION` placeholder):

```console
helm upgrade -n keptn --set image.tag=$VERSION automatic-octo-telegram chart/
```

### Uninstall

To delete a deployed *automatic-octo-telegram*, use the file `deploy/*.yaml` files from this repository and delete the Kubernetes resources:

```console
helm uninstall -n keptn automatic-octo-telegram
```

## Development

Development can be conducted using any GoLang compatible IDE/editor (e.g., Jetbrains GoLand, VSCode with Go plugins).

It is recommended to make use of branches as follows:

* `master` contains the latest potentially unstable version
* `release-*` contains a stable version of the service (e.g., `release-0.1.0` contains version 0.1.0)
* create a new branch for any changes that you are working on, e.g., `feature/my-cool-stuff` or `bug/overflow`
* once ready, create a pull request from that branch back to the `master` branch

When writing code, it is recommended to follow the coding style suggested by the [Golang community](https://github.com/golang/go/wiki/CodeReviewComments).

### Where to start

If you don't care about the details, your first entrypoint is [eventhandlers.go](eventhandlers.go). Within this file 
 you can add implementation for pre-defined Keptn Cloud events.
 
To better understand all variants of Keptn CloudEvents, please look at the [Keptn Spec](https://github.com/keptn/spec).
 
If you want to get more insights into processing those CloudEvents or even defining your own CloudEvents in code, please 
 look into [main.go](main.go) (specifically `processKeptnCloudEvent`), [chart/values.yaml](chart/values.yaml),
 consult the [Keptn docs](https://keptn.sh/docs/) as well as existing [Keptn Core](https://github.com/keptn/keptn) and
 [Keptn Contrib](https://github.com/keptn-contrib/) services.

### Common tasks

* Build the binary: `go build -ldflags '-linkmode=external' -v -o automatic-octo-telegram`
* Run tests: `go test -race -v ./...`
* Build the docker image: `docker build . -t raffy23/automatic-octo-telegram:dev` (Note: Ensure that you use the correct DockerHub account/organization)
* Run the docker image locally: `docker run --rm -it -p 8080:8080 raffy23/automatic-octo-telegram:dev`
* Push the docker image to DockerHub: `docker push raffy23/automatic-octo-telegram:dev` (Note: Ensure that you use the correct DockerHub account/organization)
* Deploy the service using `kubectl`: `kubectl apply -f deploy/`
* Delete/undeploy the service using `kubectl`: `kubectl delete -f deploy/`
* Watch the deployment using `kubectl`: `kubectl -n keptn get deployment automatic-octo-telegram -o wide`
* Get logs using `kubectl`: `kubectl -n keptn logs deployment/automatic-octo-telegram -f`
* Watch the deployed pods using `kubectl`: `kubectl -n keptn get pods -l run=automatic-octo-telegram`
* Deploy the service using [Skaffold](https://skaffold.dev/): `skaffold run --default-repo=your-docker-registry --tail` (Note: Replace `your-docker-registry` with your container image registry (defaults to ghcr.io/raffy23/automatic-octo-telegram); also make sure to adapt the image name in [skaffold.yaml](skaffold.yaml))


### Testing Cloud Events

We have dummy cloud-events in the form of [RFC 2616](https://ietf.org/rfc/rfc2616.txt) requests in the [test-events/](test-events/) directory. These can be easily executed using third party plugins such as the [Huachao Mao REST Client in VS Code](https://marketplace.visualstudio.com/items?itemName=humao.rest-client).

## Automation

### GitHub Actions: Automated Pull Request Review

This repo uses [reviewdog](https://github.com/reviewdog/reviewdog) for automated reviews of Pull Requests. 

You can find the details in [.github/workflows/reviewdog.yml](.github/workflows/reviewdog.yml).

### GitHub Actions: Unit Tests

This repo has automated unit tests for pull requests. 

You can find the details in [.github/workflows/CI.yml](.github/workflows/CI.yml).

### GH Actions/Workflow: Build Docker Images

This repo uses GH Actions and Workflows to test the code and automatically build docker images.

Docker Images are automatically pushed based on the configuration done in [.ci_env](.ci_env) and the two [GitHub Secrets](https://github.com/raffy23/automatic-octo-telegram/settings/secrets/actions)
* `REGISTRY_USER` - your DockerHub username
* `REGISTRY_PASSWORD` - a DockerHub [access token](https://hub.docker.com/settings/security) (alternatively, your DockerHub password)

## How to release a new version of this service

It is assumed that the current development takes place in the master branch (either via Pull Requests or directly).

To make use of the built-in automation using GH Actions for releasing a new version of this service, you should

* branch away from master to a branch called `release-x.y.z` (where `x.y.z` is your version),
* write release notes in the [releasenotes/](releasenotes/) folder,
* check the output of GH Actions builds for the release branch, 
* verify that your image was built and pushed to DockerHub with the right tags,
* update the image tags in [deploy/service.yaml], and
* test your service against a working Keptn installation.

If any problems occur, fix them in the release branch and test them again.

Once you have confirmed that everything works and your version is ready to go, you should

* create a new release on the release branch using the [GitHub releases page](https://github.com/raffy23/automatic-octo-telegram/releases), and
* merge any changes from the release branch back to the master branch.

## License

Please find more information in the [LICENSE](LICENSE) file.
