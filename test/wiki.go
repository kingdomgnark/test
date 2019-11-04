package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "hireme123"
  dbname   = "postgres"
) //postgresql server info

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//Call to ParseForm makes form fields available.
    err := r.ParseForm()
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	address := r.PostFormValue("address")
	city := r.PostFormValue("city")
	state := r.PostFormValue("state")
	zip := r.PostFormValue("zip")
	
	http.Redirect(w, r, "/view/"+zip, http.StatusFound)
	          
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

  
	sqlStatement := `
	INSERT INTO addresses (Address, City, State, Zip)
	VALUES ($1, $2, $3, $4)`
	id := 0
	err = db.QueryRow(sqlStatement, address, city, state, zip).Scan(&id)
	//post to postgresql server
	
	sqlDeleted := `
	delete from cacheweather where ti<NOW()-interval 30 minute`
	
	sqlScan := `SELECT ZIP FROM cacheweather WHERE ZIP = zip`
    err := db.QueryRow(sqlStmt, zip).Scan(&zip)
    if err != nil {
        if err != sql.ErrNoRows {
            // a real error happened! you should change your function return
            // to "(bool, error)" and return "false, err" here
            log.Print(err)
        }

        s := strconv.Itoa(zip)
		u, err := url.Parse("api.openweathermap.org")
		if err != nil {
		log.Fatal(err)
		}
		u.Scheme = "http"
		u.Host = "api.openweathermap.org"
		u.Path = "data/2.5/forecast"
		q := u.Query()
		q.Set("zip", s)
		q.Set("APPID","12ea4c57a813ac94155a0cabe7788b15")
		q.set("units","imperial")
		u.RawQuery = q.Encode()
		// insert get request and store json return as type and in sql
		// use json to create new html file
		//forward user to new html file
		ut, err := url.Parse("api.openweathermap.org")
		if err != nil {
		log.Fatal(err)
		}
		ut.Scheme = "http"
		ut.Host = "api.openweathermap.org"
		ut.Path = "data/2.5/weather"
		q := ut.Query()
		q.Set("zip", s)
		q.Set("APPID","12ea4c57a813ac94155a0cabe7788b15")
		q.set("units","imperial")
		ut.RawQuery = q.Encode()

    }

    //Forward the user to view the cached html file
	
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}


func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	
	log.Fatal(http.ListenAndServe(":8080", nil))

}