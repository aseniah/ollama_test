using System;
using System.Collections.Generic;
using System.Linq;

var lines = File.ReadAllLines("input/text.txt");

var words = new Dictionary<string, int>();

foreach (var line in lines)
{
    var text = string.Concat(line); // Ensure no extra newlines interfere
    foreach (var word in text.Split())
    {
        var cleanWord = word.Trim();
        if (string.IsNullOrEmpty(cleanWord)) continue;

        cleanWord = cleanWord.ToLowerInvariant();
        cleanWord = Regex.Replace(cleanWord, @"[^a-z]", string.Empty);

        if (!string.IsNullOrEmpty(cleanWord))
        {
            if (words.ContainsKey(cleanWord))
            {
                words[cleanWord]++;
            }
            else
            {
                words[cleanWord] = 1;
            }
        }
    }
}

var sortedWords = words
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key)
    .ToList();

foreach (var item in sortedWords)
{
    Console.WriteLine($"{item.Key}: {item.Value}");
}