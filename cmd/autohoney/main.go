// Tool to automatically manage honeypot instances.
//
// Usage:
//
// autohoney <command>
//
// For more information, please see README.md.

package main

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path"
	"text/template"
	"time"
)
