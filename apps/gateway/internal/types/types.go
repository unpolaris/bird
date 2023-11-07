// Code generated by goctl. DO NOT EDIT.
package types

type BirdAddReq struct {
	BirdName    string `json:"birdName"`
	BirdType    int64  `json:"birdType"`
	Description string `json:"description"`
	PicUrl      string `json:"picUrl"`
}

type BirdAddResp struct {
	BirdName string `json:"birdName"`
}

type BirUpdateReq struct {
	BirdId      int64  `json:"birdId"`
	BirdName    string `json:"birdName"`
	BirdType    int64  `json:"birdType"`
	Description string `json:"description"`
	PicUrl      string `json:"picUrl"`
}

type BirUpdateResp struct {
	BirdId int64 `json:"birdId"`
}

type BirListReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
	BirdType int64 `json:"birdType"`
}

type BirListResp struct {
	List []BirListData `json:"list"`
}

type BirListData struct {
	BirdId      int64  `json:"birdId"`
	BirdName    string `json:"birdName"`
	BirdType    int64  `json:"birdType"`
	Description string `json:"description"`
	PicUrl      string `json:"picUrl"`
}

type BirdSearchReq struct {
	Keyword string `json:"keyword"`
}

type BirGetReq struct {
	BirdId int64 `json:"birdId"`
}

type BirGetResp struct {
	BirdId      int64  `json:"birdId"`
	BirdName    string `json:"birdName"`
	BirdType    int64  `json:"birdType"`
	Description string `json:"description"`
	PicUrl      string `json:"picUrl"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
}

type BirDeleteReq struct {
	BirdId int64 `json:"birdId"`
}

type BirDeleteResp struct {
	BirdId int64 `json:"birdId"`
}
