module.exports = {
	inHours: function(a, b) {
        return parseInt((b - a) / (60 * 60))
    },

    inDays: function(a, b) {
        return parseInt((b - a) / (24 * 60 * 60))
    }
}