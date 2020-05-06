/*
 * Types for service.
 * Wang Ruichao (793160615@qq.com)
 */

package srv

// roomSt is the structure of room.
type roomSt struct {
	RID     string `json:"id" bson:"_id"`
	Token   string `json:"token" bson:"token"`
	Passwd  string `json:"passwd" bson:"passwd"`
	Ct      int64  `json:"ct" bson:"ct"`
	UserCnt int64  `json:"user_cnt" bson:"user_cnt"`
}

// userSt is the structure of user.
type userSt struct {
	UID     string `json:"id" bson:"_id"`
	Usrname string `json:"username" bson:"username"`
	Passwd  string `json:"passwd" bson:"passwd"`
	Ct      int64  `json:"ct" bson:"ct"`
	Status  string `json:"status" bson:"status"`
	Lit     int64  `json:"lit" bson:"lit"` // Login time
	Lot     int64  `json:"lot" bson:"lot"` // Logout time
}
