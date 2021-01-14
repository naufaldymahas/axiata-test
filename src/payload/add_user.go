package payload

type AddUser struct {
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	BirthPlace string `json:"birth_place"`
	BirthDate  string `json:"birth_date"`
}
