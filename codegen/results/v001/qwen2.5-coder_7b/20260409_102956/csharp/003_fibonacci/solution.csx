using System;
using System.Collections.Generic;

var Args = new List<string> { "10" }; // Replace with actual command-line arguments if running in a script

if (Args.Count == 0)
{
    Console.WriteLine("No arguments provided.");
    return;
}

if (!int.TryParse(Args[0], out int N) || N < 1)
{
    Console.WriteLine("Please provide a positive integer.");
    return;
}

long a = 1, b = 1;

while (a <= N)
{
    Console.WriteLine(a);
    long temp = a + b;
    a = b;
    b = temp;
}