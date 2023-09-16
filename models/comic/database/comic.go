package database

import "ujk-golang/models/comic/request"

type Comic struct {
	Id int					`json:"id" gorm:"primaryKey autoIncrement"`
	Title string			`json:"title"`
	Author string			`json:"author"`
	Genre string			`json:"genre"`
	Category string			`json:"category"`
	Date_Published string	`json:"date_published"`
	Completed bool			`json:"completed"`
}

func (comic *Comic) MapFromAddingComic(comicAdd request.ComicAdd){
	comic.Title = comicAdd.Title
	comic.Author = comicAdd.Author
	comic.Genre = comicAdd.Genre
	comic.Category = comicAdd.Category
	comic.Date_Published = comicAdd.Date_Published
	comic.Completed = comicAdd.Completed
}

func (comic *Comic) MapFromUpdateComic(comicAdd request.ComicUpdate){
	comic.Title = comicAdd.Title
	comic.Author = comicAdd.Author
	comic.Genre = comicAdd.Genre
	comic.Category = comicAdd.Category
	comic.Date_Published = comicAdd.Date_Published
	comic.Completed = comicAdd.Completed
}