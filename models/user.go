package models

// UserExternal Model
//
// The User holds information about permission and other stuff.
//
// swagger:model User
type UserExternal struct {
	// The user id.
	//
	// read only: true
	// required: true
	// example: 25
	ID uint `json:"id"`
	// The user name. For login.
	//
	// required: true
	// example: unicorn
	Name string `binding:"required" json:"name" query:"name" form:"name"`
	// If the user is an administrator.
	//
	// example: true
	Admin bool `json:"admin" form:"admin" query:"admin"`
}

// UserExternalWithPass Model
//
// The UserWithPass holds information about the credentials and other stuff.
//
// swagger:model UserWithPass
type UserExternalWithPass struct {
	UserExternal
	UserExternalPass
}

// UserExternalPass Model
//
// The Password for updating the user.
//
// swagger:model UserPass
type UserExternalPass struct {
	// The user password. For login.
	//
	// required: true
	// example: nrocinu
	Pass string `json:"pass,omitempty" form:"pass" query:"pass" binding:"required"`
}
