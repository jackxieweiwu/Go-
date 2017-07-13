package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

const (
	UPLOAD_DIR = ".\\uploads"
)

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var listHtml string

	for _, fileInfo := range fileInfoArr {
		imgid := fileInfo.Name()
		listHtml += "<li><a href=\"/view?id=" + imgid + "\">" + imgid + "</a></li>"
	}
	log.Printf("list:", listHtml)
	io.WriteString(w, "<html><body><ol>"+listHtml+"</ol></body></html>")

}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "\\" + imageId
	log.Println("imagePath:", imagePath)

	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w,
			`<html>
				<body>
					<form method="POST" action="/upload" enctype="multipart/form-data">
						Choose an image to upload: <input name="image" type="file" /><input type="submit" value="Upload" />
					</form>
				</body>
			</html>`)

		return
	} else if r.Method == "POST" {

		//调用r.FormFile()方法会返回3个值，各个值的类型分别是multipart.File，*multipart.FileHeader和error
		f, h, err := r.FormFile("image") //对应与<input name="image" type="file" />中的name
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := h.Filename
		log.Println("filename:", filename)
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + "\\" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer t.Close()

		if _, err = io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("StatusFound:", http.StatusFound)
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
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

func main() {
	//http.HandleFunc("/", listHandler)
	//http.HandleFunc("/view", viewHandler)
	//http.HandleFunc("/upload", uploadHandler)

	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
