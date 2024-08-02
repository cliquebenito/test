package models

type Input struct {
	ID []int `json:"id" binding:"required"`
}

type Data struct {
	ID                        int    `json:"id"`
	UID                       string `json:"uid"`
	Domain                    string `json:"domain"`
	CN                        string `json:"cn"`
	Department                string `json:"department"`
	Title                     string `json:"title"`
	Who                       string `json:"who"`
	LogonCount                string `json:"logon_count"`
	NumLogons7                string `json:"num_logons_7"`
	NumShare7                 string `json:"num_share_7"`
	NumFile7                  string `json:"num_file_7"`
	NumAD7                    string `json:"num_ad_7"`
	NumN7                     string `json:"num_n_7"`
	NumLogons14               string `json:"num_logons_14"`
	NumShare14                string `json:"num_share_14"`
	NumFile14                 string `json:"num_file_14"`
	NumAD14                   string `json:"num_ad_14"`
	NumN14                    string `json:"num_n_14"`
	NumLogons30               string `json:"num_logons_30"`
	NumShare30                string `json:"num_share_30"`
	NumFile30                 string `json:"num_file_30"`
	NumAD30                   string `json:"num_ad_30"`
	NumN30                    string `json:"num_n_30"`
	NumLogons150              string `json:"num_logons_150"`
	NumShare150               string `json:"num_share_150"`
	NumFile150                string `json:"num_file_150"`
	NumAD150                  string `json:"num_ad_150"`
	NumN150                   string `json:"num_n_150"`
	NumLogons365              string `json:"num_logons_365"`
	NumShare365               string `json:"num_share_365"`
	NumFile365                string `json:"num_file_365"`
	NumAD365                  string `json:"num_ad_365"`
	NumN365                   string `json:"num_n_365"`
	HasUserPrincipalName      string `json:"has_user_principal_name"`
	HasMail                   string `json:"has_mail"`
	HasPhone                  string `json:"has_phone"`
	FlagDisabled              string `json:"flag_disabled"`
	FlagLockout               string `json:"flag_lockout"`
	FlagPasswordNotRequired   string `json:"flag_password_not_required"`
	FlagPasswordCantChange    string `json:"flag_password_cant_change"`
	FlagDontExpirePassword    string `json:"flag_dont_expire_password"`
	OwnedFiles                string `json:"owned_files"`
	NumMailboxes              string `json:"num_mailboxes"`
	NumMemberOfGroups         string `json:"num_member_of_groups"`
	NumMemberOfIndirectGroups string `json:"num_member_of_indirect_groups"`
	MemberOfIndirectGroupsIDs string `json:"member_of_indirect_groups_ids"`
	MemberOfGroupsIDs         string `json:"member_of_groups_ids"`
	IsAdmin                   string `json:"is_admin"`
	IsService                 string `json:"is_service"`
}
