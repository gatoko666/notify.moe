'use strict'

let aero = require('aero')
let database = require('../modules/database')
let arn = require('../lib')

database(aero, function(error) {
	let tasks = []

    arn.scan('Anime', function(anime) {
		tasks.push(arn.remove('Anime', anime.id))
    }, function() {
		console.log('Waiting...')
		Promise.all(tasks).then(() => console.log(`Finished deleting ${tasks.length} anime`))
    })
})