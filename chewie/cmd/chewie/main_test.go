package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun_NoArgs(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{}, &stdout, &stderr, nil)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "chewie - Chewie CLI") {
		t.Errorf("expected help text in stdout, got: %q", stdout.String())
	}
}

func TestRun_ShortHelp(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"-h"}, &stdout, &stderr, nil)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "chewie - Chewie CLI") {
		t.Errorf("expected help text in stdout, got: %q", stdout.String())
	}
}

func TestRun_LongHelp(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--help"}, &stdout, &stderr, nil)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "chewie - Chewie CLI") {
		t.Errorf("expected help text in stdout, got: %q", stdout.String())
	}
}

func TestRun_ShortVersion(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"-v"}, &stdout, &stderr, nil)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), version) {
		t.Errorf("expected version %q in stdout, got: %q", version, stdout.String())
	}
}

func TestRun_LongVersion(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--version"}, &stdout, &stderr, nil)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), version) {
		t.Errorf("expected version %q in stdout, got: %q", version, stdout.String())
	}
}

func TestRun_InvalidFlag(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"--unknown"}, &stdout, &stderr, nil)
	if code != 2 {
		t.Errorf("expected exit code 2, got %d", code)
	}
}

func TestRun_HelpContainsUsage(t *testing.T) {
	var stdout, stderr bytes.Buffer
	run([]string{"-h"}, &stdout, &stderr, nil)
	out := stdout.String()
	for _, expected := range []string{"USAGE:", "OPTIONS", "DESCRIPTION:", "EXAMPLES:", "EXIT CODES:"} {
		if !strings.Contains(out, expected) {
			t.Errorf("help text missing section %q", expected)
		}
	}
}
