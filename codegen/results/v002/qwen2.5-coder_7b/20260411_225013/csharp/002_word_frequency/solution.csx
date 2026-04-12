using System;
using System.Collections.Generic;
using System.Linq;

var lines = File.ReadAllLines("input/text.txt");
var words = string.Join(" ", lines)
                .ToLower()
                .Where(char.IsLetterOrDigit or char.IsWhiteSpace)
                .Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

var wordCount = new Dictionary<string, int>();
foreach (var word in words)
{
    if (wordCount.ContainsKey(word))
        wordCount[word]++;
    else
        wordCount[word] = 1;
}

var sortedWords = wordCount.OrderByDescending(kvp => kvp.Value).ThenBy(kvp => kvp.Key);

foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}