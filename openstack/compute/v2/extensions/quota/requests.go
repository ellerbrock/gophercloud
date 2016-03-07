package quotas

import (
	"github.com/rackspace/gophercloud"
)

// Get returns public data about a previously created Quota.
func Get(client *gophercloud.ServiceClient, tenantID string) GetResult {
	var res GetResult
	_, res.Err = client.Get(getURL(client, tenantID), &res.Body, nil)
	return res
}
