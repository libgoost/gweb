package gweb

type Config struct {
	ListenPort string `json:"listen_port"`
	MetricPort string `json:"metric_port"`
	RepoRoot   string `json:"repo_root"`
}
