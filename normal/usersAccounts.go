package main

import (
	"fmt"
	"sort"
)

/*
You have two lists of numbers: users and accounts.
Each users[i] is the ID of a user, and accounts[i] is the ID of an account that user has.
A user can have multiple accounts, and an account can be shared by multiple users.
Your goal is to write a function to find all the users who have the exact same set of accounts as at least one other user.
The function should return a list of the IDs of these users.

Example 1:

users 	 = 	[1, 2, 3, 4, 2, 5, 1]

accounts = 	[1, 1, 2, 2, 3, 4, 3]

Output: [[1, 2], [3, 4]]

Explanation:

Users 1 and 2 both have accounts 1 and 3.
Users 3 and 4 both have account 2.
Example 2:

users 	 = [1, 2, 3, 4, 2, 1, 1]

accounts = [1, 1, 2, 2, 3, 4, 3]

Output: [[3, 4]]

Explanation:

Users 1 and 2 both have accounts 1 and 3, but user 1 also has account 4, so they don't have the exact same set of accounts.
Users 3 and 4 both have account 2.
*/

func main() {
	users := []int{1, 2, 3, 4, 2, 5, 1, 2, 1}
	accounts := []int{3, 1, 2, 2, 11, 4, 1, 3, 11}
	fmt.Println(findUsersWithSameAccounts(users, accounts)) // [1 2] [3 4]
	users = []int{1, 2, 3, 4, 2, 1, 1}
	accounts = []int{1, 1, 2, 2, 3, 4, 3}
	fmt.Println(findUsersWithSameAccounts(users, accounts)) // [3 4]
	users = []int{1, 2, 3, 4, 5, 6}
	accounts = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(findUsersWithSameAccounts(users, accounts)) // []
}

func findUsersWithSameAccounts(users, accounts []int) [][]int {
	userToAccounts := make(map[int][]int)
	for i, user := range users {
		account := accounts[i]
		userToAccounts[user] = append(userToAccounts[user], account)
	}

	// Sort accounts for each user
	for _, accs := range userToAccounts {
		sort.Ints(accs)
	}

	accToUsers := make(map[string][]int,0)
	for users, accs := range userToAccounts{
		accountString := fmt.Sprintf("%v", accs)
		if _, ok:= accToUsers[accountString]; !ok {
			accToUsers[accountString] = make([]int,0)
		}
		accToUsers[accountString] = append(accToUsers[accountString], users)
	}

	result := make([][]int,0)
	for _, users := range accToUsers {
		if len(users) > 1 {
			result = append(result, users)
		}
	}
	
	return result
}
