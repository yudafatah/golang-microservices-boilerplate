/*
 * @File: daos.user.go
 * @Description: Implements User Auth functions for MongoDB
 */
package daos

import (
	"github.com/yudafatah/golang-microservices-boilerplate/tree/main/Identity/common"
	"github.com/yudafatah/golang-microservices-boilerplate/tree/main/Identity/internal/databases"
	"github.com/yudafatah/golang-microservices-boilerplate/tree/main/Identity/internal/models"
	"github.com/yudafatah/golang-microservices-boilerplate/tree/main/Identity/utils"
	"gopkg.in/mgo.v2/bson"
)

// Auth manages User authentication
type Auth struct {
	utils *utils.Utils
}

// GetAll gets the list of Users
// func (u *Auth) GetAll() ([]models.User, error) {
// 	sessionCopy := databases.Database.MgDbSession.Copy()
// 	defer sessionCopy.Close()

// 	// Get a collection to execute the query against.
// 	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

// 	var users []models.User
// 	err := collection.Find(bson.M{}).All(&users)
// 	return users, err
// }

// GetByUname finds a User by its username
func (u *Auth) GetByUname(uname string) (models.User, error) {
	var err error
	// err = u.utils.ValidateObjectID(id)
	if err != nil {
		return models.User{}, err
	}

	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	var user models.User
	err = collection.Find(bson.M{"username":uname}).One(&user)
	return user, err
}

// DeleteByID finds a User by its id
// func (u *User) DeleteByID(id string) error {
// 	var err error
// 	err = u.utils.ValidateObjectID(id)
// 	if err != nil {
// 		return err
// 	}

// 	sessionCopy := databases.Database.MgDbSession.Copy()
// 	defer sessionCopy.Close()

// 	// Get a collection to execute the query against.
// 	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

// 	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
// 	return err
// }

// Login User
func (u *Auth) Login(name string, password string) (models.User, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	var user models.User

	// conditional query
	conditions := bson.M{
		"password": password,
		 "$or": []bson.M{
			 {"username":name}, 
			 {"email":name},
		 }}

	err := collection.Find(conditions).One(&user)
	return user, err
}

// Insert adds a new User into database'
func (u *Auth) Insert(user models.User) error {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	err := collection.Insert(&user)
	return err
}

// Delete remove an existing User
// func (u *User) Delete(user models.User) error {
// 	sessionCopy := databases.Database.MgDbSession.Copy()
// 	defer sessionCopy.Close()

// 	// Get a collection to execute the query against.
// 	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

// 	err := collection.Remove(&user)
// 	return err
// }

// Update modifies an existing User
// func (u *User) Update(user models.User) error {
// 	sessionCopy := databases.Database.MgDbSession.Copy()
// 	defer sessionCopy.Close()

// 	// Get a collection to execute the query against.
// 	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

// 	err := collection.UpdateId(user.ID, &user)
// 	return err
// }