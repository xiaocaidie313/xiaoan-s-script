package content

// ResponseContent 内容响应
type ResponseContent struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

// Article 文章
type Article struct {
	ID          int    `json:"id"`
	Author      string `json:"author"`
	Content     string `json:"content"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Name        string `json:"name"`
	PublishedAt int64  `json:"published_at"`
	Tags        []string `json:"tags"`
	URL         string `json:"url"`
}

// Video 视频
type Video struct {
	ID          int      `json:"id"`
	Author      string   `json:"author"`
	Cover       string   `json:"cover"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	PublishedAt int64    `json:"published_at"`
	Tags        []string `json:"tags"`
	URL         string   `json:"url"`
}

// Podcast 播客
type Podcast struct {
	ID          int                `json:"id"`
	Author      string             `json:"author"`
	Channel     string             `json:"channel"`
	Cover       string             `json:"cover"`
	Description string             `json:"description"`
	Highlights  []PodcastHighlight `json:"highlights"`
	Name        string             `json:"name"`
	PublishedAt int64              `json:"published_at"`
	Status      int                `json:"status"`
	Tags        []string           `json:"tags"`
	URL         string             `json:"url"`
}

// Comic 漫画
type Comic struct {
	ID          int      `json:"id"`
	Author      string   `json:"author"`
	Cover       string   `json:"cover"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	PublishedAt int64    `json:"published_at"`
	Tags        []string `json:"tags"`
}

// ComicChapters 漫画章节列表
type ComicChapters struct {
	List  []ComicChapter `json:"list"`
	Total int            `json:"total"`
}

// ComicChapter 漫画章节
type ComicChapter struct {
	ID          int    `json:"id"`
	ComicID     int    `json:"comic_id"`
	ChapterNo   int    `json:"chapter_no"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ComicPages 漫画页面
type ComicPages struct {
	List  []ComicPage `json:"list"`
	Total int         `json:"total"`
}

// ComicPage 漫画页面
type ComicPage struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Order int    `json:"order"`
}

// ArticleList 文章列表
type ArticleList struct {
	List  []Article `json:"list"`
	Total int       `json:"total"`
}

// VideoList 视频列表
type VideoList struct {
	List  []Video `json:"list"`
	Total int     `json:"total"`
}

// PodcastList 播客列表
type PodcastList struct {
	List  []Podcast `json:"list"`
	Total int       `json:"total"`
}

// ComicList 漫画列表
type ComicList struct {
	List  []Comic `json:"list"`
	Total int     `json:"total"`
}

// RootComments 根评论
type RootComments struct {
	List  []Comment `json:"list"`
	Total int       `json:"total"`
}

// SubComments 子评论
type SubComments struct {
	List  []Comment `json:"list"`
	Total int       `json:"total"`
}

// Comment 评论
type Comment struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	Avatar      string `json:"avatar"`
	CommentText string `json:"comment_text"`
	CreatedAt   int64  `json:"created_at"`
}
