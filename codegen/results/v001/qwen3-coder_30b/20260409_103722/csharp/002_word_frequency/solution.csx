using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

var words = new Dictionary<string, int>();

var text = File.ReadAllText("input/text.txt");
var matches = Regex.Matches(text, @"[a-zA-Z]+");

foreach (Match match in matches)
{
    var word = match.Value.ToLower();
    if (words.ContainsKey(word))
        words[word]++;
    else
        words[word] = 1;
}

foreach (var kvp in words.OrderByDescending(x => x.Value).ThenBy(x => x.Key))
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}