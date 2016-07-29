package drudsub

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/drud/drud-go/secrets"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/pubsub"
)

var project = os.Getenv("DRUDSUB_PROJECT")
var jwtPath = os.Getenv("DRUDSUB_JWT")
var gitToken = os.Getenv("GITHUB_TOKEN")
var vaultHost = os.Getenv("VAULT_ADDR")

// Connection to pubsub/nats
type Connection struct {
	// The pubsub connection client.
	Client *pubsub.Client

	// The pubsub context.
	Context context.Context
}

// JWT content struct.
type JWT struct {
	ProjectID string `json:"project_id"`
}

// GetJWTByes returns jwt from file or vault
func GetJWTByes() (jbytes []byte, err error) {
	if jwtPath != "" {
		// read contents of jwt file
		jbytes, err = ioutil.ReadFile(jwtPath)
	} else if gitToken != "" {
		// get jwt from vault
		jbytes, err = secrets.GetJWT(gitToken, vaultHost, project)
	}
	return
}

// Connect to drudsub backing service.
func (c *Connection) Connect() error {
	// read contents of jwt file or use vault
	jbytes, err := GetJWTByes()
	if err != nil {
		return err
	}

	// get the project id from the jwt
	jwtStruct := JWT{}
	err = json.Unmarshal(jbytes, &jwtStruct)
	if err != nil {
		return err
	}

	// instantiate google conf using jwt contents
	conf, err := google.JWTConfigFromJSON(
		jbytes,
		pubsub.ScopeCloudPlatform,
		pubsub.ScopePubSub,
	)
	if err != nil {
		return err
	}

	// create a google cloud context
	c.Context = cloud.NewContext(jwtStruct.ProjectID, conf.Client(oauth2.NoContext))
	// instantiate a client for workign with pub sub
	c.Client, err = pubsub.NewClient(c.Context, jwtStruct.ProjectID, cloud.WithTokenSource(conf.TokenSource(c.Context)))

	return err
}
