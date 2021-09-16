package models

type RespMsg struct {
	Msg string `json:"msg"`
}

type SharedReq struct {
	FileCid string `json:"cid"`
	Shared  bool   `json:"shared"`
}
