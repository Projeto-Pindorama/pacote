/*
* This contains a reimplementation of pfmt(), which is a tool for making
* SVID4-style error reporting. In original Sun pkgtools, one would use progerr()
* instead of pfmt(), although the messages looked identical at some point.
* These pkgtools will use this pfmt() implementation instead of log.Fatal or
* something like.
* Temporary reference: http://www.sco.com/developers/devspecs/vol1a.pdf
 */
package pfmt

import (
	"fmt"
	"io"
)

func Pfmt(stream io.Writer, report string, message string) {
	fmt.Fprintf(stream, "%s: %s", reportType(report), message)
}

func reportType(report string) string {
	switch report {
	/* The default messages/report flags, nothing new  */
	case "MM_HALT":
		return "HALT"
	case "MM_ERROR":
		return "ERROR"
	case "MM_WARNING":
		return "WARNING"
	case "MM_INFO":
		return "INFO"
	/* Action: this will print a string saying that something has to be
	* fixed or done */
	case "MM_ACTION":
		return "TO FIX"
	/* Catalog access: I think this would be meant for gettext and
	* i18n. Since we haven't implemented i18n yet, we don't have a specific
	* way to tell "Hey, don't translate this error message, ok? Print it
	* verbatim.", so I think this will be a no-op for now. */
	case "MM_GET":
	case "MM_NOGET":
	/* These are used to set the output format. Basically, if "MM_NOSTD" is
	* set, pfmt() shall work as a generic printf() funcion. I had an idea of
	* how to implement this, but I fear breaking the rest, so I will keep as
	* a no-op for now too. */
	case "MM_STD":
	case "MM_NOSTD":
	}
}
