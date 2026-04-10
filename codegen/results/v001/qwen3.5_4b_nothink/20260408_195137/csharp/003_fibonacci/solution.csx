if Args.Count != 1 or int.TryParse(Args[0], out int n) or n < 1 {
    return;
}

let sequence = [1L, 1L]
let fib = sequence

for (a, b) <- zip(fib, fib) do
    if a > n then break
    yield b

for num <- Fib do
    print $"" + num + "\n"
return