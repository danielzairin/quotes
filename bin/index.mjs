#!/usr/bin/env node
import fs from "node:fs";
import os from "node:os";
/** @typedef {import("./types.ts").Quote} Quote */
const displayProbability = Number(process.argv[2] || "0.05"); // 5% chance by default
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
if (displayProbability > Math.random()) {
    console.log(getRandomQuote());
}
