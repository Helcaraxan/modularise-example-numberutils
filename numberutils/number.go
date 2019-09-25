package numberutils

import (
	"fmt"

	"github.com/Helcaraxan/modularise-example-numberutils/internal/random"
)

func PrintRandomNumber() {
	fmt.Println(random.GenerateRandomNumber())
}
