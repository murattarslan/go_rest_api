package main

type Response struct {
	Result         Result
	Result_message ResultMessage
}

type Result struct {
	Name string
	Id   int
}

type ResultMessage struct {
	Message string
	Title   string
	Status  string
}

type Body struct {
	Id   int
	Name string
}
