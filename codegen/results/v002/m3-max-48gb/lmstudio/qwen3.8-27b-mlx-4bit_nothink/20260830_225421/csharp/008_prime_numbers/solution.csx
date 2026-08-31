using System;
using System.Numerics;

if (Args.Count != 1)
{
    Console.Error.WriteLine("Usage: program <N>");
    return;
}

if (!int.TryParse(Args[0], out int N))
{
    Console.Error.WriteLine("Invalid integer argument.");
    return;
}

if (N < 2)
{
    return;
}

bool[] isComposite = new bool[N + 1];
isComposite[0] = true;
isComposite[1] = true;

int sqrtN = (int)Math.Sqrt(N);
for (int i = 2; i <= sqrtN; i++)
{
    if (!isComposite[i])
    {
        for (int j = i * i; j <= N; j += i)
        {
            isComposite[j] = true;
        }
    }
}

for (int i = 2; i <= N; i++)
{
    if (!isComposite[i])
    {
        Console.WriteLine(i);
    }
}