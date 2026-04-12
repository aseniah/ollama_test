#r "System.Text.Json"

using System;
using System.Text.Json;

if (Args.Count != 1)
{
    Console.Error.WriteLine("Usage: Please provide exactly one integer argument N.");
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    Console.Error.WriteLine("Error: The argument must be a valid integer.");
    return;
}

if (n < 1)
{
    return;
}

int a = 1;
int b = 1;

while (a <= n)
{
    Console.WriteLine(a);
    int next = a + b;
    a = b;
    b = next;
}