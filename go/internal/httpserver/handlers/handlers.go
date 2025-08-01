package handlers

import (
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	autogen_client "github.com/kagent-dev/kagent/go/internal/autogen/client"
	"github.com/kagent-dev/kagent/go/internal/database"
)

// Handlers holds all the HTTP handler components
type Handlers struct {
	Health      *HealthHandler
	ModelConfig *ModelConfigHandler
	Model       *ModelHandler
	Provider    *ProviderHandler
	Sessions    *SessionsHandler
	Agents      *AgentsHandler
	Tools       *ToolsHandler
	ToolServers *ToolServersHandler
	Memory      *MemoryHandler
	Feedback    *FeedbackHandler
	Namespaces  *NamespacesHandler
}

// Base holds common dependencies for all handlers
type Base struct {
	KubeClient         client.Client
	AutogenClient      autogen_client.Client
	DefaultModelConfig types.NamespacedName
	DatabaseService    database.Client
}

// NewHandlers creates a new Handlers instance with all handler components
func NewHandlers(kubeClient client.Client, autogenClient autogen_client.Client, defaultModelConfig types.NamespacedName, dbService database.Client, watchedNamespaces []string) *Handlers {
	base := &Base{
		KubeClient:         kubeClient,
		AutogenClient:      autogenClient,
		DefaultModelConfig: defaultModelConfig,
		DatabaseService:    dbService,
	}

	return &Handlers{
		Health:      NewHealthHandler(),
		ModelConfig: NewModelConfigHandler(base),
		Model:       NewModelHandler(base),
		Provider:    NewProviderHandler(base),
		Sessions:    NewSessionsHandler(base),
		Agents:      NewAgentsHandler(base),
		Tools:       NewToolsHandler(base),
		ToolServers: NewToolServersHandler(base),
		Memory:      NewMemoryHandler(base),
		Feedback:    NewFeedbackHandler(base),
		Namespaces:  NewNamespacesHandler(base, watchedNamespaces),
	}
}
