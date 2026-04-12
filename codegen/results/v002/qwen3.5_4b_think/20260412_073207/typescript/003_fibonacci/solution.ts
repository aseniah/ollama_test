const args = process.argv;

const n = parseInt(args[2]);

if (n < 1) {
    process.exit(0);
}

let a = 1;
let b = 1;

console.log(a);

while (b <= n) {
    console.log(b);
    const next = a + b;
    a = b;
    b = next;
}