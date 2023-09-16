package response

type ComicUpdate struct {
	Title string			`json:"title"`
	Author string			`json:"author"`
	Genre string			`json:"genre"`
	Category string			`json:"category"`
	Date_Published string	`json:"date_published"`
	Completed bool			`json:"completed"`
}