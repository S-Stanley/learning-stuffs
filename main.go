package main

import (
	"fmt"
	"math/rand"
	// "time"
)

var NUMBER_OF_QUESTION int = 5;

func getData() map[string]string {
	data := map[string]string {
		"etre": "ser",
		"etre (etre situé)": "estar",
		"avoir": "tener",
		"faire": "hacer",
		"aller": "ir",
		"pouvoir": "poder",
		"savoir": "saber",
		"mettre": "poner",
		"dire": "decir",
		"vouloir": "querer",
		"parler": "hablar",
		"donner": "dar",
		"voir (regarder)": "ver",
		"manger": "comer",
		"prendre": "tomar",
		"vivre": "vivir",
		"avoir besoin de": "necessitar",
		"rester": "quedar",
		"venir": "venir",
	};
	return data;
}

func getMapKeys(data map[string]string) []string {
	var output []string;
	output = make([]string, len(data));
	var i int = 0;
	for key := range data {
		output[i] = key;
		i++;
	}
	return output
}

func getUserInput() string {
	var user_input string;
	fmt.Scanln(&user_input);
	return (user_input);
}

func play(data map[string]string) {
	var points = 0;
	for i := 0; i < NUMBER_OF_QUESTION; i++ {
		var key = getMapKeys(data)[rand.Intn(len(data))];
		var answer = data[key];
		fmt.Printf("Quel est la traduction de %s?\n", key);
		var userInput string = getUserInput();
		if  userInput == answer {
			fmt.Println("good");
			points += 1;
		} else {
			fmt.Printf("bad, the answer was %s\n", answer);
		}
	}
	fmt.Printf("You made %d on %d\n", points, NUMBER_OF_QUESTION);
}

// func main(){
// 	rand.Seed(time.Now().Unix());
// 	var data = getData();
// 	play(data);
// }