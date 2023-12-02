package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"app/models"
	"github.com/gin-gonic/gin"
)

// Upload File godoc
// @ID upload_file
// @Router /upload [POST]
// @Summary Upload file
// @Description Upload file
// @Tags File
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Param table_slug query string true "table_slug"
// @Success 200 {object} Response{} "File"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) Upload(c *gin.Context){
	fmt.Println("ok")
	file, _ := c.FormFile("file")
	data,err := file.Open()
	if err!=nil{
		log.Println(err)
	}
	jsonFile,err := io.ReadAll(data)
	if err!=nil{
		log.Println("Err while reading file: " , err)
	}
	fmt.Println(jsonFile)
	table := c.Query("table_slug")

	if table == "country"{
	var countries []models.CreateCountry
	json.Unmarshal(jsonFile,&countries)
	for _,country := range countries{
		resp,err:= h.strg.Country().Create(&country)
		if err!=nil{
			handleResponse(c,http.StatusInternalServerError,err)
		}
		handleResponse(c,http.StatusOK,resp)
	}
	}
	if table == "city"{
		var cities []models.CreateCity
		json.Unmarshal(jsonFile,&cities)
		for _,city :=range cities{
			resp,err := h.strg.City().Create(&city)
			if err!=nil{
				handleResponse(c,http.StatusInternalServerError,err)
			}

		handleResponse(c,http.StatusOK,resp)
		}
	}

	if table == "airport"{
		var aeroports []models.CreateAeroport
		json.Unmarshal(jsonFile,&aeroports)
		for _,aeroport :=range aeroports{
			resp,err := h.strg.Aeroport().Create(&aeroport)
			if err!=nil{
				handleResponse(c,http.StatusInternalServerError,err)
			}

		handleResponse(c,http.StatusOK,resp)
		}
	}
	
	
}