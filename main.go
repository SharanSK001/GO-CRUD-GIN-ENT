package main

import (
	"context"
	"log"
	"net/http"

	"entdemo/ent"
	"entdemo/ent/blog"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	db *ent.Client
}

var svr Server

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseJSON(c *gin.Context, httpCode, errCode int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, 400
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, 500
	}
	if !check {
		return http.StatusBadRequest, 400
	}

	return http.StatusOK, 200
}
func connectdb() {
	//Postgres Initiation
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=blog password=1505 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	svr.db = client
	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

// Create API for Blog Creation
// Gets Input from the user
// Returns Response what is inserted
func CreateBlog(c *gin.Context) {
	type createbloginput struct {
		Name        string `json:"name"`
		Title       string `json:"Title"`
		Email       string `json:"email"`
		Phonenumber int    `json:"number"`
		BlogContent string `json:"content"`
		Githublink  string `json:"link"`
	}
	var input createbloginput
	httpCode, errCode := BindAndValid(c, &input)
	if errCode != 200 {
		ResponseJSON(c, httpCode, errCode, "invalid param", nil)
		return
	}
	//Create Blog API
	b, err := svr.db.Blog.
		Create().
		SetName(input.Name).
		SetTitle(input.Title).
		SetEmail(input.Email).
		SetPhonenumber(input.Phonenumber).
		SetBlogcontent(input.BlogContent).
		SetGithublink(input.Githublink).
		Save(context.Background())
	if err != nil {
		ResponseJSON(c, http.StatusOK, 500, "create user failed: "+err.Error(), nil)
		return
	}
	type ResponseData struct {
		Name        string `json:"name"`
		Title       string `json:"title"`
		Email       string `json:"email"`
		Phonenumber int    `json:"number"`
		Blogcontent string `json:"content"`
		Githublink  string `json:"link"`
	}
	var resp ResponseData
	resp.Name = input.Name
	resp.Title = input.Title
	resp.Email = input.Email
	resp.Phonenumber = input.Phonenumber
	resp.Blogcontent = input.BlogContent
	resp.Githublink = b.Githublink
	ResponseJSON(c, http.StatusOK, 200, "Success", resp)

}
func CreateUser(c *gin.Context) {
	type createuserinput struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	var input createuserinput
	httpCode, errCode := BindAndValid(c, &input)
	if errCode != 200 {
		ResponseJSON(c, httpCode, errCode, "invalid param", nil)
		return
	}
	inp, err := svr.db.User.
		Create().
		SetUsername(input.UserName).
		SetPassword(input.Password).
		Save(context.Background())
	if err != nil {
		ResponseJSON(c, http.StatusOK, 500, "create user failed: "+err.Error(), nil)
		return
	}
	type ResponseData struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	var resp ResponseData
	resp.UserName = inp.Username
	resp.Password = inp.Password
	ResponseJSON(c, http.StatusOK, 200, "Success", resp)

}

// Get a Blog API
// Gets Title of the blog
// Responds with the details of the Blog
func GetBlog(c *gin.Context) {
	Title := c.Param("title")
	if len(Title) == 0 {
		ResponseJSON(c, 200, 400, "invalid param", nil)
	}
	//Get Param is Title
	titles, _ := svr.db.Blog.
		Query().
		Where(blog.Title(Title)).
		First(context.Background())
	if titles == nil {
		ResponseJSON(c, http.StatusOK, 500, "Title doesn't exist", nil)
		return
	}
	type ResponseData struct {
		Name        string `json:"name"`
		Title       string `json:"title"`
		Email       string `json:"email"`
		Phonenumber int    `json:"number"`
		Blogcontent string `json:"content"`
		Githublink  string `json:"link"`
	}
	var resp ResponseData
	resp.Name = titles.Name
	resp.Title = titles.Title
	resp.Email = titles.Email
	resp.Phonenumber = titles.Phonenumber
	resp.Blogcontent = titles.Blogcontent
	resp.Githublink = titles.Githublink
	ResponseJSON(c, http.StatusOK, 200, "Success", resp)

}

// Update API to update any fields in the blog
// Responds with a successfull Message after Updating
func UpdateBlog(c *gin.Context) {
	type updatebloginput struct {
		Name        string `json:"name"`
		Title       string `json:"Title"`
		Email       string `json:"email"`
		Phonenumber int    `json:"number"`
		BlogContent string `json:"content"`
		Githublink  string `json:"link"`
	}
	var input updatebloginput
	httpCode, errCode := BindAndValid(c, &input)
	if errCode != 200 {
		ResponseJSON(c, httpCode, errCode, "invalid param", nil)
		return
	}
	//Update API
	updates, _ := svr.db.Blog.
		Update().
		SetName(input.Name).
		SetTitle(input.Title).
		SetEmail(input.Email).
		SetPhonenumber(input.Phonenumber).
		SetBlogcontent(input.BlogContent).
		SetGithublink(input.Githublink).
		Save(context.Background())
	if updates == 0 {
		ResponseJSON(c, http.StatusOK, 500, "update blog failed", nil)
		return
	}
	upd, _ := svr.db.Blog.
		Query().
		Where(blog.Title(input.Title)).
		First(context.Background())

	if upd == nil {
		ResponseJSON(c, http.StatusOK, 500, "blog doesn't exist", nil)
		return
	}
	type ResponseData struct {
		Name        string `json:"name"`
		Title       string `json:"title"`
		Email       string `json:"email"`
		Phonenumber int    `json:"number"`
		Blogcontent string `json:"content"`
		Githublink  string `json:"link"`
	}
	var resp ResponseData
	resp.Name = input.Name
	resp.Title = input.Title
	resp.Email = input.Email
	resp.Phonenumber = input.Phonenumber
	resp.Blogcontent = input.BlogContent
	resp.Githublink = input.Githublink

	ResponseJSON(c, http.StatusOK, 200, "Success", resp)

}

// Delete API
// Deletes Blog with corresponding Title
func DeleteBlog(c *gin.Context) {
	Title := c.Param("title")
	if len(Title) == 0 {
		ResponseJSON(c, 200, 400, "Invalid", nil)
		return
	}

	_, err := svr.db.Blog.
		Delete().
		Where(blog.Title(Title)).
		Exec(context.Background())
	if err != nil {
		ResponseJSON(c, 200, 500, "Delete Blog Failed", nil)
		return
	}

	ResponseJSON(c, 200, 200, "Successfully Deleted", nil)
}

func main() {
	//connecting DB
	connectdb()
	// Establishing GIN
	r := gin.Default()
	//Create a blog
	//r.POST("/blog/create",CreateBlog)
	r.POST("/blog/create", gin.BasicAuth(gin.Accounts{"admin": "admin"}), CreateBlog)
	// Create a user
	r.POST("/blog/create/user", CreateUser)
	//Read a Blog
	r.GET("/blog/:title", GetBlog)
	//Update a Blog
	//r.POST("blog/update",UpdateBlog)
	r.POST("blog/update", gin.BasicAuth(gin.Accounts{"admin": "admin"}), UpdateBlog)
	//Delete a Blog
	//r.DELETE("blog/:title",DeleteBlog)
	r.DELETE("blog/:title", gin.BasicAuth(gin.Accounts{"admin": "admin"}), DeleteBlog)
	//Running it on PORT 8080
	r.Run(":8080")
}
