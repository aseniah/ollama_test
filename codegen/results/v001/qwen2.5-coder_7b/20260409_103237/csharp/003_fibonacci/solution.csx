using System;
using System.Collections.Generic;

if (Args.Count != 1)
{
    Console.WriteLine("Usage: fibonacci.csx <integer>");
    return;
}

if (!int.TryParse(Args[0], out int n) || n < 1)
{
    Console.WriteLine("Please provide a positive integer.");
    return;
}

int a = 1, b = 1;

while (a <= n)
{
    Console.WriteLine(a);
    int temp = a + b;
    a = b;
    b = temp;
}