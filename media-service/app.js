var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');

var indexRouter = require('./routes/index');
var usersRouter = require('./routes/users');
var mediaRouter = require('./routes/media');

var app = express();

app.use(logger('dev'));
app.use(express.json({limit : '50mb'}));
app.use(express.urlencoded({ extended: false, limit : '50mb' }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter);
app.use('/users', usersRouter);
app.use('/media', mediaRouter);

module.exports = app;
