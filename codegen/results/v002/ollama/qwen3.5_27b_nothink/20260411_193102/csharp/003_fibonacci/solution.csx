if (Args.Count != 1 || !int.TryParse(Args[0], out int n))
{
    return; // Or throw an exception depending on desired behavior for invalid input
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