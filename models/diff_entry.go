package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
DiffEntry diff entry

swagger:model DiffEntry
*/
type DiffEntry struct {

	/* Blame blame
	 */
	Blame []BlameLine `json:"blame,omitempty"`

	/* ByteData byte data
	 */
	ByteData []strfmt.Base64 `json:"byteData,omitempty"`

	/* ChangeType change type
	 */
	ChangeType string `json:"changeType,omitempty"`

	/* Charset charset
	 */
	Charset string `json:"charset,omitempty"`

	/* CommitID commit id
	 */
	CommitID string `json:"commitId,omitempty"`

	/* Data data
	 */
	Data string `json:"data,omitempty"`

	/* Deletions deletions
	 */
	Deletions int32 `json:"deletions,omitempty"`

	/* File file
	 */
	File bool `json:"file,omitempty"`

	/* FlattenPath flatten path
	 */
	FlattenPath string `json:"flatten_path,omitempty"`

	/* Image image
	 */
	Image bool `json:"image,omitempty"`

	/* Insertions insertions
	 */
	Insertions int32 `json:"insertions,omitempty"`

	/* Language language
	 */
	Language string `json:"language,omitempty"`

	/* LastCommit last commit
	 */
	LastCommit Commit `json:"lastCommit,omitempty"`

	/* Mime mime
	 */
	Mime string `json:"mime,omitempty"`

	/* Mode mode
	 */
	Mode int32 `json:"mode,omitempty"`

	/* ModeStr mode str
	 */
	ModeStr string `json:"modeStr,omitempty"`

	/* Name name
	 */
	Name string `json:"name,omitempty"`

	/* NormalizedDiffStat normalized diff stat
	 */
	NormalizedDiffStat NormalizedDiffStat `json:"normalizedDiffStat,omitempty"`

	/* ObjectID object id
	 */
	ObjectID string `json:"objectId,omitempty"`

	/* Path path
	 */
	Path string `json:"path,omitempty"`

	/* Size size
	 */
	Size int64 `json:"size,omitempty"`

	/* Submodule submodule
	 */
	Submodule bool `json:"submodule,omitempty"`

	/* SubmoduleCommitID submodule commit id
	 */
	SubmoduleCommitID string `json:"submoduleCommitId,omitempty"`

	/* SubmodulePath submodule path
	 */
	SubmodulePath string `json:"submodulePath,omitempty"`

	/* SubmoduleURl submodule u rl
	 */
	SubmoduleURl string `json:"submoduleURl,omitempty"`

	/* Symlink symlink
	 */
	Symlink bool `json:"symlink,omitempty"`

	/* Text text
	 */
	Text bool `json:"text,omitempty"`

	/* Tree tree
	 */
	Tree bool `json:"tree,omitempty"`
}

// Validate validates this diff entry
func (m *DiffEntry) Validate(formats strfmt.Registry) error {

	var res []error

	if err := m.validateChangeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}

var diffEntryChangeTypeEnum []interface{}

func (m *DiffEntry) validateChangeTypeEnum(path, location string, value string) error {
	if diffEntryChangeTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["ADD","MODIFY","DELETE","RENAME","COPY"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			diffEntryChangeTypeEnum = append(diffEntryChangeTypeEnum, v)
		}
	}
	return validate.Enum(path, location, value, diffEntryChangeTypeEnum)
}

func (m *DiffEntry) validateChangeType(formats strfmt.Registry) error {

	if err := m.validateChangeTypeEnum("changeType", "body", m.ChangeType); err != nil {
		return err
	}

	return nil
}
