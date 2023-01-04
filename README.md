You need to enable Youtube API v3 on Google's Developer Console.

In addition, you will need to go to the "Credentials" section of left-hand
panel and click on credentials.

Click on "CREATE CREDENTIALS" and create an API key. Store this API Key
in a safe place. We will be setting this API Key in environment variables
so we can use it to authenticate ourselves and retrive transcripts from
Youtube through their API.

More detailed explanation on [Youtube Data API authorization options](https://developers.google.com/youtube/registering_an_application)

[Quota impact per Documentation](https://developers.google.com/youtube/v3/docs/captions/list#usage)
- A call to this method has a quota cost of 50 units.

Required parameters
part	string
The part parameter specifies the caption resource parts that the API response will include.

The list below contains the part names that you can include in the parameter value:
id
snippet

svideoId	string
The videosssId parameter specifies the YouTube video ID of the video for which the API should return caption tracks.

TODO:
- add tests
- add more endpoints that could be useful
- determine available permission scopes for automation with api key. Some endpoints such as captions download blocked, api key is not enough. You need Oauth 2.0 verification.


```text 
curl \
  'https://youtube.googleapis.com/youtube/v3/captions?part=snippet&videoId=M7FIvfx5J10&key=[YOUR_API_KEY]' \
  --header 'Authorization: Bearer [YOUR_ACCESS_TOKEN]' \
  --header 'Accept: application/json' \
  --compressed

```


example of cli command
```text
go run *.go transcript -videoId=ORdWE_ffirg -outputFileName=jamesJaniGreatCryptoScam.txt
```