package main

import (
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

func getPersonByIDs(ids []int) ([]models.Person, error) {
	filename := "data.json"
	file, err := os.Open(filename)
	if err != nil {
		return []models.Person{}, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return []models.Person{}, err
	}
	var people []models.Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		return []models.Person{}, err
	}
	var result []models.Person
	for _, person := range people {
		for _, id := range ids {
			if person.ID == id {
				result = append(result, person)
				break
			}
		}
	}

	if len(result) == 0 {
		return nil, errors.New("People not found")
	}

	return result, nil
}

func main() {
	router := gin.Default()
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("Error initializing config", err)

	}
	router.GET("/get-items", getItemsHandler)
	router.Run(viper.GetString("port"))
}
