package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/scyna/go/example/contacts/proto"
	"github.com/scyna/go/scyna"
)

func ListFriend(c *scyna.Service, request *proto.ListFriendRequest) {
	c.LOG.Info("Receive ListFriendRequest")

	if validation.Validate(request.Email, validation.Required, is.Email) != nil {
		c.Error(scyna.REQUEST_INVALID)
		return
	}

	if err, user := Repository.GetByEmail(c.LOG, request.Email); err != nil {
		c.Error(USER_NOT_EXISTED)
	} else {
		if err, users := Repository.ListFriend(c.LOG, user.ID); err != nil {
			c.Error(err)
		} else {
			result := make([]*proto.User, len(users))
			for i, u := range users {
				result[i] = u.ToDTO()
			}
			c.Done(&proto.ListFriendResponse{
				Friends: result,
			})
		}
	}
}
