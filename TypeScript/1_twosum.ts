// 1. Two Sum
// 140ms 44.78mb

const nums: number[] = [3,3];
const target: number = 6

const twoSum = (nums: number[], target: number) => {
    let pos: number = 0;
    let length: number = 1;
    let ans: number[] = [];

    while(true){
        if(length === nums.length){
            pos += 1;
            length = pos + 1;
            continue;
        }else if(nums[pos] + nums[length] == target){
            ans.push(pos, length);
            return ans
        }else{
            length += 1;
        }
    }
}

console.log(twoSum(nums, target));