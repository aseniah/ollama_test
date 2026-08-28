using System;

// Get the argument
if (Args.Count < 1)
{
    return;
}

int n = int.Parse(Args[0]);

// If N < 1, print nothing
if (n < 1)
{
    return;
}

int a = 1;
int b = 1;

// Print first Fibonacci number (1)
Console.WriteLine(a);

// Print second Fibonacci number (1) if N >= 1
Console.WriteLine(b);

// Generate subsequent Fibonacci numbers
while (true)
{
    int next = a + b;
    if (next > n)
    {
        break;
    }
    Console.WriteLine(next);
    a = b;
    b = next;
}