package translations

import (
	"embed"
	"encoding/json"
	"io/fs"
	"strings"

	"golang.org/x/text/language"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

//go:generate goi18n extract -format=json -outdir ./languages -sourceLanguage ja

type LangCode string

const (
	LangCodeZhTw LangCode = "zh-TW"
	LangCodeJa   LangCode = "ja"
	LangCodeEn   LangCode = "en"

	LangCodeDefault LangCode = LangCodeJa
)

var languages = map[LangCode]language.Tag{
	LangCodeZhTw: language.MustParse(string(LangCodeZhTw)),
	LangCodeJa:   language.MustParse(string(LangCodeJa)),
	LangCodeEn:   language.MustParse(string(LangCodeEn)),
}

// init will panic upon error.
func init() {
	bundle := i18n.NewBundle(languages[LangCodeDefault])

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	languageFiles, err := getLanguageFiles()
	if err != nil {
		panic(err)
	}

	for _, languageFile := range languageFiles {
		_, err := bundle.LoadMessageFileFS(localeFS, languageFile)
		if err != nil {
			panic(err)
		}
	}

	initLocalizers(bundle, languages)
	parseStaticMessages()
}

//go:embed languages/active.*.json
var localeFS embed.FS

func getLanguageFiles() ([]string, error) {
	var files []string

	err := fs.WalkDir(localeFS, "languages", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasPrefix(d.Name(), "active.") && strings.HasSuffix(d.Name(), ".json") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

var localizers = map[LangCode]*i18n.Localizer{}

func initLocalizers(b *i18n.Bundle, languages map[LangCode]language.Tag) {
	localizers = make(map[LangCode]*i18n.Localizer, len(languages))
	for langCode := range languages {
		localizers[langCode] = i18n.NewLocalizer(b, string(langCode))
	}
}

// Messages to be initialize using init().
var staticMessages = map[LangCode]map[MessageID]string{}

func GetStatisMessage(lc string) map[MessageID]string {
	if _, ok := languages[LangCode(lc)]; ok {
		return staticMessages[LangCode(lc)]
	}

	return staticMessages[LangCode(LangCodeDefault)]
}

func parseStaticMessages() {
	for langCode, localizer := range localizers {
		if _, ok := staticMessages[langCode]; !ok {
			staticMessages[langCode] = make(map[MessageID]string, len(_staticMessages))
		}

		for msgKey, value := range _staticMessages {
			message := localizer.MustLocalize(value)
			staticMessages[langCode][msgKey] = message
		}
	}
}
