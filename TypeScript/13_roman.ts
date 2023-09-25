// Test case
const s: string = 'III';
const s2: string = 'LVIII';

const romanToInt = (s: string) => {
    const romanNumbers = new Map([
        ['I',1],
        ['V',5],
        ['X',10],
        ['L',50],
        ['C',100],
        ['D',500],
        ['M',1000]
    ]);

    let s_split: string[] = s.split("");
    let ans: number = 0;

    for(let i = 0; i < s_split.length; i++){
        if(s_split[i] === 'I' && s_split[i + 1] === 'V'){
            ans += 4
            continue;
        } else if (s_split[i] === 'I' && s_split[i + 1] === 'X'){
            ans += 9
            continue;
        }

        if(s_split[i] === 'X' && s_split[i + 1] === 'L'){
            ans += 40
            continue;
        } else if (s_split[i] === 'X' && s_split[i + 1] === 'C'){
            ans += 90
            continue;
        }

        if(s_split[i] === 'C' && s_split[i + 1] === 'D'){
            ans += 400
            continue;
        } else if (s_split[i] === 'C' && s_split[i + 1] === 'M'){
            ans += 900
            continue;
        }

    }

}