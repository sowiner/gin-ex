package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	app := gin.Default()
	// default ping json
	app.GET("/ping",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// -----------------------------------------------------------------------------------------
	// mutil file
	app.POST("/mfile",func(c *gin.Context){
		form,_ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
		}
		c.String(http.StatusOK,fmt.Sprintf("%d files is be upload",len(files)))
	})
	// curl -X POST http://localhost:8080/upload \
	//   -F "upload[]=@/Users/appleboy/test1.zip" \
	//   -F "upload[]=@/Users/appleboy/test2.zip" \
	//   -H "Content-Type: multipart/form-data"


	// -------------------------------------------------------------------------------------------
	// single file
	app.POST("/sfile",func(c *gin.Context){
		ufile,_ := c.FormFile("upload")
		log.Println(ufile.Filename)
		c.String(http.StatusOK,fmt.Sprintf("%s file is be upload",ufile.Filename))
	})
	// curl -X POST http://localhost:8080/upload \
	//   -F "file=@/Users/appleboy/test.zip" \
	//   -H "Content-Type: multipart/form-data"
	
	//------------------------------------------------------------------------------------------------- 
	//  AsciiJSON
	app.GET("/asciijson",func(c *gin.Context){
		data := map[string]interface{}{
			"lang":"GO語言",
			"tag":"<br>",
		}
		c.AsciiJSON(http.StatusOK,data)
		// show {"lang":"GO\u8a9e\u8a00","tag":"\u003cbr\u003e"}
	})
	

	
	
	// ----------------------------------------------------------------------------------------------
	// run server
	app.Run()
}