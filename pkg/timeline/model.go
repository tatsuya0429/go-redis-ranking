package timeline

type Post struct {
	Title       string `faker:"word,lang=jpn"`
	Description string `faker:"sentense,lang=jpn,len=140"`
	CreatedAt   int64  `faker:"unix_time"`
}

type PostJson struct {
	Title       string
	Description string
}
