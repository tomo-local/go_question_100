package main

type UserFetcher interface {
	Fetch(userID int) (string, error)
}

func GetUser(fetcher UserFetcher, id int) (string, error) {
	return fetcher.Fetch(id)
}

func main() {

}
