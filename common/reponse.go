package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessReponse(data interface{}, paging interface{}, filter interface{}) *successRes {
	return &successRes{data, paging, filter}
}

func SimpleSuccessReponse(data interface{}) *successRes {
	return &successRes{data, nil, nil}
}
