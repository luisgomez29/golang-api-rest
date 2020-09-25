package auto

import "github.com/luisgomez29/golang-api-rest/models"

var users = []models.User{
	{FirstName: "Luis", LastName: "Gómez", Email: "luis.gomez@usantoto.edu.co", Password: "123456"},
	{FirstName: "Carlos", LastName: "Pérez", Email: "calos@gmail.com", Password: "567890"},
	{FirstName: "Oscar", LastName: "Gómez", Email: "oscar@gmail.com", Password: "567890"},
}

var products = []models.Product{
	{UserID: 1, Name: "Computador"},
	{UserID: 1, Name: "Xbox One X"},
	{UserID: 2, Name: "Mouse"},
}
