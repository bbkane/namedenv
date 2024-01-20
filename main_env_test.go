package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.bbkane.com/warg"
)

func TestBuildApp(t *testing.T) {
	app := buildApp()

	if err := app.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestEnvCreate(t *testing.T) {
	updateGolden := os.Getenv("NAMEDENV_TEST_UPDATE_GOLDEN") != ""

	dbFile, err := os.CreateTemp(os.TempDir(), "namedenv-test-")
	require.NoError(t, err)
	err = dbFile.Close()
	require.NoError(t, err)

	t.Log("dbFile:", dbFile.Name())

	const zeroTime = "0001-01-01T00:00:00Z"

	tests := []struct {
		name            string
		app             *warg.App
		args            []string
		expectActionErr bool
	}{
		{
			name: "01_envCreate",
			app:  buildApp(),
			args: []string{
				"namedenv", "env", "create",
				"--sqlite-dsn", dbFile.Name(),
				"--name", "env_name",
				"--create-time", zeroTime,
				"--update-time", zeroTime,
			},
			expectActionErr: false,
		},
		{
			name: "02_envShow",
			app:  buildApp(),
			args: []string{
				"namedenv", "env", "show",
				"--sqlite-dsn", dbFile.Name(),
				"--name", "env_name",
				"--timezone", "utc",
			},
			expectActionErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			warg.GoldenTest(
				t,
				*tt.app,
				warg.GoldenTestOpts{
					UpdateGolden:    updateGolden,
					ExpectActionErr: false,
				},
				warg.OverrideArgs(tt.args),
				warg.OverrideLookupFunc(warg.LookupMap(nil)),
			)
		})
	}
}

func TestEnvDelete(t *testing.T) {
	updateGolden := os.Getenv("NAMEDENV_TEST_UPDATE_GOLDEN") != ""

	dbFile, err := os.CreateTemp(os.TempDir(), "namedenv-test-")
	require.NoError(t, err)
	err = dbFile.Close()
	require.NoError(t, err)

	t.Log("dbFile:", dbFile.Name())

	const zeroTime = "0001-01-01T00:00:00Z"

	tests := []struct {
		name            string
		app             *warg.App
		args            []string
		expectActionErr bool
	}{
		{
			name: "01_envCreate",
			app:  buildApp(),
			args: []string{
				"namedenv", "env", "create",
				"--sqlite-dsn", dbFile.Name(),
				"--name", "env_name",
				"--create-time", zeroTime,
				"--update-time", zeroTime,
			},
			expectActionErr: false,
		},
		{
			name: "02_envShow",
			app:  buildApp(),
			args: []string{
				"namedenv", "env", "show",
				"--sqlite-dsn", dbFile.Name(),
				"--name", "env_name",
				"--timezone", "utc",
			},
			expectActionErr: false,
		},
		{
			name: "03_envDelete",
			app:  buildApp(),
			args: []string{
				"namedenv", "env", "delete",
				"--sqlite-dsn", dbFile.Name(),
				"--name", "env_name",
			},
			expectActionErr: false,
		},
		{
			name: "04_envShow",
			app:  buildApp(),
			args: []string{
				"namedenv", "env", "show",
				"--sqlite-dsn", dbFile.Name(),
				"--name", "env_name",
				"--timezone", "utc",
			},
			expectActionErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			warg.GoldenTest(
				t,
				*tt.app,
				warg.GoldenTestOpts{
					UpdateGolden:    updateGolden,
					ExpectActionErr: tt.expectActionErr,
				},
				warg.OverrideArgs(tt.args),
				warg.OverrideLookupFunc(warg.LookupMap(nil)),
			)
		})
	}
}
