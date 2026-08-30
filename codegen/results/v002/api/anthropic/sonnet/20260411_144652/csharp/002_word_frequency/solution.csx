var text = File.ReadAllText("input/text.txt");
var words = text.Split(new char[]{' ','\t','\n','\r'}, StringSplitOptions.RemoveEmptyEntries);

var freq = new Dictionary<string, int>();
foreach (var w in words)
{
    var clean = new string(w.ToLower().Where(char.IsLetter).ToArray());
    if (clean.Length == 0) continue;
    freq[clean] = freq.GetValueOrDefault(clean, 0) + 1;
}

foreach (var kvp in freq.OrderByDescending(x => x.Value).ThenBy(x => x.Key))
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}
