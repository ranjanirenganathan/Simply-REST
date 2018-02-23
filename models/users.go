package models


type User struct {

	DocumentBase `json:",inline" bson:",inline"`
	Name    	    string   		`json:"name"  bson:"name"`
	Status  	    string   		`json:"status"  bson:"status"`
	Description     string   		`json:"description"  bson:"description"`
	Address         Address 		`json:"address" bson:"address"`
}


type Address struct {
	Name    	    string   		`json:"name"    bson:"name"`
	Email    	    string   		`json:"email"   bson:"email"`
	Phone   	    string   		`json:"phone"   bson:"phone"`
	Fax   		    string   		`json:"fax"     bson:"fax"`
	Address 	    string 		    `json:"address" bson:"address"`
	City   		    string 			`json:"city"    bson:"city"`
	State           string		    `json:"state"   bson:"state"`
	PostalCode   	string    	    `json:"postalCode" bson:"postalCode"`
	Country		    string			`json:"country" bson:"country"`
}




