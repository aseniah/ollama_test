#r "System.Text.Json"
using System;
using System.Text.Json;
using System.Text.Json.Nodes;

// Validate argument count
if (Args.Count != 1)
{
    Console.Error.WriteLine("Usage: dotnet-script fib.csx <N>");
    Environment.Exit(1);
}

// Parse the integer argument
if (!int.TryParse(Args[0], out int n))
{
    Console.Error.WriteLine("Error: Argument must be a valid integer.");
    Environment.Exit(1);
}

// If N < 1, print nothing
if (n < 1)
{
    return;
}

// Generate Fibonacci numbers
long a = 1;
long b = 1;

while (a <= n)
{
    Console.WriteLine(a);
    long next = a + b;
    a = b;
    b = next;
}