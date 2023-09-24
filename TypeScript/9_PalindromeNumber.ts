// Test Patterns
const x : number = 121
const x2 : number = -121
const x3 : number = 10

const isPalindrome = (x: number) => {
    if(x.toString() === x.toString().split("").reverse().join("")){
        return true;
    } else {
        return false;
    }
}

console.log(isPalindrome(x));

// result 145ms 51.88mb