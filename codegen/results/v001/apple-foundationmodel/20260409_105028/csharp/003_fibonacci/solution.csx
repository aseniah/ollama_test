using System;
using System.Text.Json;
using System.Text.Json.Nodes;

// Function to compute the nth Fibonacci number
static long Fibonacci(int n)
{
    if (n <= 0) return 0;
    if (n == 1 || n == 2) return 1;

    long a = 0, b = 1, c = 0;
    for (int i = 3; i <= n; i++)
    {
        c = a + b;
        a = b;
        b = c;
    }
    return c;
}

// Function to print Fibonacci numbers up to N
static void PrintFibonacciUpTo(int N)
{
    if (N < 1) return;

    long fib = 1;
    while (fib <= N)
    {
        Console.WriteLine(fib);
        fib = Fibonacci(fib + 1);
    }
}

// Main entry point
if (Args.Length == 1)
{
    int N = Convert.ToInt32(Args[0]);
    PrintFibonacciUpTo(N);
}