package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Comment struct {
	Name string
	Text string
}

type Post struct {
	Url       string
	Title     string
	Date      string
	Category  string
	Tag       []string
	Body      string
	ShortBody string
	Comments  []Comment
}

var posts []Post

func init() {
	p1 := Post{
		Url:      "why-passwords-matter",
		Title:    "Why Passwords Still Matter in 2004",
		Date:     "Tuesday, March 09, 2004",
		Category: "Security",
		Tag:      []string{"passwords", "security", "basics"},
		Body: `
		<p>I keep seeing people use "123456" and "password" as their passwords. Folks, we need to talk. With all the worms going around these days — Blaster, Sasser, MyDoom — a weak password is basically leaving your front door wide open with a neon sign saying "ROB ME."</p>
		<p>A good password should be at least 8 characters long (I'd go for 12 if you can remember it), mix upper and lowercase letters, throw in some numbers, and if the site allows it, a special character like ! or @. Don't use your dog's name. Don't use your birthday. Don't use the name of your favourite football team.</p>
		<p>I've started writing mine down on a piece of paper and keeping it in my desk drawer. Some security guys say that's a bad idea but I think the bigger threat right now is someone guessing "password123" from the internet, not someone breaking into my house specifically to steal my sticky note.</p>
		<p>Anyway. Stay safe out there. And patch your Windows — yes, I know Windows Update is annoying. Do it anyway.</p>
		`,
		ShortBody: `
		<p>After much deliberation I picked up a Linksys WRT54G from Best Buy for $79. The guy at the store tried to sell me the $120 Netgear one but honestly I didn't see the point. This thing does 802.11g which is supposed to be five times faster than my old 802.11b setup. Five times!</p>
		`,
		Comments: []Comment{
			{"Labib", "No CMNT"},
			{"Anon-01", "NICE, ATTRACTIVE!!!"},
			{"Anon-02", "Fahhhhhhh!!"},
			{"Nasir P", "CHANDABAZ....."},
		},
	}

	p2 := Post{
		Url:      "my-new-router",
		Title:    "I Finally Upgraded My Router",
		Date:     "Saturday, February 21, 2004",
		Category: "Hardware",
		Tag:      []string{"router", "wifi", "networking", "linksys"},
		Body: `
        <p>After much deliberation I picked up a Linksys WRT54G from Best Buy for $79. The guy at the store tried to sell me the $120 Netgear one but honestly I didn't see the point. This thing does 802.11g which is supposed to be five times faster than my old 802.11b setup. Five times!</p>
        <p>Setting up WEP encryption was a bit of a pain — the interface is not exactly user-friendly — but I got there in the end. You have to type in a hex key which is pretty intimidating the first time. I used a random one I found on a website that generates them for you.</p>
        <p>The signal now reaches the garage which means I can finally use the laptop out there while I work on the car. My wife is less thrilled because she says I'll never come inside now. She's probably right.</p>
        <p>One thing I'll say: the documentation is terrible. I spent two hours trying to figure out why the connection kept dropping before I realised I had the MTU set wrong. There's a forum thread about it at Broadbandreports.com if you run into the same issue.</p>
      	`,
		ShortBody: `
        <p>After much deliberation I picked up a Linksys WRT54G from Best Buy for $79. The guy at the store tried to sell me the $120 Netgear one but honestly I didn't see the point. This thing does 802.11g which is supposed to be five times faster than my old 802.11b setup. Five times!</p>
        `,
	}

	p3 := Post{
		Url:      "ie6-vs-firefox",
		Title:    "IE6 vs This New Firefox Thing",
		Date:     "Monday, January 05, 2004",
		Category: "Software",
		Tag:      []string{"browser", "firefox", "ie6", "internet"},
		Body: `
        <p>My buddy Mike has been going on about Firefox 0.8 for weeks now. I've been a loyal Internet Explorer user since IE4 so I was skeptical. But he finally convinced me to try it and I have to say... it's pretty good actually.</p>
        <p>The built-in popup blocker alone is worth it. The internet has gotten absolutely overrun with those pop-up ads lately. Half the time I open a website I get five windows launching at once, some of them pretending to be Windows error messages trying to get you to install stuff. Very annoying.</p>
        <p>Firefox also seems faster loading pages, though that might just be placebo. And I like that the tabs are built right into the browser — I know IE has that with third-party things like Crazy Browser but having it built in feels cleaner.</p>
        <p>That said, some websites just plain don't work right in Firefox. Banking sites mostly. And some sites actually pop up a message saying "This site requires Internet Explorer." So I'm keeping both installed for now. IE6 for the stubborn sites, Firefox for general browsing.</p>
      	`,
		ShortBody: `
        <p>My buddy Mike has been going on about Firefox 0.8 for weeks now. I've been a loyal Internet Explorer user since IE4 so I was skeptical. But he finally convinced me to try it and I have to say... it's pretty good actually.</p>
        `,
	}

	posts = append(posts, []Post{p1, p2, p3}...)
}

// {{add $i 1}} — used for 1-based comment numbering
var funcMap = template.FuncMap{
	"add": func(a, b int) int { return a + b },
}
var homeTmpl = template.Must(
	template.New("blog.html").Funcs(funcMap).ParseFiles("templates/blog.html"),
)
var postTmpl = template.Must(
	template.New("post.html").Funcs(funcMap).ParseFiles("templates/post.html"),
)
var SearchQueryV string = "SearchQuery"
var PostsV string = "Posts"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./frontend/blog.html")
	if r.Method != http.MethodGet {
		http.Error(w, "Fahhhh", 400)
		return
	}

	searchItem := r.URL.Query().Get("search")
	fmt.Println("/", r.URL.RawPath, r.Method, searchItem)

	pageContent := map[string]any{
		SearchQueryV: template.HTML(searchItem),
		PostsV:       posts,
	}

	homeTmpl.Execute(
		w,
		pageContent,
	)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	postUrl := r.URL.Query().Get("blog")
	fmt.Println("/post", r.URL.RawPath, r.Method, postUrl)

	if r.Method == http.MethodGet {

		var found Post
		for _, p := range posts {
			if p.Url == postUrl {
				found = p
				break
			}
		}

		pageContent := map[string]any{
			"Post":     found,
			"Comments": found.Comments,
		}

		postTmpl.Execute(
			w,
			pageContent,
		)

	} else if r.Method == http.MethodPost {

	} else {
		http.Error(w, "Fahh...", 400)
	}
}

func main() {
	_ = funcMap

	fs := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/post", postHandler)

	log.Println("Server running at http://127.0.0.1:8080")
	http.ListenAndServe(":8080", mux)
}
