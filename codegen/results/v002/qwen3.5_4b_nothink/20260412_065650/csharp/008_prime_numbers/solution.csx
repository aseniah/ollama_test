let args = System.Runtime.InteropServices.RuntimeEnvironment.GetCommandLineArgs()
let n = int TryParse args.[1]
if n >= 2 then
  for i in [2..n] do
    let isPrime =
      [1..i - 1]
      .fold (true, fun (acc, d) -> acc && (i % d) <> 0)
    if isPrime then
      print i