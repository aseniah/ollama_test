const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 1) {
    process.exit(0);
}

if (N >= 1) {
    console.log(1);
}

if (N >= 1) {
    let a = 1;
    let b = 1;
    while (true) {
        const c = a + b;
        if (c > N) break;
        console.log(c);
        a = b;
        b = c;
    }
}