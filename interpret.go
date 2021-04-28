package leetcode

func interpret(command string) string {
	res := ""
	for i := range command {
		if (command[i] >= 'A' && command[i] <= 'Z') || (command[i] >= 'a' && command[i] <= 'z') {
			res = res + string(command[i])
		} else if command[i] == '(' && command[i+1] == ')' {
			res = res + "o"
		}
	}
	return res
}
