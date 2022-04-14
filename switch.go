package main;

import "fmt";

func main() {
	var day_number int = 4;

	switch day_number {
	case 1:
		fmt.Println("Saturday");
		break;
	case 2:
		fmt.Println("Sunday");
		break;
	case 3:
		fmt.Println("Monday");
		break;
	case 4:
		fmt.Println("Tuesday");
		break;
	case 5:
		fmt.Println("Wednesday");
		break;
	case 6:
		fmt.Println("Thursday");
		break;
	case 7:
		fmt.Println("Friday");
		break;
	}
}