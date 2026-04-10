const [,, nStr] = process.argv;
const n = parseInt(nStr, 10);

if (n < 1) {
    console.log('');
    return;
}

let a = 1, b = 1;
while (b <= n) {
    console.log(b);
    [a, b] = [b, a + b];
}