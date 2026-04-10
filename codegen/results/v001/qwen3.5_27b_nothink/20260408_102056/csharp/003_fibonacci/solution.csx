#r "System.Text.Json"

using System;
using System.Text.Json;
using System.Text.Json.Nodes;

if (Args.Count != 1)
{
    Console.Error.WriteLine("Error: Exactly one integer argument is required.");
    Environment.Exit(1);
}

if (!int.TryParse(Args[0], out int n))
{
    Console.Error.WriteLine("Error: Argument must be a valid integer.");
    Environment.Exit(1);
}

if (n < 1)
{
    return;
}

long a = 1;
long b = 1;

while (a <= n)
{
    Console.WriteLine(a);
    long next = a + b;
    a = b;
    b = next;
}