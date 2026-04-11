function fizzbuzz(n: number): string[] {
    const result: string[] = [];
    for (let i = 0; i < n; i++) {  // Bug: should start at 1, use i <= n
        if (i % 15 === 0) result.push("FizzBuzz");
        else if (i % 3 === 0) result.push("Fizz");
        else if (i % 5 === 0) result.push("Buzz");
        else result.push(String(i));
    }
    return result;
}
fizzbuzz(15).forEach(v => console.log(v));
