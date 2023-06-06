package bluesky

import (
	"fmt"
	"regexp"
	"strings"

	appbsky "github.com/bluesky-social/indigo/api/bsky"
)

func DetectFacets(text string) []*appbsky.RichtextFacet {
	facets := []*appbsky.RichtextFacet{}
	{
		// mentions
		regex := regexp.MustCompile(`(^|\s|\()(@)([a-zA-Z0-9.-]+)(\b)`)
		matches := regex.FindAllStringSubmatch(text, -1)
		indexes := regex.FindAllStringIndex(text, -1)
		for index, match := range matches {
			if len(match) == 0 {
				continue
			}

			domain := match[3]
			fmt.Println(domain)
			if !isValidDomain(domain) && !strings.HasSuffix(domain, ".test") {
				continue
			}

			facets = append(facets, &appbsky.RichtextFacet{
				Index: &appbsky.RichtextFacet_ByteSlice{
					ByteStart: int64(indexes[index][0]),
					ByteEnd:   int64(indexes[index][1]),
				},
				Features: []*appbsky.RichtextFacet_Features_Elem{
					{
						RichtextFacet_Mention: &appbsky.RichtextFacet_Mention{
							LexiconTypeID: "app.bsky.richtext.facet#mention",
							Did:           domain,
						},
					},
				},
			})
		}
	}

	{
		// links
		regex := regexp.MustCompile("(?is)" + `(^|\s|\()((https?:\/\/[\S]+)|((?P<domain>[a-z][a-z0-9]*(\.[a-z0-9]+)+)[\S]*))`)
		matches := regex.FindAllStringSubmatch(text, -1)
		indexes := regex.FindAllStringIndex(text, -1)
		for index, match := range matches {
			if len(match) == 0 {
				continue
			}

			uri := match[2]
			if !strings.HasPrefix(uri, "http") {
				domain := match[regex.SubexpIndex("domain")]
				if len(domain) == 0 || !isValidDomain(domain) {
					continue
				}

				uri = "https://" + uri
			}

			fmt.Println(uri)

			start := int64(indexes[index][0])
			end := int64(indexes[index][1])

			// if (/[.,;!?]$/.test(uri)) {
			// 	uri = uri.slice(0, -1)
			// 	index.end--
			//   }
			//   if (/[)]$/.test(uri) && !uri.includes('(')) {
			// 	uri = uri.slice(0, -1)
			// 	index.end--
			//   }

			facets = append(facets, &appbsky.RichtextFacet{
				Index: &appbsky.RichtextFacet_ByteSlice{
					ByteStart: start,
					ByteEnd:   end,
				},
				Features: []*appbsky.RichtextFacet_Features_Elem{
					{
						RichtextFacet_Link: &appbsky.RichtextFacet_Link{
							LexiconTypeID: "app.bsky.richtext.facet#link",
							Uri:           uri,
						},
					},
				},
			})
		}
	}

	if len(facets) == 0 {
		return nil
	}

	return facets
}

// TODO: actually check if it is valid
func isValidDomain(inp string) bool {
	return true
}
