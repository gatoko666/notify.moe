'use strict'

let Promise = require('bluebird')

arn.get = Promise.promisify(function(set, key, callback) {
	arn.db.get({
		ns: 'arn',
		set: set,
		key: key
	}, function(error, record, metadata, key) {
		if(callback)
			callback(error.code !== 0 ? error : undefined, record)
		else if(error.code !== 0)
			console.error(error.stack)
	})
})

arn.set = Promise.promisify(function(set, key, obj, callback) {
	let aerospikeKey = {
		ns: 'arn',
		set: set,
		key: key
	}

	arn.db.put(aerospikeKey, obj, function(error) {
		if(callback)
			callback(error.code !== 0 ? error : undefined, obj)
		else if(error.code !== 0)
			console.error(error.stack)
	})
})

arn.remove = Promise.promisify(function(set, key, callback) {
	arn.db.remove({
		ns: 'arn',
		set: set,
		key: key
	}, function(error, key) {
		if(callback)
			callback(error.code !== 0 ? error : undefined, key)
		else if(error.code !== 0)
			console.error(error.stack)
	})
})

arn.forEach = Promise.promisify(function(set, func, callback) {
	let scan = arn.db.query('arn', set, {
	    concurrent: true,
	    nobins: false
	})

	let stream = scan.execute()

	stream.on('data', function(record) {
		func(record.bins)
	})

	stream.on('error', function(error) {
		console.error('Error occured while scanning:', error, error.stack)
	})

	if(callback)
		stream.on('end', callback)
})

arn.filter = (set, func) => {
	let records = []

	return arn.forEach(set, record => {
		if(func(record))
			records.push(record)
	}).then(() => records)
}