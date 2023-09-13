package config

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestApplication_Add(t *testing.T) {
	testCases := []struct {
		caseName      string
		name          string
		configuration any
		expected      any
	}{
		{"add string config", "key_string", "string value", "string value"},
		{"add integer config", "key_integer", 100, 100},
		{"add float config", "key_float", 10.12, 10.12},
		{"add boolean config", "key_boolean", true, true},
	}

	appConfig := &Application{
		viper: viper.New(),
	}

	for _, tc := range testCases {
		appConfig.Add(tc.name, tc.configuration)
		actual := appConfig.viper.Get(tc.name)

		assert.Equal(
			t,
			actual,
			tc.expected,
			fmt.Sprintf("failed on [%s]", tc.caseName),
		)
	}
}

func TestApplication_Env(t *testing.T) {
	testCases := []struct {
		caseName     string
		name         string
		defaultValue any
		expected     any
	}{
		{"get existing string config", "key_string", nil, "string value"},
		{"get existing integer config", "key_integer", nil, 100},
		{"get existing float config", "key_float", nil, 10.12},
		{"get existing boolean config", "key_boolean", nil, true},
		{"get non-existing config with default", "key_random", "default string", "default string"},
	}

	appConfig := &Application{
		viper: viper.New(),
	}

	func() {
		appConfig.viper.Set("key_string", "string value")
		appConfig.viper.Set("key_integer", 100)
		appConfig.viper.Set("key_float", 10.12)
		appConfig.viper.Set("key_boolean", true)
	}()

	for _, tc := range testCases {
		actual := appConfig.Env(tc.name, tc.defaultValue)

		assert.Equal(
			t,
			actual,
			tc.expected,
			fmt.Sprintf("failed on [%s]", tc.caseName),
		)
	}
}

func TestApplication_Inspect(t *testing.T) {
	appConfig := &Application{
		viper: viper.New(),
	}

	func() {
		appConfig.viper.Set("key_string", "string value")
		appConfig.viper.Set("key_integer", 100)
		appConfig.viper.Set("key_float", 10.12)
		appConfig.viper.Set("key_boolean", true)
	}()

	expected := map[string]interface{}{
		"key_string":  "string value",
		"key_integer": 100,
		"key_boolean": true,
		"key_float":   10.12,
	}
	actual := appConfig.Inspect()

	assert.Equal(
		t,
		actual,
		expected,
		"failed on [Inspecting Application Config]",
	)
}

func TestNewApplication(t *testing.T) {
	file, err := os.CreateTemp("", "test-*.env")
	if nil != err {
		log.Fatal("Environment temporary file not created.")
	}
	defer file.Close()

	file.Write([]byte("CONFIG_1=test\nCONFIG_2=1\nCONFIG_3=true"))

	appConfig := NewApplication(file.Name())
	expected := map[string]interface{}{
		"config_1": "test",
		"config_2": "1",
		"config_3": "true",
	}
	actual := appConfig.Inspect()
	assert.Equal(
		t,
		actual,
		expected,
		"failed on [Creating New Application]",
	)

}
