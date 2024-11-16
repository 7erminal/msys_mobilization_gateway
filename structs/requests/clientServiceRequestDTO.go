package requests

type ClientServiceRequest struct {
	Client            int64
	Service           int64
	ClientServicePath string `orm:"size(100)"`
	CreatedBy         int
	ModifiedBy        int
}
