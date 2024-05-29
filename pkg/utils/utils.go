package utils

import (
	"encoding/json"
	"net/http"
	//"net/http"
)

func ParseBody(r *http.Request,x interface{}){
	  err:=json.NewDecoder(r.Body).Decode(&x)
		if err!=nil{
			return 
		}
}