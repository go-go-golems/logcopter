// Package doc contains embedded Markdown help entries for downstream tools.
//
// Applications that already depend on Glazed can load the documents through
// AddDocToHelpSystem. Non-Glazed consumers can keep mounting FS directly as a
// regular fs.FS.
package doc

import (
	"embed"

	"github.com/go-go-golems/glazed/pkg/help"
)

// FS contains the embedded logcopter help documents.
//
// Glazed and other downstream tools can mount this filesystem directly instead
// of locating a checkout of the logcopter repository at runtime.
//
//go:embed topics/*.md tutorials/*.md
var FS embed.FS

// AddDocToHelpSystem loads the embedded logcopter help documents into a Glazed
// help system.
func AddDocToHelpSystem(helpSystem *help.HelpSystem) error {
	return helpSystem.LoadSectionsFromFS(FS, ".")
}
