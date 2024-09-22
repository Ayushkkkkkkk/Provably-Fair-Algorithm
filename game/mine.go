package game

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
func WinFromLuck(remainingMines int, totalmines int, mineMatrix [][]bool) [][]bool {
	var allPairs []Pair
	for i := 0; i < len(mineMatrix); i++ {
		for j := 0; j < len(mineMatrix[0]); j++ {
			allPairs = append(allPairs, Pair{first: i, second: j})
		}
	}
	rand.Shuffle(len(allPairs), func(i, j int) {
		allPairs[i], allPairs[j] = allPairs[j], allPairs[i]
	})

	for i := 0; i < totalmines; i++ {
		x := allPairs[i].first
		y := allPairs[i].second
		mineMatrix[x][y] = false
	}
	return mineMatrix
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
