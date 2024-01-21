package main

import (
	"fmt"
)

type Dimensions struct {
	Height         int `bson:"height"`
	Weight         int `bson:"weight"`
	ShoulderLength int `bson:"shoulder_length"`
}

type Preferences struct {
	Accessibility         string                `bson:"accessibility"`
	Occupation            string                `bson:"occupation"`
	PreferencesController PreferencesController `bson:"preferences_controller"`
}

type Person struct {
	Name        string                 `bson:"name"`
	Age         int                    `bson:"age"`
	Gender      string                 `bson:"gender"`
	Dimensions  Dimensions             `bson:"dimensions"`
	Preferences Preferences            `bson:"preferences"`
	Settings    map[string]interface{} `bson:"settings"`
}

type PreferencesController struct {
	DarkMode string `bson:"dark_mode"`
	Friendly string `bson:"friendly"`
}

func main() {
	person := &Person{}
	fields := Fields(person, "bson")

	fmt.Println(fields)
	//for _, field := range fields {
	//	tag := field.Tag("bson")
	//	fmt.Println(tag)
	//}
}
