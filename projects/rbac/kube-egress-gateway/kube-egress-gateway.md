
kube-egress-gateway  provides a scalable and cost-efficient way to configure fixed source IP for Kubernetes pod egress traffic on Azure.
kube-egress-gateway components run in kubernetes clusters, either managed (Azure Kubernetes Service, AKS) or unmanaged, utilize one or more dedicated kubernetes nodes as pod egress gateways and route pod outbound traffic to gateway via wireguard tunnel.

Compared with existing methods, for example, creating dedicated kubernetes nodes with NAT gateway or instance level public ip address and only scheduling pods with such requirement on these nodes, kube-egress-gateway provides a more cost-efficient method as pods requiring different egress IPs can share the same gateway and can be scheduled on any regular worker node. 

## Usage

### Operating system requirements

* Only Linux (Ubuntu or Azure Linux) based nodes are supported. Other Linux distributions have not been tested but may work. Windows node or pod support is not available at this time.

### Deploy a Static Egress Gateway

To deploy a static egress gateway, you need to create a StaticGatewayConfiguration CR:
```yaml
apiVersion: egressgateway.kubernetes.azure.com/v1alpha1
kind: StaticGatewayConfiguration
metadata:
  name: myStaticEgressGateway
  namespace: myNamespace
spec:
  gatewayVmssProfile:
    vmssResourceGroup: myResourceGroup
    vmssName: myGatewayVMSS
    publicIpPrefixSize: 31
  provisionPublicIps: true
  publicIpPrefixId: /subscriptions/mySubscriptionID/resourcegroups/myResourceGroup/providers/Microsoft.Network/publicipprefixes/myPIPPrefix
  defaultRoute: staticEgressGateway
  excludeCidrs:
    - 10.244.0.0/16
    - 10.245.0.0/16
```
StaticGatewayConfiguration is a namespaced resource, meaning a static egress gateway can only be used by pods in the same namespace. There are two **required** configurations: 

* `gatewayVmssProfile`: gateway vmss information:
  * `vmssName`: Name of the Azure VirtualMachineScaleSet (VMSS) to be used as gateway nodepool.
  * `vmssResourceGroup`: Azure resource group of gateway VMSS.
  * `publicIpPrefixSize`: Length of the public IP prefix to be installed on the gateway nodepool as egress. In above example, 31 means a `/31` pip prefix, which contains 2 public IPs, will be installed. Note that gateway VMSS instance count cannot exceed this size. The gateway VMSS can only have 1 or 2 instances if public IP prefix size is 31. Likewise, at most 4 instances are allowed if public IP prefix size is 30. Otherwise, kube-egress-gateway operator will report error. At the time of writing, [Azure](https://learn.microsoft.com/en-us/azure/virtual-network/ip-services/public-ip-address-prefix#prefix-sizes) only supports prefix sizes `/28-/31`.
* `provisionPublicIps`: true if egress gateway needs Internet access. A public IP prefix will be associated with the gateway VMSS secondary IPConfiguration.

Three **optional** configurations:
* `publicIpPrefixId`: BYO public IP prefix is supported. Users can provide Azure resource ID of their own public IP prefix in this field. Make sure kube-egress-gateway operator has access to the prefix. If not provided, a system generated prefix will be provisioned. `provisionPublicIps` must be true.
* `defaultRoute`: Enum, either `staticEgressGateway` or `azureNetworking`. Set it to be `staticEgressGateway` if traffic by default should be routed to the egress gateway or `azureNetworking` if traffic should be routed to pods' `eth0` by default like regular pods. Default value is `staticEgressGateway`.
* `excludeCidrs`: List of destination network CIDRs that should bypass the default route and flow via the other network interface. That is, if `defaultRoute` is `staticEgressGateway`, cidrs set in `excludeCidrs` will be routed via pod's `eth0` interface. For example, traffic within the cluster like pod-pod traffic and pod-service traffic should not be routed to the egress gateway and can be set here. On the other hand, if `defaultRoute` is `azureNetworking`, then only cidrs set in `excludeCidrs` will be routed to the egress gateway.

kube-egress-gateway reconcilers manage the setup and resources and report the egress public IP prefix (private IPs) in `StaticGatewayConfiguration` status:
```yaml
apiVersion: egressgateway.kubernetes.azure.com/v1alpha1
kind: StaticGatewayConfiguration
metadata:
  name: myStaticEgressGateway
  namespace: myNamespace
spec:
  ...
status:
  egressIpPrefix: 1.2.3.4/31 # example public IP prefix output, this will be pods' egress IPNet
```
If `provisionPublicIps` is false, `egressIpPrefix` will be a list of private IPs configured on the corresponding gateway VMSS instance secondary ipConfigurations, e.g. `10.0.1.8,10.0.1.9`.

### Deploy a Pod using Static Egress Gateway

Constructing a pod to use a static egress gateway is simple: just add pod annotation `kubernetes.azure.com/static-gateway-configuration: <StaticGatewayConfiguration name>`. Only name is required here because kube-egress-gateway CNI plugin always assume the gateway is in the same namespace as the pod. Note that existing pods must be recreated to enable egress gateway because CNI plugin can only take effect when pod is being created. 


