package main

type MsgType int

const (
	LOGIN        = 1
	BROADCAST    = 2
	CREATE_GROUP = 3
	JOIN_GROUP   = 4
	LOGOUT       = 99
)

type Body interface{}

type LoginBody struct {
	Body
	Uid  string
	Name string
}

type MsgBody struct {
	Body
	Gid string
	Msg string
}

type GroupBody struct {
	Body
	Gid  string
	Name string
}
