package mysql

import (
	"context"
	"strings"

	"gorm.io/gorm"
)

type channelDao struct {
	db *gorm.DB
}

func NewChannelDao() *channelDao {
	return &channelDao{
		db: db.Table("channel"),
	}
}

type Channel struct {
	ID        int64  `json:"id"`
	ChannelID string `json:"channel_id"`
	Address   string `json:"address"`
}

func (dao *channelDao) Create(ctx context.Context, obj *Channel) error {
	err := dao.db.Create(obj).Error
	return err
}

func (dao *channelDao) Update(ctx context.Context, obj *Channel) error {
	err := dao.db.Where("id=?", obj.ID).Updates(obj).Error
	return err
}

func (dao *channelDao) Select(ctx context.Context, where *ChannelWhereParams) ([]*Channel, error) {
	offset, limit := where.Offset, where.Limit
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 1000
	}
	channels := make([]*Channel, 0)
	conditions, values := where.buildExpression()
	err := dao.db.Where(conditions, values...).Offset(offset).Limit(limit).Find(&channels).Error // ignore_security_alert
	return channels, err
}

type ChannelWhereParams struct {
	ID *int64

	Offset int
	Limit  int
}

func (params *ChannelWhereParams) buildExpression() (string, []interface{}) {
	// build
	conditions, values := make([]string, 0), make([]interface{}, 0)
	if params.ID != nil {
		conditions = append(conditions, "id = ?")
		values = append(values, params.ID)
	}
	return strings.Join(conditions, " and "), values
}
