package sega_cmd

import (
	"errors"
	"fmt"
)

func ChoiceCmd(text string) (bool, error) {
	var choiceCmd string

	fmt.Printf("%s y/n", text)

	_, err := fmt.Scan(&choiceCmd)

	if err != nil {
		return false, err
	}

	if choiceCmd == "y" {
		return true, nil
	} else if choiceCmd != "y" {
		return false, nil
	}

	return false, errors.New("неверный ввод: нужно ввести y или n")
}
