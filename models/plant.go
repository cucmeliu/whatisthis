package models

type Plant struct {
	User_token string
	Img_base64 string
}

type PlantResult struct {
	Log_id int64       `json:"log_id"`
	Result []oneResult `json:"result"`
}

type oneResult struct {
	Name  string  `json:"name"`
	Score float32 `json:"score"`
}

// 返回格式为：
// {
// 	"log_id": "6200002018523288590",
// 	"result": [
// 		{
// 			"score": 0.7758851647377,
// 			"name": "日本樱花"
// 		},
// 		{
// 			"score": 0.091255366802216,
// 			"name": "樱花"
// 		},
// 		{
// 			"score": 0.052016720175743,
// 			"name": "樱花树"
// 		},
// 		{
// 			"score": 0.025772722437978,
// 			"name": "山樱花"
// 		},
// 		{
// 			"score": 0.011653803288937,
// 			"name": "东京樱花"
// 		}
// 	]
// }
