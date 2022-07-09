package model

import proto "github.com/scyna/go/manager/.proto/generated"

type Client struct {
	Organization string
	ID           string
	Secret       string
}

func (c *Client) FromDTO(o *proto.Client) {
	/*TODO*/
}

func (c *Client) ToDTO() *proto.Client {
	/*TODO*/
	return nil
}
