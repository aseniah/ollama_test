const main = () => {
    // Read command line argument
    const arg = process.argv[2];
    
    if (!arg) {
        return;
    }

    const n = parseInt(arg, 10);

    // If N is not a valid number or less than 1, print nothing
    if (isNaN(n) || n < 1) {
        return;
    }

    let a = 1;
    let b = 1;

    // Print the first Fibonacci number if it is within range
    if (a <= n) {
        console.log(a);
    }

    // Generate subsequent numbers
    while (b <= n) {
        console.log(b);
        const next = a + b;
        a = b;
        b = next;
    }
};

main();