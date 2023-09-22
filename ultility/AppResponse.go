package ultility

// struct khai bao thanh cong
type SuccessRes struct {
	Data interface{} `json:"nofi"`
}

func SimpleSuccessRes(data interface{}) *SuccessRes {
	return &SuccessRes{Data: data}
}
