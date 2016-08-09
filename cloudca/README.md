#Resources

##cloudca_vpc
Create a vpc.

###Example usage
```
resource "cloudca_vpc" "my_instance" {
	service_code = "compute-east"
	environment_name = "dev"
	name = "test-vpc"
	description = "This is a test vpc"
	vpc_id = "8b46e2d1-bbc4-4fad-b3bd-1b25fcba4cec"
	network_offering = "Standard Tier"
	network_acl = "default_allow"
}
```
###Argument Reference
The following arguments are supported:
- service_code - (Required) Service code
- environment_name - (Required) Name of environment
- name - (Required) Name of the VPC
- description - (Required) Description of the VPC
- vpc_offering - (Required) The name of the VPC offering to use for the vpc
- network_domain - (Optional) A custom DNS suffix at the level of a network

###Attribute Reference
- id - ID of VPC.

##cloudca_tier
Create a tier.

###Example usage
```
resource "cloudca_tier" "my_instance" {
	service_code = "compute-east"
	environment_name = "dev"
	name = "test-tier"
	description = "This is a test tier"
	vpc_id = "8b46e2d1-bbc4-4fad-b3bd-1b25fcba4cec"
	network_offering = "Standard Tier"
	network_acl = "default_allow"
}
```
###Argument Reference
The following arguments are supported:
- service_code - (Required) Service code
- environment_name - (Required) Name of environment
- name - (Required) Name of the tier
- description - (Required) Description of the tier
- vpc_id - (Required) The ID of the vpc where the tier should be created
- network_offering - (Required) The name of the network offering to use for the tier
- network_acl - (Required) The name of the network ACL to use for the tier

###Attribute Reference
- id - ID of tier.

##cloudca_instance
Create and starts an instance.

###Example usage
```
resource "cloudca_instance" "my_instance" {
	service_code = "compute-east"
	environment_name = "dev"
	name = "test-instance"
	network_id = "672016ef-05ee-4e88-b68f-ac9cc462300b"
	template = "CentOS 6.7 base (64bit)"
	compute_offering = "1vCPU.512MB"
	ssh_key_name = "my_ssh_key"
}
```
###Argument Reference
The following arguments are supported:
- service_code - (Required) Service code
- environment_name - (Required) Name of environment
- name - (Required) Name of instance
- network_id - (Required) The ID of the network where the instance should be created
- template - (Required) Name of template to use for the instance
- compute_offering - (Required) Name of the compute offering to use for the instance
- user_date - (Optional) User data to add to the instance
- ssh_key_name - (Optional) Name of the SSH key pair to attach to the instance. Mutually exclusive with public_key.
- public_key - (Optional) Public key to attach to the instance. Mutually exclusive with ssh_key_name.
- purge - (Optional) If true, then it will purge the instance on destruction

###Attribute Reference
- id - ID of instance.

##cloudca_publicip
Acquires a public IP in a specific VPC. If you update any of the fields in the resource, then it will release this IP and recreate it.

###Example usage
```
resource "cloudca_publicip" "my_publicip" {
	service_code = "compute-east"
	environment_name = "dev"
	vpc_id = "8b46e2d1-bbc4-4fad-b3bd-1b25fcba4cec"
}
```
###Argument Reference
The following arguments are supported:
- service_code - (Required) Service code
- environment_name - (Required) Name of environment
- vpc_id - (Required) The ID of the vpc to acquire the public IP

###Attribute Reference
- id - The public IP ID.
- ip_address - The public IP address