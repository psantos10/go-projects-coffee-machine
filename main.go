package main

import (
	"fmt"
)

const (
	espressoRequiredWater         = 250
	espressoRequiredCoffeeBeans   = 16
	espressoPrice                 = 4
	latteRequiredWater            = 350
	latteRequiredMilk             = 75
	latteRequiredCoffeeBeans      = 20
	lattePrice                    = 7
	cappuccinoRequiredWater       = 200
	cappuccinoRequiredMilk        = 100
	cappuccinoRequiredCoffeeBeans = 12
	cappuccinoPrice               = 6
)

type CoffeeMachine struct {
	waterQty       int
	milkQty        int
	coffeeBeansQty int
	disposableCups int
	money          int
}

func (cm CoffeeMachine) printState() {
	fmt.Println()
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", cm.waterQty)
	fmt.Printf("%d ml of milk\n", cm.milkQty)
	fmt.Printf("%d g of coffee beans\n", cm.coffeeBeansQty)
	fmt.Printf("%d disposable cups\n", cm.disposableCups)
	fmt.Printf("$%d of money\n", cm.money)
	fmt.Println()
}

func (cm *CoffeeMachine) buyCoffee(coffeeType int) {
	switch coffeeType {
	case 1:
		cm.waterQty -= 250
		cm.coffeeBeansQty -= 16
		cm.money += 4
	case 2:
		cm.waterQty -= 350
		cm.milkQty -= 75
		cm.coffeeBeansQty -= 20
		cm.money += 7
	case 3:
		cm.waterQty -= 200
		cm.milkQty -= 100
		cm.coffeeBeansQty -= 12
		cm.money += 6
	}

	cm.disposableCups--
}

func (cm *CoffeeMachine) fillCoffeeMachine(waterQty, milkQty, coffeeBeansQty, disposableCups int) {
	cm.waterQty += waterQty
	cm.milkQty += milkQty
	cm.coffeeBeansQty += coffeeBeansQty
	cm.disposableCups += disposableCups
}

func (cm *CoffeeMachine) takeMoney() {
	fmt.Printf("I gave you $%d\n", cm.money)
	cm.money = 0
}

func (cm CoffeeMachine) hasEnoughResources(coffeeType, quantity int) (bool, string) {
	if cm.disposableCups < quantity {
		return false, "disposable cups"
	}

	switch coffeeType {
	case 1:
		if cm.waterQty < espressoRequiredWater*quantity {
			return false, "water"
		}

		if cm.coffeeBeansQty < espressoRequiredCoffeeBeans*quantity {
			return false, "coffee beans"
		}

		return true, ""
	case 2:
		if cm.waterQty < latteRequiredWater*quantity {
			return false, "water"
		}

		if cm.milkQty < latteRequiredMilk*quantity {
			return false, "milk"
		}

		if cm.coffeeBeansQty < latteRequiredCoffeeBeans*quantity {
			return false, "coffee beans"
		}

		return true, ""
	case 3:
		if cm.waterQty < cappuccinoRequiredWater*quantity {
			return false, "water"
		}

		if cm.milkQty < cappuccinoRequiredMilk*quantity {
			return false, "milk"
		}

		if cm.coffeeBeansQty < cappuccinoRequiredCoffeeBeans*quantity {
			return false, "coffee beans"
		}

		return true, ""
	default:
		return false, ""
	}
}

func processCoffeeMachineAction(action string, coffeeMachine *CoffeeMachine) {
	switch action {
	case "buy":
		var coffeeType int
		fmt.Print("\nWhat do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:\n> ")
		fmt.Scan(&coffeeType)

		ok, resource := coffeeMachine.hasEnoughResources(coffeeType, 1)

		if !ok {
			fmt.Printf("Sorry, not enough %s!\n\n", resource)
			return
		}

		fmt.Printf("I have enough resources, making you a coffee!\n\n")
		coffeeMachine.buyCoffee(coffeeType)
	case "fill":
		var waterQty, milkQty, coffeeBeansQty, disposableCups int
		fmt.Print("Write how many ml of water do you want to add:\n> ")
		fmt.Scan(&waterQty)

		fmt.Print("Write how many ml of milk do you want to add:\n> ")
		fmt.Scan(&milkQty)

		fmt.Print("Write how many grams of coffee beans do you want to add:\n> ")
		fmt.Scan(&coffeeBeansQty)

		fmt.Print("Write how many disposable cups of coffee do you want to add:\n> ")
		fmt.Scan(&disposableCups)

		coffeeMachine.fillCoffeeMachine(waterQty, milkQty, coffeeBeansQty, disposableCups)
	case "take":
		coffeeMachine.takeMoney()
	case "remaining":
		coffeeMachine.printState()
	default:
		fmt.Println("Invalid action")
	}
}

func main() {
	var action string
	var coffeeMachine CoffeeMachine = CoffeeMachine{
		waterQty:       400,
		milkQty:        540,
		coffeeBeansQty: 120,
		disposableCups: 9,
		money:          550,
	}

	for {
		fmt.Print("Write action (buy, fill, take, remaining, exit):\n> ")
		fmt.Scan(&action)

		if action == "exit" {
			break
		}

		processCoffeeMachineAction(action, &coffeeMachine)
	}
}
