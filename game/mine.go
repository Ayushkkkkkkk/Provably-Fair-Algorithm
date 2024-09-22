package game

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type mineStruct struct {
	minesMatrix [][]bool
	totalMines  int
}

type Pair struct {
	first  int
	second int
}

func mineRoutes(r *gin.Engine) {
	mineGroup := r.Group("/api/mine")
	mineGroup.POST("/clicked/:totalmines/:minesMatrix/:clickedCoords", ConfigureTOGiveMineOrDiamond)
}

// all win condition 100% win rate condtion
// check and put the mines in the random postion which wasn't clicked
func winCondtion(totalMines int, minesMatrix [][]bool) [][]bool {
	var filter []Pair
	count := 0
	for i := 0; i < len(minesMatrix); i++ {
		for j := 0; j < len(minesMatrix[i]); j++ {
			if count == totalMines {
				break
			}
			if minesMatrix[i][j] == true {
				filter = append(filter, Pair{first: i, second: j})
				count++
			}
		}
	}

	for i := 0; i < len(filter); i++ {
		minesMatrix[filter[i].first][filter[i].second] = false
	}

	return minesMatrix

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
	minesMatrixStr := r.Param("minesMatrix")
	clickedCoordsStr := r.Param("clickedCoords")

	totalMines, err := strconv.Atoi(totalMinesStr)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "invalid value for totalMines"})
		return
	}

	var minesMatrix [][]bool
	err = json.Unmarshal([]byte(minesMatrixStr), &minesMatrix)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "invalid minesMatrix format"})
		return
	}

	var coords [][]Pair
	err = json.Unmarshal([]byte(clickedCoordsStr), &coords)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "invalid clickedCoords format"})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"totalMines":    totalMines,
		"minesMatrix":   minesMatrix,
		"clickedCoords": coords,
	})

}
