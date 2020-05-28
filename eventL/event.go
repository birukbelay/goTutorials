//Event ..returns single event
func (eri *EventRepoImpl) Event(id int) (entity.Event, error) {
    row := eri.conn.QueryRow("SELECT * FROM events WHERE id = $1", id)
    event := entity.Event{}
    err := rows.Scan(&event.EId, &event.Name, &event.Details, &event.Image, &event.City, &event.Country, &event.Place, &event.Coordinates, &event.host, &event.IsPassed, &event.Rating, &event.PostedDate, &event.price)
    if err != nil {
        return event, err
    }

    return event, nil
}


//Post gy
func (eri *EventRepoImpl) Post(event entity.Event) error {
    _, err := eri.conn.Exec("INSERT INTO events (EId, Name, Details, Image, City, Country, Place, Coordinates, host, IsPassed, Rating, PostedDate, price) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",event.EId, event.Name, event.Details, event.Image, event.City, event.Country, event.Place, event.Coordinates, event.host, event.IsPassed, event.Rating, event.PostedDate, event.price)
    if err != nil {
        return errors.New("Insertion has failed")
    }
    return nil


func (rri *ReviewRepoImpl) EventReviews(id int) ([]entity.Review, error) {

	query := "SELECT * FROM review WHERE Event_id = $1"
	rows, err := rri.conn.Query(query, id)

	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	rvws := []entity.Review{}

	for rows.Next() {
		review := entity.Review{}
		err = rows.Scan(&review.ID, &review.Rating, &review.ReviewedAt, &review.UserID, &review.EventID, &review.Message)
		if err != nil {
			return nil, err
		}

		rvws = append(rvws, review)
	}

	return rvws, nil
}


func GetPost(id int) (post Post, err error) {
post = Post{}
post.Comments = []Comment{}
err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
rows, err := Db.Query("select id, content, author from comments where post_id = $1", id)
if err != nil {
return
}
for rows.Next() {
comment := Comment{Post: &post}
err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
if err != nil {
return
}
post.Comments = append(post.Comments, comment)
}
rows.Close()
return
}
