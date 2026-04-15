using System;

if (Args.Count < 1)
{
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    return;
}

if (n < 1)
{
    return;
}

int a = 1;
int b = 1;

Console.WriteLine(a);
Console.WriteLine(b);

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