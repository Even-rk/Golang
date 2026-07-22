package main

import (
	"fmt"

	"gorm.io/gorm"
)

// 1.Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	// 查当前文章数量
	var user User
	tx.Model(&User{}).Where("id = ?", p.UserID).First(&user)
	// 更新用户的文章数量统计字段 ，增加 1
	tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", user.PostCount+1)
	return nil
}

// 2.为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 查当前评论数量
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	fmt.Println("当前评论数量:", count)
	// 如果评论数量为 0，则更新文章的评论状态为 "无评论", 否则更新为评论数量
	if count == 0 {
		tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论")
	} else {
		tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", count)
	}
	return nil
}
