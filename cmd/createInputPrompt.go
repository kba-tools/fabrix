package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/vineshtk/fabrix/pkg/configs"
)

var (
	domainName            string
	numberOfOrganisations string
	channelName           string
	fabricVersion         string
)

func inputFromUser() {

	// theme := huh.ThemeBase16()
	// theme.Form = lipgloss.Style{}

	orgPeers := make(map[string]int)

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Domain Name").
				Placeholder("Enter the domain name here!").
				Value(&domainName),

			huh.NewInput().
				Title("Channel Name").
				Placeholder("Enter the channel name here!").
				Value(&channelName),
		),

		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("2.5.4", "2.5.9", "3.0")...).
				Title("Fabric Version").
				Description("Choose your fabric version").
				Validate(func(t string) error {
					if t == "3.0" {
						return fmt.Errorf("Uh oh! sorry we don't support 3.0 right now")
					}
					return nil
				}).
				Value(&fabricVersion),

			huh.NewInput().
				Title("Number of organisations").
				Placeholder("Enter the number of organisations here!").
				Validate(func(t string) error {
					if _, err := strconv.Atoi(t); err != nil {
						return fmt.Errorf("please enter a valid integer")
					}
					return nil
				}).
				Value(&numberOfOrganisations),
		),
	).WithLayout(huh.LayoutColumns(2)).WithShowHelp(true).WithTheme(huh.ThemeCharm()).Run()

	if err != nil {
		if err == huh.ErrUserAborted {
			os.Exit(0)
			fmt.Println("see you!!!")
		}
		fmt.Println(err)
		os.Exit(1)
	}

	numOrgs, err := strconv.Atoi(numberOfOrganisations)
	if err != nil {
		return
	}

	for numOrgs > 0 {

		peerNum := ""
		orgName := ""

		err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title(fmt.Sprintf("Organisation %d Name", numOrgs)).
					Placeholder("Enter the organisation name here").
					Value(&orgName),

				huh.NewInput().
					Title(fmt.Sprintf("Number of peers of organisation %d", numOrgs)).
					Placeholder("Enter number of peers").
					Value(&peerNum).
					Validate(func(t string) error {
						if _, err := strconv.Atoi(t); err != nil {
							return fmt.Errorf("please enter a valid integer")
						}
						return nil
					}),
			),
		).WithShowHelp(true).WithTheme(huh.ThemeCharm()).Run()

		if err != nil {
			if err == huh.ErrUserAborted {
				os.Exit(0)
				fmt.Println("see you!!!")
			}
			fmt.Println(err)
			os.Exit(1)
		}

		peerInt, _ := strconv.Atoi(peerNum)
		orgPeers[orgName] = peerInt
		numOrgs -= 1

	}

	fmt.Println("new org details",
		orgPeers,
		domainName,
		numberOfOrganisations,
		channelName,
		fabricVersion)

	orgsNumber, _ := strconv.Atoi(numberOfOrganisations)

	configs.CreateConfigs(domainName, orgPeers, channelName, version, orgsNumber)

}
