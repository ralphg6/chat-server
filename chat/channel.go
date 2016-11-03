package chat

type Channel chan ChannelStruct

type ChannelStruct struct {
	Uid   int     `json:'uid',omitempty`
	Type  MsgType `json:'msg_type'`
	Body  Body    `json:'body'`
	Teste string
}
