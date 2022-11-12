package stripr

import (
	"flag"
	"fmt"
	"os"

	"github.com/aosasona/stripr/types"
	"github.com/aosasona/stripr/utils"
)

func main() {

	var err error

	targetPath := flag.String("target", ".", "The directory or file to read")
	showStats := flag.Bool("show-stats", false, "Show the number of files and lines that will be affected")
	skipCheck := flag.Bool("skip-check", false, "Skip the confirmation prompt before stripping comments")

	flag.Parse()

	defer os.Exit(0)

	stripr, err := CreateCMD(targetPath, Stripr{
		ShowStats: *showStats,
		SkipCheck: *skipCheck,
		Args:      flag.Args(),
	})
	_, err = stripr.Run()

	if err != nil {
		switch err.(type) {
		case *types.CustomError:
			utils.Terminate(err)
		default:
			utils.Terminate(&types.CustomError{Message: fmt.Sprintf("An error occurred: %s", err.Error())})
		}
	}
}
