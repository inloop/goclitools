package goclitools

import (
	"fmt"
	"testing"
)

func TestDependencyCheck(t *testing.T) {

	dep := Dependency{CheckCmd: "which ls", Name: "ls"}

	res, err := dep.Check()

	if err != nil {
		t.Fatalf("expected nil error, received: %s", err.Error())
	}
	if !res {
		t.Fatal("expected true result")
	}

}

func TestInvalidDependencyCheck(t *testing.T) {

	dep := Dependency{CheckCmd: "which blah2", Name: "blah2"}

	res, err := dep.Check()

	if err != nil {
		t.Fatalf("expected nil error, received: %s", err.Error())
	}
	if res {
		t.Fatal("expected false result")
	}

}

func TestDependencyInstallation(t *testing.T) {

	dep := Dependency{CheckCmd: "which blah", Name: "blah", InstallScripts: []DependencyScript{
		DependencyScriptString{Fn: "ln -s /usr/local/bin/go /usr/local/bin/blah"},
	}, UninstallScripts: []DependencyScript{
		DependencyScriptString{Fn: "rm /usr/local/bin/blah"},
	}}

	if err := dep.Install(); err != nil {
		fmt.Println(nil, err, err != nil, err == nil)
		t.Fatalf("expected nil error, received: %s", err.Error())
	}

	if err := dep.Uninstall(); err != nil {
		fmt.Println(err, err != nil, err == nil)
		t.Fatalf("expected nil error, received: %s", err.Error())
	}
}
