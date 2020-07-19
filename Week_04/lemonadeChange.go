package main

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, v := range bills {
		switch v {
		case 5:
			five++
		case 10:
			five--
			ten++
		case 20:
			if ten > 0 {
				ten--
				five--
			} else {
				five -= 3
			}
		}

		if five < 0 || ten < 0 {
			return false
		}
	}

	return true
}

func main() {

}
