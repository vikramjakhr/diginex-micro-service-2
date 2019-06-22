# Diginext Micro Service 2

## QuickStart

```bash
$ helm install ./diginex-micro-service-2 --name app2 --namespace dev
```

## Introduction

This chart bootstraps diginex-micro-service-2 deployment and service on a Kubernetes cluster using the Helm Package manager.

## Prerequisites

- Kubernetes 1.4+

## Installing the Chart

To install the chart with the release name `app2`:

```bash
$ helm install --name app2 ./diginex-micro-service-2
```

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `app2` deployment:

```bash
$ helm delete app2 --purge
```

The command removes all the Kubernetes components associated with the chart and deletes the release.