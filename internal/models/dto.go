package models

type Input struct {
	ID []int `json:"id" binding:"required"`
}

type Data struct {
	ID                        int      `json:"id" db:"id"`
	UID                       int      `json:"uid" db:"uid"`
	Domain                    string   `json:"domain" db:"domain"`
	CN                        string   `json:"cn" db:"cn"`
	Department                string   `json:"department" db:"department"`
	Title                     string   `json:"title" db:"title"`
	Who                       string   `json:"who" db:"who"`
	LogonCount                string   `json:"logon_count" db:"logon_count"`
	NumLogons7                int      `json:"num_logons_7" db:"num_logons_7"`
	NumShare7                 int      `json:"num_share_7" db:"num_share_7"`
	NumFile7                  int      `json:"num_file_7" db:"num_file_7"`
	NumAD7                    int      `json:"num_ad_7" db:"num_ad_7"`
	NumN7                     int      `json:"num_n_7" db:"num_n_7"`
	NumLogons14               int      `json:"num_logons_14" db:"num_logons_14"`
	NumShare14                int      `json:"num_share_14" db:"num_share_14"`
	NumFile14                 int      `json:"num_file_14" db:"num_file_14"`
	NumAD14                   int      `json:"num_ad_14" db:"num_ad_14"`
	NumN14                    int      `json:"num_n_14" db:"num_n_14"`
	NumLogons30               int      `json:"num_logons_30" db:"num_logons_30"`
	NumShare30                int      `json:"num_share_30" db:"num_share_30"`
	NumFile30                 int      `json:"num_file_30" db:"num_file_30"`
	NumAD30                   int      `json:"num_ad_30" db:"num_ad_30"`
	NumN30                    int      `json:"num_n_30" db:"num_n_30"`
	NumLogons150              int      `json:"num_logons_150" db:"num_logons_150"`
	NumShare150               int      `json:"num_share_150" db:"num_share_150"`
	NumFile150                int      `json:"num_file_150" db:"num_file_150"`
	NumAD150                  int      `json:"num_ad_150" db:"num_ad_150"`
	NumN150                   int      `json:"num_n_150" db:"num_n_150"`
	NumLogons365              int      `json:"num_logons_365" db:"num_logons_365"`
	NumShare365               int      `json:"num_share_365" db:"num_share_365"`
	NumFile365                int      `json:"num_file_365" db:"num_file_365"`
	NumAD365                  int      `json:"num_ad_365" db:"num_ad_365"`
	NumN365                   int      `json:"num_n_365" db:"num_n_365"`
	HasUserPrincipalName      int      `json:"has_user_principal_name" db:"has_user_principal_name"`
	HasMail                   int      `json:"has_mail" db:"has_mail"`
	HasPhone                  int      `json:"has_phone" db:"has_phone"`
	FlagDisabled              bool     `json:"flag_disabled" db:"flag_disabled"`
	FlagLockout               bool     `json:"flag_lockout" db:"flag_lockout"`
	FlagPasswordNotRequired   bool     `json:"flag_password_not_required" db:"flag_password_not_required"`
	FlagPasswordCantChange    bool     `json:"flag_password_cant_change" db:"flag_password_cant_change"`
	FlagDontExpirePassword    bool     `json:"flag_dont_expire_password" db:"flag_dont_expire_password"`
	OwnedFiles                int      `json:"owned_files" db:"owned_files"`
	NumMailboxes              int      `json:"num_mailboxes" db:"num_mailboxes"`
	NumMemberOfGroups         int      `json:"num_member_of_groups" db:"num_member_of_groups"`
	NumMemberOfIndirectGroups int      `json:"num_member_of_indirect_groups" db:"num_member_of_indirect_groups"`
	MemberOfIndirectGroupsIDs int      `json:"member_of_indirect_groups_ids" db:"member_of_indirect_groups_ids"`
	MemberOfGroupsIDs         []string `json:"member_of_groups_ids" db:"member_of_groups_ids"`
	IsAdmin                   []string `json:"is_admin" db:"is_admin"`
	IsService                 bool     `json:"is_service" db:"is_service"`
}
