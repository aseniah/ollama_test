function main(): void {
    const N = parseInt(process.argv[2], 10);

    if (isNaN(N) || N < 1) {
        return;
    }

    // Generate Fibonacci numbers starting with 1, 1, 2, 3, 5, ...
    // Print all that are <= N
    if (N >= 1) {
        console.log(1);
        if (N >= 1) {
            // Second 1
            console.log(1);
        }
    }

    let a = 1;
    let b = 1;
    while (true) {
        const next = a + b;
        if (next > N) {
            break;
        }
        console.log(next);
        a = b;
        b = next;
    }
}

main();