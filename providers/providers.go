package providers

import (
	"fmt"
	"os"

	"github.com/aandrew-me/tgpt/v2/providers/blackboxai"
	"github.com/aandrew-me/tgpt/v2/providers/duckduckgo"
	"github.com/aandrew-me/tgpt/v2/providers/groq"
	"github.com/aandrew-me/tgpt/v2/providers/koboldai"
	"github.com/aandrew-me/tgpt/v2/providers/ollama"
	"github.com/aandrew-me/tgpt/v2/providers/openai"
	"github.com/aandrew-me/tgpt/v2/providers/opengpts"
	"github.com/aandrew-me/tgpt/v2/providers/phind"
	"github.com/aandrew-me/tgpt/v2/structs"
	http "github.com/bogdanfinn/fhttp"
)

var availableProviders = []string{
	"", "opengpts", "ollama", "openai", "phind", "koboldai", "blackboxai", "groq", "duckduckgo",
}

func GetMainText(line string, provider string, input string) string {
	if provider == "blackboxai" {
		return blackboxai.GetMainText(line)
	} else if provider == "groq" {
		return groq.GetMainText(line)
	} else if provider == "duckduckgo" {
		return duckduckgo.GetMainText(line)
	} else if provider == "koboldai" {
		return koboldai.GetMainText(line)
	} else if provider == "ollama" {
		return ollama.GetMainText(line)
	} else if provider == "opengpts" {
		return opengpts.GetMainText(line, input)
	} else if provider == "openai" {
		return openai.GetMainText(line)
	} else if provider == "phind" {
		return phind.GetMainText(line)
	}

	return phind.GetMainText(line)
}

func NewRequest(input string, params structs.Params, extraOptions structs.ExtraOptions) (*http.Response, error) {
	validProvider := false
	for _, str := range availableProviders {
		if str == params.Provider {
			validProvider = true
			break
		}
	}
	if !validProvider {
		fmt.Fprintln(os.Stderr, "Invalid provider")
		os.Exit(1)
	}

	if params.Provider == "blackboxai" {
		return blackboxai.NewRequest(input, params, extraOptions.PrevMessages)
	} else if params.Provider == "groq" {
		return groq.NewRequest(input, params, extraOptions.PrevMessages)
	} else if params.Provider == "duckduckgo" {
		return duckduckgo.NewRequest(input, params, extraOptions.PrevMessages)
	} else if params.Provider == "koboldai" {
		return koboldai.NewRequest(input, params, extraOptions.PrevMessages)
	} else if params.Provider == "ollama" {
		return ollama.NewRequest(input, params, extraOptions.PrevMessages)
	} else if params.Provider == "opengpts" {
		return opengpts.NewRequest(input, params, extraOptions)
	} else if params.Provider == "openai" {
		return openai.NewRequest(input, params, extraOptions.PrevMessages)
	} else if params.Provider == "phind" {
		return phind.NewRequest(input, params, extraOptions.PrevMessages)
	}

	return phind.NewRequest(input, params, extraOptions.PrevMessages)
}
