/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateInfoOrder struct {
	Key
	Attributes CreateInfoOrderAttributes `json:"attributes"`
}
type CreateInfoOrderResponse struct {
	Data     CreateInfoOrder `json:"data"`
	Included Included        `json:"included"`
}

type CreateInfoOrderListResponse struct {
	Data     []CreateInfoOrder `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustCreateInfoOrder - returns CreateInfoOrder from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateInfoOrder(key Key) *CreateInfoOrder {
	var createInfoOrder CreateInfoOrder
	if c.tryFindEntry(key, &createInfoOrder) {
		return &createInfoOrder
	}
	return nil
}
