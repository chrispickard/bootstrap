package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

func bindEnvs(param string, usage string) {
	err := viper.BindEnv(param)
	if err != nil {
		panic(err)
	}
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	pflag.String(param, "", usage)
}

func init() {
	bindEnvs("jenkins-url", "The jenkins url")
	bindEnvs("jenkins-user", "The jenkins user")
	bindEnvs("jenkins-pass", "The jenkins password. Insecure, favor interactive use instead")
	bindEnvs("dcos-url", "The URL for the dcos master")
	bindEnvs("docker-registry", "The URL for the docker registry")
	bindEnvs("image-name", "The name to take the resulting image with")
	bindEnvs("git-url", "The git url you have pushed the code to")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	// the questions to ask
	qs := []*survey.Question{
		{
			Name: "jenkinsURL",
			Prompt: &survey.Input{
				Message: "What's the Jenkins URL?",
				Default: viper.GetString("jenkins-url")},
			Validate: survey.Required,
		},
		{
			Name: "jenkinsUser",
			Prompt: &survey.Input{Message: "Jenkins User?",
				Default: viper.GetString("jenkins-user")},
			Validate: survey.Required,
		},
		{
			Name:     "jenkinsPass",
			Prompt:   &survey.Password{Message: "Jenkins Pass"},
			Validate: survey.Required,
		},
		{
			Name: "dcosURL",
			Prompt: &survey.Input{Message: "DCOS URL?",
				Default: viper.GetString("dcos-url")},
			Validate: survey.Required,
		},
		{
			Name: "dockerRegistry",
			Prompt: &survey.Input{Message: "DOCKER REGISTRY?",
				Default: viper.GetString("docker-registry")},
			Validate: survey.Required,
		},
		{
			Name: "imageName",
			Prompt: &survey.Input{Message: "Image name?",
				Default: viper.GetString("image-name")},
			Validate: survey.Required,
		},
		{
			Name: "gitUrl",
			Prompt: &survey.Input{Message: "Git URL?",
				Default: viper.GetString("git-url")},
			Validate: survey.Required,
		},
	}
	// the answers will be written to this struct
	answers := struct {
		JenkinsURL     string
		JenkinsUser    string
		JenkinsPass    string
		DCOSURL        string
		DockerRegistry string
		ImageName      string
		GitURL         string
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("answers: %+v\n", answers)
}
