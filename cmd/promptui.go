package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetInput(label string, pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	// templates := &promptui.PromptTemplates{
	// 	Prompt:  "{{ . }} ",
	// 	Valid:   "{{ . | green }} ",
	// 	Invalid: "{{ . | red }} ",
	// 	Success: "{{ . | bold }} ",
	// }

	prompt := promptui.Prompt{
		Label: pc.label,
		// Templates:   templates,
		Validate:    validate,
		HideEntered: true,
	}

	result, err := prompt.Run()
	if err != nil {
		// color.Red("Prompt failed %v", err)
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	// color.Green("%s: %s", label, result)
	fmt.Printf("%s: %s\n", label, result)

	return result
}

func promptGetSelect(label string, pc promptContent) string {
	items := []string{"nestjs", "go-micro"}

	prompt := promptui.Select{
		Label:        pc.label,
		Items:        items,
		HideSelected: true,
	}

	_, result, err := prompt.Run()

	if err != nil {
		// color.Red("Prompt failed %v", err)
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	// color.Green("%s: %s", label, result)
	fmt.Printf("%s: %s\n", label, result)

	return result
}
