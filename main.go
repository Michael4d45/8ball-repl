package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

func getAnswer(text string) string {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	return answers[rand.Intn(len(answers))]
}

func main() {
	rand.Seed(42)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("I am a Magic 8 ball, ask me a question")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1) // for windows

		matched, _ := regexp.MatchString(`.*\?$`, text)
		if matched {
			fmt.Println(getAnswer(text))
		}

		if strings.Compare("exit", text) == 0 {
			break
		}
	}
}
