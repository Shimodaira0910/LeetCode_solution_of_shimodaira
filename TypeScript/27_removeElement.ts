// Test case
const nums1: number[] = [3,2,2,3];
const val: number = 2;

const removeElement = (numb: number[], val: number) => {
    let index: number = 0;
    for(let num of numb){
        if(num !== val){
            numb[index++] = num;
        }
    }
    return index;
}

console.log(removeElement(nums1, val));