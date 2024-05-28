import "source-map-support/register";

import express from "express";
const app = express();
const port = 9090;

// Function to check if a number is prime
const isPrime = (n: number): boolean => {
  if (n <= 1) return false;
  if (n === 2) return true;
  if (n % 2 === 0) return false;
  for (let i = 3; i <= Math.sqrt(n); i += 2) {
    if (n % i === 0) return false;
  }
  return true;
};

// Function to get primes in a range
const getPrimes = (start: number, end: number): number[] => {
  const primes: number[] = [];
  for (let i = start; i <= end; i++) {
    if (isPrime(i)) primes.push(i);
  }
  return primes;
};

app.get("/test_node_20", (req, res) => {
  const start = parseInt(req.query.start as string, 10);
  const end = parseInt(req.query.end as string, 10);

  if (isNaN(start) || isNaN(end)) {
    res.status(400).send("Invalid start or end parameter");
    return;
  }

  const primes = getPrimes(start, end);
  res.json({ primes });
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
