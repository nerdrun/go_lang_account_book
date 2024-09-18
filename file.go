package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	domains "account.com/test/domains"
	services "account.com/test/services"
)

func FileOptionPrompt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose option")
	fmt.Println("l - Load a file")
	fmt.Println("c - Create a file")
	fmt.Println("d - Delete a file")
	opt, _ := GetInput("---> ", reader)
	switch opt {
	case "l":
		files, err := services.LookupFolder()
		if err != nil {
			fmt.Println(err)
			FileOptionPrompt()
		}
		index := selectFile(files, reader)
		selectedFile := (*files)[index]
		f, err := services.LoadFile(selectedFile.Name())
		if err != nil {
			panic(err)
		}
		buf, err := io.ReadAll(f)

		if err != nil {
			panic(err)
		}

		var account domains.Account
		err = json.Unmarshal(buf, &account)
		if err != nil {
			panic(err)
		}
		opt := selectMenu(reader)
		displayItems(opt, &account)
	case "c":
		createFile(reader)
		FileOptionPrompt()
	case "d":
		services.DeleteFile()
	default:
		FileOptionPrompt()
	}
}

func itemFormat(name string, value float32) {
	fmt.Printf("%-25s%.2f\n", name+":", value)
}

func displayInputItems(input *domains.Input) {
	fmt.Printf("Your Input\n\n")

	income := &input.Income
	fmt.Printf("Income %s\n", strings.Repeat("=", 22))
	itemFormat(income.Salary.Name, income.Salary.Value)
	itemFormat(income.Tip.Name, income.Tip.Value)
	itemFormat(income.Bonus.Name, income.Bonus.Value)
	itemFormat(income.Commission.Name, income.Commission.Value)
	itemFormat(income.Other.Name, income.Other.Value)
	fmt.Println()

	other := &input.Other
	fmt.Printf("Other %s\n", strings.Repeat("=", 23))
	itemFormat(other.Transferred.Name, other.Transferred.Value)
	itemFormat(other.Interest.Name, other.Interest.Value)
	itemFormat(other.Dividend.Name, other.Dividend.Value)
	itemFormat(other.Gift.Name, other.Gift.Value)
	itemFormat(other.Refund.Name, other.Refund.Value)
	itemFormat(other.Installment.Name, other.Installment.Value)
	itemFormat(other.Balance.Name, other.Balance.Value)
}

func displayOutputItems(output *domains.Output) {
	fmt.Println("Your Output\n\n")

	children := output.Children
	fmt.Printf("Children %s\n", strings.Repeat("=", 20))
	itemFormat(children.Activity.Name, children.Activity.Value)
	itemFormat(children.PocketMoney.Name, children.PocketMoney.Value)
	itemFormat(children.Medical.Name, children.Medical.Value)
	itemFormat(children.Childcare.Name, children.Childcare.Value)
	itemFormat(children.Clothes.Name, children.Clothes.Value)
	itemFormat(children.School.Name, children.School.Value)
	itemFormat(children.Toy.Name, children.Toy.Value)
	itemFormat(children.Other.Name, children.Other.Value)
	fmt.Println()

	dept := output.Debt
	fmt.Printf("Dept %s\n", strings.Repeat("=", 24))
	itemFormat(dept.CreditCard.Name, dept.CreditCard.Value)
	itemFormat(dept.StudentLoan.Name, dept.StudentLoan.Value)
	itemFormat(dept.OtherLoan.Name, dept.OtherLoan.Value)
	itemFormat(dept.FederalTax.Name, dept.FederalTax.Value)
	itemFormat(dept.StateTax.Name, dept.StateTax.Value)
	itemFormat(dept.Other.Name, dept.Other.Value)
	fmt.Println()

	education := output.Education
	fmt.Printf("Education %s\n", strings.Repeat("=", 19))
	itemFormat(education.Tuition.Name, education.Tuition.Value)
	itemFormat(education.Book.Name, education.Book.Value)
	itemFormat(education.Lesson.Name, education.Lesson.Value)
	itemFormat(education.Other.Name, education.Other.Value)
	fmt.Println()

	activity := output.Activity
	fmt.Printf("Activity %s\n", strings.Repeat("=", 20))
	itemFormat(activity.Book.Name, activity.Book.Value)
	itemFormat(activity.Concert.Name, activity.Concert.Value)
	itemFormat(activity.Games.Name, activity.Games.Value)
	itemFormat(activity.Hobby.Name, activity.Hobby.Value)
	itemFormat(activity.Movie.Name, activity.Movie.Value)
	itemFormat(activity.Music.Name, activity.Music.Value)
	itemFormat(activity.Outdoor.Name, activity.Outdoor.Value)
	itemFormat(activity.Photography.Name, activity.Photography.Value)
	itemFormat(activity.Exercise.Name, activity.Exercise.Value)
	itemFormat(activity.Theater.Name, activity.Theater.Value)
	itemFormat(activity.TV.Name, activity.TV.Value)
	itemFormat(activity.Other.Name, activity.Other.Value)
	fmt.Println()

	dailyLife := output.DailyLife
	fmt.Printf("Daily life %s\n", strings.Repeat("=", 18))
	itemFormat(dailyLife.Grocery.Name, dailyLife.Grocery.Value)
	itemFormat(dailyLife.DiningOut.Name, dailyLife.DiningOut.Value)
	itemFormat(dailyLife.Goods.Name, dailyLife.Goods.Value)
	itemFormat(dailyLife.Clothes.Name, dailyLife.Clothes.Value)
	itemFormat(dailyLife.Laundry.Name, dailyLife.Laundry.Value)
	itemFormat(dailyLife.Beauty.Name, dailyLife.Beauty.Value)
	itemFormat(dailyLife.Subscription.Name, dailyLife.Subscription.Value)
	itemFormat(dailyLife.Other.Name, dailyLife.Other.Value)
	fmt.Println()

	gift := output.Gift
	fmt.Printf("Gift %s\n", strings.Repeat("=", 24))
	itemFormat(gift.Gift.Name, gift.Gift.Value)
	itemFormat(gift.Donation.Name, gift.Donation.Value)
	itemFormat(gift.Other.Name, gift.Other.Value)
	fmt.Println()

	health := output.Health
	fmt.Printf("Health %s\n", strings.Repeat("=", 22))
	itemFormat(health.GeneralCare.Name, health.GeneralCare.Value)
	itemFormat(health.SpecializedCare.Name, health.SpecializedCare.Value)
	itemFormat(health.Medication.Name, health.Medication.Value)
	itemFormat(health.Other.Name, health.Other.Value)
	fmt.Println()

	housing := output.Housing
	fmt.Printf("Housing %s\n", strings.Repeat("=", 21))
	itemFormat(housing.Mortgage.Name, housing.Mortgage.Value)
	itemFormat(housing.PropertyTax.Name, housing.PropertyTax.Value)
	itemFormat(housing.Furniture.Name, housing.Furniture.Value)
	itemFormat(housing.Garden.Name, housing.Garden.Value)
	itemFormat(housing.Supply.Name, housing.Supply.Value)
	itemFormat(housing.Maintenance.Name, housing.Maintenance.Value)
	itemFormat(housing.Renovation.Name, housing.Renovation.Value)
	itemFormat(housing.Moving.Name, housing.Moving.Value)
	itemFormat(housing.Other.Name, housing.Other.Value)
	fmt.Println()

	insurance := output.Insurance
	fmt.Printf("Insurance %s\n", strings.Repeat("=", 19))
	itemFormat(insurance.Automobile.Name, insurance.Automobile.Value)
	itemFormat(insurance.Health.Name, insurance.Health.Value)
	itemFormat(insurance.Home.Name, insurance.Home.Value)
	itemFormat(insurance.Life.Name, insurance.Life.Value)
	itemFormat(insurance.Other.Name, insurance.Other.Value)
	fmt.Println()

	pet := output.Pet
	fmt.Printf("Pet %s\n", strings.Repeat("=", 25))
	itemFormat(pet.PetFood.Name, pet.PetFood.Value)
	itemFormat(pet.VeterinaryCare.Name, pet.VeterinaryCare.Value)
	itemFormat(pet.Toy.Name, pet.Toy.Value)
	itemFormat(pet.Supply.Name, pet.Supply.Value)
	itemFormat(pet.Other.Name, pet.Other.Value)
	fmt.Println()

	technology := output.Technology
	fmt.Printf("Technology %s\n", strings.Repeat("=", 18))
	itemFormat(technology.Hosting.Name, technology.Hardware.Value)
	itemFormat(technology.Service.Name, technology.Service.Value)
	itemFormat(technology.Hardware.Name, technology.Hardware.Value)
	itemFormat(technology.Software.Name, technology.Software.Value)
	itemFormat(technology.Other.Name, technology.Other.Value)
	fmt.Println()

	transportation := output.Transportation
	fmt.Printf("Transportation %s\n", strings.Repeat("=", 14))
	itemFormat(transportation.Fuel.Name, transportation.Fuel.Value)
	itemFormat(transportation.Payment.Name, transportation.Payment.Value)
	itemFormat(transportation.Repair.Name, transportation.Repair.Value)
	itemFormat(transportation.Registration.Name, transportation.Registration.Value)
	itemFormat(transportation.Supply.Name, transportation.Supply.Value)
	itemFormat(transportation.Transportation.Name, transportation.Transportation.Value)
	itemFormat(transportation.Other.Name, transportation.Other.Value)
	fmt.Println()

	travel := output.Travel
	fmt.Printf("Travel %s\n", strings.Repeat("=", 22))
	itemFormat(travel.Airfare.Name, travel.Airfare.Value)
	itemFormat(travel.Hotel.Name, travel.Hotel.Value)
	itemFormat(travel.Food.Name, travel.Food.Value)
	itemFormat(travel.Transportation.Name, travel.Transportation.Value)
	itemFormat(travel.Activity.Name, travel.Activity.Value)
	itemFormat(travel.Other.Name, travel.Other.Value)
	fmt.Println()

	utility := output.Utility
	fmt.Printf("Utility %s\n", strings.Repeat("=", 21))
	itemFormat(utility.Phone.Name, utility.Phone.Value)
	itemFormat(utility.TV.Name, utility.TV.Value)
	itemFormat(utility.Internet.Name, utility.Internet.Value)
	itemFormat(utility.Electricity.Name, utility.Electricity.Value)
	itemFormat(utility.HeatingGas.Name, utility.HeatingGas.Value)
	itemFormat(utility.WaterSewage.Name, utility.WaterSewage.Value)
	itemFormat(utility.TrashDisposal.Name, utility.TrashDisposal.Value)
	itemFormat(utility.Other.Name, utility.Other.Value)
	fmt.Println()
}

func displayItems(opt string, account *domains.Account) {
	// output := account.Output
	switch opt {
	case "1":
		displayInputItems(&account.Input)
	case "2":
		displayOutputItems(&account.Output)
	}
}

func selectMenu(r *bufio.Reader) string {
	fmt.Println("What do you want to see?")
	fmt.Println("1) input")
	fmt.Println("2) output")
	opt, _ := GetInput(": ", r)
	return opt
}

func selectFile(files *[]fs.DirEntry, r *bufio.Reader) int {
	opt, _ := GetInput("Please select a file: ", r)
	index, err := services.SelectFile(opt, files, r)
	if err != nil {
		fmt.Println(err)
		return selectFile(files, r)
	}
	return index
}

func createFile(r *bufio.Reader) {
	opt, _ := GetInput("Please type a file name you want to create : ", r)
	err := services.CreateFile(opt, r)
	if err != nil {
		fmt.Println(err)
		createFile(r)
	}
}
