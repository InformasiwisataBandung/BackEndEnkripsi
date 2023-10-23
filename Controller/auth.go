package Controller

import (
	"fmt"
	"github.com/whatsauth/watoken"
)

var pvtKey = "fbd4a28176db98361b4fb8936e2a1cb499bfe6a3760a3a0726fba735c6edac75513f6d70886d4abd6c475403e025afb4a053cc988cb8ba31ef062847e5e8b4d6"
var pbcKey = "513f6d70886d4abd6c475403e025afb4a053cc988cb8ba31ef062847e5e8b4d6"

func Auth() {
	userid := "salman"
	tokenstring, _ := watoken.Encode(userid, pvtKey)
	fmt.Println(tokenstring)
	fmt.Println(pvtKey)
	fmt.Println(pbcKey)
	//decode token to get userid
	useridstring := watoken.DecodeGetId(pbcKey, tokenstring)
	if useridstring == "" {
		fmt.Println("expire token")
	}
}
