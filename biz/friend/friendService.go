package friend

import "github.com/gin-gonic/gin"

type FriendService struct {
	Data *Friend
}

func (srv *FriendService) GetFriendList(c *gin.Context) error {
	uid := srv.Data.Uid
	fList, err := srv.Data.GetFriend(uid)
	if err != nil {
		return err
	}
	srv.Data = fList
	return nil
}

func (srv *FriendService) AddFriend(c *gin.Context) error {
	return nil
}

func (srv *FriendService) RemoveFriend(c *gin.Context) error {
	return nil
}

func (srv *FriendService) BlockFriend(c *gin.Context) error {
	return nil
}
