package main

import(

	//For printing to console and scanln command for input
	"fmt"
	//For executing the clear console command 
	"os"
	"os/exec"
	//For string validation when user inputs a value
	"strings"
)

/*
All Entities in the game are represented by boolean values 
If the entity boolean equals true they are on the left side of the river
If the entity boolean equals false they are on the right side of the river
*/
var chicken bool
var corn bool
var fox bool
var man bool

/*
Boolean values to handle game logic
gameOver - Is the game over?
youWon - Did the player win the game? i.e are all the entity values false meaning they are all on the right side of the river
first - boolean to handle if this is the first time running the game, if it is then the previously chosen value will not be printed 
since the user has not chosen a value yet

lastChoice - the previous choice the user made, declared outside of the printChoices function to ensure it does not get redeclared 
every time the user makes a choice 
*/

var gameOver bool
var youWon bool
var first bool
var lastChoice string

//Main Function

func main(){
	Init()
	for(!gameOver){
		printIntroText()
		printChoices()
		clear()
	}
	if(youWon){
		fmt.Println("***********************************")
		fmt.Println("**************YOU WON**************")
		fmt.Println("***********************************")
	}else{
		fmt.Println("Game over... Try again? (y/n)")
		var answer string
		fmt.Scanln(&answer)
		if(answer == "y"){
			restart()
		}
	}
}

//Init function - sets the values to their initial value at the start of the game
func Init(){
	first = true
	gameOver = false
	youWon = false
    corn = true
    chicken = true
    fox = true
    man = true
}


//Prints gameplay info and overview to the player
func printIntroText(){
	fmt.Println("The controls to this game are simple... Enter the item you want to move to the other side and press the Enter key...")
	fmt.Println("If you move an item or an animal the man has to go with them so don't leave the chicken alone with the corn or the fox alone with the chicken...")
}


//Prints the choices the player can make as well as displaying the current state of the entities
func printChoices(){
	var choice string
	if(!first){
		printLastChoice(lastChoice)
	}
	manResultat := fmt.Sprintf("Man: %s, ", checkSide(man))
	cornResultat := fmt.Sprintf("Corn: %s, ", checkSide(corn))
	chickenResultat := fmt.Sprintf("Chicken: %s, ", checkSide(chicken))
	foxResultat := fmt.Sprintf("Fox: %s", checkSide(fox))
    fmt.Println(manResultat + cornResultat + chickenResultat, foxResultat)
	fmt.Println("Who do you want to move?")
	fmt.Scanln(&choice)
	lastChoice = choice
	fmt.Println(gameMove(choice))
	clear()
	first = false
}


//Checks what side the entity is on and returns a string to represent the left and right side of the river
func checkSide(entity bool) string{
	if(entity == true){
		return "Left"
	}else{
		return "Right"
	}
}


//Game Logic
func gameMove(input string) string{
	input = strings.ToLower(input)
	toPrint := fmt.Sprintf("You moved the %s to the other side...", input)
	switch input{
		case "man":
			man = !man
		case "corn":
			if(isOnSameSide(corn)){
				corn = !corn
				man = !man
			}else{
				toPrint = fmt.Sprintf("The man is not on the same side of the river as the %s, move the man to the other side", input)
			}
			
		case "chicken":
			if(isOnSameSide(chicken)){
				chicken = !chicken
				man = !man
			}else{
				toPrint = fmt.Sprintf("The man is not on the same side of the river as the %s, move the man to the other side", input)
			}
		case "fox":
			if(isOnSameSide(fox)){
				fox = !fox
				man = !man
			}else{
				toPrint = fmt.Sprintf("The man is not on the same side of the river as the %s, move the man to the other side", input)
			}
		default:
			toPrint = fmt.Sprintf("%s is not a valid input... Try again", input)
	}
	checkGameState()
	return toPrint
}
//Print the last choice the player made TODO: Clean this up and detangle game logic from previous function
func printLastChoice(input string){ 
	input = strings.ToLower(input)
	toPrint := fmt.Sprintf("You moved %s to the other side...", input)
	switch input{
		case "man":
			
		case "corn":
			if(isOnSameSide(corn)){
			
			}else{
				toPrint = fmt.Sprintf("The man is not on the same side of the river as %s, move the man to the other side", input)
			}
			
		case "chicken":
			if(isOnSameSide(chicken)){
			
			}else{
				toPrint = fmt.Sprintf("The man is not on the same side of the river as %s, move the man to the other side", input)
			}
		case "fox":
			if(isOnSameSide(fox)){
			
			}else{
				toPrint = fmt.Sprintf("The man is not on the same side of the river as %s, move the man to the other side", input)
			}
		default:
			toPrint = fmt.Sprintf("%s is not a valid input... Try again", input)
	}
	fmt.Println(toPrint)
}

//Checks if an entity is the same side as the man, therefore the entity can be moved
func isOnSameSide(entity bool) bool{
	if (entity == man){
		return true
	}else{
		return false
	}
}


//Checks at the end of every round if the game has ended or if the user has won
func checkGameState(){
	if((chicken == corn && man != chicken) || (fox == chicken && man != fox)){
		gameOver = true
	}else if(man == false && corn == false && chicken == false && fox == false){
		gameOver = true
		youWon = true
	}
}


//Clears the console every round
func clear(){
	cmdName := "clear"
	cmd := exec.Command(cmdName)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//Restart the game function TODO: DOES NOT WORK PLEASE FIX
func restart(){
	cmdName := "go run main.go"
	cmd := exec.Command(cmdName)
	cmd.Stdout = os.Stdout
	cmd.Run()
}