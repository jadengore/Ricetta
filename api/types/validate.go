package types

import (
	"fmt"
	"github.com/jadengore/goconfig"
	"github.com/mccoyst/validate"
	"reflect"
	"regexp"
)

type vc struct {
	// Config Constants
	MAX_HANDLE_LENGTH       int
	MIN_PASS_LENGTH         int
	MAX_PASS_LENGTH         int
	MAX_RECIPE_TITLE_LENGTH int
	MAX_RECIPE_NOTES_LENGTH int
	MAX_INGREDIENT_LENGTH   int
	MAX_TAG_LENGTH          int
	MAX_STEP_LENGTH         int
	AUTH_TOKEN_EXPIRES      int64

	// Regex Constants
	HANDLE_REGEX    *regexp.Regexp
	EMAIL_REGEX     *regexp.Regexp
	URL_REGEX       *regexp.Regexp
	TIME_UNIT_REGEX *regexp.Regexp
}

// Main Ricetta Validator
// Validator function found in respective data
// package in types

type RicettaValidator struct {
	Validator validate.V
	Constants vc
}

func NewValidator(config *goconfig.ConfigFile) *RicettaValidator {
	vd := RicettaValidator{}
	vd.Constants = initializeConstants(config)
	vd.Validator = validate.V{
		"handle":      vd.validateHandle,
		"email":       vd.validateEmail,
		"password":    vd.validatePassword,
		"timeunit":    vd.validateTimeUnit,
		"time":        vd.validateTime,
		"recipetitle": vd.validateRecipeTitle,
		"recipenotes": vd.validateRecipeNotes,
		"ingredient":  vd.validateIngredient,
		"url":         vd.validateURL,
		"tag":         vd.validateTag,
		"step":        vd.validateStep,

		// Required field Validation
		"existence": vd.validateExistence,
	}
	return &vd
}

func initializeConstants(config *goconfig.ConfigFile) vc {
	c := vc{}
	// Constants set by configuration
	c.MAX_HANDLE_LENGTH, _ = config.GetInt("global", "handle-length")
	c.MIN_PASS_LENGTH, _ = config.GetInt("global", "min-pass")
	c.MAX_PASS_LENGTH, _ = config.GetInt("global", "max-pass")
	c.MAX_RECIPE_TITLE_LENGTH, _ = config.GetInt("global", "max-recipe-title")
	c.MAX_RECIPE_NOTES_LENGTH, _ = config.GetInt("global", "max-recipe-notes")
	c.MAX_INGREDIENT_LENGTH, _ = config.GetInt("global", "max-ingredient-length")
	c.MAX_TAG_LENGTH, _ = config.GetInt("global", "max-tag-length")
	c.MAX_STEP_LENGTH, _ = config.GetInt("global", "max-step-length")

	// Using int64 for query layer
	c.AUTH_TOKEN_EXPIRES, _ = config.GetInt64("global", "auth-token-expires")

	// Regular expression constants
	c.HANDLE_REGEX = regexp.MustCompile(`^[\p{L}\p{M}][\d\p{L}\p{M}]*$`)
	c.EMAIL_REGEX = regexp.MustCompile(`^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	c.URL_REGEX = regexp.MustCompile(`(http(s)?:\/\/.)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&\/\/=]*)`)
	c.TIME_UNIT_REGEX = regexp.MustCompile(`(hr|min|sec|day|week)s?$`)
	return c
}

func (v RicettaValidator) validateExistence(i interface{}) error {
	if reflect.ValueOf(i).IsNil() {
		return fmt.Errorf("Required field omitted")
	} else {
		return nil
	}
}
