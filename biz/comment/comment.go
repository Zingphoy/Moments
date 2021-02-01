package model

type Comment struct {
	Id       interface{}              `json:"-"`
	Aid      int64                    `json:"aid"`
	Comment  []map[string]interface{} `json:"comment"`
	LikeList []int                    `json:"like_list"`
}
