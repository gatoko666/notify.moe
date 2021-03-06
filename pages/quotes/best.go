package quotes

import (
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Best renders the quotes page ordered by the most favorites first.
func Best(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	quotes := arn.FilterQuotes(func(track *arn.Quote) bool {
		return !track.IsDraft
	})

	arn.SortQuotesPopularFirst(quotes)

	// Limit the number of displayed quotes
	loadMoreIndex := 0

	if len(quotes) > maxQuotes {
		quotes = quotes[:maxQuotes]
		loadMoreIndex = maxQuotes
	}

	return ctx.HTML(components.Quotes(quotes, loadMoreIndex, user))
}

// BestFrom renders the quotes from the given index.
func BestFrom(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	index, err := ctx.GetInt("index")

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Invalid start index", err)
	}

	allQuotes := arn.FilterQuotes(func(track *arn.Quote) bool {
		return !track.IsDraft
	})

	if index < 0 || index >= len(allQuotes) {
		return ctx.Error(http.StatusBadRequest, "Invalid start index (maximum is "+strconv.Itoa(len(allQuotes))+")", nil)
	}

	arn.SortQuotesPopularFirst(allQuotes)

	quotes := allQuotes[index:]

	if len(quotes) > maxQuotes {
		quotes = quotes[:maxQuotes]
	}

	nextIndex := index + maxQuotes

	if nextIndex >= len(allQuotes) {
		// End of data - no more scrolling
		ctx.Response().Header().Set("X-LoadMore-Index", "-1")
	} else {
		// Send the index for the next request
		ctx.Response().Header().Set("X-LoadMore-Index", strconv.Itoa(nextIndex))
	}

	return ctx.HTML(components.QuotesScrollable(quotes, user))
}
