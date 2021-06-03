package binding

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/id"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/pt"
	"github.com/go-playground/locales/pt_BR"
	"github.com/go-playground/locales/ru"
	"github.com/go-playground/locales/zh"
	zh_TW "github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	es_translations "github.com/go-playground/validator/v10/translations/es"
	fr_translations "github.com/go-playground/validator/v10/translations/fr"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
	pt_translations "github.com/go-playground/validator/v10/translations/pt"
	pt_BR_translations "github.com/go-playground/validator/v10/translations/pt_BR"
	ru_translations "github.com/go-playground/validator/v10/translations/ru"
	tr_translations "github.com/go-playground/validator/v10/translations/tr"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	zh_TW_translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

const (
	LANG_EN    = "en"
	LANG_ES    = "es"
	LANG_FR    = "fr"
	LANG_ID    = "id"
	LANG_JA    = "ja"
	LANG_PT    = "pt"
	LANG_PT_BR = "pt_BR"
	LANG_RU    = "ru"
	LANG_TR    = "tr"
	LANG_ZH    = "zh"
	LANG_ZH_TW = "zh_TW"
)

var (
	ValidatorList map[string]*defaultValidator
	uni_en        *ut.UniversalTranslator
	uni_es        *ut.UniversalTranslator
	uni_fr        *ut.UniversalTranslator
	uni_id        *ut.UniversalTranslator
	uni_ja        *ut.UniversalTranslator
	uni_pt        *ut.UniversalTranslator
	uni_pt_BR     *ut.UniversalTranslator
	uni_ru        *ut.UniversalTranslator
	uni_tr        *ut.UniversalTranslator
	uni_zh        *ut.UniversalTranslator
	uni_zh_TW     *ut.UniversalTranslator
)

func init() {
	ValidatorList = map[string]*defaultValidator{
		LANG_EN:    new(defaultValidator),
		LANG_ES:    new(defaultValidator),
		LANG_FR:    new(defaultValidator),
		LANG_ID:    new(defaultValidator),
		LANG_JA:    new(defaultValidator),
		LANG_PT:    new(defaultValidator),
		LANG_PT_BR: new(defaultValidator),
		LANG_RU:    new(defaultValidator),
		LANG_TR:    new(defaultValidator),
		LANG_ZH:    new(defaultValidator),
		LANG_ZH_TW: new(defaultValidator),
	}
	for lang, value := range ValidatorList {
		switch lang {
		case LANG_ES:
			value.Engine()
			l := es.New()
			uni_es = ut.New(l, l)
			trans, _ := uni_en.GetTranslator("es")
			es_translations.RegisterDefaultTranslations(value.GetValidate(), trans)
			break
		case LANG_FR:
			value.Engine()
			l := fr.New()
			uni_fr = ut.New(l, l)
			trans, _ := uni_fr.GetTranslator("fr")
			fr_translations.RegisterDefaultTranslations(value.GetValidate(), trans)
			break
		case LANG_ID:
			value.Engine()
			l := id.New()
			uni_id = ut.New(l, l)
			trans, _ := uni_id.GetTranslator("id")
			id_translations.RegisterDefaultTranslations(value.GetValidate(), trans)
			break
		case LANG_JA:
			value.Engine()
			l := ja.New()
			uni_ja = ut.New(l, l)
			trans, _ := uni_ja.GetTranslator("ja")
			ja_translations.RegisterDefaultTranslations(value.GetValidate(), trans)
			break
		case LANG_PT:
			value.Engine()
			l := pt.New()
			uni_pt = ut.New(l, l)
			trans, _ := uni_pt.GetTranslator("pt")
			pt_translations.RegisterDefaultTranslations(value.GetValidate(), trans)
			break
		case LANG_PT_BR:
			value.Engine()
			l := pt_BR.New()
			uni_pt_BR = ut.New(l, l)
			trans, _ := uni_pt_BR.GetTranslator("pt_BR")
			pt_BR_translations.RegisterDefaultTranslations(value.GetValidate(), trans)
			break
		case LANG_RU:
			value.Engine()
			l := ru.New()
			uni_ru = ut.New(l, l)
			trans, _ := uni_ru.GetTranslator("ru")
			ru_translations.RegisterDefaultTranslations(value.GetValidate(), trans)
			break
		case LANG_TR:
			value.Engine()
			l := ru.New()
			uni_tr = ut.New(l, l)
			trans, _ := uni_tr.GetTranslator("tr")
			tr_translations.RegisterDefaultTranslations(value.GetValidate(), trans)
			break
		case LANG_ZH:
			value.Engine()
			l := zh.New()
			uni_zh = ut.New(l, l)
			trans, _ := uni_zh.GetTranslator("zh")
			zh_translations.RegisterDefaultTranslations(value.GetValidate(), trans)
			break
		case LANG_ZH_TW:
			value.Engine()
			l := zh_TW.New()
			uni_zh_TW = ut.New(l, l)
			trans, _ := uni_zh_TW.GetTranslator("zh_TW")
			zh_TW_translations.RegisterDefaultTranslations(value.GetValidate(), trans)
			break
		default:
			value.Engine()
			l := en.New()
			uni_en = ut.New(l, l)
			trans, _ := uni_en.GetTranslator("en")
			en_translations.RegisterDefaultTranslations(value.GetValidate(), trans)
			break
		}
	}
}