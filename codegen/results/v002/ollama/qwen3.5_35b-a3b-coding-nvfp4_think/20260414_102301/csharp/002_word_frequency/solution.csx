var text = File.ReadAllText("input/text.txt");
var words = text.ToLower()
    .SelectMany(c => char.IsLetter(c) ? new[] { c } : new[] { ' ' })
    .Where(c => c != ' ' || char.IsLetter(c))
    .GroupAdjacent(c => c)
    .Where(g => g.First() != ' ')
    .Select(g => new string(g.ToArray()))
    .Where(w => w.Length > 0)
    .GroupBy(w => w)
    .Select(g => new { word = g.Key, count = g.Count() })
    .OrderByDescending(x => x.count)
    .ThenBy(x => x.word)
    .Select(x => $"{x.word}: {x.count}");

foreach (var line in words)
{
    Console.WriteLine(line);
}