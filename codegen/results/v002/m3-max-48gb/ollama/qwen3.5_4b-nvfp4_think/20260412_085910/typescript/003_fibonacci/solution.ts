const n = parseInt(process.argv[2]);

if (n >= 1) {
    let a = 1;
    let b = 1;

    console.log(a);

    if (b <= n) {
        console.log(b);
    }

    while (true) {
        let c = a + b;
        if (c > n) {
            break;
        }
        console.log(c);
        a = b;
        b = c;
    }
}