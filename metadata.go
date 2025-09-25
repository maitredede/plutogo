package plutogo

/*
#cgo pkg-config: plutobook

#include <plutobook.h>
*/
import "C"

// PdfMetadata Defines different metadata fields for a PDF document
type PdfMetadata int

const (
	PdfMetadataTitle            PdfMetadata = C.PLUTOBOOK_PDF_METADATA_TITLE
	PdfMetadataAuthor           PdfMetadata = C.PLUTOBOOK_PDF_METADATA_AUTHOR
	PdfMetadataSubject          PdfMetadata = C.PLUTOBOOK_PDF_METADATA_SUBJECT
	PdfMetadataKeywords         PdfMetadata = C.PLUTOBOOK_PDF_METADATA_KEYWORDS
	PdfMetadataCreator          PdfMetadata = C.PLUTOBOOK_PDF_METADATA_CREATOR
	PdfMetadataCreationDate     PdfMetadata = C.PLUTOBOOK_PDF_METADATA_CREATION_DATE
	PdfMetadataModificationDate PdfMetadata = C.PLUTOBOOK_PDF_METADATA_MODIFICATION_DATE
)
