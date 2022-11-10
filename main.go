package main

func main() {

	srv := server.new(":8080")

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
