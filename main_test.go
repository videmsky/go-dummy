package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hello World!!" {
		t.Fatal("Test fail")
	}
}

func TestHealthCheck(t *testing.T) {
	if healthcheck() != "Health OK!" {
		t.Fatal("Test fail")
	}
}

func TestLivenessCheck(t *testing.T) {
	if livenesscheck() != "I am alive!!!" {
		t.Fatal("Test fail")
	}
}
