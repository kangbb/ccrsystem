var express = require('express')
var app = express()

app.use(express.static('./dist'))

app.listen(2333)
console.log('Server running at http://127.0.0.1:2333')