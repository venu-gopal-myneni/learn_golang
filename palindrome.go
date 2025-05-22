package main

import "strings"
import "fmt"
import "strconv"
func isPalindrome(x int) bool {
    // x_str := string(x)
    x_str := strconv.Itoa(x)

    var x_rev  []string
    for i := len(x_str)-1;i>=0;i--{
        x_rev = append(x_rev,string(x_str[i]))
    }
    x_rev_new := strings.Join(x_rev, "")
    if x_str == x_rev_new {
        return true
    }
    return false
    
}