const input = process.argv[2];

// Check if the argument was provided
if (input === undefined) {
    process.exit(1);
}

const N = parseInt(input, 10);

// Check if parsing resulted in a valid number
if (isNaN(N)) {
    process.exit(1);
}

// If N is less than 2, there are no prime numbers to print
if (N < 2) {
    process.exit(0);
}

function isPrime(num: number): boolean {
    if (num <= 1) return false;
    if (num <= 3) return true;
    if (num % 2 === 0 || num % 3 === 0) return false;
    
    for (let i = 5; i * i <= num; i += 6) {
        if (num % i === 0 || num % (i + 2) === 0) return false;
    }
    return true;
}

for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
        console.log(i);
    }
}