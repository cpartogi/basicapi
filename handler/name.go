package handler

import (
	"basicapi/db"
	"basicapi/models"
	"basicapi/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetData(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {

		responseError := utils.ResponseError{
			Code:    405,
			Message: "method not allowed",
		}

		resjson, _ := json.Marshal(&responseError)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, resjson)
		return
	}

	if req.URL.Query().Get("id") == "" {
		var names []models.NameList

		for _, name := range db.Namedb {
			names = append(names, name)
		}

		response := utils.ResponseSuccess{
			Code: 200,
			Data: names,
		}

		resjson, _ := json.Marshal(&response)
		utils.ReturnJsonResponse(res, http.StatusOK, resjson)
		return
	} else {
		id := req.URL.Query()["id"][0]

		if len(id) == 1 {
			//check valid id
			_, err := strconv.Atoi(id)

			if err != nil {
				response := utils.ResponseError{
					Code:    400,
					Message: "invalid or empty ID: " + id,
				}

				resjson, _ := json.Marshal(&response)

				utils.ReturnJsonResponse(res, http.StatusBadRequest, resjson)
				return
			}

			name, ok := db.Namedb[id]
			if !ok {
				responseError := utils.ResponseError{
					Code:    404,
					Message: "resource with ID " + id + " not exist",
				}

				resjson, _ := json.Marshal(&responseError)

				utils.ReturnJsonResponse(res, http.StatusNotFound, resjson)
				return
			}

			response := utils.ResponseSuccess{
				Code: 200,
				Data: name,
			}

			resjson, _ := json.Marshal(&response)

			utils.ReturnJsonResponse(res, http.StatusOK, resjson)
			return

		} else {
			splitId := strings.Split(id, ",")
			var names []models.NameList

			for _, idname := range splitId {
				name, ok := db.Namedb[idname]
				if ok {
					names = append(names, name)
				}
			}
			response := utils.ResponseSuccess{
				Code: 200,
				Data: names,
			}

			resjson, _ := json.Marshal(&response)
			utils.ReturnJsonResponse(res, http.StatusOK, resjson)
			return

		}
	}
}
