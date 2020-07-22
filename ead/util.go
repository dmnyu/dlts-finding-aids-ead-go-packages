package ead

import (
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strings"
)

func getConvertedTextWithTags(text string) ([]byte, error) {
	decoder := xml.NewDecoder(strings.NewReader(text))

	var result string
	needClosingTag := true
	for {
		token, err := decoder.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch token := token.(type) {
		case xml.StartElement:
			switch token.Name.Local {
			default:
				result += fmt.Sprintf("<span class=\"ead-%s\">", token.Name.Local)
			case "emph":
				{
					var render string
					for i := range token.Attr {
						if token.Attr[i].Name.Local == "render" {
							render = token.Attr[i].Value
							break
						}
					}
					result += fmt.Sprintf("<span class=\"%s\">", "ead-emph ead-emph-" + render)
				}
			}

		case xml.EndElement:
			if needClosingTag {
				result += "</span>"
			} else {
				// Reset
				needClosingTag = true
			}
		case xml.CharData:
			result += string(token)
		}
	}

	return []byte(result), nil
}

func regexpReplaceAllLiteralStringInNameWithRoleSlice( nameWithRoleSlice []NameWithRole, re *regexp.Regexp, replacementString string ) {
	nameWithRoleSliceWithSubfieldDelimitersConverted := nameWithRoleSlice[:0]
	for _, nameWithRole := range nameWithRoleSlice {
		nameWithRole.Value = re.ReplaceAllLiteralString(nameWithRole.Value, replacementString)
		nameWithRoleSliceWithSubfieldDelimitersConverted = append(
			nameWithRoleSliceWithSubfieldDelimitersConverted,
			nameWithRole,
		)
	}
}

func regexpReplaceAllLiteralStringInTextSlice( textSlice []string, re *regexp.Regexp, replacementString string ) {
	nameWithRoleSliceWithSubfieldDelimitersConverted := textSlice[:0]
	for _, text := range textSlice {
		nameWithRoleSliceWithSubfieldDelimitersConverted = append(
			nameWithRoleSliceWithSubfieldDelimitersConverted,
			re.ReplaceAllLiteralString(text, replacementString),
		)
	}
}

