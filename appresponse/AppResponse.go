package appresponse

// struct khai bao thanh cong
type SuccessRes struct {
	Data   interface{} `json:"nofi"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessRes(data, page, filter interface{}) *SuccessRes {
	return &SuccessRes{Data: data, Paging: page, Filter: filter}
}

func SimpleSuccessRes(data interface{}) *SuccessRes {
	return &SuccessRes{Data: data, Paging: nil, Filter: nil}
}
