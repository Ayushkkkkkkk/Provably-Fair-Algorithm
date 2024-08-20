package game

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type mineStruct struct {
	totalMines     byte
	remainingMines byte
}

func mineRoutes(r *gin.Engine) {
	mineGroup := r.Group("/api/mine")
	mineGroup.POST("/clicked/:totalmines/:remainingMines", ConfigureTOGiveMineOrDiamond)
}

// all win condition 100% win rate condtion

func winCondtion(totalMines, remainingMines byte) string {

}

// all lose condition  losing algortithm 0 percent win condition
func loseCondtion(totalmines, remainingMines byte) string {

}

// probabilty condtion  fair play alogrithm
func WinFromLuck(totalmines byte) {
	minesLocation := make([]byte, totalmines)
	for i := 0; i < int(totalmines); i++ {
		dx := rand.Intn(int(totalmines))
		dy := rand.Intn(int(totalmines))
	}
}

func ConfigureTOGiveMineOrDiamond(r *gin.Context) {
	totalMinesStr := r.Param("totalmines")
	remainingMinesStr := r.Param("remainingMines")

	totalMines, err := strconv.Atoi(totalMinesStr)

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "invalid value from the mines"})
		return
	}

	remainingMines, err := strconv.Atoi(remainingMinesStr)

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "invalid value from the remaining mines"})
	}

	if remainingMines == 0 {
		return
	}

}
