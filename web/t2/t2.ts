let isDone: boolean = false;

let howMuch: number = 0;

let name: string = 'foo bar';

console.log(
    `hello ${name}`
);

let list: number[] = [1, 2, 3];
let list1: Array<number> = [1, 2, 3];

let x: [string, number];
x = ['hello', 10];

enum Color {
    Red, Green, Blue
}

let c: Color = Color.Blue;
console.log(Color[c]);

let notSure: any = 0;
notSure = 'maybe a string';
notSure = false;

function func1(): void {
}

let u: undefined = undefined;
let n: null = null;

function func2(msg: string): never {
    throw new Error(msg);
}

let someValue: any = 'foo bar';
console.log((<string>someValue).length);
console.log((someValue as string).length);

if (false) {
    function foo() {
        return a; // runtime error
    }

    foo();
    let a;
}


let input = [1, 2];
let [first, second] = input;
