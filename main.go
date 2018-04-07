package main

import (
	"fmt"

	"github.com/spf13/viper"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

func defaultForJenkins() string {
	fmt.Println("jenkins url:", viper.AllKeys())
	return viper.GetString("JENKINS_URL")
}

func init() {
	err := viper.BindEnv("JENKINS_URL")
	if err != nil {
		panic(err)
	}
}

func main() {
	// the questions to ask
	qs := []*survey.Question{
		{
			Name: "jenkinsURL",
			Prompt: &survey.Input{
				Message: "What's the Jenkins URL?",
				Default: defaultForJenkins()},
			Validate: survey.Required,
		},
		{
			Name:     "jenkinsUser",
			Prompt:   &survey.Input{Message: "Jenkins User?"},
			Validate: survey.Required,
		},
		{
			Name:     "jenkinsPass",
			Prompt:   &survey.Password{Message: "Jenkins Pass"},
			Validate: survey.Required,
		},
		{
			Name:     "dcosURL",
			Prompt:   &survey.Input{Message: "DCOS URL?"},
			Validate: survey.Required,
		},
		{
			Name:     "dockerRegistry",
			Prompt:   &survey.Input{Message: "DOCKER REGISTRY?"},
			Validate: survey.Required,
		},
		{
			Name:     "imageName",
			Prompt:   &survey.Input{Message: "Image name?"},
			Validate: survey.Required,
		},
		{
			Name:     "gitUrl",
			Prompt:   &survey.Input{Message: "Git URL?"},
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

	fmt.Printf("%s %s.", answers.JenkinsURL, answers.JenkinsUser)
}
