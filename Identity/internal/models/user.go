/*
 * @File: models.user.go
 * @Description: Defines User model
 */
package models

import (
	"errors"
	"github.com/yudafatah/golang-microservices-boilerplate/tree/main/Identity/common"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// User information
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	LastName string        `bson:"last_name" json:"last_name" example:"fatah"`
	FirstName string        `bson:"first_name" json:"first_name" example:"yuda"`
	Username string `bson:"username" json:"username" example:"yudafatah"`
	CreatedTime time.Time `bson:"created_time" json:"created_time" example:"2020-12-31T17:00:00.000+00:00"`
	UpdatedTime time.Time `bson:"updated_time" json:"updated_time" example:"2020-12-31T17:00:00.000+00:00"`
	OldHashPassword string `bson:"old_hash_password" json:"old_hash example:"sdfsdf3f4f4"`
	HashPassword string `bson:"hash_password" json:"hash_password" example:"sfsdfsdfsd7fds8f"`
	Email string `bson:"email" json:"email" example:"yudafatah@gmail.com"`
	Address string `bson:"address" json:"address" example:"california"`
	About string `bson:"about" json:"about" example:"this is about"`
	PhoneNumber string `bson:"phone_number" json:"phone_number" example:"+628888"`
	Country string `bson:"country" json:"country" example:"Indonesia"`
	IsActive  bool `bson:"is_active" json:"is_active" example:"true"`
	IsDeleted bool `bson:"is_deleted" json:"is_deleted" example:"false"`
	ProfilePicture string `bson:"profile_picture" json:"profile_picture" example:"img.com/img.png"`
	ProfileCoverPicture string `bson:"profile_cover_picture" json:"profile_cover_picture" example:"img.com/img.png"`
}

// AddUser information
type AddUser struct {
	Name     string `json:"name" example:"User Name"`
	Password string `json:"password" example:"User Password"`
}

// Validate user
func (a AddUser) Validate() error {
	switch {
	case len(a.Name) == 0:
		return errors.New(common.ErrNameEmpty)
	case len(a.Password) == 0:
		return errors.New(common.ErrPasswordEmpty)
	default:
		return nil
	}
}