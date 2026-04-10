using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;

var text = File.ReadAllText("input/text.txt");
var words = text.ToLowerInvariant()
    .SelectMany(char.GetUnicodeCategory, (c, _) =>
    {
        var cat = char.GetUnicodeCategory(c);
        return c >= 'a' && c <= 'z' && c >= '0' && c <= '9' ? c.ToString() : string.Empty;
    })
    .Where(w => !string.IsNullOrEmpty(w))
    .OrderByDescending(w => w) // Sort by letter content descending? No, we need to count then sort.
    .ToList();

var counts = words.ToDictionary(
    w => w,
    _ => new List<string> { w }.GroupBy(x => x).Select(g => g.Count()).ToList()[0]
);

// Wait, the above dictionary logic is flawed. Let's redo it properly.
// 1. Lowercase
// 2. Filter: only a-z
// 3. Group and Count
// 4. Sort by count desc, then word asc

var lowerText = text.ToLowerInvariant();
var wordsList = lowerText
    .Aggregate(new List<char>(), (list, c) =>
    {
        var cat = char.GetUnicodeCategory(c);
        if (cat == charCategory.LowerLetter || cat == charCategory.UpperLetter)
        {
            list.Add(c);
        }
        return list;
    })
    .ToList();

var frequencyMap = new Dictionary<string, int>();
foreach (var word in wordsList)
{
    frequencyMap[word] = frequencyMap.GetValueOrDefault(word, 0) + 1;
}

var sortedWords = frequencyMap
    .OrderByDescending(pair => pair.Value)
    .ThenBy(pair => pair.Key)
    .ToList();

Console.WriteLine(string.Join("\n", sortedWords.Select(pair => $"{pair.Key}: {pair.Value}"));