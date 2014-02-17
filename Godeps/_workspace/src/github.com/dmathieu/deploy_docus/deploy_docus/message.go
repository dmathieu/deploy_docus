package deploy_docus

type Message struct {
	Id         int    `form:"id"`
	Sha        string `form:"sha"`
	Repository *Repository
}
