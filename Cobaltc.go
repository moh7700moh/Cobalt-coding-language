package main

import (
	"fmt"
	"regexp"
)

// TokenType represents the categorized identity of a piece of code
type TokenType string

const (
	TokenKeyword    TokenType = "KEYWORD"
	TokenIdentifier TokenType = "IDENTIFIER"
	TokenNumber     TokenType = "NUMBER"
	TokenOperator   TokenType = "OPERATOR"
	TokenAssign     TokenType = "ASSIGN"
	TokenLBrace     TokenType = "LBRACE"
	TokenRBrace     TokenType = "RBRACE"
	TokenString     TokenType = "STRING"
)

type Token struct {
	Type  TokenType
	Value string
}

// Cobalt's reserved keywords map for O(1) lightning-fast lookups
var keywords = map[string]bool{
	"unless":    true,
	"let":       true,
	"mut":       true,
	"fn":        true,
	"use":       true,
	"else":      true,
	"printline": true, // Added your brand new printline keyword!
}

func Tokenize(source string) []Token {
	var tokens []Token

	// Updated regex pattern to handle double-quoted strings alongside standard tokens
	pattern := `(?P<STRING>"[^"]*")|(?P<NUMBER>\d+)|(?P<ID>[a-zA-Z_][a-zA-Z0-9_]*)|(?P<ASSIGN>=)|(?P<BRACE>[{}])|(?P<OP>[><!+\-*/])`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(source, -1)
	subnames := re.SubexpNames()

	for _, match := range matches {
		for i, value := range match {
			if i == 0 || value == "" {
				continue
			}
			
			kind := subnames[i]
			tokenType := TokenType(kind)

			// Context-aware step: separate system identifiers from native language keywords
			if kind == "ID" {
				if keywords[value] {
					tokenType = TokenKeyword
				} else {
					tokenType = TokenIdentifier
				}
			} else if kind == "BRACE" {
				if value == "{" {
					tokenType = TokenLBrace
				} else {
					tokenType = TokenRBrace
				}
			} else if kind == "OP" {
				tokenType = TokenOperator
			} else if kind == "ASSIGN" {
				tokenType = TokenAssign
			} else if kind == "STRING" {
				tokenType = TokenString
			}

			tokens = append(tokens, Token{Type: tokenType, Value: value})
			break
		}
	}
	return tokens
}

func main() {
	// Updated test string to utilize printline with a text argument!
	cobaltCode := `unless temperature > 1115 { printline "Safe temperature range" }`
	tokens := Tokenize(cobaltCode)

	fmt.Println("--- Cobalt Go Lexer Output ---")
	for _, tok := range tokens {
		fmt.Printf("[%s: '%s']\n", tok.Type, tok.Value)
	}
}
