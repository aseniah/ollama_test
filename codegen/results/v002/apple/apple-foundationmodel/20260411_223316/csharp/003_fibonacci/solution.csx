using System;

// Function to calculate Fibonacci numbers up to a given limit
Func<int, int> FibonacciLimit(int limit) =>
    var fibs = new List<int> { 1, 1 };

    while (fibs[fibs.Count - 1] <= limit) {
        fibs.Add(fibs[fibs.Count - 1] + fibs[fibs.Count - 2]);
    }

    return fibs[fibs.Count - 1];

// Main entry point
int N = GetIntArg();

if (N < 1)
    return;

var largestFibonacci = FibonacciLimit(N);

// Print Fibonacci numbers up to the largest one <= N
foreach (var fib in FibonacciLimit(N).TakeWhile(f => f <= N))
    Console.WriteLine(fib);