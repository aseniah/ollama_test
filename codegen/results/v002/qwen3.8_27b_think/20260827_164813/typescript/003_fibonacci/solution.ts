const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 1) {
    process.exit(0);
}

const fibs: number[] = [];
let a = 1, b = 1;

while (a <= N) {
    fibs.push(a);
    const next = a + b;
    a = b;
    b = next;
}

for (const f of fibs) {
    console.log(f);
}