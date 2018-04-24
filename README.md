# SMNetwork

[![wercker status](https://app.wercker.com/status/79e99a29dee6182bbb880fc88b1b9e88/s/master "wercker status")](https://app.wercker.com/project/byKey/79e99a29dee6182bbb880fc88b1b9e88)

## Prereqs
 - Have [docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/) installed
 - Have [docker-compose](https://docs.docker.com/compose/install/) installed
 - Have [Wercker CLI](https://devcenter.wercker.com/docs/cli/installation) installed
 - Have [Terraform](https://www.terraform.io/intro/getting-started/install.html) installed
 - Have [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) installed

## Setup
### Install terraform-provider-oci
- [Download Plugin](https://github.com/oracle/terraform-provider-oci/releases)
- Unpack and move to ~/.terraform.d/plugins/
### Download terraform terraform-kubernetes-installer
- Clone repo
- Rename tfvars file
```
$ mv terraform.example.tfvars terraform.tfvars
```
- find tenancy, user and compartment ocid as well as provide public pem file
- instructions on how to create [pem file](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm)
- create env file
```
$ vi env
export TF_VAR_tenancy_ocid="<oci-tenancy-ocid>"
export TF_VAR_user_ocid="<oci-user-ocid>"
export TF_VAR_compartment_ocid="<oci-compartment-ocid>"
export TF_VAR_fingerprint="<fingerprint>"
export TF_VAR_region="us-phoenix-1"
export TF_VAR_private_key_path="<path to private pem file>"
export TF_VAR_ssh_public_key=$(cat <path to public ssh key>)
export TF_VAR_ssh_private_key=$(cat <path to private ssh key>)
export TF_VAR_ssh_public_key_openssh=$(cat <path to private ssh key>)
export DOCKER_ID_USER="<docker-username>"

$ chmod +x env
```
 - initialize your environment variables
 ```
$ . env
```

- move to terraform-kubernetes-installer
- apply terraform plan
```
$ terraform plan
$ terraform apply
```
### Access Kubernetes cluster
- export path to kubeconfig
```
$ export KUBECONFIG=`pwd`/generated/kubeconfig
```
- test access to kubernetes cluster
```
$ kubectl cluster-info
Kubernetes master is running at https://<master-public-ip-address>:443
KubeDNS is running at https://<master-public-ip-address>:443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.

```
### Create Wercker file in your rep
- TODO
### Create Docker hub repo
- TODO
### Create docker-push pipeline in wercker
- TODO
### Create Deploy pipeline in wercker
- input kubenetes authentication info
```
$ cat generated/kubeconfig
```
- locate the certificate-client-data and client-key-data
- copy the 2 data and input into the following commands
```
$ echo "<certificate-client-data>" | base64 -d
$ echo "<client-key-data>" | base64 -d
```
- from each respective output, enter in wercker application for $KUBE_ADMIN and $KUBE_ADMIN_KEY