package twitter

var twitterMap = map[string]string{
	"rajatjindal/goodfirstissue": "rajatjindal1983",
	"openfaas":                   "openfaas",
	"helm":                       "HelmPack",
	"asyncy":                     "asyncy",
}

//GetTwitterHandleMap returns the twitter map
func GetTwitterHandleMap() map[string]string {
	return twitterMap
}
