using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

var text = File.ReadAllText("input/text.txt");
var words = Regex.Split(text.ToLower(), @"\W+").Where(word => !string.IsNullOrEmpty(word));

var wordCount = new Dictionary<string, int>();
foreach (var word in words)
{
    if (wordCount.ContainsKey(word))
    {
        wordCount[word]++;
    }
    else
    {
        wordCount[word] = 1;
    }
}

var sortedWordCount = wordCount.OrderByDescending(pair => pair.Value).ThenBy(pair => pair.Key);

foreach (var item in sortedWordCount)
{
    Console.WriteLine($"{item.Key}: {item.Value}");
}