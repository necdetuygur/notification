package services

import (
	"net/http"
	"notification/functions"
	"strings"

	"github.com/labstack/echo/v4"
)

func Folder(c echo.Context) error {
	file := c.Param("file")
	if file == "" || file == "/" {
		file = "index.html"
	}
	if strings.Contains(file, ".png") {
		c.Response().Header().Set("Content-Type", "image/png")
	} else if strings.Contains(file, ".svg") {
		c.Response().Header().Set("Content-Type", "image/svg+xml")
	} else if strings.Contains(file, ".js") {
		c.Response().Header().Set("Content-Type", "text/javascript; charset=UTF-8")
	} else if strings.Contains(file, ".css") {
		c.Response().Header().Set("Content-Type", "text/css; charset=UTF-8")
	} else if strings.Contains(file, ".json") {
		c.Response().Header().Set("Content-Type", "application/json; charset=UTF-8")
	} else {
		c.Response().Header().Set("Content-Type", "text/html; charset=UTF-8")
	}

	/**/
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// functions.Pr("Getwd :::::::::::::----> " + dir)
	// files, err := ioutil.ReadDir("./")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, f := range files {
	// 	functions.Pr(f.Name())
	// }
	/**/

	fullPath := strings.ReplaceAll(("./docs/" + file), "//", "/")
	// functions.Pr("fullPath ::::::::::: ----> " + fullPath)
	fileContent := functions.ReadFile(fullPath)
	// for k, v := range config.ALL {
	// 	fileContent = strings.ReplaceAll(fileContent, "<?"+k+"?>", v)
	// 	if strings.Contains(file, ".html") ||
	// 		strings.Contains(file, ".js") ||
	// 		strings.Contains(file, ".css") ||
	// 		strings.Contains(file, ".json") {
	// 		fileContent = functions.Minifier(fileContent)
	// 	}
	// }
	c.Response().Header().Set("cache-control", "public,max-age=31536000,immutable")
	return c.String(http.StatusOK, fileContent)
}
