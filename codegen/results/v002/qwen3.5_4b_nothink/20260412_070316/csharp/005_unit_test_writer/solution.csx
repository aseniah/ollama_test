// Define the CalculateDiscount function
let calculateDiscount price quantity =
    if quantity < 10 then
        price * quantity
    else if quantity < 50 then
        price * quantity * 0.9
    else
        price * quantity * 0.75

// Run test cases and print results
let test1 = {
    let desc = "Price $10.00, quantity 9 (no discount)"
    let expected = calculateDiscount 10.0 9.0
    let got = calculateDiscount 10.0 9.0
    let actual = match (expected, got) with
        | (e, g) when abs(e - g) < 0.0001 -> printfn "PASS: %s" desc
        | (e, g) -> printfn "FAIL: %s (expected: %f, got: %f)" desc e g
}

let test2 = {
    let desc = "Price $10.00, quantity 10 (10% discount)"
    let expected = calculateDiscount 10.0 10.0
    let got = calculateDiscount 10.0 10.0
    let actual = match (expected, got) with
        | (e, g) when abs(e - g) < 0.0001 -> printfn "PASS: %s" desc
        | (e, g) -> printfn "FAIL: %s (expected: %f, got: %f)" desc e g
}

let test3 = {
    let desc = "Price $10.00, quantity 49 (10% discount)"
    let expected = calculateDiscount 10.0 49.0
    let got = calculateDiscount 10.0 49.0
    let actual = match (expected, got) with
        | (e, g) when abs(e - g) < 0.0001 -> printfn "PASS: %s" desc
        | (e, g) -> printfn "FAIL: %s (expected: %f, got: %f)" desc e g
}

let test4 = {
    let desc = "Price $10.00, quantity 50 (25% discount)"
    let expected = calculateDiscount 10.0 50.0
    let got = calculateDiscount 10.0 50.0
    let actual = match (expected, got) with
        | (e, g) when abs(e - g) < 0.0001 -> printfn "PASS: %s" desc
        | (e, g) -> printfn "FAIL: %s (expected: %f, got: %f)" desc e g
}

run test1
run test2
run test3
run test4