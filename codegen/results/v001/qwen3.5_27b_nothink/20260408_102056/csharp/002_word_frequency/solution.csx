#r "System.Text.Json"
#r "System.Collections"
#r "System.Linq"

using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text;

var filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    Console.Error.WriteLine($"Error: File '{filePath}' not found.");
    return;
}

var text = File.ReadAllText(filePath);
var frequencyMap = new Dictionary<string, int>();

// Split by non-letter characters, filter out empty strings, convert to lowercase
var words = text
    .Where(char.IsLetter)
    .ToList()
    .Aggregate(new List<string>(), (list, ch) => 
    {
        if (list.Count > 0)
        {
            if (char.IsLetter(list[list.Count - 1] as string? is null ? (char?)null : (char?)list[list.Count-1]!.ToString()?[0]))
            {
                 // This aggregation approach is complex for simple splitting.
                 // Let's switch to a cleaner Regex-free approach or simple iteration.
                 return list;
            }
        }
        return list;
    });

// Cleaner approach without complex aggregation:
var wordsList = new List<string>();
var currentWord = new StringBuilder();

foreach (char c in text)
{
    if (char.IsLetter(c))
    {
        currentWord.Append(char.ToLower(c));
    }
    else
    {
        if (currentWord.Length > 0)
        {
            wordsList.Add(currentWord.ToString());
            currentWord.Clear();
        }
    }
}
if (currentWord.Length > 0)
{
    wordsList.Add(currentWord.ToString());
}

foreach (var word in wordsList)
{
    if (frequencyMap.ContainsKey(word))
    {
        frequencyMap[word]++;
    }
    else
    {
        frequencyMap[word] = 1;
    }
}

var sortedWords = frequencyMap
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

foreach (var kv in sortedWords)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}