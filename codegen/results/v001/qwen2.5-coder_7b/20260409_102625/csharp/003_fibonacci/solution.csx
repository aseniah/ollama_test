using System;

if (Args.Count != 1)
{
    Console.WriteLine("Please provide a single integer argument.");
    return;
}

if (!int.TryParse(Args[0], out int N) || N < 1)
{
    Console.WriteLine();
    return;
}

int a = 1, b = 1;
while (a <= N)
{
    Console.WriteLine(a);
    int temp = a;
    a = b;
    b = temp + b;
}