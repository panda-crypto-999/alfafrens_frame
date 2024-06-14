package mysql

import (
	"context"
	"fmt"
	"testing"

	"github.com/panda-crypto-999/alfafrens_frame/common/utils"
)

func TestCreateChannel(t *testing.T) {
	err := NewChannelDao().Create(context.Background(), &Channel{
		ChannelID: "123",
		Address:   "0x",
	})
	fmt.Println(err)
}

func TestSelectChannel(t *testing.T) {
	id := int64(1)
	channels, err := NewChannelDao().Select(context.Background(), &ChannelWhereParams{
		ID: &id,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(utils.EncodeIgnoreErr(channels))
}

func TestUpdateChannel(t *testing.T) {

}
