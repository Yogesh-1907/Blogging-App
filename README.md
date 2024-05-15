A blogging application that supports the pagination of blog posts needs to be developed so that users can view a limited number of posts at a time.

Features :

Retrieves all blog posts from the database.
Retrieves and displays paginated blog posts.

Technologies Used :

Golang: The backend of the application is developed using the Go programming language.
Gin Framework: Gin is a web framework for Go that provides routing, middleware, and more for building web applications.
GORM ORM: GORM is an ORM library for Go that simplifies database operations and provides a convenient way to interact with the database.
Middleware: The application uses middleware like Helmet for securing HTTP headers and CORS for handling Cross-Origin Resource Sharing.

Endpoints
The application exposes the following endpoints:
/posts: Retrieves all blog posts.
/posts/paginated: Retrieves and displays paginated blog posts.
/posts/paginated-by-title: Retrieves and displays paginated blog posts order by title.
/posts/paginated-by-author-id: Retrieves and displays paginated blog posts order by authorId.


Clone repository -> Add dependency using go mod tidy -> Add your database related fields in DSN -> execute table script in database -> run main.go file and then consume_api.go file
main.go files run on 8080 port and consume_api.go which consumes API from 8080 runs on 8081 port.
