package models

import (
	"net/http"

	"time"

	"github.com/vmihailenco/taskq/v3"
)

// HandlerInterface defines the methods a Handler should define
type HandlerInterface interface {
	ServerVersionHandler(w http.ResponseWriter, r *http.Request)

	ProviderMiddleware(http.Handler) http.Handler
	AuthMiddleware(http.Handler) http.Handler
	SessionInjectorMiddleware(func(http.ResponseWriter, *http.Request, *Preference, *User, Provider)) http.Handler

	ProviderHandler(w http.ResponseWriter, r *http.Request)
	ProvidersHandler(w http.ResponseWriter, r *http.Request)
	ProviderUIHandler(w http.ResponseWriter, r *http.Request)
	ProviderCapabilityHandler(w http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	ProviderComponentsHandler(w http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)

	TokenHandler(w http.ResponseWriter, r *http.Request, provider Provider, fromMiddleWare bool)
	LoginHandler(w http.ResponseWriter, r *http.Request, provider Provider, fromMiddleWare bool)
	LogoutHandler(w http.ResponseWriter, req *http.Request, provider Provider)
	UserHandler(w http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)

	K8SConfigHandler(w http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	GetContextsFromK8SConfig(w http.ResponseWriter, req *http.Request)
	KubernetesPingHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)

	LoadTestHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	LoadTestUsingSMPHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	CollectStaticMetrics(config *SubmitMetricsConfig) error
	FetchResultsHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	FetchAllResultsHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	GetResultHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	GetSMPServiceMeshes(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)

	FetchSmiResultsHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)

	MeshAdapterConfigHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	MeshOpsHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	GetAllAdaptersHandler(w http.ResponseWriter, req *http.Request, provider Provider)
	EventStreamHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	AdapterPingHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)

	GrafanaConfigHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	GrafanaBoardsHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	GrafanaQueryHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	GrafanaQueryRangeHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	GrafanaPingHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	SaveSelectedGrafanaBoardsHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)

	ScanPromGrafanaHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	ScanPrometheusHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	ScanGrafanaHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	PrometheusConfigHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	GrafanaBoardImportForPrometheusHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	PrometheusQueryHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	PrometheusQueryRangeHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	PrometheusPingHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	PrometheusStaticBoardHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	SaveSelectedPrometheusBoardsHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)

	LoadTestPrefencesHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	AnonymousStatsHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)

	UserTestPreferenceHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	UserTestPreferenceStore(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	UserTestPreferenceGet(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	UserTestPreferenceDelete(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)

	SavePerformanceProfileHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	GetPerformanceProfilesHandler(w http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	GetPerformanceProfileHandler(w http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	DeletePerformanceProfileHandler(w http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)

	SessionSyncHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)

	PatternFileHandler(rw http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	ExportPatternFile(rw http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	ImportPatternFile(rw http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	ImportPatternFileGithub(rw http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	OAMRegisterHandler(rw http.ResponseWriter, r *http.Request)
	PatternFileRequestHandler(rw http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	DeleteMesheryPatternHandler(rw http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	GetMesheryPatternHandler(rw http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)

	GraphqlSystemHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)

	ExtensionsEndpointHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	LoadExtensionFromPackage(w http.ResponseWriter, req *http.Request, provider Provider) error

	SaveScheduleHandler(w http.ResponseWriter, req *http.Request, prefObj *Preference, user *User, provider Provider)
	GetSchedulesHandler(w http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	GetScheduleHandler(w http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
	DeleteScheduleHandler(w http.ResponseWriter, r *http.Request, prefObj *Preference, user *User, provider Provider)
}

// HandlerConfig holds all the config pieces needed by handler methods
type HandlerConfig struct {
	// SessionName   string
	// RefCookieName string

	// SessionStore sessions.Store

	AdapterTracker AdaptersTrackerInterface
	QueryTracker   QueryTrackerInterface

	Queue taskq.Queue

	KubeConfigFolder string

	GrafanaClient         *GrafanaClient
	GrafanaClientForQuery *GrafanaClient

	PrometheusClient         *PrometheusClient
	PrometheusClientForQuery *PrometheusClient

	Providers              map[string]Provider
	ProviderCookieName     string
	ProviderCookieDuration time.Duration
}

// SubmitMetricsConfig is used to store config used for submitting metrics
type SubmitMetricsConfig struct {
	TestUUID, ResultID, PromURL string
	StartTime, EndTime          time.Time
	// TokenKey,
	TokenVal string
	Provider Provider
}
