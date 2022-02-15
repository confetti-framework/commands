package config

import (
	"github.com/confetti-framework/support/env"
	"golang.org/x/text/language"
	"os"
	"time"
)

var App = struct {
	Name,
	Url,
	AssetUrl,
	LineSeparator,
	Key,
	Env string
	OsArgs   []string
	Port     int
	Cipher   string
	Debug    bool
	Timezone *time.Location
	Locale,
	FallbackLocale,
	FakerLocale language.Tag
}{

	/*
	   |--------------------------------------------------------------------------
	   | Application Name
	   |--------------------------------------------------------------------------
	   |
	   | This value is the name of your application. This value is used when the
	   | framework needs to place the application's name in a notification or
	   | any other location as required by the application or its packages.
	   |
	*/
	Name: env.StringOr("APP_NAME", "Confetti"),

	/*
		|--------------------------------------------------------------------------
		| Encryption Key
		|--------------------------------------------------------------------------
		|
		| This key is used by the Confetti encrypter service and should be set
		| to a random, 32 character string, otherwise these encrypted strings
		| will not be safe. Please do this before deploying an application!
		|
	*/
	Key: env.StringOr("APP_KEY", ""),

	/*
		|--------------------------------------------------------------------------
		| Cipher type
		|--------------------------------------------------------------------------
		|
		| Source encrypted values are encrypted using OpenSSL and this
		| cipher. Furthermore, all encrypted values are signed with a message
		| authentication code (MAC) to detect any modifications to the encrypted
		| string.
		|
	*/
	Cipher: "AES-256-CBC",

	/*
	   |--------------------------------------------------------------------------
	   | Application IsEnvironment
	   |--------------------------------------------------------------------------
	   |
	   | This value determines the "environment" your application is currently
	   | running in. This may determine how you prefer to configure various
	   | services the application utilizes. Set this in your ".env" file.
	   |
	*/
	Env: env.StringOr("APP_ENV", "production"),

	/*
	   |--------------------------------------------------------------------------
	   | Application Command-line Arguments
	   |--------------------------------------------------------------------------
	   |
	   | Args hold the command-line arguments, starting with the program name.
	   |
	*/
	OsArgs: os.Args,

	/*
	   |--------------------------------------------------------------------------
	   | Application Debug Mode
	   |--------------------------------------------------------------------------
	   |
	   | When your application is in debug mode, detailed error messages with
	   | stack traces will be shown on every error that occurs within your
	   | application. If disabled, a simple generic error page is shown.
	   |
	*/
	Debug: env.BoolOr("APP_DEBUG", false),

	/*
		|--------------------------------------------------------------------------
		| Application Timezone
		|--------------------------------------------------------------------------
		|
		| Here you may specify the default timezone for your application, which
		| will be used by the PHP date and date-time functions. We have gone
		| ahead and set this to a sensible default for you out of the box.
		|
	*/
	Timezone: time.UTC,

	/*
		|--------------------------------------------------------------------------
		| Application Locale Configuration
		|--------------------------------------------------------------------------
		|
		| The application locale determines the default locale that will be used
		| by the translation service provider. You are free to set this value
		| to any of the locales which will be supported by the application.
		|
	*/
	Locale: env.LangOr("LOCALE", language.AmericanEnglish),

	/*
		|--------------------------------------------------------------------------
		| Application Fallback Locale
		|--------------------------------------------------------------------------
		|
		| The fallback locale determines the locale to use when the current one
		| is not available. You may change the value to correspond to any of
		| the language folders that are provided through your application.
		|
	*/
	FallbackLocale: env.LangOr("FALLBACK_LOCALE", language.AmericanEnglish),

	/*
		|--------------------------------------------------------------------------
		| Faker Locale
		|--------------------------------------------------------------------------
		|
		| This locale will be used by the Faker PHP library when generating fake
		| data for your database seeds. For example, this will be used to get
		| localized telephone numbers, street address information and more.
		|
	*/
	FakerLocale: env.LangOr("FAKE_LOCALE", language.AmericanEnglish),
}
