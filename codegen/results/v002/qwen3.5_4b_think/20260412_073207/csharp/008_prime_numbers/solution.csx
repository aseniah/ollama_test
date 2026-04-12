using System;

int n = int.Parse(Args[0]);

if (n < 2)
{
    return;
}

bool[] isComposite = new bool[n + 1];

int i = 2;
while (i * i <= n)
{
    if (!isComposite[i])
    {
        int j = i * i;
        while (j <= n)
        {
            isComposite[j] = true;
            j += i;
        }
    }
    i++;
}

for (int k = 2; k <= n; k++)
{
    if (!isComposite[k])
    {
        Console.WriteLine(k);
    }
}