var cors = require('cors')
var express = require('express');
var app = express();

app.use(cors());
app.options('*', cors());

var Redis = require('ioredis');
var client = new Redis({
    host: 'redis',
    port: 6379
});

client.on("error", function (err) {
    console.error("Redis error", err);
});

app.get('/', function (req, res) {
    res.redirect('/index.html');
});

app.get('/json', function (req, res) {
    client.hlen('wallet', function (err, coins) {
        console.log("hlen wallet err: ", err)
        client.get('hashes', function (err, hashes) {
            console.log("get hashes err: ", err)
            var now = Date.now() / 1000;
            res.json( {
                coins: coins,
                hashes: hashes,
                now: now
            });
        });
    });
});

app.use(express.static('files'));

var server = app.listen(8080, function () {
    console.log('WEBUI running on port 8080');
});
