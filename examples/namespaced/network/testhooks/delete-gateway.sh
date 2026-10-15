#!/usr/bin/env bash
set -aeuo pipefail

# Delete the gateway chain, waiting for each kind, before the Routers and Networks the gateways are attached to are deleted.
${KUBECTL} delete gatewayconnectiontunnel.network.m.upcloud.com --all --all-namespaces
${KUBECTL} delete gatewayconnection.network.m.upcloud.com --all --all-namespaces
${KUBECTL} delete gateway.network.m.upcloud.com --all --all-namespaces
