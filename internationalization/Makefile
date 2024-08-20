i18n/init:
	@goi18n extract -format=json -outdir ./translations/languages -sourceLanguage ja
	# @goi18n merge -sourceLanguage ja -outdir ./translations/languages -format json ./translations/languages/active.*.json
	@goi18n merge -sourceLanguage ja -outdir ./translations/languages -format json ./translations/languages/active.ja.json ./translations/languages/translate.zh-TW.json


goi18n merge -sourceLanguage ja -outdir ./translations/languages -format json ./translations/languages/active.ja.json ./translations/languages/translate.en.json

goi18n merge -sourceLanguage ja -outdir ./translations/languages -format json ./translations/languages/active.*.json ./translations/languages/translate.*.json 