// fibonacci.ts

function generateFibonacciUpToN(N: number): void {
    if (N < 1) {
        return;
    }

    let a = 0;
    let b = 1;

    while (b <= N) {
        console.log(b);
        [a, b] = [b, a + b];
    }
}

const N = parseInt(process.argv[2]);
generateFibonacciUpToN(N);