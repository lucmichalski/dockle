package credential

import (
	"fmt"
	"os"
	"path/filepath"

	deckodertypes "github.com/goodwithtech/deckoder/types"

	"github.com/goodwithtech/dockle/pkg/log"

	"github.com/goodwithtech/dockle/pkg/types"
)

type CredentialAssessor struct{}

func (a CredentialAssessor) Assess(fileMap deckodertypes.FileMap) ([]*types.Assessment, error) {
	log.Logger.Debug("Start scan : credential files")
	assesses := []*types.Assessment{}
	reqFiles := a.RequiredFiles()
	for filename := range fileMap {
		basename := filepath.Base(filename)
		// check exist target files
		for _, reqFilename := range reqFiles {
			if reqFilename == basename {
				assesses = append(
					assesses,
					&types.Assessment{
						Code:     types.AvoidCredential,
						Filename: filename,
						Desc:     fmt.Sprintf("Suspicious filename found : %s ", filename),
					})
				break
			}
		}
	}
	return assesses, nil
}

func (a CredentialAssessor) RequiredFiles() []string {
	return []string{"credentials.json", "credential.json", "credentials", "credential"}
}

func (a CredentialAssessor) RequiredPermissions() []os.FileMode {
	return []os.FileMode{}
}
