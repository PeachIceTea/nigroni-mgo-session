package nigronimgosession

import "github.com/globalsign/mgo"

type contextKey int

// Define keys that support equality.
const KEY contextKey = 0

type NMS struct {
	DB      *mgo.Database
	Session *mgo.Session
}
