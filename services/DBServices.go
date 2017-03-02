package services

import (
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"my-blog/models"
)

const DB_URL = "127.0.0.1:27017"
const DB_NAME = "blog"

var (
	session *mgo.Session
)

//开启数据库，得到session
func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(DB_URL)
		session.SetMode(mgo.Monotonic, true)
		if err != nil {
			beego.Error("db disconnect:", err)
		}
	}
	return session.Clone()
}

//选择数据库
func WitchCollection(collectionName string) *mgo.Collection {
	session := GetSession()
	defer session.Close()
	c := session.DB(DB_NAME).C(collectionName)
	return c
}

//增加博客
func AddNewBlog(blog models.Blog) {
	//c := WitchCollection("blogs")
	//session := GetSession()
	session, err := mgo.Dial(DB_URL)
	beego.Info("++++++++++")
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		beego.Info("db disconnect:", err)
	}
	defer session.Close()
	c := session.DB(DB_NAME).C("blogs")
	err = c.Insert(&blog)
	if err != nil {
		beego.Info("insert fail:", err)
	}
}

//目前所有的博客
func GetAllBlogs() []models.Blog {
	session, err := mgo.Dial(DB_URL)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		beego.Info("db disconnect:", err)
	}
	defer session.Close()
	c := session.DB(DB_NAME).C("blogs")
	var blogs []models.Blog
	c.Find(nil).All(&blogs)
	return blogs
}

//得到具体的博客
func GetBlog(uuid string) models.Blog {
	session, err := mgo.Dial(DB_URL)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		beego.Info("db disconnect:", err)
	}
	defer session.Close()
	c := session.DB(DB_NAME).C("blogs")
	var blog models.Blog
	c.Find(bson.M{"uuid": uuid}).One(&blog)
	return blog
}

func AddComment(blogid string, conment models.BlogComment) {
	session, err := mgo.Dial(DB_URL)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		beego.Info("db disconnect:", err)
	}
	defer session.Close()
	c := session.DB(DB_NAME).C("blogs")
	var blog models.Blog
	c.Find(bson.M{"uuid": blogid}).One(&blog)
	{
		comments := blog.Comment
		if comments == nil {
			comments = make([]models.BlogComment, 1)
		}
		comments = append(comments, conment)
		c.Update(bson.M{"uuid": blogid}, bson.M{"$set": bson.M{"comment": comments}})
	}
}

func AddContactMsg(contactMsg models.ContactMe)  {
	session, err := mgo.Dial(DB_URL)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		beego.Info("db disconnect:", err)
	}
	defer session.Close()
	c := session.DB(DB_NAME).C("contactmsg")
	err = c.Insert(&contactMsg)
	if err != nil {
		beego.Info("insert fail:", err)
	}
}