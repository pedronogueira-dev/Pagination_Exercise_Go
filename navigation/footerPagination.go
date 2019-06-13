package navigation

import (
	"fmt"
	"strconv"
	"strings"
)

// GetPagination returns a string representing the footer pagination, given the following parameters:
// current page, total pages, number of pages to display on either end of the interval (boundaries) and surrounding the current page (around).
// An error is returned if the current page is out of bounds, i.e., it's index is less than 1 or greater than the total number of pages.
func GetPagination(current int, total int, boundaries int, around int) (string, error) {
	var ind int
	var res strings.Builder

	if current < 1 || current > total {
		return "", fmt.Errorf("Error: current page index (current_page = %v) is out of bounds", current)
	}

	if boundaries > 0 || current-around <= 1 {
		ind = 1

	} else {
		res.WriteString("...")
		ind = current - around
	}

	for ind <= total {
		if ind > boundaries && ind < current-around {
			res.WriteString(" ...")
			ind = current - around
			continue
		} else if ind > current+around && ind < total+1-boundaries {
			res.WriteString(" ...")
			ind = total + 1 - boundaries
			continue
		}
		if ind != 1 {
			res.WriteString(" ")
		}
		res.WriteString(strconv.Itoa(ind))

		ind++
	}

	return res.String(), nil
}
