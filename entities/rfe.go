package entities

import (
	"time"
)

type Moderator struct {
	moderatorId int32
}
type Customer struct {
	customerId int32
	email string
}

type RFE struct {
	RfeId int32
	createdTimestamp	time.Time
	submitter Customer
	moderatedBy Moderator
	product string
	moderated bool
	idea string
	Description string
	status string
	votes int32
}

type Voting struct {
	rfeId int32
	customerID int32

}

/*

(RFE) STATUS
- Accepted (Moderated)
- Rejected
- Pending
- Working On
- Delivered

 */

 /*

 Products
 - SmartShop
 - Chop Chop
 - GOL
  */

  /*

  Subcat

  SmartShop App
  SmartShop Toolbox

   */