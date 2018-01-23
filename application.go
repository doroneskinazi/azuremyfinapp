package main

import (

	// "github.com/heroku/myfinapp/controllers"

	"fmt"

	"github.com/heroku/myfinapp/models"
	"github.com/kataras/iris"
	_ "github.com/lib/pq"
)

// var db *sql.DB

func transEntryHandler(ctx iris.Context) {
	var t models.MoneyTrans
	trans, err := t.GetAllTrans()
	if err != nil {
		fmt.Println("err: ", err)
	}
	ctx.ViewData("MoneyTrans", trans)
	ctx.ViewData("TransCategory", models.Lists.TransCategories)
	ctx.View("transEntry.html")
}

func transListHandler(ctx iris.Context) {
	var t models.MoneyTrans
	trans, err := t.GetAllTrans()
	if err != nil {
		fmt.Println("err: ", err)
	}
	ctx.ViewData("MoneyTrans", trans)
	ctx.View("transList.html")

}

func main() {
	app := iris.New()

	// serve our app in public, public folder
	// contains the client-side vue.js application,
	// no need for any server-side template here,
	// actually if you're going to just use vue without any
	// back-end services, you can just stop afer this line and start the server.
	app.StaticWeb("/", "./public")
	tmpl := iris.HTML("./public/views", ".html")
	tmpl.Reload(true) // reload templates on each request (development mode)
	app.RegisterView(tmpl)
	models.ConnectDB()
	// if !models.connectDB() {
	// 	fmt.Println("Unable to connect to database")
	// }

	// dbinfo := fmt.Sprintf("port=5432 user=%s password=%s dbname=%s sslmode=disable", "postgres", "andy", "postgres")
	// var err error
	// db, err = sql.Open("postgres", dbinfo)
	// if err != nil {
	// 	fmt.Printf("err: %+v ", err)
	// }

	// fmt.Println("err, Connection: ", err, db)

	// defer db.Close()

	// configure the http sessions.
	// sess := sessions.New(sessions.Config{
	// 	Cookie: "iris_session",
	// })

	// configure the websocket server.
	// ws := websocket.New(websocket.Config{})

	// create a sub router and register the client-side library for the iris websockets,
	// you could skip it but iris websockets supports socket.io-like API.
	// todosRouter := app.Party("/todos")
	// http://localhost:8080/todos/iris-ws.js
	// serve the javascript client library to communicate with
	// the iris high level websocket event system.
	// myFinApp.Any("/iris-ws.js", websocket.ClientHandler())

	// create our mvc application targeted to /todos relative sub path.
	// myFinApp := mvc.New(myFinAppRouter)

	// any dependencies bindings here...
	// myFinApp.Register(
	// todo.NewMemoryService(),
	// 	sess.Start,
	// 	ws.Upgrade,
	// )

	app.Get("/transEntry", transEntryHandler)
	app.Get("/transList", transListHandler)

	// controllers registration here...
	// todosApp.Handle(new(controllers.TodoController))
	// port := os.Getenv("PORT")

	port := "8080"

	// start the web server at http://localhost:8080
	app.Run(iris.Addr(":"+port), iris.WithoutVersionChecker)

}
