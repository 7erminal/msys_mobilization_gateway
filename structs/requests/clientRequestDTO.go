package requests

type AddClientRequest struct {
	ClientName string `orm:"size(100)"`
	ClientCode string `orm:"size(100)"`
	CreatedBy  int
	ModifiedBy int
}
