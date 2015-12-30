'use strict'

let arn = require('../../../lib')

module.exports = {
	get: function(request, response) {
		let nick = request.params[0]

		if(!nick)
			return response.end()

		let anilistNick = 'Akyoto'

		arn.AniList.getAnimeList(anilistNick, function(error, watching) {
			watching.forEach(function(entry) {
				entry.animeProvider = {
					url: null,
					nextEpisodeUrl: null,
					videoUrl: null
				}
			})

			let json = {
				listProvider: 'AniList',
				listUrl: arn.AniList.getAnimeListUrl(anilistNick),
				watching
			}

			response.setHeader('Content-Type', 'application/json')
			response.end(JSON.stringify(json))
		})
	}
}