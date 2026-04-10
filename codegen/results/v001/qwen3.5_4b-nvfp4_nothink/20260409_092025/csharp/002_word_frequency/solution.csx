import System.IO;
import System.Collections.Generic;
import System.Text.Json;
import System.Text.Json.Nodes;

// Read file
String content = File.ReadAllText("input/text.txt");

// Parse words: split by non-letter characters, filter empty strings, convert to lowercase
string[] rawWords = Regex.Split(content, "[^a-zA-Z]");
Dictionary<string, int> wordCounts = new Dictionary<string, int>();

foreach (var word in rawWords)
{
    if (!string.IsNullOrWhiteSpace(word))
    {
        string cleanWord = System.Text.RegularExpressions.Regex.Match(word, @"^[a-z]+$").Value;
        if (!string.IsNullOrEmpty(cleanWord))
        {
            if (wordCounts.ContainsKey(cleanWord))
            {
                wordCounts[cleanWord]++;
            }
            else
            {
                wordCounts.Add(cleanWord, 1);
            }
        }
    }
}

// Sort: primary by count descending, secondary by word ascending
var sortedWords = wordCounts.Keys.ToList();
sortedWords.Sort((w1, w2) =>
{
    if (wordCounts[w1] != wordCounts[w2])
        return wordCounts[w2].CompareTo(wordCounts[w1]); // Descending count
    return string.Compare(w1, w2, System.StringComparison.OrdinalIgnoreCase); // Ascending alphabetically
});

// Output results
foreach (var word in sortedWords)
{
    Console.WriteLine($"{word}: {wordCounts[word]}");
}