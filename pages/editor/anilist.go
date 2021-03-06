package editor

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxAniListEntries = 70

// AniList ...
func AniList(ctx *aero.Context) string {
	missing := arn.FilterAnime(func(anime *arn.Anime) bool {
		return anime.GetMapping("anilist/anime") == ""
	})

	sort.Slice(missing, func(i, j int) bool {
		a := missing[i]
		b := missing[j]

		aPop := a.Popularity.Total()
		bPop := b.Popularity.Total()

		if aPop == bPop {
			return a.Title.Canonical < b.Title.Canonical
		}

		return aPop > bPop
	})

	if len(missing) > maxAniListEntries {
		missing = missing[:maxAniListEntries]
	}

	return ctx.HTML(components.AniListMissingMapping(missing))
}
