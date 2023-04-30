package main

import (
	"github.com/gin-gonic/gin"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
	"net/http"
	"log"
	cors "github.com/rs/cors/wrapper/gin"
)

type Component struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    ImageURL string `json:"image_url"`
    Likes    int    `json:"likes"`
}

func main() {
	 // Create a new router
	 r := gin.Default()

	 // Add CORS middleware
	 cors := cors.New(cors.Options{
		 AllowedOrigins: []string{"*"}, // replace "*" with your frontend domain
		 AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		 AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		 Debug:          false,
	 })

	 r.Use(cors)
 
	 components, err := getComponents()
	 if err != nil {
		fmt.Println(err)
	 }
	 	
	 // Add route handlers
	 r.GET("/sample-data", func(c *gin.Context) {
		c.JSON(http.StatusOK, components)
	})
 
	 // Start the server
	 log.Println("Starting server on port 8080")
	 log.Fatal(http.ListenAndServe(":8080", r))
}

func sampleDataHandler(w http.ResponseWriter, r *http.Request) {
    // Set the content type header to JSON
    w.Header().Set("Content-Type", "application/json")

	components, err := getComponents()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(components)
	}
    // Encode the components slice to JSON and write it to the response
    err = json.NewEncoder(w).Encode(components)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func getComponents() ([]Component, error) {
	// Open the JSON file
	file, err := os.Open("sample-data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file contents
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Decode the JSON data into a slice of Component structs
	var components []Component
	err = json.Unmarshal(contents, &components)
	if err != nil {
		return nil, err
	}

	return components, nil

}