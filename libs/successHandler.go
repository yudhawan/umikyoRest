package libs

import (
	"encoding/json"
	"net/http"
)

func SendSuccessResponse(w http.ResponseWriter, msg string) {
	// m :=""
	// if msg!="" {
	// 	m="Success"
	// }
	json.NewEncoder(w).Encode("s")
}
