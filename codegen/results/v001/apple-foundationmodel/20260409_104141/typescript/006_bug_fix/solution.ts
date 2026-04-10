function fizzbuzz(n: number): string[] {
    const result: string[] = [];
    for (let i = 1; i <= n; i++) {  // Corrected: loop should start at 1 and go up to n
        if (i % 15 === 0) result.push("FizzBuzz");
        else if (i % 3 === 0) result.push("Fizz");
        else if (i % 5 === 0) result.push("Buzz");
        else result.push(String(i));
    }
    return result;
}

console.log(fizzbuzz(15).join("\n"));