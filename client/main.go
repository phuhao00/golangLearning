package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type UpdateRequestParam struct {
	Id             int64    `json:"id omitempty" validate:"gte=0"`
	Title          string   `json:"title omitempty"`                          //标题
	Content        string   `json:"content omitempty"`                        //内容
	Images         []string `json:"images omitempty"`                         //图片
	LayoutType     int      `json:"layout_type omitempty" validate:"gte=0"`   //图片布局方式
	Tags           []string `json:"tags omitempty"`                           //标签
	Channel        int64    `json:"channels omitempty" validate:"gte=0"`      //频道
	Abstract       string   `json:"abstract omitempty"`                       //摘要
	RelatedTeams   []string `json:"related_teams omitempty"`                  //关联的球队
	RelatedPlayers []string `json:"related_players omitempty"`                //关联的球员
	IsCommented    int      `json:"is_commented omitempty" validate:"gte=0" ` //是否允许评论
	CopyRight      int      `json:"copy_right omitempty" validate:"gte=0"`    //版权转载限制
	CopyRightUrl   string   `json:"copy_right_url omitempty"`                 //版权链接
	UserId         int64    `json:"user_id omitempty" validate:"gte=0"`       //不需要添
	UserName       string   `json:"user_name omitempty"`                      //不需要填//
}

func main() {
	client := resty.New()
	resp, err := client.R().
		SetBody(UpdateRequestParam{
			Id:             66,
			Title:          "hhhh",
			Content:        "iiiiii45456546545456455454ooouuu",
			Images:         []string{"exercita888tion"},
			LayoutType:     0,
			Tags:           []string{"exercitation"},
			Channel:        0,
			Abstract:       "",
			RelatedTeams:   []string{"exercitation"},
			RelatedPlayers: []string{"exercitation"},
			IsCommented:    0,
			CopyRight:      0,
			CopyRightUrl:   "",
			UserId:         18,
			UserName:       "",
		}).
		Post("http://127.0.0.1:8088/content/article/draft-update")

	fmt.Println(string(resp.Body()), err)

}
