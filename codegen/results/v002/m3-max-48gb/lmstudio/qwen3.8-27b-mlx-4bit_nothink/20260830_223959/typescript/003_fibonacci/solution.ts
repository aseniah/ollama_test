const N = parseInt(process.argv[2] ?? "", 10);

if (isNaN(N) || N < 1) {
    process.exit(0);
}

const fibs: number[] = [];
let a = 1;
let b = 1;

fibs.push(a);
fibs.push(b);

let next: number;
while (true) {
    next = a + b;
    if (next > N) break;
    fibs.push(next);
    a = b;
    b = next;
}

for (const f of fibs) {
    console.log(f);
}