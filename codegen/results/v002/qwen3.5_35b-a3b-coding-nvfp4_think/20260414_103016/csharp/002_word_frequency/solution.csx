using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;

var lines = File.ReadAllLines("input/text.txt");

var wordCounts = new Dictionary<string, int>();

foreach (var line in lines)
{
    var words = line.Split(new[] { ' ', ',', '.', ';', ':', '!', '?', '\'', '"', '(', ')', '[', ']', '{', '}' }, StringSplitOptions.RemoveEmptyEntries);
    
    foreach (var word in words)
    {
        var lowerWord = word.ToLower();
        var cleanedWord = new string(lowerWord.Where(char.IsLetter).ToArray());
        
        if (!string.IsNullOrEmpty(cleanedWord))
        {
            if (wordCounts.ContainsKey(cleanedWord))
            {
                wordCounts[cleanedWord]++;
            }
            else
            {
                wordCounts[cleanedWord] = 1;
            }
        }
    }
}

var sortedWords = wordCounts
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key)
    .ToList();

foreach (var word in sortedWords)
{
    Console.WriteLine($"{word.Key}: {word.Value}");
}