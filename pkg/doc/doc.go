// Package doc contains embedded Markdown help entries for downstream tools.
//
// The package deliberately does not import Glazed. Applications that already
// depend on Glazed can mount FS with their own help system, while non-Glazed
// consumers can read the embedded files as a regular fs.FS.
package doc

import "embed"

// FS contains the embedded logcopter help documents.
//
// Glazed and other downstream tools can mount this filesystem directly instead
// of locating a checkout of the logcopter repository at runtime.
//
//go:embed topics/*.md tutorials/*.md
var FS embed.FS
