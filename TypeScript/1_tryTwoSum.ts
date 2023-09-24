// try too fast and save memory
// 

// Test Patterns
const newNums: number[] = [2,7,11,15];
const newTarget: number = 9;

const twoSumNew = (nums: number[], target: number) => {
    const map = new Map<number, number>();
    let result: number[] = [];

    for(let i = 0; i <= nums.length ; i++){
        map.set(i, nums[i]);
    }

    for(const [key, value] of map){
        let targetNum = target - value;
        if(nums.includes(targetNum)){
            
        }
    }


}

console.log(twoSumNew(newNums, newTarget));

