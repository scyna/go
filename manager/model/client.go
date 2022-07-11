package model

import proto "github.com/scyna/go/manager/.proto/generated"

type Client struct {
	Organization string `db:"org_code"`
	ID           string `db:"id"`
	Secret       string `db:"secret"`
}

func (c *Client) FromDTO(o *proto.Client) {
	c.Secret = o.Secret
	c.Organization = o.Organization
	c.ID = o.Id
}

func (c *Client) ToDTO() *proto.Client {
	return &proto.Client{
		Organization: c.Organization,
		Id:           c.ID,
		Secret:       "", // NO RETURN
	}
}
