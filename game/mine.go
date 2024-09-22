package game

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

type mineStruct struct {
	minesMatrix    [][]bool
	totalMines     int
	remainingMines int
}

type Pair struct {
	first  int
	second int
}

func mineRoutes(r *gin.Engine) {
	mineGroup := r.Group("/api/mine")
	mineGroup.POST("/clicked/:totalmines/:remainingMines/:minesMatrix", ConfigureTOGiveMineOrDiamond)
}

// all win condition 100% win rate condtion

func winCondtion() string {
	return "win"
}

// all lose condition  losing algortithm 0 percent win condition

func loseCondtion() string {
	return "lose"
}

// probabilty condtion  fair play alogrithm
func WinFromLuck(remainingMines int, totalmines int, mineMatrix [][]bool) {
	mpp := make(map[Pair]int)
	count := 0
	for i := 0; i < totalmines; i++ {
		x := rand.Intn(len(mineMatrix))
		y := rand.Intn(len(mineMatrix))
		if value, exists := mpp[Pair{first: x, second: y}]; exists {
			continue
		} else {
			count++
			mineMatrix[x-1][y-1] = false
			mpp[Pair{first: x, second: y}]++
		}
	}
}

func ConfigureTOGiveMineOrDiamond(r *gin.Context) {
	totalMinesStr := r.Param("totalmines")
	remainingMinesStr := r.Param("remainingMines")
	minesMatrix := r.Param("minesMatrix")
	var matrix [][]bool

	err := json.Unmarshal([]byte(minesMatrix), &matrix)
	totalMines, err := strconv.Atoi(totalMinesStr)

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "invalid value from the mines"})
		return
	}

	remainingMines, err := strconv.Atoi(remainingMinesStr)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "invalid value from the remaining mines"})
	}
	WinFromLuck(int(remainingMines), int(totalMines), matrix)

	if remainingMines == 0 {
		return
	}

}
