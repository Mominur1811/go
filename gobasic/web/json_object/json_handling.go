package json_object

import (
	"encoding/json"
	"net/http"
)

func JsonDecoding(r *http.Request, v interface{}) error {

	return json.NewDecoder(r.Body).Decode(&v)

	//switch val := v.(type){

	//case *db.User:return json.NewDecoder(r.Body).Decode(&v)
	//case **db.User: return json.NewDecoder(r.Body).Decode(&v)

	//}

}
