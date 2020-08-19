package repo

type Paper struct {
	Pid     string    `bson:"pid" json:"pid"`
	Title   string    `bson:"title" json:"title"`
	Content []Content `bson:"content" json:"content"`
	Updated int64     `bson:"updated" json:"updated"` /* time.Time.second */
}

type Content struct {
	Schema string `bson:"schema" json:"schema"`
	Text   string `bson:"text" json:"text"`
	Attach string `bson:"attach" json:"attach"`
	Check  int    `bson:"check" json:"check"`
}
