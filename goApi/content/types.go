package content

// Tag 标签类型
type Tag = string

// AddArticleParams 添加文章参数
type AddArticleParams struct {
	Author      string `json:"author"`
	Content     string `json:"content"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Name        string `json:"name"`
	PublishedAt int64  `json:"published_at"`
	Tags        []Tag  `json:"tags"`
	URL         string `json:"url"`
}

// ModifyArticleParams 修改文章参数
type ModifyArticleParams struct {
	AddArticleParams
	ArticleID int `json:"article_id"`
}

// AddVideoParams 添加视频参数
type AddVideoParams struct {
	Author      string `json:"author"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Name        string `json:"name"`
	PublishedAt int64  `json:"published_at"`
	Tags        []Tag  `json:"tags"`
	URL         string `json:"url"`
}

// ModifyVideoParams 修改视频参数
type ModifyVideoParams struct {
	AddVideoParams
	VideoID int `json:"video_id"`
}

// PodcastHighlight 播客高亮
type PodcastHighlight struct {
	Highlight string `json:"highlight"`
	Second    int    `json:"second"`
}

// AddPodcastParams 添加播客参数
type AddPodcastParams struct {
	Author      string            `json:"author"`
	Channel     string            `json:"channel"`
	Cover       string            `json:"cover"`
	Description string            `json:"description"`
	Highlights  []PodcastHighlight `json:"highlights"`
	Name        string            `json:"name"`
	PublishedAt int64             `json:"published_at"`
	Status      int               `json:"status"`
	Tags        []Tag             `json:"tags"`
	URL         string            `json:"url"`
}

// ModifyPodcastParams 修改播客参数
type ModifyPodcastParams struct {
	AddPodcastParams
	PodcastID int `json:"podcast_id"`
}

// AddComicParams 添加漫画参数
type AddComicParams struct {
	Author      string `json:"author"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Name        string `json:"name"`
	PublishedAt int64  `json:"published_at"`
	Tags        []Tag  `json:"tags"`
}

// ModifyComicParams 修改漫画参数
type ModifyComicParams struct {
	AddComicParams
	ComicID int `json:"comic_id"`
}

// AddComicChapterParams 添加漫画章节参数
type AddComicChapterParams struct {
	ComicID     int      `json:"comic_id"`
	ChapterNo   int      `json:"chapter_no"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	PageURLs    []string `json:"page_urls"`
	PublishedAt int64    `json:"published_at"`
	Status      int      `json:"status"`
}

// ModifyComicChapterParams 修改漫画章节参数
type ModifyComicChapterParams struct {
	AddComicChapterParams
	ComicChapterID int `json:"comic_chapter_id"`
}

// GetComicChapterParams 获取漫画章节参数
type GetComicChapterParams struct {
	ComicID  int `json:"comic_id"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// GetComicPageParams 获取漫画页面参数
type GetComicPageParams struct {
	ComicChapterID int `json:"comic_chapter_id"`
	Page           int `json:"page"`
	PageSize       int `json:"page_size"`
}

// AddCommentParams 添加评论参数
type AddCommentParams struct {
	Avatar         string `json:"avatar"`
	CommentText    string `json:"comment_text"`
	ContentID      int    `json:"content_id"`
	ContentType    string `json:"content_type"`
	ParentID       int    `json:"parent_id"`
	ReplyCommentID int    `json:"reply_comment_id"`
	ReplyUserID    int    `json:"reply_user_id"`
	Status         int    `json:"status"`
	UserName       string `json:"user_name"`
}

// GetRootCommentParams 获取根评论参数
type GetRootCommentParams struct {
	ContentType string `json:"content_type"`
	ContentID   int    `json:"content_id"`
	Page        int    `json:"page"`
	PageSize    int    `json:"page_size"`
}

// GetSubCommentParams 获取子评论参数
type GetSubCommentParams struct {
	ContentType     string `json:"content_type"`
	ContentID       int    `json:"content_id"`
	ParentCommentID int    `json:"parent_comment_id"`
	Page            int    `json:"page"`
	PageSize        int    `json:"page_size"`
}

// ContentActionParams 内容通用参数
type ContentActionParams struct {
	ContentID   int    `json:"content_id"`
	ContentType string `json:"content_type"`
}

// CursorPageParams 游标分页参数
type CursorPageParams struct {
	PageSize int `json:"page_size"`
	Cursor   int `json:"cursor"`
}
