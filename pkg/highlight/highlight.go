package highlight

import (
	"os"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)


func HighlightCode(code, language string) string {
	lexer := lexers.Get(language)
	if lexer== nil {
		lexer = lexers.Fallback
	}
	
	iterator,err := lexer.Tokenise(nil, code)
	if err!=nil {
		return code 
	}
	
	style := styles.Get("monokai")
	if style==nil {
		style = styles.Fallback
	}
	
	formatter := formatters.Get("terminal256")
	if formatter==nil {
		return code
	}
	var highlightedCode string 
	formatter.Format(os.Stdout, style, iterator)
	return highlightedCode
}