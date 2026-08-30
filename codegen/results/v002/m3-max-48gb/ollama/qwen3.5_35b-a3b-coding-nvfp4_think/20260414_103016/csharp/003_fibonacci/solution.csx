using System;

// Get the argument
if (Args.Count < 1)
{
    return;
}

// Parse the integer
if (!int.TryParse(Args[0], out int n))
{
    return;
}

// If N < 1, print nothing
if (n < 1)
{
    return;
}

// Generate Fibonacci numbers
int a = 1;
int b = 1;

// Print first 1
Console.WriteLine(a);

// Print second 1
Console.WriteLine(b);

// Continue generating Fibonacci numbers
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