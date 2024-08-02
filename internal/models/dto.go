package models

type Input struct {
	ID []int `json:"id" binding:"required"`
}

//	type Person struct {
//		ID         int    `json:"id"`
//		Name       string `json:"name"`
//		Occupation string `json:"occupation"`
//		Born       string `json:"born"`
//	}
type Person struct {
	Number      int    `json:"number"`
	ID          int    `json:"id"`
	UID         string `json:"uid"`
	Domain      string `json:"domain"`
	CN          string `json:"cn"`
	Department  string `json:"department"`
	Title       string `json:"title"`
	Who         string `json:"who"`
	LogonCount  int    `json:"logon_count"`
	NumLogons7  int    `json:"num_logons_7"`
	NumShare7   int    `json:"num_share_7"`
	NumFile7    int    `json:"num_file_7"`
	NumAd7      int    `json:"num_ad_7"`
	NumN7       int    `json:"num_n_7"`
	NumLogons14 int    `json:"num_logons_14"`
	NumShare14  int    `json:"num_share_14"`
	NumFile14   int    `json:"num_file_14"`
	NumAd14     int    `json:"num_ad_14"`
	NumN14      int    `json:"num_n_14"`
}

type User struct {
	ID                        int
	UID                       string
	Domain                    string
	CN                        string
	Department                string
	Title                     string
	Who                       string
	LogonCount                int
	NumLogons7                int
	NumShare7                 int
	NumFile7                  int
	NumAD7                    int
	NumN7                     int
	NumLogons14               int
	NumShare14                int
	NumFile14                 int
	NumAD14                   int
	NumN14                    int
	NumLogons30               int
	NumShare30                int
	NumFile30                 int
	NumAD30                   int
	NumN30                    int
	NumLogons150              int
	NumShare150               int
	NumFile150                int
	NumAD150                  int
	NumN150                   int
	NumLogons365              int
	NumShare365               int
	NumFile365                int
	NumAD365                  int
	NumN365                   int
	HasUserPrincipalName      bool
	HasMail                   bool
	HasPhone                  bool
	FlagDisabled              bool
	FlagLockout               bool
	FlagPasswordNotRequired   bool
	FlagPasswordCantChange    bool
	FlagDontExpirePassword    bool
	OwnedFiles                int
	NumMailboxes              int
	NumMemberOfGroups         int
	NumMemberOfIndirectGroups int
	MemberOfIndirectGroupsIDs string
	MemberOfGroupsIDs         string
	IsAdmin                   bool
	IsService                 bool
}
