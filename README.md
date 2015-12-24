# Vanguard
[![Build Status](https://travis-ci.org/vampirewalk/vanguard.svg?branch=master)](https://travis-ci.org/vampirewalk/vanguard)
[![Coverage Status](https://coveralls.io/repos/vampirewalk/vanguard/badge.svg?branch=master&service=github)](https://coveralls.io/github/vampirewalk/vanguard?branch=master)
[![Twitter: @vampirewalk666](https://img.shields.io/badge/contact-%40vampirewalk-blue.svg)](https://twitter.com/vampirewalk666)

Vanguard is a slack bot. It reports the state of github repository when somebody mentions.

![Screenshot](https://raw.githubusercontent.com/vampirewalk/vanguard/master/screenshot.png)

# Installation
* [Create a new bot user integration.](https://my.slack.com/services/new/bot)

* Get token.

* Install bot.

```
$ go get github.com/vampirewalk/vanguard
$ echo "SLACK_TOKEN=YOUR_REAL_TOKEN" > .env
$ go run main.go
```
# Deploy to Heroku
* [Install the Heroku Toolbelt.](https://devcenter.heroku.com/articles/getting-started-with-go#set-up)
* ```$ heroku login```
* Deploy the app.

```
$ go get github.com/vampirewalk/vanguard
$ cd $GOPATH/src/github.com/vampirewalk/vanguard
$ heroku create
$ heroku config:set SLACK_TOKEN=YOUR_REAL_TOKEN
$ git push heroku master
```
* [Invite bot to channel.](https://get.slack.help/hc/en-us/articles/201980108-Inviting-team-members-to-a-channel) Your bot has been named at creating integration step.

