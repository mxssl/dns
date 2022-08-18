//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

// BinaryName ...
var BinaryName = "dns"

// DockerRegistry ...
var DockerRegistry = "mxssl"

// Build the app
func Build() error {
	if err := sh.Run("go", "build", "-v", "-o", BinaryName); err != nil {
		return err
	}
	fmt.Printf("%s is successfully built\n", BinaryName)
	return nil
}

// Lint the app
func Lint() error {
	if err := sh.RunV("golangci-lint", "run"); err != nil {
		return err
	}
	return nil
}

// Clean delete compiled binary
func Clean() error {
	if !fileExists(BinaryName) {
		return fmt.Errorf("cannnot delete binary: %s doesn't exist", BinaryName)
	}
	if err := sh.Run("rm", "-f", BinaryName); err != nil {
		return err
	}
	fmt.Printf("%s is successfully deleted\n", BinaryName)
	return nil
}

// DockerBuild build container with latest tag
func DockerBuild() error {
	containerWithTag := fmt.Sprintf(DockerRegistry + "/" + BinaryName + ":" + "latest")
	fmt.Printf("Image: %s\n", containerWithTag)

	if err := sh.RunV("docker", "build", "--tag", containerWithTag, "."); err != nil {
		return err
	}

	return nil
}

// DockerTestRun test run latest container
func DockerTestRun() error {
	containerWithTag := fmt.Sprintf(DockerRegistry + "/" + BinaryName + ":" + "latest")
	fmt.Printf("Image: %s\n", containerWithTag)

	if err := sh.RunV("docker", "container", "run", containerWithTag, "dns", "a", "google.com"); err != nil {
		return err
	}

	return nil
}

// DockerRelease build and push container to the registry
func DockerRelease() error {
	fmt.Printf("Registry: %s\n", DockerRegistry)
	tag, err := getLastGitTag()
	if err != nil {
		return err
	}
	fmt.Printf("Git tag: %s\n", tag)

	containerWithTag := fmt.Sprintf(DockerRegistry + "/" + BinaryName + ":" + tag)

	if err := sh.RunV("docker", "build", "--tag", containerWithTag, "."); err != nil {
		return err
	}

	if err := sh.RunV("docker", "push", containerWithTag); err != nil {
		return err
	}

	return nil
}

// GitHubReleaseDryRun goreleaser dry run
func GitHubReleaseDryRun() error {
	tag, err := getLastGitTag()
	if err != nil {
		return err
	}
	fmt.Printf("Git tag: %s\n", tag)

	_, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		fmt.Println("env variable GITHUB_TOKEN is not set\n")
	}

	if err := sh.RunV("goreleaser", "release", "--rm-dist", "--snapshot", "--skip-publish"); err != nil {
		return err
	}
	return nil
}

// GitHubRelease run goreleaser
func GitHubRelease() error {
	tag, err := getLastGitTag()
	if err != nil {
		return err
	}
	fmt.Printf("Git tag: %s\n", tag)

	_, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		fmt.Println("env variable GITHUB_TOKEN is not set\n")
	}

	if err := sh.RunV("goreleaser", "release", "--rm-dist"); err != nil {
		return err
	}
	return nil
}

// check that file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// obtain latest git tag
func getLastGitTag() (string, error) {
	tag, err := sh.Output("git", "describe", "--abbrev=0", "--tags")
	if err != nil {
		return "", nil
	}
	return tag, nil
}
