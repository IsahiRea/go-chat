package main

// Message struct for chat messages
type Message struct {
	Room string `json:"room"`
	User string `json:"user"`
	Text string `json:"text"`
	//Timestamp time.Time `json:"timestamp"`
}

// User struct for authenticated users
type User struct {
	Username string `json:"username"`
	//Password string `json:"password"`
	//Token    string `json:"token"`
}

type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
