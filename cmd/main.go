package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var emp models.User
	var employees []models.User

	for _, each := range csvData {
		emp.ID, _ = strconv.Atoi(each[0])
		emp.UID = each[1]
		emp.Domain = each[2]
		emp.CN = each[3]
		emp.Department = each[4]
		emp.Title = each[5]
		emp.Who = each[6]
		emp.LogonCount, _ = strconv.Atoi(each[7])
		emp.NumLogons7, _ = strconv.Atoi(each[8])
		emp.NumShare7, _ = strconv.Atoi(each[9])
		emp.NumFile7, _ = strconv.Atoi(each[10])
		emp.NumAD7, _ = strconv.Atoi(each[11])
		emp.NumN7, _ = strconv.Atoi(each[12])
		emp.NumLogons14, _ = strconv.Atoi(each[13])
		emp.NumShare14, _ = strconv.Atoi(each[14])
		emp.NumFile14, _ = strconv.Atoi(each[15])
		emp.NumAD14, _ = strconv.Atoi(each[16])
		emp.NumN14, _ = strconv.Atoi(each[17])
		emp.NumLogons30, _ = strconv.Atoi(each[18])
		emp.NumShare30, _ = strconv.Atoi(each[19])
		emp.NumFile30, _ = strconv.Atoi(each[20])
		emp.NumAD30, _ = strconv.Atoi(each[21])
		emp.NumN30, _ = strconv.Atoi(each[22])
		emp.NumLogons150, _ = strconv.Atoi(each[23])
		emp.NumShare150, _ = strconv.Atoi(each[24])
		emp.NumFile150, _ = strconv.Atoi(each[25])
		emp.NumAD150, _ = strconv.Atoi(each[26])
		emp.NumN150, _ = strconv.Atoi(each[27])
		emp.NumLogons365, _ = strconv.Atoi(each[28])
		emp.NumShare365, _ = strconv.Atoi(each[29])
		emp.NumFile365, _ = strconv.Atoi(each[30])
		emp.NumAD365, _ = strconv.Atoi(each[31])
		emp.NumN365, _ = strconv.Atoi(each[32])
		emp.HasUserPrincipalName, _ = strconv.ParseBool(each[33])
		emp.HasMail, _ = strconv.ParseBool(each[34])
		emp.HasPhone, _ = strconv.ParseBool(each[35])
		emp.FlagDisabled, _ = strconv.ParseBool(each[36])
		emp.FlagLockout, _ = strconv.ParseBool(each[37])
		emp.FlagPasswordNotRequired, _ = strconv.ParseBool(each[38])
		emp.FlagPasswordCantChange, _ = strconv.ParseBool(each[39])
		emp.FlagDontExpirePassword, _ = strconv.ParseBool(each[40])
		emp.OwnedFiles, _ = strconv.Atoi(each[41])
		emp.NumMailboxes, _ = strconv.Atoi(each[42])
		emp.NumMemberOfGroups, _ = strconv.Atoi(each[43])
		emp.NumMemberOfIndirectGroups, _ = strconv.Atoi(each[44])
		emp.MemberOfIndirectGroupsIDs = each[45]
		emp.MemberOfGroupsIDs = each[46]
		emp.IsAdmin, _ = strconv.ParseBool(each[47])
		emp.IsService, _ = strconv.ParseBool(each[48])

		employees = append(employees, emp)
	}

	jsonData, err := json.Marshal(employees)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))

	jsonFile, err := os.Create("newdata.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
	return nil
}

func getPersonByIDs(ids []int) ([]models.Person, error) {
	err := convertCsvtoJson()
	if err != nil {
		return []models.Person{}, err
	}

	filename := "newdata.json"
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
		return nil, errors.New("Info not found")
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
