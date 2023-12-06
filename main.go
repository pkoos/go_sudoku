package main

import (
	"fmt"
	// "text/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type grid struct {
	ID string `json:"id"`
	Difficulty string `json:"difficulty"`
	Start [][]int `json:"start"`
	Solution [][]int `json:"solution"`
}

var sudoku_grids = []grid{
	{
		ID: "1", 
		Difficulty: "Easy", 
		Start: [][]int{
			{0, 3, 0, 7, 0, 6, 0, 8, 0},
			{0, 0, 6, 3, 0, 0, 2, 0, 7},
			{0, 0, 8, 0, 0, 0, 6, 0, 3},
			{0, 9, 0, 0, 7, 0, 8, 3, 0},
			{8, 1, 0, 0, 6, 0, 0, 4, 5},
			{3, 0, 7, 8, 4, 0, 1, 9, 0},
			{0, 0, 0, 0, 5, 4, 3, 7, 0},
			{5, 8, 0, 0, 3, 0, 0, 2, 9},
			{0, 0, 0, 2, 8, 0, 5, 0, 0},
		},
		Solution: [][]int{
			{1, 3, 5, 7, 2, 6, 9, 8, 4},
			{9, 4, 6, 3, 1, 8, 2, 5, 7},
			{7, 2, 8, 4, 9, 5, 6, 1, 3},
			{6, 9, 4, 5, 7, 1, 8, 3, 2},
			{8, 1, 2, 9, 6, 3, 7, 4, 5},
			{3, 5, 7, 8, 4, 2, 1, 9, 6},
			{2, 6, 9, 1, 5, 4, 3, 7, 8},
			{5, 8, 1, 6, 3, 7, 4, 2, 9},
			{4, 7, 3, 2, 8, 9, 5, 6, 1},
		},
	},
	{
		ID: "2",
		Difficulty: "Easy",
		Start: [][]int{
			{5, 6, 8, 0, 9, 3, 0, 0, 7},
			{3, 4, 0, 0, 0, 7, 0, 0, 5},
			{0, 9, 7, 5, 0, 4, 6, 0, 3},
			{7, 0, 0, 0, 1, 2, 5, 0, 0},
			{0, 1, 9, 0, 0, 8, 7, 6, 0},
			{0, 8, 0, 7, 0, 0, 0, 3, 0},
			{0, 2, 0, 0, 3, 5, 8, 0, 0},
			{0, 5, 1, 2, 0, 6, 3, 0, 0},
			{0, 0, 0, 9, 0, 0, 0, 0, 0},
		},
		Solution: [][]int{
			{5, 6, 8, 1, 9, 3, 4, 2, 7},
			{3, 4, 2, 8, 6, 7, 9, 1, 5},
			{1, 9, 7, 5, 2, 4, 6, 8, 3},
			{7, 3, 4, 6, 1, 2, 5, 9, 8},
			{2, 1, 9, 3, 5, 8, 7, 6, 4},
			{6, 8, 5, 7, 4, 9, 1, 3, 2},
			{9, 2, 6, 4, 3, 5, 8, 7, 1},
			{8, 5, 1, 2, 7, 6, 3, 4, 9},
			{4, 7, 3, 9, 8, 1, 2, 5, 6},
		},
	},
	{
		ID: "3",
		Difficulty: "Easy",
		Start: [][]int{
			{0, 0, 2, 0, 0, 5, 0, 8, 7},
			{0, 8, 5, 0, 0, 0, 4, 9, 0},
			{0, 7, 4, 9, 2, 0, 0, 0, 0},
			{0, 0, 0, 0, 9, 0, 0, 0, 0},
			{4, 0, 6, 1, 0, 7, 0, 0, 0},
			{0, 5, 3, 0, 8, 4, 0, 0, 0},
			{0, 0, 0, 0, 7, 2, 9, 4, 8},
			{0, 4, 9, 8, 3, 0, 0, 7, 5},
			{0, 2, 7, 0, 4, 0, 6, 3, 1},
		},
		Solution: [][]int{
			{9, 6, 2, 4, 1, 5, 3, 8, 7},
			{1, 8, 5, 7, 6, 3, 4, 9, 2},
			{3, 7, 4, 9, 2, 8, 5, 1, 6},
			{2, 1, 8, 3, 9, 6, 7, 5, 4},
			{4, 9, 6, 1, 5, 7, 8, 2, 3},
			{7, 5, 3, 2, 8, 4, 1, 6, 9},
			{5, 3, 1, 6, 7, 2, 9, 4, 8},
			{6, 4, 9, 8, 3, 1, 2, 7, 5},
			{8, 2, 7, 5, 4, 9, 6, 3, 1},
		},
	},
}

func getGrids(c * gin.Context) {
	c.IndentedJSON(http.StatusOK, sudoku_grids)
}

func postGrids(c *gin.Context) {
	var newGrid grid
	if err := c.BindJSON(&newGrid); err != nil {
		return
	}

	sudoku_grids = append(sudoku_grids, newGrid)
	c.IndentedJSON(http.StatusCreated, newGrid)
}

func getGridByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range sudoku_grids {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sudoku grid not found"})
}

func displayGrid(c *gin.Context) {
	id := c.Param("id")
	var sudoku_grid grid
	for _, s_grid := range sudoku_grids {
		if s_grid.ID == id {
			sudoku_grid = s_grid
		}
	}
	// fmt.Println(sudoku_grid)
	var start_html strings.Builder
	start_html.WriteString("<div class=\"sudoku_grid start_grid\"><h2>Starting Grid</h2>")
	start_html.WriteString(genGrid(sudoku_grid.Start))
	start_html.WriteString("</div>")
	
	var solved_html strings.Builder
	solved_html.WriteString("<div class=\"sudoku_grid end_grid\"><h2>Solved Grid</h2>")
	solved_html.WriteString(genGrid(sudoku_grid.Solution))
	solved_html.WriteString("</div>")
	
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf("%s%s", start_html.String(), solved_html.String())))
}

func genGrid(grid [][]int) string {
	var html strings.Builder
	for row_idx := 0; row_idx < len(grid); row_idx++ {
		if row_idx % 3 == 0 {
			html.WriteString("<p></p>")
		}
		html.WriteString("<div>")
		for col_idx := 0; col_idx < len(grid[row_idx]); col_idx++ {
			if col_idx % 3 == 0 && col_idx != 0 {
				html.WriteString("&nbsp;&nbsp;&nbsp;")
			}

			if grid[row_idx][col_idx] == 0 {
				html.WriteString("_ ")
			} else {
				html.WriteString(fmt.Sprintf("%d ", grid[row_idx][col_idx]))
			}
		}
		html.WriteString("</div>")
	}
	return html.String()
}

func getHomepage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func getTest(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", gin.H{})
}

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLFiles("index.html", "test.html")
	router.GET("/", getHomepage)
	router.GET("/test", getTest)
	router.GET("/grids", getGrids)
	router.GET("/grid/:id", getGridByID)
	router.GET("/dgrid/:id", displayGrid)
	router.POST("/grids", postGrids)
	

	router.Run("localhost:8080")
}
