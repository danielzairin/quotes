#!/usr/bin/env node
import fs from "node:fs";
import os from "node:os";

/** @typedef {import("./types.ts").Quote} Quote */

const PATH_TO_QUOTES = os.homedir() + "/quotes.json";

const rawQuotes = fs.readFileSync(PATH_TO_QUOTES);

/** @type {Array<Quote>} */
const quotes = JSON.parse(rawQuotes.toString());

/**
 * @returns {Quote} A random quote
 */
export function getRandomQuote() {
  return quotes[Math.floor(Math.random() * quotes.length)];
}

console.log(getRandomQuote());
