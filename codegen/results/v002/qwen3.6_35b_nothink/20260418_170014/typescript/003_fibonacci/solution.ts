const n = parseInt(process.argv[2], 10);

if (n < 1) {
    process.exit(0);
}

let a = 1;
let b = 1;

while (a <= n) {
    console.log(a);
    const next = a + b;
    a = b;
    b = next;
}