var N = int.Parse(Args[0]);

if (N < 1)
{
    return;
}

var a = 1;
var b = 1;

while (a <= N)
{
    Console.WriteLine(a);
    var next = a + b;
    a = b;
    b = next;
}