package calc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func getNCalculateFine() {
	fmt.Print("please provide fine total: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	total, err := strconv.ParseFloat(input.Text(), 64)
	if err != nil {
		log.Fatalf("fine total can only be numeric and comma character [0-9,], %s", err.Error())
	}

	fmt.Println("please input each player point, separate each of them with a space")
	input.Scan()
	pointList := input.Text()
	pointBoard := strings.Split(pointList, " ")

	var (
		totalPoint float64
		pointArr   = make([]float64, 0)
	)

	for i := range pointBoard {
		if pointBoard[i] == "" {
			continue
		}
		point, err := strconv.ParseFloat(pointBoard[i], 64)
		if err != nil {
			log.Fatalf("player point can only be numeric and comma character [0-9,], %s", err.Error())
		}
		point = 1 / point
		pointArr = append(pointArr, point)
		totalPoint += point
	}

	if totalPoint == 0 {
		log.Fatal("no player score was provided")
	}

	fmt.Println("each player fine fee:")
	var res string
	for j := range pointArr {
		res += fmt.Sprintf("%.2f, ", (total*pointArr[j])/totalPoint)
	}
	fmt.Println(res[:len(res)-2])
}

// CreateCommand - Create new formula command
func CreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "calc",
		Short: "calculate the fine fee",
		Long:  "calculate the fine fee with the input point and the fine formula we have.. use 'fine formula command to check it out'",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			getNCalculateFine()
		},
	}
}
