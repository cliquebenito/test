package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"test/config"
	"test/internal/models"
	"test/pkg"
)

func getItemsHandler(c *gin.Context) {

	var input models.Input
	if err := c.BindJSON(&input); err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	people, err := getPersonByIDs(input.ID)
	if err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, people)
}

func convertCsvtoJson() error {
	csvFile, err := os.Open("ueba(3).csv")
	if err != nil {
		return err

	}
	defer func() {
		if err := csvFile.Close(); err != nil {
			return
		}
	}()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		return err
	}

	employees := pkg.FillIn(csvData)
	jsonData, err := json.Marshal(employees)
	if err != nil {
		return err
	}

	jsonFile, err := os.Create("newdata.json")
	if err != nil {
		return err
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			return
		}
	}()
	if _, err = jsonFile.Write(jsonData); err != nil {
		return err
	}
	if err = jsonFile.Close(); err != nil {
		return err
	}
	return nil
}

func getPersonByIDs(ids []int) ([]models.Data, error) {
	err := convertCsvtoJson()
	if err != nil {
		return []models.Data{}, err
	}

	filename := "newdata.json"
	file, err := os.Open(filename)
	if err != nil {
		return []models.Data{}, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return []models.Data{}, err
	}
	var people []models.Data
	err = json.Unmarshal(data, &people)
	if err != nil {
		return []models.Data{}, err
	}
	var result []models.Data
	for _, person := range people {
		for _, id := range ids {
			if person.ID == id {
				result = append(result, person)
				break
			}
		}
	}

	if len(result) == 0 {
		return []models.Data{}, errors.New("Info not found")
	}

	return result, nil
}

func main() {
	router := gin.Default()
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("Error initializing config", err)

	}
	router.GET("/get-items", getItemsHandler)
	if err := router.Run(viper.GetString("port")); err != nil {
		logrus.Fatalf("Error starting server", err)
	}
}
