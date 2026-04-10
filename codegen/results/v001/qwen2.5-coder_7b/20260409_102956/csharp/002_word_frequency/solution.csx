using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

var fileContent = File.ReadAllText("input/text.txt");
var words = Regex.Matches(fileContent.ToLower(), @"\b\w+\b")
                 .Cast<Match>()
                 .Select(m => m.Value)
                 .GroupBy(word => word)
                 .OrderByDescending(group => group.Count())
                 .ThenBy(group => group.Key);

foreach (var group in words)
{
    Console.WriteLine($"{group.Key}: {group.Count()}");
}