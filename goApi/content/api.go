package content

import (
	"xiaoan-s-script/goApi"
)

// AddArticle 添加文章
func AddArticle(client *goApi.Client, params AddArticleParams) (*goApi.CommonResponse[ResponseContent], error) {
	var result goApi.CommonResponse[ResponseContent]
	err := client.Post("/api/content/add-article", params, &result)
	return &result, err
}

// GetArticleInfo 获取文章详情
func GetArticleInfo(client *goApi.Client, id int) (*goApi.CommonResponse[Article], error) {
	params := map[string]int{"article_id": id}
	var result goApi.CommonResponse[Article]
	err := client.Get("/api/content/get-article-content", params, &result)
	return &result, err
}

// ModifyArticle 修改文章
func ModifyArticle(client *goApi.Client, params ModifyArticleParams) (*goApi.CommonResponse[ResponseContent], error) {
	var result goApi.CommonResponse[ResponseContent]
	err := client.Post("/api/content/modify-article", params, &result)
	return &result, err
}

// DeleteArticle 删除文章
func DeleteArticle(client *goApi.Client, id int) (*goApi.CommonResponse[map[string]any], error) {
	params := map[string]int{"article_id": id}
	var result goApi.CommonResponse[map[string]any]
	err := client.Delete("/api/content/delete-article", params, &result)
	return &result, err
}

// AddVideo 添加视频
func AddVideo(client *goApi.Client, params AddVideoParams) (*goApi.CommonResponse[ResponseContent], error) {
	var result goApi.CommonResponse[ResponseContent]
	err := client.Post("/api/content/add-video", params, &result)
	return &result, err
}

// GetVideoContent 获取视频详情
func GetVideoContent(client *goApi.Client, id int) (*goApi.CommonResponse[Video], error) {
	params := map[string]int{"video_id": id}
	var result goApi.CommonResponse[Video]
	err := client.Get("/api/content/get-video-content", params, &result)
	return &result, err
}

// ModifyVideo 修改视频
func ModifyVideo(client *goApi.Client, params ModifyVideoParams) (*goApi.CommonResponse[ResponseContent], error) {
	var result goApi.CommonResponse[ResponseContent]
	err := client.Post("/api/content/modify-video", params, &result)
	return &result, err
}

// DeleteVideo 删除视频
func DeleteVideo(client *goApi.Client, id int) (*goApi.CommonResponse[map[string]any], error) {
	params := map[string]int{"video_id": id}
	var result goApi.CommonResponse[map[string]any]
	err := client.Delete("/api/content/delete-video", params, &result)
	return &result, err
}

// AddPodcast 添加播客
func AddPodcast(client *goApi.Client, params AddPodcastParams) (*goApi.CommonResponse[ResponseContent], error) {
	var result goApi.CommonResponse[ResponseContent]
	err := client.Post("/api/content/add-podcast", params, &result)
	return &result, err
}

// GetPodcastContent 获取播客内容
func GetPodcastContent(client *goApi.Client, id int) (*goApi.CommonResponse[Podcast], error) {
	params := map[string]int{"podcast_id": id}
	var result goApi.CommonResponse[Podcast]
	err := client.Get("/api/content/get-podcast-content", params, &result)
	return &result, err
}

// ModifyPodcast 修改播客
func ModifyPodcast(client *goApi.Client, params ModifyPodcastParams) (*goApi.CommonResponse[ResponseContent], error) {
	var result goApi.CommonResponse[ResponseContent]
	err := client.Post("/api/content/modify-podcast", params, &result)
	return &result, err
}

// DeletePodcast 删除播客
func DeletePodcast(client *goApi.Client, id int) (*goApi.CommonResponse[map[string]any], error) {
	params := map[string]int{"podcast_id": id}
	var result goApi.CommonResponse[map[string]any]
	err := client.Delete("/api/content/delete-podcast", params, &result)
	return &result, err
}

// AddComic 添加漫画
func AddComic(client *goApi.Client, params AddComicParams) (*goApi.CommonResponse[ResponseContent], error) {
	var result goApi.CommonResponse[ResponseContent]
	err := client.Post("/api/content/add-comic", params, &result)
	return &result, err
}

// AddComicChapter 添加漫画章节
func AddComicChapter(client *goApi.Client, params AddComicChapterParams) (*goApi.CommonResponse[ResponseContent], error) {
	var result goApi.CommonResponse[ResponseContent]
	err := client.Post("/api/content/add-comic-chapter", params, &result)
	return &result, err
}

// GetComic 获取漫画
func GetComic(client *goApi.Client, id int) (*goApi.CommonResponse[Comic], error) {
	params := map[string]int{"comic_id": id}
	var result goApi.CommonResponse[Comic]
	err := client.Get("/api/content/get-comic", params, &result)
	return &result, err
}

// GetComicChapter 获取漫画章节
func GetComicChapter(client *goApi.Client, params GetComicChapterParams) (*goApi.CommonResponse[ComicChapters], error) {
	var result goApi.CommonResponse[ComicChapters]
	err := client.Get("/api/content/get-comic-chapter", params, &result)
	return &result, err
}

// GetComicPage 获取漫画页面
func GetComicPage(client *goApi.Client, params GetComicPageParams) (*goApi.CommonResponse[ComicPages], error) {
	var result goApi.CommonResponse[ComicPages]
	err := client.Get("/api/content/get-comic-page", params, &result)
	return &result, err
}

// ModifyComic 修改漫画
func ModifyComic(client *goApi.Client, params ModifyComicParams) (*goApi.CommonResponse[ResponseContent], error) {
	var result goApi.CommonResponse[ResponseContent]
	err := client.Post("/api/content/modify-comic", params, &result)
	return &result, err
}

// ModifyComicChapter 修改漫画章节
func ModifyComicChapter(client *goApi.Client, params ModifyComicChapterParams) (*goApi.CommonResponse[ResponseContent], error) {
	var result goApi.CommonResponse[ResponseContent]
	err := client.Post("/api/content/modify-comic-chapter", params, &result)
	return &result, err
}

// DeleteComicChapter 删除漫画章节
func DeleteComicChapter(client *goApi.Client, comicChapterID, comicID int) (*goApi.CommonResponse[map[string]any], error) {
	params := map[string]int{"comic_chapter_id": comicChapterID, "comic_id": comicID}
	var result goApi.CommonResponse[map[string]any]
	err := client.Delete("/api/content/delete-comic-chapter", params, &result)
	return &result, err
}

// DeleteComic 删除漫画
func DeleteComic(client *goApi.Client, id int) (*goApi.CommonResponse[map[string]any], error) {
	params := map[string]int{"comic_id": id}
	var result goApi.CommonResponse[map[string]any]
	err := client.Delete("/api/content/delete-comic", params, &result)
	return &result, err
}

// AddComment 添加评论
func AddComment(client *goApi.Client, params AddCommentParams) (*goApi.CommonResponse[ResponseContent], error) {
	var result goApi.CommonResponse[ResponseContent]
	err := client.Post("/api/content/add-comment", params, &result)
	return &result, err
}

// GetRootComment 获取根评论
func GetRootComment(client *goApi.Client, params GetRootCommentParams) (*goApi.CommonResponse[RootComments], error) {
	var result goApi.CommonResponse[RootComments]
	err := client.Get("/api/content/get-root-comment", params, &result)
	return &result, err
}

// GetSubComment 获取子评论
func GetSubComment(client *goApi.Client, params GetSubCommentParams) (*goApi.CommonResponse[SubComments], error) {
	var result goApi.CommonResponse[SubComments]
	err := client.Get("/api/content/get-sub-comment", params, &result)
	return &result, err
}

// DeleteComment 删除评论
func DeleteComment(client *goApi.Client, commentID int) (*goApi.CommonResponse[map[string]any], error) {
	params := map[string]int{"comment_id": commentID}
	var result goApi.CommonResponse[map[string]any]
	err := client.Delete("/api/content/delete-comment", params, &result)
	return &result, err
}

// CollectContent 收藏内容
func CollectContent(client *goApi.Client, params ContentActionParams) (*goApi.CommonResponse[map[string]any], error) {
	var result goApi.CommonResponse[map[string]any]
	err := client.Post("/api/content/collect", params, &result)
	return &result, err
}

// LikeContent 点赞内容
func LikeContent(client *goApi.Client, params ContentActionParams) (*goApi.CommonResponse[map[string]any], error) {
	var result goApi.CommonResponse[map[string]any]
	err := client.Post("/api/content/like", params, &result)
	return &result, err
}

// UncollectContent 取消收藏
func UncollectContent(client *goApi.Client, params ContentActionParams) (*goApi.CommonResponse[map[string]any], error) {
	var result goApi.CommonResponse[map[string]any]
	err := client.Post("/api/content/uncollect", params, &result)
	return &result, err
}

// UnlikeContent 取消点赞
func UnlikeContent(client *goApi.Client, params ContentActionParams) (*goApi.CommonResponse[map[string]any], error) {
	var result goApi.CommonResponse[map[string]any]
	err := client.Post("/api/content/unlike", params, &result)
	return &result, err
}

// GetNewArticles 获取最新文章
func GetNewArticles(client *goApi.Client, params CursorPageParams) (*goApi.CommonResponse[ArticleList], error) {
	var result goApi.CommonResponse[ArticleList]
	err := client.Get("/api/content/get-new-articles", params, &result)
	return &result, err
}

// GetNewComics 获取最新漫画
func GetNewComics(client *goApi.Client, params CursorPageParams) (*goApi.CommonResponse[ComicList], error) {
	var result goApi.CommonResponse[ComicList]
	err := client.Get("/api/content/get-new-comics", params, &result)
	return &result, err
}

// GetNewPodcasts 获取最新播客
func GetNewPodcasts(client *goApi.Client, params CursorPageParams) (*goApi.CommonResponse[PodcastList], error) {
	var result goApi.CommonResponse[PodcastList]
	err := client.Get("/api/content/get-new-podcasts", params, &result)
	return &result, err
}

// GetNewVideos 获取最新视频
func GetNewVideos(client *goApi.Client, params CursorPageParams) (*goApi.CommonResponse[VideoList], error) {
	var result goApi.CommonResponse[VideoList]
	err := client.Get("/api/content/get-new-videos", params, &result)
	return &result, err
}
