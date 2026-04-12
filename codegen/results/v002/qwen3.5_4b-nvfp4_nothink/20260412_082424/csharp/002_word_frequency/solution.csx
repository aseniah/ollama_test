using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;

var lines = File.ReadAllLines("input/text.txt");
var words = new Dictionary<string, int>();

foreach (var line in lines)
{
    var text = line.ToLowerInvariant();
    // Remove punctuation: replace with non-greedy match of anything not a letter, or split on non-letters
    var cleanText = new string(text.Where(c => char.IsLetterOrDigit(c)).ToArray());
    
    // Tokenize by whitespace
    foreach (var word in cleanText.Split())
    {
        if (word.Length > 0)
        {
            if (!char.IsLetter(word[0])) continue; // Skip numbers/single non-letters
            words[word]++;
        }
    }
}

var sortedWords = words.ToOrderedByDescending(w => w.Value, s => string.Compare(s.Key, w.Key, StringComparison.OrdinalIgnoreCase));

foreach (var item in sortedWords)
{
    Console.WriteLine($"{item.Key}: {item.Value}");
}

class OrderedDictionary<TKey, TValue> : System.Collections.IEqualityComparer, IOrderedEnumerable<string> 
{
    // We'll use LINQ instead of this complexity since .NET std doesn't have ordered dict without packages
    public Dictionary<string, int> D { get; } = new Dictionary<string, int>();
}

// Re-evaluate sorting logic properly without creating custom classes
var wordsList = words.Keys.ToList();
wordsList.Sort((a, b) => string.Compare(a, b, StringComparison.OrdinalIgnoreCase)); // Alphabetically order to group ties

var sortedResult = wordsList.Where(w => w != null).OrderByDescending(w => (int)(w == "" ? -1 : words[w].Value)).ThenBy(w => w);
var finalWords = new List<(string word, int count)>();
foreach (var word in wordsList)
{
    if (!finalWords.Any(x => x.word == word)) 
    {
        // Check count descending by iterating and sorting manually
        var count = words[word];
        finalWords.Add((word, count));
    }
}

// Actually just sort the list of tuples directly
var tupleList = wordsList.Select(w => (w, words[w])).ToList();
tupleList.Sort((a, b) => 
    a.Item2 != b.Item2 ? b.Item2.CompareTo(a.Item2) : string.Compare(a.Item1, b.Item1, StringComparison.OrdinalIgnoreCase));

foreach (var item in tupleList)
{
    Console.WriteLine($"{item.Item1}: {item.Item2}");
}