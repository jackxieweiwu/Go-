package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

const (
	ListDir      = 0x0001
	STATIC_DIR   = ".\\public\\"
	UPLOAD_DIR   = ".\\uploads"
	TEMPLATE_DIR = ".\\views"
)

var templates map[string]*template.Template = make(map[string]*template.Template)

func init() {
	log.Println("============INIT==============")

	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)

	var templateName, templatePath string

	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		log.Println("templateName:", templateName)

		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}

		templatePath = TEMPLATE_DIR + "\\" + templateName
		log.Println("Loading template:", templatePath)

		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}

	for key, _ := range templates {
		log.Println("template key:", key)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	err := templates[tmpl].Execute(w, locals)
	check(err)
}

func isExists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	return os.IsExist(err)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, "upload.html", nil)
	} else if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)

		filename := h.Filename
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + "\\" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer t.Close()

		_, err = io.Copy(t, f)
		check(err)

		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "\\" + imageId

	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}

	var listHtml string
	listHtml = "<html><head><title>美女图片</title></head><body><center><h1>美女</h1><img src=\"/img?id=" + imageId + "\" /></center></body></html>"
	io.WriteString(w, listHtml)

	//w.Header().Set("Content-Type", "image")
	//http.ServeFile(w, r, imagePath)
}
func imgHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "\\" + imageId

	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	check(err)

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	log.Println("Images:", images)

	locals["images"] = images
	renderHtml(w, "list.html", locals)
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)

				//或者输出自定义的50X错误页面
				//w.WriteHeader(http.StatusInternalServerError)
				//renderHtml(w,"error",e)

				//login
				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()

		fn(w, r)
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]

		if (flags & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", STATIC_DIR, 0)
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	mux.HandleFunc("/img", safeHandler(imgHandler))

	err := http.ListenAndServe(":8088", mux)

	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
