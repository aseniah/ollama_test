import * as process from 'process';

const N = parseInt(process.argv[2], 10);

if (N < 1) {
    process.exit(0);
}

let a = 1;
let b = 1;

while (a <= N) {
    console.log(a);
    const next = a + b;
    a = b;
    b = next;
}