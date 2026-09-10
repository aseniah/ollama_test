using System.Text;
using System.Linq;

var text = File.ReadAllText("input/text.txt");

// Split into words by whitespace
var words = text.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

// For each token, keep only letters, convert to lowercase, skip empty
var cleaned = words
    .Select(w => new string(w.Where(c => char.IsLetter(c)).ToArray()).ToLower())
    .Where(w => w.Length > 0);

// Count frequencies
var counts = cleaned
    .GroupBy(w => w)
    .Select(g => (Word: g.Key, Count: g.Count()))
    .OrderByDescending(x => x.Count)
    .ThenBy(x => x.Word, StringComparer.Ordinal)
    .ToList();

// Output
foreach (var (word, count) in counts)
{
    Console.WriteLine($"{word}: {count}");
}