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

	var emp models.Data
	var employees []models.Data

	for _, each := range csvData {
		emp.ID, _ = strconv.Atoi(each[0])
		emp.UID = each[1]
		emp.Domain = each[2]
		emp.CN = each[3]
		emp.Department = each[4]
		emp.Title = each[5]
		emp.Who = each[6]
		emp.LogonCount = each[7]
		emp.NumLogons7 = each[8]
		emp.NumShare7 = each[9]
		emp.NumFile7 = each[10]
		emp.NumAD7 = each[11]
		emp.NumN7 = each[12]
		emp.NumLogons14 = each[13]
		emp.NumShare14 = each[14]
		emp.NumFile14 = each[15]
		emp.NumAD14 = each[16]
		emp.NumN14 = each[17]
		emp.NumLogons30 = each[18]
		emp.NumShare30 = each[19]
		emp.NumFile30 = each[20]
		emp.NumAD30 = each[21]
		emp.NumN30 = each[22]
		emp.NumLogons150 = each[23]
		emp.NumShare150 = each[24]
		emp.NumFile150 = each[25]
		emp.NumAD150 = each[26]
		emp.NumN150 = each[27]
		emp.NumLogons365 = each[28]
		emp.NumShare365 = each[29]
		emp.NumFile365 = each[30]
		emp.NumAD365 = each[31]
		emp.NumN365 = each[32]
		emp.HasUserPrincipalName = each[33]
		emp.HasMail = each[34]
		emp.HasPhone = each[35]
		emp.FlagDisabled = each[36]
		emp.FlagLockout = each[37]
		emp.FlagPasswordNotRequired = each[38]
		emp.FlagPasswordCantChange = each[39]
		emp.FlagDontExpirePassword = each[40]
		emp.OwnedFiles = each[41]
		emp.NumMailboxes = each[42]
		emp.NumMemberOfGroups = each[43]
		emp.NumMemberOfIndirectGroups = each[44]
		emp.MemberOfIndirectGroupsIDs = each[45]
		emp.MemberOfGroupsIDs = each[46]
		emp.IsAdmin = each[47]
		emp.IsService = each[48]
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
