function isPrime(num: number): boolean {
    if (num <= 1) return false;
    if (num <= 3) return true;

    if (num % 2 === 0 || num % 3 === 0) return false;

    for (let i = 5; i * i <= num; i += 6) {
        if (num % i === 0 || num % (i + 2) === 0) return false;
    }

    return true;
}

const N = parseInt(process.argv[2], 10);

if (!isNaN(N)) {
    for (let number = 2; number <= N; number++) {
        if (isPrime(number)) {
            console.log(number);
        }
    }
}