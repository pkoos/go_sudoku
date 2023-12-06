package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type grid struct {
	ID string `json:"id"`
	Difficulty string `json:"difficulty"`
	Start [SUDOKU_HEIGHT][SUDOKU_WIDTH]int `json:"start"`
	Solution [SUDOKU_HEIGHT][SUDOKU_WIDTH]int `json:"solution"`
}

type technique struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Steps [][SUDOKU_HEIGHT][SUDOKU_WIDTH]int `json:"steps"`
}

const SUDOKU_HEIGHT = 9;
const SUDOKU_WIDTH = 9;

var techniques []technique

var sudoku_grids = []grid{
	{
		ID: "1", 
		Difficulty: "Easy", 
		Start: [SUDOKU_HEIGHT][SUDOKU_WIDTH]int{
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
		Solution: [SUDOKU_HEIGHT][SUDOKU_WIDTH]int{
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
		Start: [SUDOKU_HEIGHT][SUDOKU_WIDTH]int{
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
		Solution: [SUDOKU_HEIGHT][SUDOKU_WIDTH]int{
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
		Start: [SUDOKU_HEIGHT][SUDOKU_WIDTH]int{
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
		Solution: [SUDOKU_HEIGHT][SUDOKU_WIDTH]int{
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
	var start_html strings.Builder
	start_html.WriteString("<div class=\"sudoku_grid\"><h2>Starting Grid</h2><table>")
	start_html.WriteString(genGrid(sudoku_grid.Start))
	start_html.WriteString("</table></div>")
	
	var solved_html strings.Builder
	solved_html.WriteString("<div class=\"sudoku_grid\"><h2>Solved Grid</h2><table>")
	solved_html.WriteString(genGrid(sudoku_grid.Solution))
	solved_html.WriteString("</table></div>")
	
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf("%s%s", start_html.String(), solved_html.String())))
}

func genGrid(grid [SUDOKU_HEIGHT][SUDOKU_WIDTH]int) string {
	var html strings.Builder

	for row_idx := 0; row_idx < len(grid); row_idx++ {
		html.WriteString(fmt.Sprintf("<tr class=\"row_%d\">", row_idx))
		for col_idx := 0; col_idx < len(grid[row_idx]); col_idx++ {
			html.WriteString(gridValue(grid[row_idx][col_idx], col_idx))
		}
		html.WriteString("</tr>")
	}
	return html.String()
}

func gridValue(value int, col int) string {
	if value == 0 {
		return fmt.Sprintf("<td class=\"col_%d\"></td>", col) 
	} else {
		return fmt.Sprintf("<td class=\"col_%d\">%d</td>", col, value)
	}
}

func getHomepage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func getTest(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", gin.H{})
}

func getTest2(c *gin.Context) {
	c.HTML(http.StatusOK, "test2.html", gin.H{})
}


func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLFiles("index.html", "test.html", "test2.html")
	router.GET("/", getHomepage)
	router.GET("/test", getTest)
	router.GET("/test2", getTest2)
	router.GET("/grids", getGrids)
	router.GET("/grid/:id", getGridByID)
	router.GET("/dgrid/:id", displayGrid)
	router.POST("/grids", postGrids)
	

	router.Run("localhost:8080")
}
