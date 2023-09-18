package appresponse

// struct khai bao thanh cong
type SuccessRes struct {
	Data interface{} `json:"nofi"`
}

func NewSuccessRes(data, page, filter interface{}) *SuccessRes {
	return &SuccessRes{Data: data}
}

func SimpleSuccessRes(data interface{}) *SuccessRes {
	return &SuccessRes{Data: data}
}
