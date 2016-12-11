# Slacker
A simple slack post message integration.

## Usage

```
Usage of slacker
  -a string
    	BOT avatar (default ":scream:")
  -c string
    	Slack channel to post message
  -m string
    	Message to post
  -t string
    	Slack integration token (default "")
  -u string
    	BOT name (default "Slacker")
```

> You can set the default integration token on environment variable `SLACKTOKEN`

### Example
````
$ slacker -c '#general' -u 'Bot Test' -m 'This is a test!'
```
