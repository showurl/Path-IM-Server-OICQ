// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
	Uid   string `json:"uid"`
}
