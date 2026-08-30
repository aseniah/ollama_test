using System;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

var content = File.ReadAllText("input/text.txt");
var words = Regex.Matches(content, @"[a-zA-Z]+")
                  .Select(m => m.Value.ToLower())
                  .ToList();

var wordCounts = words
    .GroupBy(w => w)
    .ToDictionary(g => g.Key, g => g.Count());

var sortedWords = wordCounts.Select(kv => (word: kv.Key, count: kv.Value))
                            .OrderByDescending(x => x.count)
                            .ThenBy(x => x.word);

foreach (var item in sortedWords) {
    System.Console.WriteLine($"{item.word}: {item.count}");
}