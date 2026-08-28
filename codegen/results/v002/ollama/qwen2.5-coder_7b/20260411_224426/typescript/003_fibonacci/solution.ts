function fibonacci(n: number): number[] {
    if (n < 1) return [];
    
    const fibs: number[] = [1, 1];
    let nextFib = fibs[fibs.length - 1] + fibs[fibs.length - 2];
    
    while (nextFib <= n) {
        fibs.push(nextFib);
        nextFib = fibs[fibs.length - 1] + fibs[fibs.length - 2];
    }
    
    return fibs;
}

const N = parseInt(process.argv[2], 10);
if (isNaN(N)) {
    console.error("Invalid input. Please provide a single integer argument.");
    process.exit(1);
}

const fibNumbers = fibonacci(N);
fibNumbers.forEach(num => console.log(num));