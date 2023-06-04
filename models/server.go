package models

//This doesn't necessarily have to be a db model
//it can persist in memory as it can suffice
//Server will contain all the user connected to a particular channel
//microservice might need some way to partition
type Server struct {
	ChannelID uint
	UserList  []User
}
