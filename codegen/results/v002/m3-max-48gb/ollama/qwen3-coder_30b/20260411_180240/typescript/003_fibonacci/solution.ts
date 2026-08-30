const N = parseInt(process.argv[2]);

if (isNaN(N) || N < 1) {
    process.exit(0);
}

let a = 1;
let b = 1;

console.log(a);
if (b <= N) {
    console.log(b);
}

while (true) {
    const next = a + b;
    if (next > N) {
        break;
    }
    console.log(next);
    a = b;
    b = next;
}