package services

type Account struct {
	Input  Input
	Output Output
}

type Item struct {
	Name  string
	Value float32
}

type Input struct {
	Income Income
	Other  OtherIncome
}

type Income struct {
	Salary     Item
	Tip        Item
	Bonus      Item
	Commission Item
	Other      Item
}

type OtherIncome struct {
	Transferred Item
	Interest    Item
	Dividend    Item
	Gift        Item
	Refund      Item
	Installment Item
	Balance     Item
}

type Output struct {
	Children       Children
	Debt           Debt
	Education      Education
	Activity       Activity
	DailyLife      DailyLife
	Gift           Gift
	Health         Health
	Housing        Housing
	Insurance      Insurance
	Pet            Pet
	Technology     Technology
	Transportation Transportation
	Travel         Travel
	Utility        Utility
}

type Children struct {
	Activity    Item
	PocketMoney Item
	Medical     Item
	Childcare   Item
	Clothes     Item
	School      Item
	Toy         Item
	Other       Item
}

type Debt struct {
	CreditCard  Item
	StudentLoan Item
	OtherLoan   Item
	FederalTax  Item
	StateTax    Item
	Other       Item
}

type Education struct {
	Tuition Item
	Book    Item
	Lesson  Item
	Other   Item
}

type Activity struct {
	Book        Item
	Concert     Item
	Games       Item
	Hobby       Item
	Movie       Item
	Music       Item
	Outdoor     Item
	Photography Item
	Exercise    Item
	Theater     Item
	TV          Item
	Other       Item
}

type DailyLife struct {
	Grocery      Item
	DiningOut    Item
	Goods        Item
	Clothes      Item
	Laundry      Item
	Beauty       Item
	Subscription Item
	Other        Item
}

type Gift struct {
	Gift     Item
	Donation Item
	Other    Item
}

type Health struct {
	GeneralCare     Item
	SpecializedCare Item
	Medication      Item
	Emergency       Item
	Other           Item
}

type Housing struct {
	Mortgage    Item
	PropertyTax Item
	Furniture   Item
	Garden      Item
	Supply      Item
	Maintenance Item
	Renovation  Item
	Moving      Item
	Other       Item
}

type Insurance struct {
	Automobile Item
	Health     Item
	Home       Item
	Life       Item
	Other      Item
}

type Pet struct {
	PetFood        Item
	VeterinaryCare Item
	Toy            Item
	Supply         Item
	Other          Item
}

type Technology struct {
	Hosting  Item
	Service  Item
	Hardware Item
	Software Item
	Other    Item
}

type Transportation struct {
	Fuel           Item
	Payment        Item
	Repair         Item
	Registration   Item
	Supply         Item
	Transportation Item
	Other          Item
}

type Travel struct {
	Airfare        Item
	Hotel          Item
	Food           Item
	Transportation Item
	Activity       Item
	Other          Item
}

type Utility struct {
	Phone         Item
	TV            Item
	Internet      Item
	Electricity   Item
	HeatingGas    Item
	WaterSewage   Item
	TrashDisposal Item
	Other         Item
}

func GenerateAccount() Account {
	input := generateInput()
	output := generateOutput()
	return Account{
		Input:  input,
		Output: output,
	}
}

func generateItem(name string) Item {
	return Item{
		Name:  name,
		Value: 0,
	}
}

func generateInput() Input {
	return Input{
		Income: Income{
			Salary:     generateItem("salary"),
			Tip:        generateItem("tip"),
			Bonus:      generateItem("bonus"),
			Commission: generateItem("commission"),
			Other:      generateItem("other"),
		},
		Other: OtherIncome{
			Transferred: generateItem("transffered"),
			Interest:    generateItem("interest"),
			Dividend:    generateItem("dividend"),
			Gift:        generateItem("gift"),
			Refund:      generateItem("refund"),
			Installment: generateItem("installment"),
			Balance:     generateItem("balance"),
		},
	}
}

func generateOutput() Output {
	return Output{
		Children: Children{
			Activity:    generateItem("activity"),
			PocketMoney: generateItem("pocket money"),
			Medical:     generateItem("medical"),
			Childcare:   generateItem("childcare"),
			Clothes:     generateItem("clothes"),
			School:      generateItem("school"),
			Toy:         generateItem("toy"),
			Other:       generateItem("other"),
		},
		Debt: Debt{
			CreditCard:  generateItem("credit card"),
			StudentLoan: generateItem("student loan"),
			OtherLoan:   generateItem("other loan"),
			FederalTax:  generateItem("federal tax"),
			StateTax:    generateItem("state tax"),
			Other:       generateItem("other"),
		},
		Education: Education{
			Tuition: generateItem("tuition"),
			Book:    generateItem("book"),
			Lesson:  generateItem("lesson"),
			Other:   generateItem("other"),
		},
		Activity: Activity{
			Book:        generateItem("book"),
			Concert:     generateItem("concert"),
			Games:       generateItem("games"),
			Hobby:       generateItem("hobby"),
			Movie:       generateItem("movie"),
			Music:       generateItem("music"),
			Outdoor:     generateItem("outdoor"),
			Photography: generateItem("photography"),
			Exercise:    generateItem("exercise"),
			Theater:     generateItem("theater"),
			TV:          generateItem("tv"),
			Other:       generateItem("other"),
		},
		DailyLife: DailyLife{
			Grocery:      generateItem("grocery"),
			DiningOut:    generateItem("dining out"),
			Goods:        generateItem("goods"),
			Clothes:      generateItem("clothes"),
			Laundry:      generateItem("laundry"),
			Beauty:       generateItem("beauty"),
			Subscription: generateItem("subscription"),
			Other:        generateItem("other"),
		},
		Gift: Gift{
			Gift:     generateItem("gift"),
			Donation: generateItem("donation"),
			Other:    generateItem("other"),
		},
		Health: Health{
			GeneralCare:     generateItem("general care"),
			SpecializedCare: generateItem("specialized care"),
			Medication:      generateItem("medication"),
			Emergency:       generateItem("emergency"),
			Other:           generateItem("other"),
		},
		Housing: Housing{
			Mortgage:    generateItem("mortgage"),
			PropertyTax: generateItem("property tax"),
			Furniture:   generateItem("furniture"),
			Garden:      generateItem("garden"),
			Supply:      generateItem("supply"),
			Maintenance: generateItem("maintenance"),
			Renovation:  generateItem("renovation"),
			Moving:      generateItem("moving"),
			Other:       generateItem("other"),
		},
		Insurance: Insurance{
			Automobile: generateItem("automobile"),
			Health:     generateItem("health"),
			Home:       generateItem("home"),
			Life:       generateItem("life"),
			Other:      generateItem("other"),
		},
		Pet: Pet{
			PetFood:        generateItem("pet food"),
			VeterinaryCare: generateItem("veterinary care"),
			Toy:            generateItem("toy"),
			Supply:         generateItem("supply"),
			Other:          generateItem("other"),
		},
		Technology: Technology{
			Hosting:  generateItem("hosting"),
			Service:  generateItem("service"),
			Hardware: generateItem("hardware"),
			Software: generateItem("software"),
			Other:    generateItem("other"),
		},
		Transportation: Transportation{
			Fuel:           generateItem("fuel"),
			Payment:        generateItem("payment"),
			Repair:         generateItem("repair"),
			Registration:   generateItem("registration"),
			Supply:         generateItem("supply"),
			Transportation: generateItem("transportation"),
			Other:          generateItem("other"),
		},
		Travel: Travel{
			Airfare:        generateItem("airfare"),
			Hotel:          generateItem("hotel"),
			Food:           generateItem("food"),
			Transportation: generateItem("transportation"),
			Activity:       generateItem("activity"),
			Other:          generateItem("other"),
		},
		Utility: Utility{
			Phone:         generateItem("phone"),
			TV:            generateItem("tv"),
			Internet:      generateItem("internet"),
			Electricity:   generateItem("electricity"),
			HeatingGas:    generateItem("heating gas"),
			WaterSewage:   generateItem("water sewage"),
			TrashDisposal: generateItem("trash disposal"),
			Other:         generateItem("other"),
		},
	}
}
