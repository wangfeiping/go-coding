package operations

// 本代码应该由go-swagger生成，
// 但可能由于go-swagger还没有开发完成，所以部分代码无法成功生成，
// 因此这部分代码为根据go-swagger工程中的示例脚本生成的代码结构为规范，手写完成的代码！

import (
	"fmt"
	
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
TweetLikeParams contains all the parameters to send to the API endpoint
for the get user by name operation typically these are written to a http.Request
*/
type TweetLikeParams struct {
	TweetId int32
}

// WriteToRequest writes these params to a swagger request
// API: POST /api/social/tweet/{tweet_id}/like
func (o *TweetLikeParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {
	// path param name
	if err := r.SetPathParam("social", "tweet"); err != nil {
		return err
	}
	if err := r.SetPathParam(fmt.Sprintf("%d", o.TweetId), "like"); err != nil {
		return err
	}
	
	return nil
}
