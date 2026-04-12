function fizzbuzz(n: number): string[] {
    const result: string[] = [];
    for (let i = 1; i <= n; i++) {
        if (i % 15 === 0) result.push("FizzBuzz");
        else if (i % 3 === 0) result.push("Fizz");
        else if (i % 5 === 0) result.push("Buzz");
        else result.push(String(i));
    }
    return result;
}

const n = Number(process.argv[2] || "15");
fizzbuzz(n).forEach(v => console.log(v));