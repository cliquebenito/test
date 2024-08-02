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
	LogonCount                int    `json:"logon_count"`
	NumLogons7                int    `json:"num_logons_7"`
	NumShare7                 int    `json:"num_share_7"`
	NumFile7                  int    `json:"num_file_7"`
	NumAD7                    int    `json:"num_ad_7"`
	NumN7                     int    `json:"num_n_7"`
	NumLogons14               int    `json:"num_logons_14"`
	NumShare14                int    `json:"num_share_14"`
	NumFile14                 int    `json:"num_file_14"`
	NumAD14                   int    `json:"num_ad_14"`
	NumN14                    int    `json:"num_n_14"`
	NumLogons30               int    `json:"num_logons_30"`
	NumShare30                int    `json:"num_share_30"`
	NumFile30                 int    `json:"num_file_30"`
	NumAD30                   int    `json:"num_ad_30"`
	NumN30                    int    `json:"num_n_30"`
	NumLogons150              int    `json:"num_logons_150"`
	NumShare150               int    `json:"num_share_150"`
	NumFile150                int    `json:"num_file_150"`
	NumAD150                  int    `json:"num_ad_150"`
	NumN150                   int    `json:"num_n_150"`
	NumLogons365              int    `json:"num_logons_365"`
	NumShare365               int    `json:"num_share_365"`
	NumFile365                int    `json:"num_file_365"`
	NumAD365                  int    `json:"num_ad_365"`
	NumN365                   int    `json:"num_n_365"`
	HasUserPrincipalName      bool   `json:"has_user_principal_name"`
	HasMail                   bool   `json:"has_mail"`
	HasPhone                  bool   `json:"has_phone"`
	FlagDisabled              bool   `json:"flag_disabled"`
	FlagLockout               bool   `json:"flag_lockout"`
	FlagPasswordNotRequired   bool   `json:"flag_password_not_required"`
	FlagPasswordCantChange    bool   `json:"flag_password_cant_change"`
	FlagDontExpirePassword    bool   `json:"flag_dont_expire_password"`
	OwnedFiles                int    `json:"owned_files"`
	NumMailboxes              int    `json:"num_mailboxes"`
	NumMemberOfGroups         int    `json:"num_member_of_groups"`
	NumMemberOfIndirectGroups int    `json:"num_member_of_indirect_groups"`
	MemberOfIndirectGroupsIDs string `json:"member_of_indirect_groups_ids"`
	MemberOfGroupsIDs         string `json:"member_of_groups_ids"`
	IsAdmin                   bool   `json:"is_admin"`
	IsService                 bool   `json:"is_service"`
}
