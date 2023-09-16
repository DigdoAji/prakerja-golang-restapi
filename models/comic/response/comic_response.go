package response

import "ujk-golang/models/comic/database"

type ComicResponse struct {
	Id int					`json:"id"`
	Title string			`json:"title"`
	Author string			`json:"author"`
	Genre string			`json:"genre"`
	Category string			`json:"category"`
	Date_Published string	`json:"date_published"`
	Completed bool			`json:"completed"`
}

func (comicResponse *ComicResponse) MapFromDatabaseComic(comicdatabase database.Comic){
	comicResponse.Id = comicdatabase.Id
	comicResponse.Title = comicdatabase.Title
	comicResponse.Author = comicdatabase.Author
	comicResponse.Genre = comicdatabase.Genre
	comicResponse.Category = comicdatabase.Category
	comicResponse.Date_Published = comicdatabase.Date_Published
	comicResponse.Completed = comicdatabase.Completed
}