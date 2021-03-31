package rest

import (
	"encoding/json"
	"fmt"
	"github.com/simeonkorchev/rentim/pkg/rent"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func PropertyHandler(service rent.PropertyService) http.Handler {
	router := httprouter.New()

	router.GET("/properties", getProperties(service))
	router.GET("/properties/:id", getProperty(service))
	return router
}

func getProperty(service rent.PropertyService) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		property, err := service.FindPropertyByID(id)
		if err != nil {
			http.Error(writer, fmt.Sprintf("can not find property with id '%v'", id), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(property)
		if err != nil {
			http.Error(writer, "can not encode properties to JSON",http.StatusInternalServerError)
			return
		}
	}
}

func getProperties(service rent.PropertyService) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		props, err := service.FindAllProperties()
		if err != nil {
			http.Error(writer, "can not find all properties", http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(props)
		if err != nil {
			http.Error(writer, "can not encode properties to JSON",http.StatusInternalServerError)
			return
		}
	}
}
