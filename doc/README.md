# cloud.ca Provider

The cloud.ca provider is used to interact with the many resources supported by [cloud.ca](https://cloud.ca/). The provider needs to be configured with the proper credentials before it can be used. Optionally with a URL pointing to a running cloud.ca API.

In order to provide the required configuration options you need to supply the value for `api_key` field.

## Example Usage

```hcl
variable "my_api_key" {}

# Configure cloud.ca Provider
provider "cloudca" {
    api_key = "${var.my_api_key}"
}

# Create an Instance
resource "frontier_instance" "instance" {
    # ...
}
```

## Argument Reference

The following arguments are supported:

- [api_key](#api_key) - (Required) This is the cloud.ca API key. It can also be sourced from the `frontier_API_KEY` environment variable.
- [api_url](#api_url) - (Optional) This is the cloud.ca API URL. It can also be sourced from the `frontier_API_URL` environment variable.

## Resources

- [**frontier_environment**](environment.md)
- [**frontier_instance**](instance.md)
- [**frontier_load_balancer_rule**](load_balancer_rule.md)
- [**frontier_network**](network.md)
- [**frontier_network_acl**](network_acl.md)
- [**frontier_network_acl_rule**](network_acl_rule.md)
- [**frontier_port_forwarding_rule**](port_forwarding_rule.md)
- [**frontier_public_ip**](public_ip.md)
- [**frontier_static_nat**](static_nat.md)
- [**frontier_ssh_key**](ssh_key.md)
- [**frontier_volume**](volume.md)
- [**frontier_vpc**](vpc.md)
