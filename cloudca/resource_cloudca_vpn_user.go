package cloudca

import (
	"fmt"

	"github.com/cloud-ca/go-cloudca"
	"github.com/cloud-ca/go-cloudca/services/cloudca"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCloudcaVpnUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudcaVpnUserCreate,
		Read:   resourceCloudcaVpnUserRead,
		Delete: resourceCloudcaVpnUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the environment where the vpn should be created",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Username of the VPN user",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Password of the VPN user",
			},
		},
	}
}

func resourceCloudcaVpnUserCreate(d *schema.ResourceData, meta interface{}) error {
	ccaResources, rerr := getResourcesForEnvironmentID(meta.(*cca.CcaClient), d.Get("environment_id").(string))
	if rerr != nil {
		return rerr
	}

	remoteAccessVpnUser := cloudca.RemoteAccessVpnUser{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}
	_, err := ccaResources.RemoteAccessVpnUser.Create(remoteAccessVpnUser)
	if err != nil {
		return fmt.Errorf("Error adding VPN user: %s", err)
	}

	// TODO: When the CMC API actually returns the ID of the created user, use it.
	// Currently there is no way to do a 'Get' based on the username, and we don't have the ID, so
	// we have to list all users and then loop through to match the username in order to find the ID.
	vpnUsers, err := ccaResources.RemoteAccessVpnUser.List()
	if err != nil {
		return fmt.Errorf("Error getting the created VPN user ID: %s", err)
	}
	var userID string
	for _, user := range vpnUsers {
		if user.Username == d.Get("username").(string) {
			userID = user.Id
		}
	}
	if userID != "" {
		d.SetId(userID)
	} else {
		return fmt.Errorf("Error finding the created VPN user ID: %s", err)
	}
	return resourceCloudcaVpnUserRead(d, meta)
}

func resourceCloudcaVpnUserRead(d *schema.ResourceData, meta interface{}) error {
	ccaResources, rerr := getResourcesForEnvironmentID(meta.(*cca.CcaClient), d.Get("environment_id").(string))
	if rerr != nil {
		return rerr
	}

	// Get the user based on the ID
	vpnUser, err := ccaResources.RemoteAccessVpnUser.Get(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	if err := d.Set("username", vpnUser.Username); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}
	return nil
}

func resourceCloudcaVpnUserDelete(d *schema.ResourceData, meta interface{}) error {
	ccaResources, rerr := getResourcesForEnvironmentID(meta.(*cca.CcaClient), d.Get("environment_id").(string))
	if rerr != nil {
		return rerr
	}
	remoteAccessVpnUser := cloudca.RemoteAccessVpnUser{
		Id:       d.Id(),
		Username: d.Get("username").(string),
	}
	if _, err := ccaResources.RemoteAccessVpnUser.Delete(remoteAccessVpnUser); err != nil {
		return handleNotFoundError("VPN UserDelete", true, err, d)
	}
	return nil
}
