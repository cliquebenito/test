package service

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"test/internal/models"
)

type Errors struct {
	Message string `json:"message"`
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

	employees, err := fillIn(csvData)
	if err != nil {
		return err
	}
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
func fillIn(csvData [][]string) ([]models.Data, error) {
	var emp models.Data
	var employees []models.Data
	for _, each := range csvData {
		emp.ID, _ = strconv.Atoi(each[0])
		emp.UID, _ = strconv.Atoi(each[1])
		emp.Domain = each[2]
		emp.CN = each[3]
		emp.Department = each[4]
		emp.Title = each[5]
		emp.Who = each[6]
		emp.LogonCount = each[7]
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
		emp.HasUserPrincipalName, _ = strconv.Atoi(each[33])
		emp.HasMail, _ = strconv.Atoi(each[34])
		emp.HasPhone, _ = strconv.Atoi(each[35])
		emp.FlagDisabled, _ = strconv.ParseBool(each[36])
		emp.FlagLockout, _ = strconv.ParseBool(each[37])
		emp.FlagPasswordNotRequired, _ = strconv.ParseBool(each[38])
		emp.FlagPasswordCantChange, _ = strconv.ParseBool(each[39])
		emp.FlagDontExpirePassword, _ = strconv.ParseBool(each[40])
		emp.OwnedFiles, _ = strconv.Atoi(each[41])
		emp.NumMailboxes, _ = strconv.Atoi(each[42])
		emp.NumMemberOfGroups, _ = strconv.Atoi(each[43])
		emp.NumMemberOfIndirectGroups, _ = strconv.Atoi(each[44])
		emp.MemberOfIndirectGroupsIDs, _ = strconv.Atoi(each[45])
		emp.MemberOfGroupsIDs = strings.Split(each[46], ";")
		emp.IsAdmin = strings.Split(each[47], ";")
		emp.IsService, _ = strconv.ParseBool(each[48])
		employees = append(employees, emp)
	}

	return employees, nil
}
