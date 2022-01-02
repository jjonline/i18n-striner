// Code generated by "i18n-stringer -type RuneOne,RuneMulti,RuneMap"; DO NOT EDIT.

package test_switch

import (
	"context"
	"fmt"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the i18n-stringer command to generate them again.
	var x [1]struct{}
	_ = x[CostRuneOT1-20]
	_ = x[CostRuneOT2-21]
	_ = x[CostRuneOT3-22]
}

const (
	_RuneOne_En_name   = "single const onesingle const twosingle const three"
	_RuneOne_ZhHk_name = "单个区间常量1单个区间常量2单个区间常量3"
)

var (
	_RuneOne_En_index   = [...]uint8{0, 16, 32, 50}
	_RuneOne_ZhHk_index = [...]uint8{0, 19, 38, 57}
)

// _transOne translate one CONST
func (i RuneOne) _transOne(locale string) string {
	i -= 20
	if i < 0 || i >= RuneOne(len(_RuneOne_En_index)-1) {
		return "RuneOne[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
	}

	switch locale {
	case "en":
		return _RuneOne_En_name[_RuneOne_En_index[i]:_RuneOne_En_index[i+1]]
	case "zh-hk":
		return _RuneOne_ZhHk_name[_RuneOne_ZhHk_index[i]:_RuneOne_ZhHk_index[i+1]]
	default:
		return ""
	}
}

// _RuneOne_supported All supported locales and text offset information
var _RuneOne_supported = map[string]int{"en": 0, "zh-hk": 1}

// _RuneOne_defaultLocale default locale
// generated pass by i18n-stringer flag -defaultlocale, Don't assign directly
var _RuneOne_defaultLocale = "en"

// _RuneOne_ctxKey Key from context.Context Value get locale
// generated pass by i18n-stringer flag -ctxkey, Don't assign directly
var _RuneOne_ctxKey = "i18nLocale"

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method Error.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the fmt.Stringer interface, so that you can output it directly by package fmt,
//  - If you understand the above mechanism then you can use this method with confidence
func (i RuneOne) String() string {
	return i._trans(_RuneOne_defaultLocale)
}

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method String.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the error interface, so that you can return the value as an error,
//  - If you understand the above mechanism then you can use this method with confidence
func (i RuneOne) Error() string {
	return i._trans(_RuneOne_defaultLocale)
}

// Wrap another error with locale set for i18n TYPE Const
//  - err another error
//  - locale i18n locale name
//  - args optional formatting component
func (i RuneOne) Wrap(err error, locale string, args ...RuneOne) *I18nRuneOneErrorWrap {
	return &I18nRuneOneErrorWrap{err: err, origin: i, locale: locale, args: args}
}

// WrapWithContext wrap another error with context.Context set for i18n TYPE Const
//  - ctx context with Value use Key from _RuneOne_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - err another error
//  - args optional formatting component
func (i RuneOne) WrapWithContext(ctx context.Context, err error, args ...RuneOne) *I18nRuneOneErrorWrap {
	return &I18nRuneOneErrorWrap{err: err, origin: i, locale: _RuneOne_localeFromCtxWithFallback(ctx), args: args}
}

// I18nRuneOneErrorWrap type i18n error wrapper
//   WARNING
//   This struct ONLY used to wrap the CONST generated by the i18n-stringer tool,
//   Pass easily obtain internationalized translations through Error, String, Translate
//   WARNING
type I18nRuneOneErrorWrap struct {
	err    error     // wrap another error
	origin RuneOne   // custom shaping type Val
	locale string    // i18n locale set
	args   []RuneOne // formatted output replacement component
}

// Translate get translated string
func (e *I18nRuneOneErrorWrap) Translate() string {
	return e.origin.Trans(e.locale, e.args...)
}

// String implement fmt.Stringer, get translated string use Translate
func (e *I18nRuneOneErrorWrap) String() string {
	return e.Translate()
}

// Error struct as error, get translated string use Translate
func (e *I18nRuneOneErrorWrap) Error() string {
	return e.Translate()
}

// Format string form inside error and TOML define
//  - this method will be formatted wrap error. Only for development and debugging
func (e *I18nRuneOneErrorWrap) Format() string {
	if e.err == nil {
		return e.Error()
	}
	return fmt.Sprintf("%s (%s)", e.Error(), e.err.Error())
}

// Value get original type value
func (e *I18nRuneOneErrorWrap) Value() RuneOne {
	return e.origin
}

// Unwrap an error. Get the error inside
func (e *I18nRuneOneErrorWrap) Unwrap() error {
	return e.err
}

// IsLocaleSupport Check if the specified locale is supported
func (i RuneOne) IsLocaleSupport(locale string) bool {
	return _RuneOne_isLocaleSupport(locale)
}

// Lang get target translate text use context.Context
//  - ctx  context with Value use Key from _RuneOne_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - args Optional placeholder replacement value
func (i RuneOne) Lang(ctx context.Context, args ...RuneOne) string {
	return i._trans(_RuneOne_localeFromCtxWithFallback(ctx), args...)
}

// Trans get target translate text use specified language locale identifier
//  - locale specified language locale identifier, need pass by IsLocaleSupport
//  - args Optional placeholder replacement value
func (i RuneOne) Trans(locale string, args ...RuneOne) string {
	if !_RuneOne_isLocaleSupport(locale) {
		locale = _RuneOne_defaultLocale
	}
	return i._trans(locale, args...)
}

func _RuneOne_isLocaleSupport(locale string) bool {
	_, ok := _RuneOne_supported[locale]
	return ok
}

// _RuneOne_localeFromCtxWithFallback retrieves and returns language locale name from context.
// It returns default locale when _RuneOne_isLocaleSupport is false
func _RuneOne_localeFromCtxWithFallback(ctx context.Context) string {
	if ctx == nil {
		return _RuneOne_defaultLocale
	}
	v := ctx.Value(_RuneOne_ctxKey)
	if v == nil {
		return _RuneOne_defaultLocale
	}
	if vv, ok := v.(string); ok && _RuneOne_isLocaleSupport(vv) {
		return vv
	}
	return _RuneOne_defaultLocale
}

// _trans trustworthy parameters inside method
func (i RuneOne) _trans(locale string, args ...RuneOne) string {
	msg := i._transOne(locale)
	if len(args) > 0 {
		var com []interface{}
		for _, arg := range args {
			com = append(com, arg._transOne(locale))
		}
		return fmt.Sprintf(msg, com...)
	}
	return msg
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the i18n-stringer command to generate them again.
	var x [1]struct{}
	_ = x[CostRuneMT1-1]
	_ = x[CostRuneMT2-2]
	_ = x[CostRuneMT3-3]
	_ = x[CostRuneMT4-1003]
	_ = x[CostRuneMT5-1004]
}

const (
	_RuneMulti_En_name_0   = "1muilt rune one2muilt rune two3muilt rune three"
	_RuneMulti_ZhHk_name_0 = "多个常量一多个常量二多个常量三"
	_RuneMulti_En_name_1   = "4muilt rune four5muilt rune five"
	_RuneMulti_ZhHk_name_1 = "多个常量四多个常量五"
)

var (
	_RuneMulti_En_index_0   = [...]uint8{0, 15, 30, 47}
	_RuneMulti_ZhHk_index_0 = [...]uint8{0, 15, 30, 45}
	_RuneMulti_En_index_1   = [...]uint8{0, 16, 32}
	_RuneMulti_ZhHk_index_1 = [...]uint8{0, 15, 30}
)

// _transOne translate one CONST
func (i RuneMulti) _transOne(locale string) string {
	switch locale {
	case "en":
		switch {
		case 1 <= i && i <= 3:
			i -= 1
			return _RuneMulti_En_name_0[_RuneMulti_En_index_0[i]:_RuneMulti_En_index_0[i+1]]
		case 1003 <= i && i <= 1004:
			i -= 1003
			return _RuneMulti_En_name_1[_RuneMulti_En_index_1[i]:_RuneMulti_En_index_1[i+1]]
		default:
			return "RuneMulti[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
		}
	case "zh-hk":
		switch {
		case 1 <= i && i <= 3:
			i -= 1
			return _RuneMulti_ZhHk_name_0[_RuneMulti_ZhHk_index_0[i]:_RuneMulti_ZhHk_index_0[i+1]]
		case 1003 <= i && i <= 1004:
			i -= 1003
			return _RuneMulti_ZhHk_name_1[_RuneMulti_ZhHk_index_1[i]:_RuneMulti_ZhHk_index_1[i+1]]
		default:
			return "RuneMulti[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
		}
	default:
		return ""
	}
}

// _RuneMulti_supported All supported locales and text offset information
var _RuneMulti_supported = map[string]int{"en": 0, "zh-hk": 1}

// _RuneMulti_defaultLocale default locale
// generated pass by i18n-stringer flag -defaultlocale, Don't assign directly
var _RuneMulti_defaultLocale = "en"

// _RuneMulti_ctxKey Key from context.Context Value get locale
// generated pass by i18n-stringer flag -ctxkey, Don't assign directly
var _RuneMulti_ctxKey = "i18nLocale"

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method Error.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the fmt.Stringer interface, so that you can output it directly by package fmt,
//  - If you understand the above mechanism then you can use this method with confidence
func (i RuneMulti) String() string {
	return i._trans(_RuneMulti_defaultLocale)
}

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method String.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the error interface, so that you can return the value as an error,
//  - If you understand the above mechanism then you can use this method with confidence
func (i RuneMulti) Error() string {
	return i._trans(_RuneMulti_defaultLocale)
}

// Wrap another error with locale set for i18n TYPE Const
//  - err another error
//  - locale i18n locale name
//  - args optional formatting component
func (i RuneMulti) Wrap(err error, locale string, args ...RuneMulti) *I18nRuneMultiErrorWrap {
	return &I18nRuneMultiErrorWrap{err: err, origin: i, locale: locale, args: args}
}

// WrapWithContext wrap another error with context.Context set for i18n TYPE Const
//  - ctx context with Value use Key from _RuneMulti_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - err another error
//  - args optional formatting component
func (i RuneMulti) WrapWithContext(ctx context.Context, err error, args ...RuneMulti) *I18nRuneMultiErrorWrap {
	return &I18nRuneMultiErrorWrap{err: err, origin: i, locale: _RuneMulti_localeFromCtxWithFallback(ctx), args: args}
}

// I18nRuneMultiErrorWrap type i18n error wrapper
//   WARNING
//   This struct ONLY used to wrap the CONST generated by the i18n-stringer tool,
//   Pass easily obtain internationalized translations through Error, String, Translate
//   WARNING
type I18nRuneMultiErrorWrap struct {
	err    error       // wrap another error
	origin RuneMulti   // custom shaping type Val
	locale string      // i18n locale set
	args   []RuneMulti // formatted output replacement component
}

// Translate get translated string
func (e *I18nRuneMultiErrorWrap) Translate() string {
	return e.origin.Trans(e.locale, e.args...)
}

// String implement fmt.Stringer, get translated string use Translate
func (e *I18nRuneMultiErrorWrap) String() string {
	return e.Translate()
}

// Error struct as error, get translated string use Translate
func (e *I18nRuneMultiErrorWrap) Error() string {
	return e.Translate()
}

// Format string form inside error and TOML define
//  - this method will be formatted wrap error. Only for development and debugging
func (e *I18nRuneMultiErrorWrap) Format() string {
	if e.err == nil {
		return e.Error()
	}
	return fmt.Sprintf("%s (%s)", e.Error(), e.err.Error())
}

// Value get original type value
func (e *I18nRuneMultiErrorWrap) Value() RuneMulti {
	return e.origin
}

// Unwrap an error. Get the error inside
func (e *I18nRuneMultiErrorWrap) Unwrap() error {
	return e.err
}

// IsLocaleSupport Check if the specified locale is supported
func (i RuneMulti) IsLocaleSupport(locale string) bool {
	return _RuneMulti_isLocaleSupport(locale)
}

// Lang get target translate text use context.Context
//  - ctx  context with Value use Key from _RuneMulti_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - args Optional placeholder replacement value
func (i RuneMulti) Lang(ctx context.Context, args ...RuneMulti) string {
	return i._trans(_RuneMulti_localeFromCtxWithFallback(ctx), args...)
}

// Trans get target translate text use specified language locale identifier
//  - locale specified language locale identifier, need pass by IsLocaleSupport
//  - args Optional placeholder replacement value
func (i RuneMulti) Trans(locale string, args ...RuneMulti) string {
	if !_RuneMulti_isLocaleSupport(locale) {
		locale = _RuneMulti_defaultLocale
	}
	return i._trans(locale, args...)
}

func _RuneMulti_isLocaleSupport(locale string) bool {
	_, ok := _RuneMulti_supported[locale]
	return ok
}

// _RuneMulti_localeFromCtxWithFallback retrieves and returns language locale name from context.
// It returns default locale when _RuneMulti_isLocaleSupport is false
func _RuneMulti_localeFromCtxWithFallback(ctx context.Context) string {
	if ctx == nil {
		return _RuneMulti_defaultLocale
	}
	v := ctx.Value(_RuneMulti_ctxKey)
	if v == nil {
		return _RuneMulti_defaultLocale
	}
	if vv, ok := v.(string); ok && _RuneMulti_isLocaleSupport(vv) {
		return vv
	}
	return _RuneMulti_defaultLocale
}

// _trans trustworthy parameters inside method
func (i RuneMulti) _trans(locale string, args ...RuneMulti) string {
	msg := i._transOne(locale)
	if len(args) > 0 {
		var com []interface{}
		for _, arg := range args {
			com = append(com, arg._transOne(locale))
		}
		return fmt.Sprintf(msg, com...)
	}
	return msg
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the i18n-stringer command to generate them again.
	var x [1]struct{}
	_ = x[ConstRuneMaT1-1000]
	_ = x[ConstRuneMaT2-2000]
	_ = x[ConstRuneMaT3-3000]
	_ = x[ConstRuneMaT4-4000]
	_ = x[ConstRuneMaT5-5000]
	_ = x[ConstRuneMaT6-6000]
	_ = x[ConstRuneMaT7-7000]
	_ = x[ConstRuneMaT8-8000]
	_ = x[ConstRuneMaT9-9000]
	_ = x[ConstRuneMaT10-10000]
	_ = x[ConstRuneMaT11-11000]
}

const (
	_RuneMap_En_name   = "1map 12map 23map 34map 45map 56map 67map 78map 89map 910map 1011map 11"
	_RuneMap_ZhHk_name = "1地图一2地图二3地图三4地图四5地图五6地图六7地图七8地图八9地图九10地图十117地图十一"
)

var (
	_RuneMap_En_map = map[RuneMap]string{
		1000:  _RuneMap_En_name[0:6],
		2000:  _RuneMap_En_name[6:12],
		3000:  _RuneMap_En_name[12:18],
		4000:  _RuneMap_En_name[18:24],
		5000:  _RuneMap_En_name[24:30],
		6000:  _RuneMap_En_name[30:36],
		7000:  _RuneMap_En_name[36:42],
		8000:  _RuneMap_En_name[42:48],
		9000:  _RuneMap_En_name[48:54],
		10000: _RuneMap_En_name[54:62],
		11000: _RuneMap_En_name[62:70],
	}
	_RuneMap_ZhHk_map = map[RuneMap]string{
		1000:  _RuneMap_ZhHk_name[0:10],
		2000:  _RuneMap_ZhHk_name[10:20],
		3000:  _RuneMap_ZhHk_name[20:30],
		4000:  _RuneMap_ZhHk_name[30:40],
		5000:  _RuneMap_ZhHk_name[40:50],
		6000:  _RuneMap_ZhHk_name[50:60],
		7000:  _RuneMap_ZhHk_name[60:70],
		8000:  _RuneMap_ZhHk_name[70:80],
		9000:  _RuneMap_ZhHk_name[80:90],
		10000: _RuneMap_ZhHk_name[90:101],
		11000: _RuneMap_ZhHk_name[101:116],
	}
)

// _transOne translate one CONST
func (i RuneMap) _transOne(locale string) string {
	switch locale {
	case "en":
		if str, ok := _RuneMap_En_map[i]; ok {
			return str
		}
		return "RuneMap[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
	case "zh-hk":
		if str, ok := _RuneMap_ZhHk_map[i]; ok {
			return str
		}
		return "RuneMap[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
	default:
		return ""
	}
}

// _RuneMap_supported All supported locales and text offset information
var _RuneMap_supported = map[string]int{"en": 0, "zh-hk": 1}

// _RuneMap_defaultLocale default locale
// generated pass by i18n-stringer flag -defaultlocale, Don't assign directly
var _RuneMap_defaultLocale = "en"

// _RuneMap_ctxKey Key from context.Context Value get locale
// generated pass by i18n-stringer flag -ctxkey, Don't assign directly
var _RuneMap_ctxKey = "i18nLocale"

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method Error.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the fmt.Stringer interface, so that you can output it directly by package fmt,
//  - If you understand the above mechanism then you can use this method with confidence
func (i RuneMap) String() string {
	return i._trans(_RuneMap_defaultLocale)
}

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method String.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the error interface, so that you can return the value as an error,
//  - If you understand the above mechanism then you can use this method with confidence
func (i RuneMap) Error() string {
	return i._trans(_RuneMap_defaultLocale)
}

// Wrap another error with locale set for i18n TYPE Const
//  - err another error
//  - locale i18n locale name
//  - args optional formatting component
func (i RuneMap) Wrap(err error, locale string, args ...RuneMap) *I18nRuneMapErrorWrap {
	return &I18nRuneMapErrorWrap{err: err, origin: i, locale: locale, args: args}
}

// WrapWithContext wrap another error with context.Context set for i18n TYPE Const
//  - ctx context with Value use Key from _RuneMap_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - err another error
//  - args optional formatting component
func (i RuneMap) WrapWithContext(ctx context.Context, err error, args ...RuneMap) *I18nRuneMapErrorWrap {
	return &I18nRuneMapErrorWrap{err: err, origin: i, locale: _RuneMap_localeFromCtxWithFallback(ctx), args: args}
}

// I18nRuneMapErrorWrap type i18n error wrapper
//   WARNING
//   This struct ONLY used to wrap the CONST generated by the i18n-stringer tool,
//   Pass easily obtain internationalized translations through Error, String, Translate
//   WARNING
type I18nRuneMapErrorWrap struct {
	err    error     // wrap another error
	origin RuneMap   // custom shaping type Val
	locale string    // i18n locale set
	args   []RuneMap // formatted output replacement component
}

// Translate get translated string
func (e *I18nRuneMapErrorWrap) Translate() string {
	return e.origin.Trans(e.locale, e.args...)
}

// String implement fmt.Stringer, get translated string use Translate
func (e *I18nRuneMapErrorWrap) String() string {
	return e.Translate()
}

// Error struct as error, get translated string use Translate
func (e *I18nRuneMapErrorWrap) Error() string {
	return e.Translate()
}

// Format string form inside error and TOML define
//  - this method will be formatted wrap error. Only for development and debugging
func (e *I18nRuneMapErrorWrap) Format() string {
	if e.err == nil {
		return e.Error()
	}
	return fmt.Sprintf("%s (%s)", e.Error(), e.err.Error())
}

// Value get original type value
func (e *I18nRuneMapErrorWrap) Value() RuneMap {
	return e.origin
}

// Unwrap an error. Get the error inside
func (e *I18nRuneMapErrorWrap) Unwrap() error {
	return e.err
}

// IsLocaleSupport Check if the specified locale is supported
func (i RuneMap) IsLocaleSupport(locale string) bool {
	return _RuneMap_isLocaleSupport(locale)
}

// Lang get target translate text use context.Context
//  - ctx  context with Value use Key from _RuneMap_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - args Optional placeholder replacement value
func (i RuneMap) Lang(ctx context.Context, args ...RuneMap) string {
	return i._trans(_RuneMap_localeFromCtxWithFallback(ctx), args...)
}

// Trans get target translate text use specified language locale identifier
//  - locale specified language locale identifier, need pass by IsLocaleSupport
//  - args Optional placeholder replacement value
func (i RuneMap) Trans(locale string, args ...RuneMap) string {
	if !_RuneMap_isLocaleSupport(locale) {
		locale = _RuneMap_defaultLocale
	}
	return i._trans(locale, args...)
}

func _RuneMap_isLocaleSupport(locale string) bool {
	_, ok := _RuneMap_supported[locale]
	return ok
}

// _RuneMap_localeFromCtxWithFallback retrieves and returns language locale name from context.
// It returns default locale when _RuneMap_isLocaleSupport is false
func _RuneMap_localeFromCtxWithFallback(ctx context.Context) string {
	if ctx == nil {
		return _RuneMap_defaultLocale
	}
	v := ctx.Value(_RuneMap_ctxKey)
	if v == nil {
		return _RuneMap_defaultLocale
	}
	if vv, ok := v.(string); ok && _RuneMap_isLocaleSupport(vv) {
		return vv
	}
	return _RuneMap_defaultLocale
}

// _trans trustworthy parameters inside method
func (i RuneMap) _trans(locale string, args ...RuneMap) string {
	msg := i._transOne(locale)
	if len(args) > 0 {
		var com []interface{}
		for _, arg := range args {
			com = append(com, arg._transOne(locale))
		}
		return fmt.Sprintf(msg, com...)
	}
	return msg
}
