using System;

if (Args.Count != 1 || !int.TryParse(Args[0], out int n) || n < 1)
{
    Console.WriteLine("Please provide a single positive integer as an argument.");
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