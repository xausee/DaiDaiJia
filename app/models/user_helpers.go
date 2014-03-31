package models

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"time"
)

func (manager *DbManager) GetUserById(userid int) (userInfo User, err error) {
	uc := manager.session.DB(DbName).C(UserCollection)

	err = uc.Find(bson.M{"id": userid}).One(&userInfo)
	if err != nil {
		fmt.Println("查询用户信息失败")
	}

	return userInfo, err
}

func (manager *DbManager) UpdateUserInfo(userid int, newUserInfo User) (err error) {
	uc := manager.session.DB(DbName).C(UserCollection)

	var oldUserInfo User
	err = uc.Find(bson.M{"id": userid}).One(&oldUserInfo)
	if err != nil {
		fmt.Println("查询用户信息失败")
	}
	// 修改一些基本的信息，并不是全部，参看修改页面的内容
	tempInfo := oldUserInfo
	tempInfo.AvatarUrl = newUserInfo.AvatarUrl
	tempInfo.NickName = newUserInfo.NickName
	tempInfo.PenName = newUserInfo.PenName
	tempInfo.Gender = newUserInfo.Gender
	tempInfo.Email = newUserInfo.Email
	tempInfo.Birth = newUserInfo.Birth
	tempInfo.FavoriteAuthor = newUserInfo.FavoriteAuthor
	tempInfo.FavoriteBook = newUserInfo.FavoriteBook
	tempInfo.Intrest = newUserInfo.Intrest
	tempInfo.Introduction = newUserInfo.Introduction

	err = uc.Update(oldUserInfo, tempInfo)
	if err != nil {
		fmt.Println("修改用户信息失败")
	}

	return err
}

func (manager *DbManager) AddUserArticle(article *UserArticle) error {
	uc := manager.session.DB(DbName).C(UserCollection)

	// 根据文章作者的ID, 查找作者的信息
	var oldUserInfo User
	err := uc.Find(bson.M{"id": article.AuthorId}).One(&oldUserInfo)
	tmpUser := oldUserInfo

	// 给新文章创建ID和创作日期并格式化
	article.Id = bson.NewObjectId().Hex()
	article.CreateTime = time.Now().Format("2006-01-02 15:04:05")

	// 追加文章到已有的集合
	as := oldUserInfo.Articles
	as = append(as, *article)
	tmpUser.Articles = as

	// 更新整个用户信息，包括新加的文章
	err = uc.Update(oldUserInfo, tmpUser)
	if err != nil {
		fmt.Println("新增文章失败")
	}

	return err
}

func (manager *DbManager) UpdateUserArticle(authorid int, newAritlce UserArticle) error {
	uc := manager.session.DB(DbName).C(UserCollection)

	// 根据文章作者的ID, 查找作者的信息
	var oldUserInfo, tmpUser User
	err := uc.Find(bson.M{"id": authorid}).One(&oldUserInfo)
	err = uc.Find(bson.M{"id": authorid}).One(&tmpUser)

	flag, index := false, 0
	for _, art := range oldUserInfo.Articles {
		if art.Id == newAritlce.Id {
			flag = true
			fmt.Println("找到指定的文章")
			break
		}
		index += 1
	}

	if flag {
		// 更新指定的文章的类型，标题和内容
		as := tmpUser.Articles
		as[index].Tag = newAritlce.Tag
		as[index].Title = newAritlce.Title
		as[index].Content = newAritlce.Content

		// 更新整个用户信息
		err = uc.Update(oldUserInfo, tmpUser)
		if err != nil {
			fmt.Println("更新失败")
		}
	}

	return err
}

func (manager *DbManager) AddArticleComment(article UserArticle, comment *Comment) error {
	uc := manager.session.DB(DbName).C(UserCollection)

	// 根据文章作者的ID, 查找作者的信息
	var oldUserInfo, tmpUser User
	err := uc.Find(bson.M{"id": article.AuthorId}).One(&oldUserInfo)
	err = uc.Find(bson.M{"id": article.AuthorId}).One(&tmpUser)

	// 给新评论创建ID，增加日期并格式化
	comment.Id = bson.NewObjectId().Hex()
	comment.Time = time.Now().Format("2006-01-02 15:04:05")

	// 追加评论到已有的集合
	cs := article.Comments
	cs = append(cs, *comment)
	article.Comments = cs

	flag, index := false, 0
	for _, art := range oldUserInfo.Articles {
		if art.Id == article.Id {
			flag = true
			fmt.Println("找到指定的文章")
			break
		}
		index += 1
	}

	if flag {
		// 更新该篇文章
		tmpUser.Articles[index] = article
	} else {
		fmt.Println("找不到指定的文章， 添加评论失败")
	}

	// 更新整个用户信息，包括新加的文章
	err = uc.Update(oldUserInfo, tmpUser)
	if err != nil {
		fmt.Println("新增文章失败")
	}

	return err
}

func (manager *DbManager) GetAllArticlesByUserId(userid int) (articles []UserArticle, err error) {
	uc := manager.session.DB(DbName).C(UserCollection)

	var userInfo User
	err = uc.Find(bson.M{"id": userid}).One(&userInfo)
	if err != nil {
		fmt.Println("查询用户信息失败")
	}
	articles = userInfo.Articles

	return articles, err
}

func (manager *DbManager) GetAllMessageByUserId(userid int) (message []Comment, err error) {
	uc := manager.session.DB(DbName).C(UserCollection)

	var userInfo User
	err = uc.Find(bson.M{"id": userid}).One(&userInfo)
	if err != nil {
		fmt.Println("查询用户信息失败")
	}
	message = userInfo.Message

	return message, err
}

func (manager *DbManager) GetArticleByUserIdAndArticleId(userid int, articleid string) (article UserArticle, err error) {
	articles, _ := manager.GetAllArticlesByUserId(userid)

	flag := false
	for _, art := range articles {
		if art.Id == articleid {
			article = art
			flag = true
			fmt.Println("找到指定的文章")
			return article, err
		}
	}

	if !flag {
		fmt.Println("找到指定的文章")
	}

	return article, err
}

func (manager *DbManager) DeleteUserArticle(userid int, articleid string) error {
	uc := manager.session.DB(DbName).C(UserCollection)

	// 根据文章作者的ID, 查找作者的信息
	var oldUserInfo User
	err := uc.Find(bson.M{"id": userid}).One(&oldUserInfo)
	tmpUser := oldUserInfo

	flag, index := false, 0
	for _, art := range oldUserInfo.Articles {
		if art.Id == articleid {
			flag = true
			fmt.Println("找到指定的文章")
			break
		}
		index += 1
	}

	if flag {
		// 删除指定的文章
		as := oldUserInfo.Articles
		as = append(as[:index], as[index+1:]...)

		// 更新整个用户信息，包括新加的文章
		tmpUser.Articles = as
		err = uc.Update(oldUserInfo, tmpUser)
		if err != nil {
			fmt.Println("更新失败")
		}
	}

	return err
}
