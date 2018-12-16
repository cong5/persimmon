package home

import (
	"bytes"
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/modules/myxmlrpc"
	"github.com/pkg/errors"
	"github.com/revel/revel"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type XmlRpcServer struct {
	AppController
}

var routine = map[string]interface{}{
	"blogger.getUsersBlogs":     getUsersBlogs,
	"metaWeblog.getPost":        getPost,
	"metaWeblog.getRecentPosts": getRecentPosts,
	"metaWeblog.newPost":        newPost,
	"metaWeblog.editPost":       editPost,
	"metaWeblog.deletePost":     deletePost,
	"metaWeblog.newMediaObject": newMediaObject,
	"metaWeblog.getCategories":  getCategories,
	"metaWeblog.getTemplate":    getTemplate,
	"metaWeblog.setTemplate":    setTemplate,
}

func (p *XmlRpcServer) ShowMessage() revel.Result {
	res := models.Res{
		Id:     0,
		Status: 200,
		List:   "",
		Item:   "",
		Info:   "XML-RPC 服务不支持 GET 请求，请使用 POST 请求",
	}
	return p.RenderJSON(res)
}

func (p *XmlRpcServer) Index() revel.Result {
	body, err := ioutil.ReadAll(p.Request.GetBody())
	if err != nil {
		log.Println("Failed to read the request body: " + err.Error())

		p := myxmlrpc.EncodeFault(err.Error())
		return myxmlrpc.ResXml(p.Bytes())
	}
	//log.Println("Request Body:\n" + string(body))

	reader := bytes.NewReader(body)
	v, err := myxmlrpc.DecodeRequest(reader)
	if err != nil {
		log.Println("Failed to parse request, " + err.Error())

		p := myxmlrpc.EncodeFault(err.Error())
		return myxmlrpc.ResXml(p.Bytes())
	}

	//find the handler to process.
	f := routine[v.Name]
	w := p.Response.GetWriter()

	if f != nil {
		//Call by reflect
		fn := reflect.ValueOf(f)
		in := make([]reflect.Value, 2)
		in[0] = reflect.ValueOf(w)
		in[1] = reflect.ValueOf(v)
		fn.Call(in)
	} else {
		log.Println("Handler not found for message: " + v.Name)

		p := myxmlrpc.EncodeFault(errors.New("Handler not found for message: " + v.Name))
		return myxmlrpc.ResXml(p.Bytes())
	}

	return myxmlrpc.ResXml("")
}

func getUsersBlogs(w http.ResponseWriter, r *myxmlrpc.MethodRequest) {
	var blogs []myxmlrpc.BlogInfo
	var BlogName = ""
	var BlogURL = ""
	allOption, _ := optionService.GetAllOption(false)
	for _, v := range allOption {
		if v.Name == "site_name" {
			BlogName = v.Value
		} else if v.Name == "site_url" {
			BlogURL = v.Value
		}

		if BlogName != "" && BlogURL != "" {
			break
		}
	}
	blogs = append(blogs, myxmlrpc.BlogInfo{BlogId: "1", URL: BlogURL, BlogName: BlogName})
	p := myxmlrpc.EncodeResponse(blogs)
	log.Println("Sent The Response:\n" + string(p.Bytes()))

	w.Write(p.Bytes())
}

func getPost(w http.ResponseWriter, r *myxmlrpc.MethodRequest) {
	//blogId := r.GetParameter(0).(string)
	//username := r.GetParameter(1)
	//password := r.GetParameter(2)

	/*
	req := r.GetParameter(3)
	godump.Dump(req)
	*/
	/*
	post, err := postService.GetPostById()
	if slug == "" || err != nil || post.Id <= 0 {
		return c.NotFound("很抱歉，没有找到这个页面.")
	}

	dataTime := time.Unix(post.CreatedAt, 0)
	obj := myxmlrpc.Post{DateCreated: dataTime.Local(), Description: post.Content, Title: post.Title}
	p := myxmlrpc.EncodeResponse(obj)
	log.Println("Sent The Response:\n" + string(p.Bytes()))

	w.Write(p.Bytes())
	*/
}

func getRecentPosts(w http.ResponseWriter, r *myxmlrpc.MethodRequest) {
	var posts []myxmlrpc.Post
	p := myxmlrpc.EncodeResponse(posts)
	log.Println("Sent The Response:\n" + string(p.Bytes()))

	w.Write(p.Bytes())
}

func newPost(w http.ResponseWriter, r *myxmlrpc.MethodRequest) {
	//blogId := r.GetParameter(0).(string)
	//username := r.GetParameter(1)
	//password := r.GetParameter(2)
	//post := r.GetParameter(3).(myxmlrpc.Struct)
	/*
		title := post["title"].(string)
		content := post["description"].(string)

		//publish := r.GetParameter(4).(bool)

		year, month, day := time.Now().Date()

		date := fmt.Sprintf("%d-%02d-%02d", year, month, day)
		path := "/_posts/" + date + "-" + title + ".html"
	*/
	//todo: stash before write and commit

	//write the post to file under the blog folder.
	//err := writePost(path, title, content)
	/*
	err := nil
	if err != nil {
		log.Panicf("Failed to write file: %s, error=%s", path, err)
		p := myxmlrpc.EncodeFault(err.Error())
		w.Write([]byte(p.Bytes()))
		log.Println("Sent The Response:\n" + string(p.Bytes()))

		return
	}

	uri := fmt.Sprintf("%d/%02d/%s", year, month, title+".html")
	*/

	/*req := r.GetParameter(3)
	godump.Dump(req)
	req1 := req.(myxmlrpc.Struct)
	godump.Dump(req1)*/
	/*
	for k, v := range req1 {
		log.Printf("k: %s, v: %s", k, v)
	}*/

	/*
	p := myxmlrpc.EncodeResponse("")
	log.Println("Sent The Response:\n" + string(p.Bytes()))
	*/
	w.Write([]byte(""))
}

func editPost(w http.ResponseWriter, r *myxmlrpc.MethodRequest) {
	newPost(w, r)
}

func deletePost(w http.ResponseWriter, r *myxmlrpc.MethodRequest) {
	log.Println("This Methos Is Not Implemented.")
}

func newMediaObject(w http.ResponseWriter, r *myxmlrpc.MethodRequest) {
	//blogid := r.GetParameter(0).(string)
	//username := r.GetParameter(1)
	//password := r.GetParameter(2)
	media := r.GetParameter(3).(myxmlrpc.Struct)

	name := media["name"].(string)
	//t := media["type"].(string)
	bits := media["bits"].([]byte)

	//prefix := "/assets/posts/"
	//remove the "Open-Live-Writer/" from the name
	name = strings.TrimLeft(name, "Open-Live-Writer/")

	//write file
	path := ""
	err := ioutil.WriteFile(path, bits, os.ModeAppend)
	if err != nil {
		p := myxmlrpc.EncodeFault(err)
		log.Println("Sent The Response:\n" + string(p.Bytes()))

		w.Write(p.Bytes())
		return
	}

	//write the media data to file under the ${blogId} folder.
	obj := myxmlrpc.MediaObjectUrl{URL: ""}
	p := myxmlrpc.EncodeResponse(obj)
	log.Println("Sent The Response:\n" + string(p.Bytes()))

	w.Write(p.Bytes())
}

func getCategories(w http.ResponseWriter, r *myxmlrpc.MethodRequest) {
	var categories []myxmlrpc.CategoryInfo
	category, _ := categoryService.GetList(999, 1, false)
	for i := 0; i < len(category); i++ {
		categories = append(categories, myxmlrpc.CategoryInfo{Description: category[i].Name, Title: category[i].Name, CategoryId: string(category[i].Id)})
	}

	p := myxmlrpc.EncodeResponse(categories)
	log.Println("Sent The Response:\n" + string(p.Bytes()))

	w.Write(p.Bytes())
}

func getTemplate(w http.ResponseWriter, r *myxmlrpc.MethodRequest) {
	log.Println("This Methos Is Not Implemented.")
}

func setTemplate(w http.ResponseWriter, r *myxmlrpc.MethodRequest) {
	log.Println("This Methos Is Not Implemented.")
}
