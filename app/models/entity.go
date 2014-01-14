package models

type User struct {
	Email    string
	Nickname string
	Gender   string
	Password []byte
}

type MockUser struct {
	Email           string
	Nickname        string
	Gender          string
	Password        string
	ConfirmPassword string
}

type LoginUser struct {
	Email    string
	Nickname string
	Password string
}

type Quotation struct {
	Title    string
	Tag      string
	Content  string
	Original string
	Author   string
}

type AncientPoetry struct {
	Title   string
	Tag     string
	Content string
	Author  string
}

type YueFu struct {
	Title   string
	Tag     string
	Content string
	Author  string
}

type TangPoems struct {
	Title   string
	Tag     string
	Content string
	Author  string
}

type SongPoems struct {
	Title   string
	Tag     string
	Content string
	Author  string
}

type YuanVerse struct {
	Title   string
	Tag     string
	Content string
	Author  string
}

type ModernPoetry struct {
	Title   string
	Tag     string
	Content string
	Author  string
}

type Essay struct {
	Title   string
	Tag     string
	Content string
	Author  string
}
