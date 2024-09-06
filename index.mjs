import fs from "node:fs";

/** @typedef {import("./types.ts").Quote} Quote */

const PATH_TO_QUOTES = "quotes.json";

const rawQuotes = fs.readFileSync(PATH_TO_QUOTES);

/** @type {Array<Quote>} */
const quotes = JSON.parse(rawQuotes.toString());

/**
 * @returns {Quote} A random quote
 */
function getRandomQuote() {
  return quotes[Math.floor(Math.random() * quotes.length)];
}

console.log(getRandomQuote());
