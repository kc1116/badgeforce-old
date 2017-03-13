package main

import "net/http"

func health(w http.ResponseWriter, r *http.Request) {

	/*healthCheckResponse is the default response for json marshalling for the /health route
	if err := json.NewEncoder(w).Encode(healthCheckResponse); err != nil {
		log.Println("JSON encoder error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}*/
}
