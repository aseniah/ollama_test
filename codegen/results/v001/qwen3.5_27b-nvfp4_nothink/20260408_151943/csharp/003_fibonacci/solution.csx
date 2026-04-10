#r "System.Runtime"
using System;
using System.Collections.Generic;

// Parse the input argument N
if (Args.Length != 1)
{
    Console.WriteLine("Usage: dotnet-script script.csx <integer>");
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    Console.WriteLine("Error: Argument must be a valid integer.");
    return;
}

// If N is less than 1, print nothing as per requirements
if (n < 1)
{
    return;
}

// Generate Fibonacci sequence up to and including the largest value <= n
int a = 1;
int b = 1;

while (a <= n)
{
    Console.WriteLine(a);
    
    // Calculate next number
    int next = a + b;
    a = b;
    b = next;
}